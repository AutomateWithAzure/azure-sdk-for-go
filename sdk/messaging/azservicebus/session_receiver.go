// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/go-amqp"
)

// SessionReceiver is a Receiver that handles sessions.
type SessionReceiver struct {
	*Receiver

	sessionID   *string
	lockedUntil time.Time
}

// SessionReceiverOptions contains options for the `Client.AcceptSessionForQueue/Subscription` or `Client.AcceptNextSessionForQueue/Subscription`
// functions.
type SessionReceiverOptions struct {
	// ReceiveMode controls when a message is deleted from Service Bus.
	//
	// `azservicebus.PeekLock` is the default. The message is locked, preventing multiple
	// receivers from processing the message at once. You control the lock state of the message
	// using one of the message settlement functions like SessionReceiver.CompleteMessage(), which removes
	// it from Service Bus, or SessionReceiver..AbandonMessage(), which makes it available again.
	//
	// `azservicebus.ReceiveAndDelete` causes Service Bus to remove the message as soon
	// as it's received.
	//
	// More information about receive modes:
	// https://docs.microsoft.com/azure/service-bus-messaging/message-transfers-locks-settlement#settling-receive-operations
	ReceiveMode ReceiveMode
}

func toReceiverOptions(sropts *SessionReceiverOptions) *ReceiverOptions {
	if sropts == nil {
		return nil
	}

	return &ReceiverOptions{
		ReceiveMode: sropts.ReceiveMode,
	}
}

func newSessionReceiver(ctx context.Context, sessionID *string, ns internal.NamespaceWithNewAMQPLinks, entity *entity, cleanupOnClose func(), options *ReceiverOptions) (*SessionReceiver, error) {
	const sessionFilterName = "com.microsoft:session-filter"
	const code = uint64(0x00000137000000C)

	sessionReceiver := &SessionReceiver{
		sessionID:   sessionID,
		lockedUntil: time.Time{},
	}

	var err error

	sessionReceiver.Receiver, err = newReceiver(ns, entity, cleanupOnClose, options, func(ctx context.Context, session internal.AMQPSession) (internal.AMQPSenderCloser, internal.AMQPReceiverCloser, error) {
		linkOptions := createLinkOptions(sessionReceiver.Receiver.receiveMode, sessionReceiver.amqpLinks.EntityPath())

		if sessionID == nil {
			linkOptions = append(linkOptions, amqp.LinkSourceFilter(sessionFilterName, code, nil))
		} else {
			linkOptions = append(linkOptions, amqp.LinkSourceFilter(sessionFilterName, code, sessionID))
		}

		_, link, err := createReceiverLink(ctx, session, linkOptions)

		if err != nil {
			return nil, nil, err
		}

		// check the session ID that came back - if we asked for a named session ID and didn't get it then
		// we failed to get the lock.
		// if we specified nil then we can _set_ our internally held session ID now that we know the value.
		receivedSessionID := link.LinkSourceFilterValue(sessionFilterName)
		asStr, ok := receivedSessionID.(string)

		if !ok || (sessionID != nil && asStr != *sessionID) {
			return nil, nil, fmt.Errorf("invalid type/value for returned sessionID(type:%T, value:%v)", receivedSessionID, receivedSessionID)
		}

		sessionReceiver.sessionID = &asStr
		return nil, link, nil
	})

	if err != nil {
		return nil, err
	}

	// temp workaround until we expose the session expiration time from the receiver in go-amqp
	if err := sessionReceiver.RenewSessionLock(ctx); err != nil {
		_ = sessionReceiver.Close(context.Background())
		return nil, err
	}

	return sessionReceiver, nil
}

// SessionID is the session ID for this SessionReceiver.
func (sr *SessionReceiver) SessionID() string {
	// return the ultimately assigned session ID for this link (anonymous will get it from the
	// link filter options, non-anonymous is set in newSessionReceiver)
	return *sr.sessionID
}

// LockedUntil is the time the lock on this session expires.
// The lock can be renewed using `SessionReceiver.RenewSessionLock`.
func (sr *SessionReceiver) LockedUntil() time.Time {
	return sr.lockedUntil
}

// RenewSessionLock renews this session's lock. The new expiration time is available
// using `LockedUntil`.
func (sr *SessionReceiver) RenewSessionLock(ctx context.Context) error {
	_, _, mgmt, _, err := sr.amqpLinks.Get(ctx)

	if err != nil {
		return err
	}

	newLockedUntil, err := mgmt.RenewSessionLock(ctx, *sr.sessionID)

	if err != nil {
		return err
	}

	sr.lockedUntil = newLockedUntil
	return nil
}

// init ensures the link was created, guaranteeing that we get our expected session lock.
func (sr *SessionReceiver) init(ctx context.Context) error {
	// initialize the links
	_, _, _, _, err := sr.amqpLinks.Get(ctx)
	return err
}
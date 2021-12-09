//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getComputeOneSkuUsages.json
func ExampleQuotaClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewQuotaClient(cred, nil)
	_, err = client.Get(ctx,
		"<subscription-id>",
		"<provider-id>",
		"<location>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/putComputeOneSkuQuotaRequest.json
func ExampleQuotaClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewQuotaClient(cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<subscription-id>",
		"<provider-id>",
		"<location>",
		"<resource-name>",
		armreservations.CurrentQuotaLimitBase{
			Properties: &armreservations.QuotaProperties{
				Name: &armreservations.ResourceName{
					Value: to.StringPtr("<value>"),
				},
				Limit: to.Int32Ptr(200),
				Unit:  to.StringPtr("<unit>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("QuotaRequestOneResourceSubmitResponse.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/patchComputeQuotaRequest.json
func ExampleQuotaClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewQuotaClient(cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<subscription-id>",
		"<provider-id>",
		"<location>",
		"<resource-name>",
		armreservations.CurrentQuotaLimitBase{
			Properties: &armreservations.QuotaProperties{
				Name: &armreservations.ResourceName{
					Value: to.StringPtr("<value>"),
				},
				Limit: to.Int32Ptr(200),
				Unit:  to.StringPtr("<unit>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("QuotaRequestOneResourceSubmitResponse.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getComputeUsages.json
func ExampleQuotaClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewQuotaClient(cred, nil)
	pager := client.List("<subscription-id>",
		"<provider-id>",
		"<location>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}
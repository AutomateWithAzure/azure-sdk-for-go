//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/create_update.json
func ExampleGatewayClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabricmesh.NewGatewayClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<gateway-resource-name>",
		armservicefabricmesh.GatewayResourceDescription{
			TrackedResource: armservicefabricmesh.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags:     map[string]*string{},
			},
			Properties: &armservicefabricmesh.GatewayResourceProperties{
				GatewayProperties: armservicefabricmesh.GatewayProperties{
					Description: to.StringPtr("<description>"),
					DestinationNetwork: &armservicefabricmesh.NetworkRef{
						Name: to.StringPtr("<name>"),
					},
					SourceNetwork: &armservicefabricmesh.NetworkRef{
						Name: to.StringPtr("<name>"),
					},
					TCP: []*armservicefabricmesh.TCPConfig{
						{
							Name: to.StringPtr("<name>"),
							Destination: &armservicefabricmesh.GatewayDestination{
								ApplicationName: to.StringPtr("<application-name>"),
								EndpointName:    to.StringPtr("<endpoint-name>"),
								ServiceName:     to.StringPtr("<service-name>"),
							},
							Port: to.Int32Ptr(80),
						}},
					HTTP: []*armservicefabricmesh.HTTPConfig{
						{
							Name: to.StringPtr("<name>"),
							Hosts: []*armservicefabricmesh.HTTPHostConfig{
								{
									Name: to.StringPtr("<name>"),
									Routes: []*armservicefabricmesh.HTTPRouteConfig{
										{
											Name: to.StringPtr("<name>"),
											Destination: &armservicefabricmesh.GatewayDestination{
												ApplicationName: to.StringPtr("<application-name>"),
												EndpointName:    to.StringPtr("<endpoint-name>"),
												ServiceName:     to.StringPtr("<service-name>"),
											},
											Match: &armservicefabricmesh.HTTPRouteMatchRule{
												Path: &armservicefabricmesh.HTTPRouteMatchPath{
													Type:    armservicefabricmesh.PathMatchTypePrefix.ToPtr(),
													Rewrite: to.StringPtr("<rewrite>"),
													Value:   to.StringPtr("<value>"),
												},
												Headers: []*armservicefabricmesh.HTTPRouteMatchHeader{
													{
														Name:  to.StringPtr("<name>"),
														Type:  armservicefabricmesh.HeaderMatchTypeExact.ToPtr(),
														Value: to.StringPtr("<value>"),
													}},
											},
										}},
								}},
							Port: to.Int32Ptr(8081),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GatewayResourceDescription.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/get.json
func ExampleGatewayClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabricmesh.NewGatewayClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<gateway-resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GatewayResourceDescription.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/delete.json
func ExampleGatewayClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabricmesh.NewGatewayClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<gateway-resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/list_byResourceGroup.json
func ExampleGatewayClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabricmesh.NewGatewayClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("GatewayResourceDescription.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/gateways/list_bySubscriptionId.json
func ExampleGatewayClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabricmesh.NewGatewayClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("GatewayResourceDescription.ID: %s\n", *v.ID)
		}
	}
}
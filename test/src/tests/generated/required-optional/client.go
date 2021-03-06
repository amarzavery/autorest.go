// Package optionalgroup implements the Azure ARM Optionalgroup service API version 1.0.0.
//
// Test Infrastructure for AutoRest
package optionalgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Optionalgroup
	DefaultBaseURI = "http://localhost"
)

// ManagementClient is the base client for Optionalgroup.
type ManagementClient struct {
	autorest.Client
	BaseURI             string
	RequiredGlobalPath  string
	RequiredGlobalQuery string
	OptionalGlobalQuery *int32
}

// New creates an instance of the ManagementClient client.
func New(requiredGlobalPath string, requiredGlobalQuery string, optionalGlobalQuery *int32) ManagementClient {
	return NewWithBaseURI(DefaultBaseURI, requiredGlobalPath, requiredGlobalQuery, optionalGlobalQuery)
}

// NewWithBaseURI creates an instance of the ManagementClient client.
func NewWithBaseURI(baseURI string, requiredGlobalPath string, requiredGlobalQuery string, optionalGlobalQuery *int32) ManagementClient {
	return ManagementClient{
		Client:              autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:             baseURI,
		RequiredGlobalPath:  requiredGlobalPath,
		RequiredGlobalQuery: requiredGlobalQuery,
		OptionalGlobalQuery: optionalGlobalQuery,
	}
}

//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// APIKey - An API key used for authenticating with a configuration store endpoint.
type APIKey struct {
	// READ-ONLY; A connection string that can be used by supporting clients for authentication.
	ConnectionString *string `json:"connectionString,omitempty" azure:"ro"`

	// READ-ONLY; The key ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The last time any of the key's properties were modified.
	LastModified *time.Time `json:"lastModified,omitempty" azure:"ro"`

	// READ-ONLY; A name for the key describing its usage.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Whether this key can only be used for read operations.
	ReadOnly *bool `json:"readOnly,omitempty" azure:"ro"`

	// READ-ONLY; The value of the key that is used for authentication purposes.
	Value *string `json:"value,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type APIKey.
func (a APIKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "connectionString", a.ConnectionString)
	populate(objectMap, "id", a.ID)
	populateTimeRFC3339(objectMap, "lastModified", a.LastModified)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "readOnly", a.ReadOnly)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type APIKey.
func (a *APIKey) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectionString":
			err = unpopulate(val, &a.ConnectionString)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, &a.ID)
			delete(rawMsg, key)
		case "lastModified":
			err = unpopulateTimeRFC3339(val, &a.LastModified)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &a.Name)
			delete(rawMsg, key)
		case "readOnly":
			err = unpopulate(val, &a.ReadOnly)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, &a.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// APIKeyListResult - The result of a request to list API keys.
type APIKeyListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string `json:"nextLink,omitempty"`

	// The collection value.
	Value []*APIKey `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type APIKeyListResult.
func (a APIKeyListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// CheckNameAvailabilityParameters - Parameters used for checking whether a resource name is available.
type CheckNameAvailabilityParameters struct {
	// REQUIRED; The name to check for availability.
	Name *string `json:"name,omitempty"`

	// REQUIRED; The resource type to check for name availability.
	Type *ConfigurationResourceType `json:"type,omitempty"`
}

// ConfigurationStore - The configuration store along with all resource properties. The Configuration Store will have all information to begin utilizing
// it.
type ConfigurationStore struct {
	TrackedResource
	// REQUIRED; The sku of the configuration store.
	SKU *SKU `json:"sku,omitempty"`

	// The managed identity information, if configured.
	Identity *ResourceIdentity `json:"identity,omitempty"`

	// The properties of a configuration store.
	Properties *ConfigurationStoreProperties `json:"properties,omitempty"`

	// READ-ONLY; Resource system metadata.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationStore.
func (c ConfigurationStore) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	c.TrackedResource.marshalInternal(objectMap)
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "systemData", c.SystemData)
	return json.Marshal(objectMap)
}

// ConfigurationStoreListResult - The result of a request to list configuration stores.
type ConfigurationStoreListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string `json:"nextLink,omitempty"`

	// The collection value.
	Value []*ConfigurationStore `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationStoreListResult.
func (c ConfigurationStoreListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// ConfigurationStoreProperties - The properties of a configuration store.
type ConfigurationStoreProperties struct {
	// Disables all authentication methods other than AAD authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// The encryption settings of the configuration store.
	Encryption *EncryptionProperties `json:"encryption,omitempty"`

	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// READ-ONLY; The creation date of configuration store.
	CreationDate *time.Time `json:"creationDate,omitempty" azure:"ro"`

	// READ-ONLY; The DNS endpoint where the configuration store API will be available.
	Endpoint *string `json:"endpoint,omitempty" azure:"ro"`

	// READ-ONLY; The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections []*PrivateEndpointConnectionReference `json:"privateEndpointConnections,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning state of the configuration store.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationStoreProperties.
func (c ConfigurationStoreProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "creationDate", c.CreationDate)
	populate(objectMap, "disableLocalAuth", c.DisableLocalAuth)
	populate(objectMap, "encryption", c.Encryption)
	populate(objectMap, "endpoint", c.Endpoint)
	populate(objectMap, "privateEndpointConnections", c.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", c.PublicNetworkAccess)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConfigurationStoreProperties.
func (c *ConfigurationStoreProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "creationDate":
			err = unpopulateTimeRFC3339(val, &c.CreationDate)
			delete(rawMsg, key)
		case "disableLocalAuth":
			err = unpopulate(val, &c.DisableLocalAuth)
			delete(rawMsg, key)
		case "encryption":
			err = unpopulate(val, &c.Encryption)
			delete(rawMsg, key)
		case "endpoint":
			err = unpopulate(val, &c.Endpoint)
			delete(rawMsg, key)
		case "privateEndpointConnections":
			err = unpopulate(val, &c.PrivateEndpointConnections)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &c.ProvisioningState)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, &c.PublicNetworkAccess)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ConfigurationStorePropertiesUpdateParameters - The properties for updating a configuration store.
type ConfigurationStorePropertiesUpdateParameters struct {
	// Disables all authentication methods other than AAD authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// The encryption settings of the configuration store.
	Encryption *EncryptionProperties `json:"encryption,omitempty"`

	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}

// ConfigurationStoreUpdateParameters - The parameters for updating a configuration store.
type ConfigurationStoreUpdateParameters struct {
	// The managed identity information for the configuration store.
	Identity *ResourceIdentity `json:"identity,omitempty"`

	// The properties for updating a configuration store.
	Properties *ConfigurationStorePropertiesUpdateParameters `json:"properties,omitempty"`

	// The SKU of the configuration store.
	SKU *SKU `json:"sku,omitempty"`

	// The ARM resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationStoreUpdateParameters.
func (c ConfigurationStoreUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", c.Identity)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// ConfigurationStoresBeginCreateOptions contains the optional parameters for the ConfigurationStores.BeginCreate method.
type ConfigurationStoresBeginCreateOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationStoresBeginDeleteOptions contains the optional parameters for the ConfigurationStores.BeginDelete method.
type ConfigurationStoresBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationStoresBeginUpdateOptions contains the optional parameters for the ConfigurationStores.BeginUpdate method.
type ConfigurationStoresBeginUpdateOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationStoresGetOptions contains the optional parameters for the ConfigurationStores.Get method.
type ConfigurationStoresGetOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationStoresListByResourceGroupOptions contains the optional parameters for the ConfigurationStores.ListByResourceGroup method.
type ConfigurationStoresListByResourceGroupOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains a nextLink element, the
	// value of the nextLink element will include a skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// ConfigurationStoresListKeysOptions contains the optional parameters for the ConfigurationStores.ListKeys method.
type ConfigurationStoresListKeysOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains a nextLink element, the
	// value of the nextLink element will include a skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// ConfigurationStoresListOptions contains the optional parameters for the ConfigurationStores.List method.
type ConfigurationStoresListOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains a nextLink element, the
	// value of the nextLink element will include a skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// ConfigurationStoresRegenerateKeyOptions contains the optional parameters for the ConfigurationStores.RegenerateKey method.
type ConfigurationStoresRegenerateKeyOptions struct {
	// placeholder for future optional parameters
}

// EncryptionProperties - The encryption settings for a configuration store.
type EncryptionProperties struct {
	// Key vault properties.
	KeyVaultProperties *KeyVaultProperties `json:"keyVaultProperties,omitempty"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info map[string]interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetails - The details of the error.
type ErrorDetails struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; Error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Error message indicating why the operation failed.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetails.
func (e ErrorDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// ErrorResponse - Error response indicates that the service is not able to process the incoming request. The reason is provided in the error message.
// Implements the error and azcore.HTTPResponse interfaces.
type ErrorResponse struct {
	raw string
	// The details of the error.
	InnerError *ErrorDetails `json:"error,omitempty"`
}

// Error implements the error interface for type ErrorResponse.
// The contents of the error text are not contractual and subject to change.
func (e ErrorResponse) Error() string {
	return e.raw
}

// KeyValue - The key-value resource along with all resource properties.
type KeyValue struct {
	// All key-value properties.
	Properties *KeyValueProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// KeyValueListResult - The result of a request to list key-values.
type KeyValueListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string `json:"nextLink,omitempty"`

	// The collection value.
	Value []*KeyValue `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type KeyValueListResult.
func (k KeyValueListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", k.NextLink)
	populate(objectMap, "value", k.Value)
	return json.Marshal(objectMap)
}

// KeyValueProperties - All key-value properties.
type KeyValueProperties struct {
	// The content type of the key-value's value. Providing a proper content-type can enable transformations of values when they are retrieved by applications.
	ContentType *string `json:"contentType,omitempty"`

	// A dictionary of tags that can help identify what a key-value may be applicable for.
	Tags map[string]*string `json:"tags,omitempty"`

	// The value of the key-value.
	Value *string `json:"value,omitempty"`

	// READ-ONLY; An ETag indicating the state of a key-value within a configuration store.
	ETag *string `json:"eTag,omitempty" azure:"ro"`

	// READ-ONLY; The primary identifier of a key-value. The key is used in unison with the label to uniquely identify a key-value.
	Key *string `json:"key,omitempty" azure:"ro"`

	// READ-ONLY; A value used to group key-values. The label is used in unison with the key to uniquely identify a key-value.
	Label *string `json:"label,omitempty" azure:"ro"`

	// READ-ONLY; The last time a modifying operation was performed on the given key-value.
	LastModified *time.Time `json:"lastModified,omitempty" azure:"ro"`

	// READ-ONLY; A value indicating whether the key-value is locked. A locked key-value may not be modified until it is unlocked.
	Locked *bool `json:"locked,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type KeyValueProperties.
func (k KeyValueProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contentType", k.ContentType)
	populate(objectMap, "eTag", k.ETag)
	populate(objectMap, "key", k.Key)
	populate(objectMap, "label", k.Label)
	populateTimeRFC3339(objectMap, "lastModified", k.LastModified)
	populate(objectMap, "locked", k.Locked)
	populate(objectMap, "tags", k.Tags)
	populate(objectMap, "value", k.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type KeyValueProperties.
func (k *KeyValueProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contentType":
			err = unpopulate(val, &k.ContentType)
			delete(rawMsg, key)
		case "eTag":
			err = unpopulate(val, &k.ETag)
			delete(rawMsg, key)
		case "key":
			err = unpopulate(val, &k.Key)
			delete(rawMsg, key)
		case "label":
			err = unpopulate(val, &k.Label)
			delete(rawMsg, key)
		case "lastModified":
			err = unpopulateTimeRFC3339(val, &k.LastModified)
			delete(rawMsg, key)
		case "locked":
			err = unpopulate(val, &k.Locked)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &k.Tags)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, &k.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// KeyValuesBeginDeleteOptions contains the optional parameters for the KeyValues.BeginDelete method.
type KeyValuesBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// KeyValuesCreateOrUpdateOptions contains the optional parameters for the KeyValues.CreateOrUpdate method.
type KeyValuesCreateOrUpdateOptions struct {
	// The parameters for creating a key-value.
	KeyValueParameters *KeyValue
}

// KeyValuesGetOptions contains the optional parameters for the KeyValues.Get method.
type KeyValuesGetOptions struct {
	// placeholder for future optional parameters
}

// KeyValuesListByConfigurationStoreOptions contains the optional parameters for the KeyValues.ListByConfigurationStore method.
type KeyValuesListByConfigurationStoreOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains a nextLink element, the
	// value of the nextLink element will include a skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// KeyVaultProperties - Settings concerning key vault encryption for a configuration store.
type KeyVaultProperties struct {
	// The client id of the identity which will be used to access key vault.
	IdentityClientID *string `json:"identityClientId,omitempty"`

	// The URI of the key vault key used to encrypt data.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
}

// LogSpecification - Specifications of the Log for Azure Monitoring
type LogSpecification struct {
	// Blob duration of the log
	BlobDuration *string `json:"blobDuration,omitempty"`

	// Localized friendly display name of the log
	DisplayName *string `json:"displayName,omitempty"`

	// Name of the log
	Name *string `json:"name,omitempty"`
}

// MetricDimension - Specifications of the Dimension of metrics
type MetricDimension struct {
	// Localized friendly display name of the dimension
	DisplayName *string `json:"displayName,omitempty"`

	// Internal name of the dimension.
	InternalName *string `json:"internalName,omitempty"`

	// Name of the dimension
	Name *string `json:"name,omitempty"`
}

// MetricSpecification - Specifications of the Metrics for Azure Monitoring
type MetricSpecification struct {
	// Only provide one value for this field. Valid values: Average, Minimum, Maximum, Total, Count.
	AggregationType *string `json:"aggregationType,omitempty"`

	// Dimensions of the metric
	Dimensions []*MetricDimension `json:"dimensions,omitempty"`

	// Localized friendly description of the metric
	DisplayDescription *string `json:"displayDescription,omitempty"`

	// Localized friendly display name of the metric
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. If set to true, then zero will be returned for time duration where no metric is emitted/published.
	FillGapWithZero *bool `json:"fillGapWithZero,omitempty"`

	// Internal metric name.
	InternalMetricName *string `json:"internalMetricName,omitempty"`

	// Name of the metric
	Name *string `json:"name,omitempty"`

	// Unit that makes sense for the metric
	Unit *string `json:"unit,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MetricSpecification.
func (m MetricSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aggregationType", m.AggregationType)
	populate(objectMap, "dimensions", m.Dimensions)
	populate(objectMap, "displayDescription", m.DisplayDescription)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "fillGapWithZero", m.FillGapWithZero)
	populate(objectMap, "internalMetricName", m.InternalMetricName)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// NameAvailabilityStatus - The result of a request to check the availability of a resource name.
type NameAvailabilityStatus struct {
	// READ-ONLY; If any, the error message that provides more detail for the reason that the name is not available.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The value indicating whether the resource name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty" azure:"ro"`

	// READ-ONLY; If any, the reason that the name is not available.
	Reason *string `json:"reason,omitempty" azure:"ro"`
}

// OperationDefinition - The definition of a configuration store operation.
type OperationDefinition struct {
	// The display information for the configuration store operation.
	Display *OperationDefinitionDisplay `json:"display,omitempty"`

	// Indicates whether the operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`

	// Origin of the operation
	Origin *string `json:"origin,omitempty"`

	// Properties of the operation
	Properties *OperationProperties `json:"properties,omitempty"`
}

// OperationDefinitionDisplay - The display information for a configuration store operation.
type OperationDefinitionDisplay struct {
	// The description for the operation.
	Description *string `json:"description,omitempty"`

	// The operation that users can perform.
	Operation *string `json:"operation,omitempty"`

	// The resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`

	// READ-ONLY; The resource provider name: Microsoft App Configuration."
	Provider *string `json:"provider,omitempty" azure:"ro"`
}

// OperationDefinitionListResult - The result of a request to list configuration store operations.
type OperationDefinitionListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string `json:"nextLink,omitempty"`

	// The collection value.
	Value []*OperationDefinition `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationDefinitionListResult.
func (o OperationDefinitionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OperationProperties - Extra Operation properties
type OperationProperties struct {
	// Service specifications of the operation
	ServiceSpecification *ServiceSpecification `json:"serviceSpecification,omitempty"`
}

// OperationsCheckNameAvailabilityOptions contains the optional parameters for the Operations.CheckNameAvailability method.
type OperationsCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// OperationsListOptions contains the optional parameters for the Operations.List method.
type OperationsListOptions struct {
	// A skip token is used to continue retrieving items after an operation returns a partial result. If a previous response contains a nextLink element, the
	// value of the nextLink element will include a skipToken parameter that specifies a starting point to use for subsequent calls.
	SkipToken *string
}

// PrivateEndpoint - Private endpoint which a connection belongs to.
type PrivateEndpoint struct {
	// The resource Id for private endpoint
	ID *string `json:"id,omitempty"`
}

// PrivateEndpointConnection - A private endpoint connection
type PrivateEndpointConnection struct {
	// The properties of a private endpoint.
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionListResult - A list of private endpoint connections
type PrivateEndpointConnectionListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string `json:"nextLink,omitempty"`

	// The collection value.
	Value []*PrivateEndpointConnection `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionListResult.
func (p PrivateEndpointConnectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// PrivateEndpointConnectionProperties - Properties of a private endpoint connection.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`

	// The resource of private endpoint.
	PrivateEndpoint *PrivateEndpoint `json:"privateEndpoint,omitempty"`

	// READ-ONLY; The provisioning status of the private endpoint connection.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionReference - A reference to a related private endpoint connection.
type PrivateEndpointConnectionReference struct {
	// The properties of a private endpoint connection.
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionsBeginCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnections.BeginCreateOrUpdate method.
type PrivateEndpointConnectionsBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnections.BeginDelete method.
type PrivateEndpointConnectionsBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsGetOptions contains the optional parameters for the PrivateEndpointConnections.Get method.
type PrivateEndpointConnectionsGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsListByConfigurationStoreOptions contains the optional parameters for the PrivateEndpointConnections.ListByConfigurationStore
// method.
type PrivateEndpointConnectionsListByConfigurationStoreOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResource - A resource that supports private link capabilities.
type PrivateLinkResource struct {
	// Private link resource properties.
	Properties *PrivateLinkResourceProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateLinkResourceListResult - A list of private link resources.
type PrivateLinkResourceListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string `json:"nextLink,omitempty"`

	// The collection value.
	Value []*PrivateLinkResource `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceListResult.
func (p PrivateLinkResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// READ-ONLY; The private link resource group id.
	GroupID *string `json:"groupId,omitempty" azure:"ro"`

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string `json:"requiredMembers,omitempty" azure:"ro"`

	// READ-ONLY; The list of required DNS zone names of the private link resource.
	RequiredZoneNames []*string `json:"requiredZoneNames,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceProperties.
func (p PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	populate(objectMap, "requiredZoneNames", p.RequiredZoneNames)
	return json.Marshal(objectMap)
}

// PrivateLinkResourcesGetOptions contains the optional parameters for the PrivateLinkResources.Get method.
type PrivateLinkResourcesGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesListByConfigurationStoreOptions contains the optional parameters for the PrivateLinkResources.ListByConfigurationStore method.
type PrivateLinkResourcesListByConfigurationStoreOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkServiceConnectionState - The state of a private link service connection.
type PrivateLinkServiceConnectionState struct {
	// The private link service connection description.
	Description *string `json:"description,omitempty"`

	// The private link service connection status.
	Status *ConnectionStatus `json:"status,omitempty"`

	// READ-ONLY; Any action that is required beyond basic workflow (approve/ reject/ disconnect)
	ActionsRequired *ActionsRequired `json:"actionsRequired,omitempty" azure:"ro"`
}

// RegenerateKeyParameters - The parameters used to regenerate an API key.
type RegenerateKeyParameters struct {
	// The id of the key to regenerate.
	ID *string `json:"id,omitempty"`
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "type", r.Type)
}

// ResourceIdentity - An identity that can be associated with a resource.
type ResourceIdentity struct {
	// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities.
	// The type 'None' will remove any
	// identities.
	Type *IdentityType `json:"type,omitempty"`

	// The list of user-assigned identities associated with the resource. The user-assigned identity dictionary keys will be ARM resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]*UserIdentity `json:"userAssignedIdentities,omitempty"`

	// READ-ONLY; The principal id of the identity. This property will only be provided for a system-assigned identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant id associated with the resource's identity. This property will only be provided for a system-assigned identity.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ResourceIdentity.
func (r ResourceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", r.PrincipalID)
	populate(objectMap, "tenantId", r.TenantID)
	populate(objectMap, "type", r.Type)
	populate(objectMap, "userAssignedIdentities", r.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// SKU - Describes a configuration store SKU.
type SKU struct {
	// REQUIRED; The SKU name of the configuration store.
	Name *string `json:"name,omitempty"`
}

// ServiceSpecification - Service specification payload
type ServiceSpecification struct {
	// Specifications of the Log for Azure Monitoring
	LogSpecifications []*LogSpecification `json:"logSpecifications,omitempty"`

	// Specifications of the Metrics for Azure Monitoring
	MetricSpecifications []*MetricSpecification `json:"metricSpecifications,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ServiceSpecification.
func (s ServiceSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "logSpecifications", s.LogSpecifications)
	populate(objectMap, "metricSpecifications", s.MetricSpecifications)
	return json.Marshal(objectMap)
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags' and a 'location'
type TrackedResource struct {
	Resource
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	t.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (t TrackedResource) marshalInternal(objectMap map[string]interface{}) {
	t.Resource.marshalInternal(objectMap)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "tags", t.Tags)
}

// UserIdentity - A resource identity that is managed by the user of the service.
type UserIdentity struct {
	// READ-ONLY; The client ID of the user-assigned identity.
	ClientID *string `json:"clientId,omitempty" azure:"ro"`

	// READ-ONLY; The principal ID of the user-assigned identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
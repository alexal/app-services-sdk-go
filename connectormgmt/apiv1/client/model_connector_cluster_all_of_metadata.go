/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.0.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
	"time"
)

// ConnectorClusterAllOfMetadata struct for ConnectorClusterAllOfMetadata
type ConnectorClusterAllOfMetadata struct {

	Owner *string `json:"owner,omitempty"`

	Name *string `json:"name,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`

	UpdatedAt *time.Time `json:"updated_at,omitempty"`

}

// NewConnectorClusterAllOfMetadata instantiates a new ConnectorClusterAllOfMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorClusterAllOfMetadata() *ConnectorClusterAllOfMetadata {
	this := ConnectorClusterAllOfMetadata{}
	return &this
}

// NewConnectorClusterAllOfMetadataWithDefaults instantiates a new ConnectorClusterAllOfMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorClusterAllOfMetadataWithDefaults() *ConnectorClusterAllOfMetadata {
	this := ConnectorClusterAllOfMetadata{}





	return &this
}


// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ConnectorClusterAllOfMetadata) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorClusterAllOfMetadata) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ConnectorClusterAllOfMetadata) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ConnectorClusterAllOfMetadata) SetOwner(v string) {
	o.Owner = &v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectorClusterAllOfMetadata) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorClusterAllOfMetadata) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectorClusterAllOfMetadata) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectorClusterAllOfMetadata) SetName(v string) {
	o.Name = &v
}


// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ConnectorClusterAllOfMetadata) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorClusterAllOfMetadata) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ConnectorClusterAllOfMetadata) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ConnectorClusterAllOfMetadata) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}


// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ConnectorClusterAllOfMetadata) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorClusterAllOfMetadata) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ConnectorClusterAllOfMetadata) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ConnectorClusterAllOfMetadata) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}


func (o ConnectorClusterAllOfMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
    
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
    
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
    
	return json.Marshal(toSerialize)
}

type NullableConnectorClusterAllOfMetadata struct {
	value *ConnectorClusterAllOfMetadata
	isSet bool
}

func (v NullableConnectorClusterAllOfMetadata) Get() *ConnectorClusterAllOfMetadata {
	return v.value
}

func (v *NullableConnectorClusterAllOfMetadata) Set(val *ConnectorClusterAllOfMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorClusterAllOfMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorClusterAllOfMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorClusterAllOfMetadata(val *ConnectorClusterAllOfMetadata) *NullableConnectorClusterAllOfMetadata {
	return &NullableConnectorClusterAllOfMetadata{value: val, isSet: true}
}

func (v NullableConnectorClusterAllOfMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorClusterAllOfMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


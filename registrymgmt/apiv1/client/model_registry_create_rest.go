/*
 * Service Registry Fleet Manager
 *
 * Managed Service Registry cloud.redhat.com API Management API that lets you create new registry instances. Registry is a datastore for standard event schemas and API designs. Service Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Registry is an Managed version of upstream project called Apicurio Registry. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.
 *
 * API version: 0.0.5
 * Contact: rhosak-eval-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registrymgmtclient

import (
	"encoding/json"
)

// RegistryCreateRest Information used to create a new Service Registry instance within a multi-tenant deployment.
type RegistryCreateRest struct {

	// User-defined Registry name. Required. Does not have to be unique.
	Name *string `json:"name,omitempty"`

	// User-provided description of the new Registry instance. Not required.
	Description *string `json:"description,omitempty"`

}

// NewRegistryCreateRest instantiates a new RegistryCreateRest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryCreateRest() *RegistryCreateRest {
	this := RegistryCreateRest{}
	return &this
}

// NewRegistryCreateRestWithDefaults instantiates a new RegistryCreateRest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryCreateRestWithDefaults() *RegistryCreateRest {
	this := RegistryCreateRest{}



	return &this
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *RegistryCreateRest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCreateRest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegistryCreateRest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegistryCreateRest) SetName(v string) {
	o.Name = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RegistryCreateRest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistryCreateRest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RegistryCreateRest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RegistryCreateRest) SetDescription(v string) {
	o.Description = &v
}


func (o RegistryCreateRest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
    
	return json.Marshal(toSerialize)
}

type NullableRegistryCreateRest struct {
	value *RegistryCreateRest
	isSet bool
}

func (v NullableRegistryCreateRest) Get() *RegistryCreateRest {
	return v.value
}

func (v *NullableRegistryCreateRest) Set(val *RegistryCreateRest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryCreateRest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryCreateRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryCreateRest(val *RegistryCreateRest) *NullableRegistryCreateRest {
	return &NullableRegistryCreateRest{value: val, isSet: true}
}

func (v NullableRegistryCreateRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryCreateRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


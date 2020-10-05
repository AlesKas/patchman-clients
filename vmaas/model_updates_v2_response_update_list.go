/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.20.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// UpdatesV2ResponseUpdateList struct for UpdatesV2ResponseUpdateList
type UpdatesV2ResponseUpdateList struct {
	AvailableUpdates *[]UpdatesResponseAvailableUpdates `json:"available_updates,omitempty"`
}

// NewUpdatesV2ResponseUpdateList instantiates a new UpdatesV2ResponseUpdateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatesV2ResponseUpdateList() *UpdatesV2ResponseUpdateList {
	this := UpdatesV2ResponseUpdateList{}
	return &this
}

// NewUpdatesV2ResponseUpdateListWithDefaults instantiates a new UpdatesV2ResponseUpdateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatesV2ResponseUpdateListWithDefaults() *UpdatesV2ResponseUpdateList {
	this := UpdatesV2ResponseUpdateList{}
	return &this
}

// GetAvailableUpdates returns the AvailableUpdates field value if set, zero value otherwise.
func (o *UpdatesV2ResponseUpdateList) GetAvailableUpdates() []UpdatesResponseAvailableUpdates {
	if o == nil || o.AvailableUpdates == nil {
		var ret []UpdatesResponseAvailableUpdates
		return ret
	}
	return *o.AvailableUpdates
}

// GetAvailableUpdatesOk returns a tuple with the AvailableUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatesV2ResponseUpdateList) GetAvailableUpdatesOk() (*[]UpdatesResponseAvailableUpdates, bool) {
	if o == nil || o.AvailableUpdates == nil {
		return nil, false
	}
	return o.AvailableUpdates, true
}

// HasAvailableUpdates returns a boolean if a field has been set.
func (o *UpdatesV2ResponseUpdateList) HasAvailableUpdates() bool {
	if o != nil && o.AvailableUpdates != nil {
		return true
	}

	return false
}

// SetAvailableUpdates gets a reference to the given []UpdatesResponseAvailableUpdates and assigns it to the AvailableUpdates field.
func (o *UpdatesV2ResponseUpdateList) SetAvailableUpdates(v []UpdatesResponseAvailableUpdates) {
	o.AvailableUpdates = &v
}

func (o UpdatesV2ResponseUpdateList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvailableUpdates != nil {
		toSerialize["available_updates"] = o.AvailableUpdates
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatesV2ResponseUpdateList struct {
	value *UpdatesV2ResponseUpdateList
	isSet bool
}

func (v NullableUpdatesV2ResponseUpdateList) Get() *UpdatesV2ResponseUpdateList {
	return v.value
}

func (v *NullableUpdatesV2ResponseUpdateList) Set(val *UpdatesV2ResponseUpdateList) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatesV2ResponseUpdateList) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatesV2ResponseUpdateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatesV2ResponseUpdateList(val *UpdatesV2ResponseUpdateList) *NullableUpdatesV2ResponseUpdateList {
	return &NullableUpdatesV2ResponseUpdateList{value: val, isSet: true}
}

func (v NullableUpdatesV2ResponseUpdateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatesV2ResponseUpdateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



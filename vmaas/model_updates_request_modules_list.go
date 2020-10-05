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

// UpdatesRequestModulesList struct for UpdatesRequestModulesList
type UpdatesRequestModulesList struct {
	ModuleName string `json:"module_name"`
	ModuleStream string `json:"module_stream"`
}

// NewUpdatesRequestModulesList instantiates a new UpdatesRequestModulesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatesRequestModulesList(moduleName string, moduleStream string, ) *UpdatesRequestModulesList {
	this := UpdatesRequestModulesList{}
	this.ModuleName = moduleName
	this.ModuleStream = moduleStream
	return &this
}

// NewUpdatesRequestModulesListWithDefaults instantiates a new UpdatesRequestModulesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatesRequestModulesListWithDefaults() *UpdatesRequestModulesList {
	this := UpdatesRequestModulesList{}
	return &this
}

// GetModuleName returns the ModuleName field value
func (o *UpdatesRequestModulesList) GetModuleName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value
// and a boolean to check if the value has been set.
func (o *UpdatesRequestModulesList) GetModuleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModuleName, true
}

// SetModuleName sets field value
func (o *UpdatesRequestModulesList) SetModuleName(v string) {
	o.ModuleName = v
}

// GetModuleStream returns the ModuleStream field value
func (o *UpdatesRequestModulesList) GetModuleStream() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ModuleStream
}

// GetModuleStreamOk returns a tuple with the ModuleStream field value
// and a boolean to check if the value has been set.
func (o *UpdatesRequestModulesList) GetModuleStreamOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModuleStream, true
}

// SetModuleStream sets field value
func (o *UpdatesRequestModulesList) SetModuleStream(v string) {
	o.ModuleStream = v
}

func (o UpdatesRequestModulesList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["module_name"] = o.ModuleName
	}
	if true {
		toSerialize["module_stream"] = o.ModuleStream
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatesRequestModulesList struct {
	value *UpdatesRequestModulesList
	isSet bool
}

func (v NullableUpdatesRequestModulesList) Get() *UpdatesRequestModulesList {
	return v.value
}

func (v *NullableUpdatesRequestModulesList) Set(val *UpdatesRequestModulesList) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatesRequestModulesList) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatesRequestModulesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatesRequestModulesList(val *UpdatesRequestModulesList) *NullableUpdatesRequestModulesList {
	return &NullableUpdatesRequestModulesList{value: val, isSet: true}
}

func (v NullableUpdatesRequestModulesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatesRequestModulesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



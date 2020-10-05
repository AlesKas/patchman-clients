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

// ErrataRequest struct for ErrataRequest
type ErrataRequest struct {
	Page *float32 `json:"page,omitempty"`
	PageSize *float32 `json:"page_size,omitempty"`
	ErrataList []string `json:"errata_list"`
	ModifiedSince *string `json:"modified_since,omitempty"`
	Type *OneOfstringarray `json:"type,omitempty"`
	Severity *OneOfstringarray `json:"severity,omitempty"`
}

// NewErrataRequest instantiates a new ErrataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrataRequest(errataList []string, ) *ErrataRequest {
	this := ErrataRequest{}
	this.ErrataList = errataList
	return &this
}

// NewErrataRequestWithDefaults instantiates a new ErrataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrataRequestWithDefaults() *ErrataRequest {
	this := ErrataRequest{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ErrataRequest) GetPage() float32 {
	if o == nil || o.Page == nil {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataRequest) GetPageOk() (*float32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ErrataRequest) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *ErrataRequest) SetPage(v float32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ErrataRequest) GetPageSize() float32 {
	if o == nil || o.PageSize == nil {
		var ret float32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataRequest) GetPageSizeOk() (*float32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ErrataRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given float32 and assigns it to the PageSize field.
func (o *ErrataRequest) SetPageSize(v float32) {
	o.PageSize = &v
}

// GetErrataList returns the ErrataList field value
func (o *ErrataRequest) GetErrataList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.ErrataList
}

// GetErrataListOk returns a tuple with the ErrataList field value
// and a boolean to check if the value has been set.
func (o *ErrataRequest) GetErrataListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrataList, true
}

// SetErrataList sets field value
func (o *ErrataRequest) SetErrataList(v []string) {
	o.ErrataList = v
}

// GetModifiedSince returns the ModifiedSince field value if set, zero value otherwise.
func (o *ErrataRequest) GetModifiedSince() string {
	if o == nil || o.ModifiedSince == nil {
		var ret string
		return ret
	}
	return *o.ModifiedSince
}

// GetModifiedSinceOk returns a tuple with the ModifiedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataRequest) GetModifiedSinceOk() (*string, bool) {
	if o == nil || o.ModifiedSince == nil {
		return nil, false
	}
	return o.ModifiedSince, true
}

// HasModifiedSince returns a boolean if a field has been set.
func (o *ErrataRequest) HasModifiedSince() bool {
	if o != nil && o.ModifiedSince != nil {
		return true
	}

	return false
}

// SetModifiedSince gets a reference to the given string and assigns it to the ModifiedSince field.
func (o *ErrataRequest) SetModifiedSince(v string) {
	o.ModifiedSince = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErrataRequest) GetType() OneOfstringarray {
	if o == nil || o.Type == nil {
		var ret OneOfstringarray
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataRequest) GetTypeOk() (*OneOfstringarray, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrataRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given OneOfstringarray and assigns it to the Type field.
func (o *ErrataRequest) SetType(v OneOfstringarray) {
	o.Type = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ErrataRequest) GetSeverity() OneOfstringarray {
	if o == nil || o.Severity == nil {
		var ret OneOfstringarray
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataRequest) GetSeverityOk() (*OneOfstringarray, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ErrataRequest) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given OneOfstringarray and assigns it to the Severity field.
func (o *ErrataRequest) SetSeverity(v OneOfstringarray) {
	o.Severity = &v
}

func (o ErrataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if true {
		toSerialize["errata_list"] = o.ErrataList
	}
	if o.ModifiedSince != nil {
		toSerialize["modified_since"] = o.ModifiedSince
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	return json.Marshal(toSerialize)
}

type NullableErrataRequest struct {
	value *ErrataRequest
	isSet bool
}

func (v NullableErrataRequest) Get() *ErrataRequest {
	return v.value
}

func (v *NullableErrataRequest) Set(val *ErrataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableErrataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableErrataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrataRequest(val *ErrataRequest) *NullableErrataRequest {
	return &NullableErrataRequest{value: val, isSet: true}
}

func (v NullableErrataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



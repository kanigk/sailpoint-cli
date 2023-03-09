/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// QueryResultFilter Allows the query results to be filtered by specifying a list of fields to include and/or exclude from the result documents.
type QueryResultFilter struct {
	// The list of field names to include in the result documents.
	Includes []string `json:"includes,omitempty"`
	// The list of field names to exclude from the result documents.
	Excludes []string `json:"excludes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryResultFilter QueryResultFilter

// NewQueryResultFilter instantiates a new QueryResultFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryResultFilter() *QueryResultFilter {
	this := QueryResultFilter{}
	return &this
}

// NewQueryResultFilterWithDefaults instantiates a new QueryResultFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryResultFilterWithDefaults() *QueryResultFilter {
	this := QueryResultFilter{}
	return &this
}

// GetIncludes returns the Includes field value if set, zero value otherwise.
func (o *QueryResultFilter) GetIncludes() []string {
	if o == nil || isNil(o.Includes) {
		var ret []string
		return ret
	}
	return o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResultFilter) GetIncludesOk() ([]string, bool) {
	if o == nil || isNil(o.Includes) {
		return nil, false
	}
	return o.Includes, true
}

// HasIncludes returns a boolean if a field has been set.
func (o *QueryResultFilter) HasIncludes() bool {
	if o != nil && !isNil(o.Includes) {
		return true
	}

	return false
}

// SetIncludes gets a reference to the given []string and assigns it to the Includes field.
func (o *QueryResultFilter) SetIncludes(v []string) {
	o.Includes = v
}

// GetExcludes returns the Excludes field value if set, zero value otherwise.
func (o *QueryResultFilter) GetExcludes() []string {
	if o == nil || isNil(o.Excludes) {
		var ret []string
		return ret
	}
	return o.Excludes
}

// GetExcludesOk returns a tuple with the Excludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResultFilter) GetExcludesOk() ([]string, bool) {
	if o == nil || isNil(o.Excludes) {
		return nil, false
	}
	return o.Excludes, true
}

// HasExcludes returns a boolean if a field has been set.
func (o *QueryResultFilter) HasExcludes() bool {
	if o != nil && !isNil(o.Excludes) {
		return true
	}

	return false
}

// SetExcludes gets a reference to the given []string and assigns it to the Excludes field.
func (o *QueryResultFilter) SetExcludes(v []string) {
	o.Excludes = v
}

func (o QueryResultFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Includes) {
		toSerialize["includes"] = o.Includes
	}
	if !isNil(o.Excludes) {
		toSerialize["excludes"] = o.Excludes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *QueryResultFilter) UnmarshalJSON(bytes []byte) (err error) {
	varQueryResultFilter := _QueryResultFilter{}

	if err = json.Unmarshal(bytes, &varQueryResultFilter); err == nil {
		*o = QueryResultFilter(varQueryResultFilter)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "includes")
		delete(additionalProperties, "excludes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryResultFilter struct {
	value *QueryResultFilter
	isSet bool
}

func (v NullableQueryResultFilter) Get() *QueryResultFilter {
	return v.value
}

func (v *NullableQueryResultFilter) Set(val *QueryResultFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryResultFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryResultFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryResultFilter(val *QueryResultFilter) *NullableQueryResultFilter {
	return &NullableQueryResultFilter{value: val, isSet: true}
}

func (v NullableQueryResultFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryResultFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


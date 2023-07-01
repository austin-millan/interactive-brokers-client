/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InlineResponse2006 struct for InlineResponse2006
type InlineResponse2006 struct {
	ApplicantId *float32 `json:"applicantId,omitempty"`
	Entities *[]IbcustEntityInfoEntities `json:"entities,omitempty"`
}

// NewInlineResponse2006 instantiates a new InlineResponse2006 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2006() *InlineResponse2006 {
	this := InlineResponse2006{}
	return &this
}

// NewInlineResponse2006WithDefaults instantiates a new InlineResponse2006 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2006WithDefaults() *InlineResponse2006 {
	this := InlineResponse2006{}
	return &this
}

// GetApplicantId returns the ApplicantId field value if set, zero value otherwise.
func (o *InlineResponse2006) GetApplicantId() float32 {
	if o == nil || o.ApplicantId == nil {
		var ret float32
		return ret
	}
	return *o.ApplicantId
}

// GetApplicantIdOk returns a tuple with the ApplicantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetApplicantIdOk() (*float32, bool) {
	if o == nil || o.ApplicantId == nil {
		return nil, false
	}
	return o.ApplicantId, true
}

// HasApplicantId returns a boolean if a field has been set.
func (o *InlineResponse2006) HasApplicantId() bool {
	if o != nil && o.ApplicantId != nil {
		return true
	}

	return false
}

// SetApplicantId gets a reference to the given float32 and assigns it to the ApplicantId field.
func (o *InlineResponse2006) SetApplicantId(v float32) {
	o.ApplicantId = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *InlineResponse2006) GetEntities() []IbcustEntityInfoEntities {
	if o == nil || o.Entities == nil {
		var ret []IbcustEntityInfoEntities
		return ret
	}
	return *o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006) GetEntitiesOk() (*[]IbcustEntityInfoEntities, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *InlineResponse2006) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []IbcustEntityInfoEntities and assigns it to the Entities field.
func (o *InlineResponse2006) SetEntities(v []IbcustEntityInfoEntities) {
	o.Entities = &v
}

func (o InlineResponse2006) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicantId != nil {
		toSerialize["applicantId"] = o.ApplicantId
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2006 struct {
	value *InlineResponse2006
	isSet bool
}

func (v NullableInlineResponse2006) Get() *InlineResponse2006 {
	return v.value
}

func (v *NullableInlineResponse2006) Set(val *InlineResponse2006) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2006) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2006) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2006(val *InlineResponse2006) *NullableInlineResponse2006 {
	return &NullableInlineResponse2006{value: val, isSet: true}
}

func (v NullableInlineResponse2006) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2006) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



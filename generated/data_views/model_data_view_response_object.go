/*
Data views

OpenAPI schema for data view endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data_views

import (
	"encoding/json"
)

// checks if the DataViewResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataViewResponseObject{}

// DataViewResponseObject struct for DataViewResponseObject
type DataViewResponseObject struct {
	DataView *DataViewResponseObjectDataView `json:"data_view,omitempty"`
}

// NewDataViewResponseObject instantiates a new DataViewResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataViewResponseObject() *DataViewResponseObject {
	this := DataViewResponseObject{}
	return &this
}

// NewDataViewResponseObjectWithDefaults instantiates a new DataViewResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataViewResponseObjectWithDefaults() *DataViewResponseObject {
	this := DataViewResponseObject{}
	return &this
}

// GetDataView returns the DataView field value if set, zero value otherwise.
func (o *DataViewResponseObject) GetDataView() DataViewResponseObjectDataView {
	if o == nil || IsNil(o.DataView) {
		var ret DataViewResponseObjectDataView
		return ret
	}
	return *o.DataView
}

// GetDataViewOk returns a tuple with the DataView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataViewResponseObject) GetDataViewOk() (*DataViewResponseObjectDataView, bool) {
	if o == nil || IsNil(o.DataView) {
		return nil, false
	}
	return o.DataView, true
}

// HasDataView returns a boolean if a field has been set.
func (o *DataViewResponseObject) HasDataView() bool {
	if o != nil && !IsNil(o.DataView) {
		return true
	}

	return false
}

// SetDataView gets a reference to the given DataViewResponseObjectDataView and assigns it to the DataView field.
func (o *DataViewResponseObject) SetDataView(v DataViewResponseObjectDataView) {
	o.DataView = &v
}

func (o DataViewResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataViewResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataView) {
		toSerialize["data_view"] = o.DataView
	}
	return toSerialize, nil
}

type NullableDataViewResponseObject struct {
	value *DataViewResponseObject
	isSet bool
}

func (v NullableDataViewResponseObject) Get() *DataViewResponseObject {
	return v.value
}

func (v *NullableDataViewResponseObject) Set(val *DataViewResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableDataViewResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableDataViewResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataViewResponseObject(val *DataViewResponseObject) *NullableDataViewResponseObject {
	return &NullableDataViewResponseObject{value: val, isSet: true}
}

func (v NullableDataViewResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataViewResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
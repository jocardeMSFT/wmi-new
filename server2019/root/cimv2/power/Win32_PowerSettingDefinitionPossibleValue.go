// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

// Win32_PowerSettingDefinitionPossibleValue struct
type Win32_PowerSettingDefinitionPossibleValue struct {
	CIM_SettingData

	//
	BinaryValue []uint8

	//
	SettingIndex uint32

	//
	StringValue string

	//
	UInt32Value uint32

	//
	UInt64Value uint64
}

// SetBinaryValue sets the value of BinaryValue for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) SetPropertyBinaryValue(value []uint8) (err error) {
	return instance.SetProperty("BinaryValue", value)
}

// GetBinaryValue gets the value of BinaryValue for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) GetPropertyBinaryValue() (value []uint8, err error) {
	retValue, err := instance.GetProperty("BinaryValue")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettingIndex sets the value of SettingIndex for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) SetPropertySettingIndex(value uint32) (err error) {
	return instance.SetProperty("SettingIndex", value)
}

// GetSettingIndex gets the value of SettingIndex for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) GetPropertySettingIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("SettingIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStringValue sets the value of StringValue for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) SetPropertyStringValue(value string) (err error) {
	return instance.SetProperty("StringValue", value)
}

// GetStringValue gets the value of StringValue for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) GetPropertyStringValue() (value string, err error) {
	retValue, err := instance.GetProperty("StringValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUInt32Value sets the value of UInt32Value for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) SetPropertyUInt32Value(value uint32) (err error) {
	return instance.SetProperty("UInt32Value", value)
}

// GetUInt32Value gets the value of UInt32Value for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) GetPropertyUInt32Value() (value uint32, err error) {
	retValue, err := instance.GetProperty("UInt32Value")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUInt64Value sets the value of UInt64Value for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) SetPropertyUInt64Value(value uint64) (err error) {
	return instance.SetProperty("UInt64Value", value)
}

// GetUInt64Value gets the value of UInt64Value for the instance
func (instance *Win32_PowerSettingDefinitionPossibleValue) GetPropertyUInt64Value() (value uint64, err error) {
	retValue, err := instance.GetProperty("UInt64Value")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
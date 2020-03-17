// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Patch struct
type Win32_Patch struct {
	Win32_MSIResource

	//
	Attributes uint16

	//
	File string

	//
	PatchSize uint32

	//
	ProductCode string

	//
	Sequence int16
}

// SetAttributes sets the value of Attributes for the instance
func (instance *Win32_Patch) SetPropertyAttributes(value uint16) (err error) {
	return instance.SetProperty("Attributes", value)
}

// GetAttributes gets the value of Attributes for the instance
func (instance *Win32_Patch) GetPropertyAttributes() (value uint16, err error) {
	retValue, err := instance.GetProperty("Attributes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFile sets the value of File for the instance
func (instance *Win32_Patch) SetPropertyFile(value string) (err error) {
	return instance.SetProperty("File", value)
}

// GetFile gets the value of File for the instance
func (instance *Win32_Patch) GetPropertyFile() (value string, err error) {
	retValue, err := instance.GetProperty("File")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPatchSize sets the value of PatchSize for the instance
func (instance *Win32_Patch) SetPropertyPatchSize(value uint32) (err error) {
	return instance.SetProperty("PatchSize", value)
}

// GetPatchSize gets the value of PatchSize for the instance
func (instance *Win32_Patch) GetPropertyPatchSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PatchSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProductCode sets the value of ProductCode for the instance
func (instance *Win32_Patch) SetPropertyProductCode(value string) (err error) {
	return instance.SetProperty("ProductCode", value)
}

// GetProductCode gets the value of ProductCode for the instance
func (instance *Win32_Patch) GetPropertyProductCode() (value string, err error) {
	retValue, err := instance.GetProperty("ProductCode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSequence sets the value of Sequence for the instance
func (instance *Win32_Patch) SetPropertySequence(value int16) (err error) {
	return instance.SetProperty("Sequence", value)
}

// GetSequence gets the value of Sequence for the instance
func (instance *Win32_Patch) GetPropertySequence() (value int16, err error) {
	retValue, err := instance.GetProperty("Sequence")
	if err != nil {
		return
	}
	value, ok := retValue.(int16)
	if !ok {
		// TODO: Set an error
	}
	return
}
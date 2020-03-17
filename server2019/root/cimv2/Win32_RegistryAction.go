// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_RegistryAction struct
type Win32_RegistryAction struct {
	CIM_Action

	//
	EntryName string

	//
	EntryValue string

	//
	key string

	//
	Registry string

	//
	Root int16
}

// SetEntryName sets the value of EntryName for the instance
func (instance *Win32_RegistryAction) SetPropertyEntryName(value string) (err error) {
	return instance.SetProperty("EntryName", value)
}

// GetEntryName gets the value of EntryName for the instance
func (instance *Win32_RegistryAction) GetPropertyEntryName() (value string, err error) {
	retValue, err := instance.GetProperty("EntryName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEntryValue sets the value of EntryValue for the instance
func (instance *Win32_RegistryAction) SetPropertyEntryValue(value string) (err error) {
	return instance.SetProperty("EntryValue", value)
}

// GetEntryValue gets the value of EntryValue for the instance
func (instance *Win32_RegistryAction) GetPropertyEntryValue() (value string, err error) {
	retValue, err := instance.GetProperty("EntryValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setkey sets the value of key for the instance
func (instance *Win32_RegistryAction) SetPropertykey(value string) (err error) {
	return instance.SetProperty("key", value)
}

// Getkey gets the value of key for the instance
func (instance *Win32_RegistryAction) GetPropertykey() (value string, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistry sets the value of Registry for the instance
func (instance *Win32_RegistryAction) SetPropertyRegistry(value string) (err error) {
	return instance.SetProperty("Registry", value)
}

// GetRegistry gets the value of Registry for the instance
func (instance *Win32_RegistryAction) GetPropertyRegistry() (value string, err error) {
	retValue, err := instance.GetProperty("Registry")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoot sets the value of Root for the instance
func (instance *Win32_RegistryAction) SetPropertyRoot(value int16) (err error) {
	return instance.SetProperty("Root", value)
}

// GetRoot gets the value of Root for the instance
func (instance *Win32_RegistryAction) GetPropertyRoot() (value int16, err error) {
	retValue, err := instance.GetProperty("Root")
	if err != nil {
		return
	}
	value, ok := retValue.(int16)
	if !ok {
		// TODO: Set an error
	}
	return
}
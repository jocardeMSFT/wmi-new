// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Appv
//////////////////////////////////////////////
package appv

// WMI_extension struct
type WMI_extension struct {
	__Win32Provider

	//
	AssemblyName string

	//
	AssemblyPath string

	//
	CLRVersion string
}

// SetAssemblyName sets the value of AssemblyName for the instance
func (instance *WMI_extension) SetPropertyAssemblyName(value string) (err error) {
	return instance.SetProperty("AssemblyName", value)
}

// GetAssemblyName gets the value of AssemblyName for the instance
func (instance *WMI_extension) GetPropertyAssemblyName() (value string, err error) {
	retValue, err := instance.GetProperty("AssemblyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAssemblyPath sets the value of AssemblyPath for the instance
func (instance *WMI_extension) SetPropertyAssemblyPath(value string) (err error) {
	return instance.SetProperty("AssemblyPath", value)
}

// GetAssemblyPath gets the value of AssemblyPath for the instance
func (instance *WMI_extension) GetPropertyAssemblyPath() (value string, err error) {
	retValue, err := instance.GetProperty("AssemblyPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCLRVersion sets the value of CLRVersion for the instance
func (instance *WMI_extension) SetPropertyCLRVersion(value string) (err error) {
	return instance.SetProperty("CLRVersion", value)
}

// GetCLRVersion gets the value of CLRVersion for the instance
func (instance *WMI_extension) GetPropertyCLRVersion() (value string, err error) {
	retValue, err := instance.GetProperty("CLRVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
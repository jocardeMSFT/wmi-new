// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_ServerFeature struct
type Win32_ServerFeature struct {
	cim.WmiInstance

	//
	ID uint32

	//
	Name string

	//
	ParentID uint32
}

// SetID sets the value of ID for the instance
func (instance *Win32_ServerFeature) SetPropertyID(value uint32) (err error) {
	return instance.SetProperty("ID", value)
}

// GetID gets the value of ID for the instance
func (instance *Win32_ServerFeature) GetPropertyID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *Win32_ServerFeature) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *Win32_ServerFeature) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *Win32_ServerFeature) SetPropertyParentID(value uint32) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *Win32_ServerFeature) GetPropertyParentID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
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

// CIM_FRU struct
type CIM_FRU struct {
	cim.WmiInstance

	//
	Caption string

	//
	Description string

	//
	FRUNumber string

	//
	IdentifyingNumber string

	//
	Name string

	//
	RevisionLevel string

	//
	Vendor string
}

// SetCaption sets the value of Caption for the instance
func (instance *CIM_FRU) SetPropertyCaption(value string) (err error) {
	return instance.SetProperty("Caption", value)
}

// GetCaption gets the value of Caption for the instance
func (instance *CIM_FRU) GetPropertyCaption() (value string, err error) {
	retValue, err := instance.GetProperty("Caption")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *CIM_FRU) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *CIM_FRU) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFRUNumber sets the value of FRUNumber for the instance
func (instance *CIM_FRU) SetPropertyFRUNumber(value string) (err error) {
	return instance.SetProperty("FRUNumber", value)
}

// GetFRUNumber gets the value of FRUNumber for the instance
func (instance *CIM_FRU) GetPropertyFRUNumber() (value string, err error) {
	retValue, err := instance.GetProperty("FRUNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdentifyingNumber sets the value of IdentifyingNumber for the instance
func (instance *CIM_FRU) SetPropertyIdentifyingNumber(value string) (err error) {
	return instance.SetProperty("IdentifyingNumber", value)
}

// GetIdentifyingNumber gets the value of IdentifyingNumber for the instance
func (instance *CIM_FRU) GetPropertyIdentifyingNumber() (value string, err error) {
	retValue, err := instance.GetProperty("IdentifyingNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *CIM_FRU) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *CIM_FRU) GetPropertyName() (value string, err error) {
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

// SetRevisionLevel sets the value of RevisionLevel for the instance
func (instance *CIM_FRU) SetPropertyRevisionLevel(value string) (err error) {
	return instance.SetProperty("RevisionLevel", value)
}

// GetRevisionLevel gets the value of RevisionLevel for the instance
func (instance *CIM_FRU) GetPropertyRevisionLevel() (value string, err error) {
	retValue, err := instance.GetProperty("RevisionLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVendor sets the value of Vendor for the instance
func (instance *CIM_FRU) SetPropertyVendor(value string) (err error) {
	return instance.SetProperty("Vendor", value)
}

// GetVendor gets the value of Vendor for the instance
func (instance *CIM_FRU) GetPropertyVendor() (value string, err error) {
	retValue, err := instance.GetProperty("Vendor")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
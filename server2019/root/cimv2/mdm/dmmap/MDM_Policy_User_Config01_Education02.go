// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_User_Config01_Education02 struct
type MDM_Policy_User_Config01_Education02 struct {
	cim.WmiInstance

	//
	AllowGraphingCalculator int32

	//
	DefaultPrinterName string

	//
	InstanceID string

	//
	ParentID string

	//
	PreventAddingNewPrinters int32

	//
	PrinterNames string
}

// SetAllowGraphingCalculator sets the value of AllowGraphingCalculator for the instance
func (instance *MDM_Policy_User_Config01_Education02) SetPropertyAllowGraphingCalculator(value int32) (err error) {
	return instance.SetProperty("AllowGraphingCalculator", value)
}

// GetAllowGraphingCalculator gets the value of AllowGraphingCalculator for the instance
func (instance *MDM_Policy_User_Config01_Education02) GetPropertyAllowGraphingCalculator() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowGraphingCalculator")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultPrinterName sets the value of DefaultPrinterName for the instance
func (instance *MDM_Policy_User_Config01_Education02) SetPropertyDefaultPrinterName(value string) (err error) {
	return instance.SetProperty("DefaultPrinterName", value)
}

// GetDefaultPrinterName gets the value of DefaultPrinterName for the instance
func (instance *MDM_Policy_User_Config01_Education02) GetPropertyDefaultPrinterName() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultPrinterName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_Education02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_Education02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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
func (instance *MDM_Policy_User_Config01_Education02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Config01_Education02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreventAddingNewPrinters sets the value of PreventAddingNewPrinters for the instance
func (instance *MDM_Policy_User_Config01_Education02) SetPropertyPreventAddingNewPrinters(value int32) (err error) {
	return instance.SetProperty("PreventAddingNewPrinters", value)
}

// GetPreventAddingNewPrinters gets the value of PreventAddingNewPrinters for the instance
func (instance *MDM_Policy_User_Config01_Education02) GetPropertyPreventAddingNewPrinters() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventAddingNewPrinters")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrinterNames sets the value of PrinterNames for the instance
func (instance *MDM_Policy_User_Config01_Education02) SetPropertyPrinterNames(value string) (err error) {
	return instance.SetProperty("PrinterNames", value)
}

// GetPrinterNames gets the value of PrinterNames for the instance
func (instance *MDM_Policy_User_Config01_Education02) GetPropertyPrinterNames() (value string, err error) {
	retValue, err := instance.GetProperty("PrinterNames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
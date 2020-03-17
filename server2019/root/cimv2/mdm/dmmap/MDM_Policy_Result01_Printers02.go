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

// MDM_Policy_Result01_Printers02 struct
type MDM_Policy_Result01_Printers02 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	PointAndPrintRestrictions string

	//
	PublishPrinters string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Printers02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Printers02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_Printers02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Printers02) GetPropertyParentID() (value string, err error) {
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

// SetPointAndPrintRestrictions sets the value of PointAndPrintRestrictions for the instance
func (instance *MDM_Policy_Result01_Printers02) SetPropertyPointAndPrintRestrictions(value string) (err error) {
	return instance.SetProperty("PointAndPrintRestrictions", value)
}

// GetPointAndPrintRestrictions gets the value of PointAndPrintRestrictions for the instance
func (instance *MDM_Policy_Result01_Printers02) GetPropertyPointAndPrintRestrictions() (value string, err error) {
	retValue, err := instance.GetProperty("PointAndPrintRestrictions")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPublishPrinters sets the value of PublishPrinters for the instance
func (instance *MDM_Policy_Result01_Printers02) SetPropertyPublishPrinters(value string) (err error) {
	return instance.SetProperty("PublishPrinters", value)
}

// GetPublishPrinters gets the value of PublishPrinters for the instance
func (instance *MDM_Policy_Result01_Printers02) GetPropertyPublishPrinters() (value string, err error) {
	retValue, err := instance.GetProperty("PublishPrinters")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
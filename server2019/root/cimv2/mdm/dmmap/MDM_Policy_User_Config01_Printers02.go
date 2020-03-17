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

// MDM_Policy_User_Config01_Printers02 struct
type MDM_Policy_User_Config01_Printers02 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	ParentID string

	//
	PointAndPrintRestrictions_User string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_Printers02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Config01_Printers02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_User_Config01_Printers02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Config01_Printers02) GetPropertyParentID() (value string, err error) {
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

// SetPointAndPrintRestrictions_User sets the value of PointAndPrintRestrictions_User for the instance
func (instance *MDM_Policy_User_Config01_Printers02) SetPropertyPointAndPrintRestrictions_User(value string) (err error) {
	return instance.SetProperty("PointAndPrintRestrictions_User", value)
}

// GetPointAndPrintRestrictions_User gets the value of PointAndPrintRestrictions_User for the instance
func (instance *MDM_Policy_User_Config01_Printers02) GetPropertyPointAndPrintRestrictions_User() (value string, err error) {
	retValue, err := instance.GetProperty("PointAndPrintRestrictions_User")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
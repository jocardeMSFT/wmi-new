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

// MDM_Policy_Result01_Handwriting02 struct
type MDM_Policy_Result01_Handwriting02 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	PanelDefaultModeDocked int32

	//
	ParentID string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Handwriting02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Handwriting02) GetPropertyInstanceID() (value string, err error) {
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

// SetPanelDefaultModeDocked sets the value of PanelDefaultModeDocked for the instance
func (instance *MDM_Policy_Result01_Handwriting02) SetPropertyPanelDefaultModeDocked(value int32) (err error) {
	return instance.SetProperty("PanelDefaultModeDocked", value)
}

// GetPanelDefaultModeDocked gets the value of PanelDefaultModeDocked for the instance
func (instance *MDM_Policy_Result01_Handwriting02) GetPropertyPanelDefaultModeDocked() (value int32, err error) {
	retValue, err := instance.GetProperty("PanelDefaultModeDocked")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Handwriting02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Handwriting02) GetPropertyParentID() (value string, err error) {
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
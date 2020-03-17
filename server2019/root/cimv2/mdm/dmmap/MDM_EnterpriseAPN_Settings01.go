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

// MDM_EnterpriseAPN_Settings01 struct
type MDM_EnterpriseAPN_Settings01 struct {
	cim.WmiInstance

	//
	AllowUserControl bool

	//
	HideView bool

	//
	InstanceID string

	//
	ParentID string
}

// SetAllowUserControl sets the value of AllowUserControl for the instance
func (instance *MDM_EnterpriseAPN_Settings01) SetPropertyAllowUserControl(value bool) (err error) {
	return instance.SetProperty("AllowUserControl", value)
}

// GetAllowUserControl gets the value of AllowUserControl for the instance
func (instance *MDM_EnterpriseAPN_Settings01) GetPropertyAllowUserControl() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowUserControl")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHideView sets the value of HideView for the instance
func (instance *MDM_EnterpriseAPN_Settings01) SetPropertyHideView(value bool) (err error) {
	return instance.SetProperty("HideView", value)
}

// GetHideView gets the value of HideView for the instance
func (instance *MDM_EnterpriseAPN_Settings01) GetPropertyHideView() (value bool, err error) {
	retValue, err := instance.GetProperty("HideView")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_EnterpriseAPN_Settings01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseAPN_Settings01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_EnterpriseAPN_Settings01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseAPN_Settings01) GetPropertyParentID() (value string, err error) {
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
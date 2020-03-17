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

// MDM_Policy_Result01_WindowsInkWorkspace02 struct
type MDM_Policy_Result01_WindowsInkWorkspace02 struct {
	cim.WmiInstance

	//
	AllowSuggestedAppsInWindowsInkWorkspace int32

	//
	AllowWindowsInkWorkspace int32

	//
	InstanceID string

	//
	ParentID string
}

// SetAllowSuggestedAppsInWindowsInkWorkspace sets the value of AllowSuggestedAppsInWindowsInkWorkspace for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) SetPropertyAllowSuggestedAppsInWindowsInkWorkspace(value int32) (err error) {
	return instance.SetProperty("AllowSuggestedAppsInWindowsInkWorkspace", value)
}

// GetAllowSuggestedAppsInWindowsInkWorkspace gets the value of AllowSuggestedAppsInWindowsInkWorkspace for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) GetPropertyAllowSuggestedAppsInWindowsInkWorkspace() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSuggestedAppsInWindowsInkWorkspace")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowWindowsInkWorkspace sets the value of AllowWindowsInkWorkspace for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) SetPropertyAllowWindowsInkWorkspace(value int32) (err error) {
	return instance.SetProperty("AllowWindowsInkWorkspace", value)
}

// GetAllowWindowsInkWorkspace gets the value of AllowWindowsInkWorkspace for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) GetPropertyAllowWindowsInkWorkspace() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsInkWorkspace")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_WindowsInkWorkspace02) GetPropertyParentID() (value string, err error) {
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
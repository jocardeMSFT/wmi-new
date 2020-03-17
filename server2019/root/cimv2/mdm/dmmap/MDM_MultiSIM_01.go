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

// MDM_MultiSIM_01 struct
type MDM_MultiSIM_01 struct {
	cim.WmiInstance

	//
	Identifier string

	//
	InstanceID string

	//
	IsEmbedded bool

	//
	ParentID string
}

// SetIdentifier sets the value of Identifier for the instance
func (instance *MDM_MultiSIM_01) SetPropertyIdentifier(value string) (err error) {
	return instance.SetProperty("Identifier", value)
}

// GetIdentifier gets the value of Identifier for the instance
func (instance *MDM_MultiSIM_01) GetPropertyIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("Identifier")
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
func (instance *MDM_MultiSIM_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_MultiSIM_01) GetPropertyInstanceID() (value string, err error) {
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

// SetIsEmbedded sets the value of IsEmbedded for the instance
func (instance *MDM_MultiSIM_01) SetPropertyIsEmbedded(value bool) (err error) {
	return instance.SetProperty("IsEmbedded", value)
}

// GetIsEmbedded gets the value of IsEmbedded for the instance
func (instance *MDM_MultiSIM_01) GetPropertyIsEmbedded() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEmbedded")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_MultiSIM_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_MultiSIM_01) GetPropertyParentID() (value string, err error) {
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
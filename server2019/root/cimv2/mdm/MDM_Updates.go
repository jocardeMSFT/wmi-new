// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Updates struct
type MDM_Updates struct {
	cim.WmiInstance

	//
	AutoUpdatePolicy uint32

	//
	key uint32
}

// SetAutoUpdatePolicy sets the value of AutoUpdatePolicy for the instance
func (instance *MDM_Updates) SetPropertyAutoUpdatePolicy(value uint32) (err error) {
	return instance.SetProperty("AutoUpdatePolicy", value)
}

// GetAutoUpdatePolicy gets the value of AutoUpdatePolicy for the instance
func (instance *MDM_Updates) GetPropertyAutoUpdatePolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoUpdatePolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setkey sets the value of key for the instance
func (instance *MDM_Updates) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", value)
}

// Getkey gets the value of key for the instance
func (instance *MDM_Updates) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
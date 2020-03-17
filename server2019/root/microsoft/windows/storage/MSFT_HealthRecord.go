// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_HealthRecord struct
type MSFT_HealthRecord struct {
	cim.WmiInstance

	//
	Name string

	//
	Units uint16
}

// SetName sets the value of Name for the instance
func (instance *MSFT_HealthRecord) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_HealthRecord) GetPropertyName() (value string, err error) {
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

// SetUnits sets the value of Units for the instance
func (instance *MSFT_HealthRecord) SetPropertyUnits(value uint16) (err error) {
	return instance.SetProperty("Units", value)
}

// GetUnits gets the value of Units for the instance
func (instance *MSFT_HealthRecord) GetPropertyUnits() (value uint16, err error) {
	retValue, err := instance.GetProperty("Units")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
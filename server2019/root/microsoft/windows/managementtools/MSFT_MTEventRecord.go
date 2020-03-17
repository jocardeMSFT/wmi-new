// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MTEventRecord struct
type MSFT_MTEventRecord struct {
	cim.WmiInstance

	//
	Description string

	//
	Id uint32

	//
	Level uint16

	//
	Log string

	//
	RecordId uint32

	//
	Source string

	//
	Timestamp string
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_MTEventRecord) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_MTEventRecord) GetPropertyDescription() (value string, err error) {
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

// SetId sets the value of Id for the instance
func (instance *MSFT_MTEventRecord) SetPropertyId(value uint32) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_MTEventRecord) GetPropertyId() (value uint32, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLevel sets the value of Level for the instance
func (instance *MSFT_MTEventRecord) SetPropertyLevel(value uint16) (err error) {
	return instance.SetProperty("Level", value)
}

// GetLevel gets the value of Level for the instance
func (instance *MSFT_MTEventRecord) GetPropertyLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("Level")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLog sets the value of Log for the instance
func (instance *MSFT_MTEventRecord) SetPropertyLog(value string) (err error) {
	return instance.SetProperty("Log", value)
}

// GetLog gets the value of Log for the instance
func (instance *MSFT_MTEventRecord) GetPropertyLog() (value string, err error) {
	retValue, err := instance.GetProperty("Log")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecordId sets the value of RecordId for the instance
func (instance *MSFT_MTEventRecord) SetPropertyRecordId(value uint32) (err error) {
	return instance.SetProperty("RecordId", value)
}

// GetRecordId gets the value of RecordId for the instance
func (instance *MSFT_MTEventRecord) GetPropertyRecordId() (value uint32, err error) {
	retValue, err := instance.GetProperty("RecordId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSource sets the value of Source for the instance
func (instance *MSFT_MTEventRecord) SetPropertySource(value string) (err error) {
	return instance.SetProperty("Source", value)
}

// GetSource gets the value of Source for the instance
func (instance *MSFT_MTEventRecord) GetPropertySource() (value string, err error) {
	retValue, err := instance.GetProperty("Source")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimestamp sets the value of Timestamp for the instance
func (instance *MSFT_MTEventRecord) SetPropertyTimestamp(value string) (err error) {
	return instance.SetProperty("Timestamp", value)
}

// GetTimestamp gets the value of Timestamp for the instance
func (instance *MSFT_MTEventRecord) GetPropertyTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("Timestamp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
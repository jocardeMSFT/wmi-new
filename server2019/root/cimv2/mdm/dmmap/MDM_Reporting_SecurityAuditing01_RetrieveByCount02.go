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

// MDM_Reporting_SecurityAuditing01_RetrieveByCount02 struct
type MDM_Reporting_SecurityAuditing01_RetrieveByCount02 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	LogCount int32

	//
	Logs string

	//
	ParentID string

	//
	StartTime string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Reporting_SecurityAuditing01_RetrieveByCount02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Reporting_SecurityAuditing01_RetrieveByCount02) GetPropertyInstanceID() (value string, err error) {
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

// SetLogCount sets the value of LogCount for the instance
func (instance *MDM_Reporting_SecurityAuditing01_RetrieveByCount02) SetPropertyLogCount(value int32) (err error) {
	return instance.SetProperty("LogCount", value)
}

// GetLogCount gets the value of LogCount for the instance
func (instance *MDM_Reporting_SecurityAuditing01_RetrieveByCount02) GetPropertyLogCount() (value int32, err error) {
	retValue, err := instance.GetProperty("LogCount")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogs sets the value of Logs for the instance
func (instance *MDM_Reporting_SecurityAuditing01_RetrieveByCount02) SetPropertyLogs(value string) (err error) {
	return instance.SetProperty("Logs", value)
}

// GetLogs gets the value of Logs for the instance
func (instance *MDM_Reporting_SecurityAuditing01_RetrieveByCount02) GetPropertyLogs() (value string, err error) {
	retValue, err := instance.GetProperty("Logs")
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
func (instance *MDM_Reporting_SecurityAuditing01_RetrieveByCount02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Reporting_SecurityAuditing01_RetrieveByCount02) GetPropertyParentID() (value string, err error) {
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

// SetStartTime sets the value of StartTime for the instance
func (instance *MDM_Reporting_SecurityAuditing01_RetrieveByCount02) SetPropertyStartTime(value string) (err error) {
	return instance.SetProperty("StartTime", value)
}

// GetStartTime gets the value of StartTime for the instance
func (instance *MDM_Reporting_SecurityAuditing01_RetrieveByCount02) GetPropertyStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
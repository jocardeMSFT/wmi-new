// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// MSFT_NetServiceStartFailed struct
type MSFT_NetServiceStartFailed struct {
	MSFT_SCMEventLogEvent

	//
	Error uint32

	//
	Service string
}

// SetError sets the value of Error for the instance
func (instance *MSFT_NetServiceStartFailed) SetPropertyError(value uint32) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *MSFT_NetServiceStartFailed) GetPropertyError() (value uint32, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetService sets the value of Service for the instance
func (instance *MSFT_NetServiceStartFailed) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *MSFT_NetServiceStartFailed) GetPropertyService() (value string, err error) {
	retValue, err := instance.GetProperty("Service")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
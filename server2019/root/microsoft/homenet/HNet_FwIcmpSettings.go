// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// HNet_FwIcmpSettings struct
type HNet_FwIcmpSettings struct {
	cim.WmiInstance

	//
	AllowInboundEchoRequest bool

	//
	AllowInboundMaskRequest bool

	//
	AllowInboundRouterRequest bool

	//
	AllowInboundTimestampRequest bool

	//
	AllowOutboundDestinationUnreachable bool

	//
	AllowOutboundParameterProblem bool

	//
	AllowOutboundSourceQuench bool

	//
	AllowOutboundTimeExceeded bool

	//
	AllowRedirect bool

	//
	Name string
}

// SetAllowInboundEchoRequest sets the value of AllowInboundEchoRequest for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowInboundEchoRequest(value bool) (err error) {
	return instance.SetProperty("AllowInboundEchoRequest", value)
}

// GetAllowInboundEchoRequest gets the value of AllowInboundEchoRequest for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowInboundEchoRequest() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowInboundEchoRequest")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowInboundMaskRequest sets the value of AllowInboundMaskRequest for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowInboundMaskRequest(value bool) (err error) {
	return instance.SetProperty("AllowInboundMaskRequest", value)
}

// GetAllowInboundMaskRequest gets the value of AllowInboundMaskRequest for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowInboundMaskRequest() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowInboundMaskRequest")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowInboundRouterRequest sets the value of AllowInboundRouterRequest for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowInboundRouterRequest(value bool) (err error) {
	return instance.SetProperty("AllowInboundRouterRequest", value)
}

// GetAllowInboundRouterRequest gets the value of AllowInboundRouterRequest for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowInboundRouterRequest() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowInboundRouterRequest")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowInboundTimestampRequest sets the value of AllowInboundTimestampRequest for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowInboundTimestampRequest(value bool) (err error) {
	return instance.SetProperty("AllowInboundTimestampRequest", value)
}

// GetAllowInboundTimestampRequest gets the value of AllowInboundTimestampRequest for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowInboundTimestampRequest() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowInboundTimestampRequest")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowOutboundDestinationUnreachable sets the value of AllowOutboundDestinationUnreachable for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowOutboundDestinationUnreachable(value bool) (err error) {
	return instance.SetProperty("AllowOutboundDestinationUnreachable", value)
}

// GetAllowOutboundDestinationUnreachable gets the value of AllowOutboundDestinationUnreachable for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowOutboundDestinationUnreachable() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowOutboundDestinationUnreachable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowOutboundParameterProblem sets the value of AllowOutboundParameterProblem for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowOutboundParameterProblem(value bool) (err error) {
	return instance.SetProperty("AllowOutboundParameterProblem", value)
}

// GetAllowOutboundParameterProblem gets the value of AllowOutboundParameterProblem for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowOutboundParameterProblem() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowOutboundParameterProblem")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowOutboundSourceQuench sets the value of AllowOutboundSourceQuench for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowOutboundSourceQuench(value bool) (err error) {
	return instance.SetProperty("AllowOutboundSourceQuench", value)
}

// GetAllowOutboundSourceQuench gets the value of AllowOutboundSourceQuench for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowOutboundSourceQuench() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowOutboundSourceQuench")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowOutboundTimeExceeded sets the value of AllowOutboundTimeExceeded for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowOutboundTimeExceeded(value bool) (err error) {
	return instance.SetProperty("AllowOutboundTimeExceeded", value)
}

// GetAllowOutboundTimeExceeded gets the value of AllowOutboundTimeExceeded for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowOutboundTimeExceeded() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowOutboundTimeExceeded")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowRedirect sets the value of AllowRedirect for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowRedirect(value bool) (err error) {
	return instance.SetProperty("AllowRedirect", value)
}

// GetAllowRedirect gets the value of AllowRedirect for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowRedirect() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowRedirect")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyName() (value string, err error) {
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
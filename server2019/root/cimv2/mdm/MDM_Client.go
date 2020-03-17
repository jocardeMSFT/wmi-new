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

// MDM_Client struct
type MDM_Client struct {
	cim.WmiInstance

	//
	DeviceClientID string

	//
	DeviceName string

	//
	DomainSID string

	//
	PlatformID string

	//
	ProcessorDescription string

	//
	UserSid string

	//
	Version string
}

// SetDeviceClientID sets the value of DeviceClientID for the instance
func (instance *MDM_Client) SetPropertyDeviceClientID(value string) (err error) {
	return instance.SetProperty("DeviceClientID", value)
}

// GetDeviceClientID gets the value of DeviceClientID for the instance
func (instance *MDM_Client) GetPropertyDeviceClientID() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceClientID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceName sets the value of DeviceName for the instance
func (instance *MDM_Client) SetPropertyDeviceName(value string) (err error) {
	return instance.SetProperty("DeviceName", value)
}

// GetDeviceName gets the value of DeviceName for the instance
func (instance *MDM_Client) GetPropertyDeviceName() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomainSID sets the value of DomainSID for the instance
func (instance *MDM_Client) SetPropertyDomainSID(value string) (err error) {
	return instance.SetProperty("DomainSID", value)
}

// GetDomainSID gets the value of DomainSID for the instance
func (instance *MDM_Client) GetPropertyDomainSID() (value string, err error) {
	retValue, err := instance.GetProperty("DomainSID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlatformID sets the value of PlatformID for the instance
func (instance *MDM_Client) SetPropertyPlatformID(value string) (err error) {
	return instance.SetProperty("PlatformID", value)
}

// GetPlatformID gets the value of PlatformID for the instance
func (instance *MDM_Client) GetPropertyPlatformID() (value string, err error) {
	retValue, err := instance.GetProperty("PlatformID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorDescription sets the value of ProcessorDescription for the instance
func (instance *MDM_Client) SetPropertyProcessorDescription(value string) (err error) {
	return instance.SetProperty("ProcessorDescription", value)
}

// GetProcessorDescription gets the value of ProcessorDescription for the instance
func (instance *MDM_Client) GetPropertyProcessorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ProcessorDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserSid sets the value of UserSid for the instance
func (instance *MDM_Client) SetPropertyUserSid(value string) (err error) {
	return instance.SetProperty("UserSid", value)
}

// GetUserSid gets the value of UserSid for the instance
func (instance *MDM_Client) GetPropertyUserSid() (value string, err error) {
	retValue, err := instance.GetProperty("UserSid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *MDM_Client) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *MDM_Client) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="DeviceClientID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_Client) SendUnenrollRequest( /* IN */ DeviceClientID string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SendUnenrollRequest", DeviceClientID)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_Client) LockWorkstation() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("LockWorkstation")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ConfigString" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_Client) ResetUserPassword( /* IN */ ConfigString string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ResetUserPassword", ConfigString)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
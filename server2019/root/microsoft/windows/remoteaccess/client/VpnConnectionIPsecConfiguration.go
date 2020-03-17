// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// VpnConnectionIPsecConfiguration struct
type VpnConnectionIPsecConfiguration struct {
	cim.WmiInstance

	//
	AuthenticationTransformConstants uint32

	//
	CipherTransformConstants uint32

	//
	DHGroup uint32

	//
	EncryptionMethod uint32

	//
	IntegrityCheckMethod uint32

	//
	PfsGroup uint32
}

// SetAuthenticationTransformConstants sets the value of AuthenticationTransformConstants for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyAuthenticationTransformConstants(value uint32) (err error) {
	return instance.SetProperty("AuthenticationTransformConstants", value)
}

// GetAuthenticationTransformConstants gets the value of AuthenticationTransformConstants for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyAuthenticationTransformConstants() (value uint32, err error) {
	retValue, err := instance.GetProperty("AuthenticationTransformConstants")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCipherTransformConstants sets the value of CipherTransformConstants for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyCipherTransformConstants(value uint32) (err error) {
	return instance.SetProperty("CipherTransformConstants", value)
}

// GetCipherTransformConstants gets the value of CipherTransformConstants for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyCipherTransformConstants() (value uint32, err error) {
	retValue, err := instance.GetProperty("CipherTransformConstants")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDHGroup sets the value of DHGroup for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyDHGroup(value uint32) (err error) {
	return instance.SetProperty("DHGroup", value)
}

// GetDHGroup gets the value of DHGroup for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyDHGroup() (value uint32, err error) {
	retValue, err := instance.GetProperty("DHGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptionMethod sets the value of EncryptionMethod for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyEncryptionMethod(value uint32) (err error) {
	return instance.SetProperty("EncryptionMethod", value)
}

// GetEncryptionMethod gets the value of EncryptionMethod for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyEncryptionMethod() (value uint32, err error) {
	retValue, err := instance.GetProperty("EncryptionMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntegrityCheckMethod sets the value of IntegrityCheckMethod for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyIntegrityCheckMethod(value uint32) (err error) {
	return instance.SetProperty("IntegrityCheckMethod", value)
}

// GetIntegrityCheckMethod gets the value of IntegrityCheckMethod for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyIntegrityCheckMethod() (value uint32, err error) {
	retValue, err := instance.GetProperty("IntegrityCheckMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPfsGroup sets the value of PfsGroup for the instance
func (instance *VpnConnectionIPsecConfiguration) SetPropertyPfsGroup(value uint32) (err error) {
	return instance.SetProperty("PfsGroup", value)
}

// GetPfsGroup gets the value of PfsGroup for the instance
func (instance *VpnConnectionIPsecConfiguration) GetPropertyPfsGroup() (value uint32, err error) {
	retValue, err := instance.GetProperty("PfsGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
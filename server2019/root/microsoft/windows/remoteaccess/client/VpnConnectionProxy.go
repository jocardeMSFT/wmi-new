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

// VpnConnectionProxy struct
type VpnConnectionProxy struct {
	cim.WmiInstance

	//
	AutoConfigurationScript string

	//
	AutoDetect bool

	//
	BypassProxyForLocal bool

	//
	ConnectionName string

	//
	ExceptionPrefix []string

	//
	ProxyServer string
}

// SetAutoConfigurationScript sets the value of AutoConfigurationScript for the instance
func (instance *VpnConnectionProxy) SetPropertyAutoConfigurationScript(value string) (err error) {
	return instance.SetProperty("AutoConfigurationScript", value)
}

// GetAutoConfigurationScript gets the value of AutoConfigurationScript for the instance
func (instance *VpnConnectionProxy) GetPropertyAutoConfigurationScript() (value string, err error) {
	retValue, err := instance.GetProperty("AutoConfigurationScript")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoDetect sets the value of AutoDetect for the instance
func (instance *VpnConnectionProxy) SetPropertyAutoDetect(value bool) (err error) {
	return instance.SetProperty("AutoDetect", value)
}

// GetAutoDetect gets the value of AutoDetect for the instance
func (instance *VpnConnectionProxy) GetPropertyAutoDetect() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoDetect")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBypassProxyForLocal sets the value of BypassProxyForLocal for the instance
func (instance *VpnConnectionProxy) SetPropertyBypassProxyForLocal(value bool) (err error) {
	return instance.SetProperty("BypassProxyForLocal", value)
}

// GetBypassProxyForLocal gets the value of BypassProxyForLocal for the instance
func (instance *VpnConnectionProxy) GetPropertyBypassProxyForLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("BypassProxyForLocal")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionName sets the value of ConnectionName for the instance
func (instance *VpnConnectionProxy) SetPropertyConnectionName(value string) (err error) {
	return instance.SetProperty("ConnectionName", value)
}

// GetConnectionName gets the value of ConnectionName for the instance
func (instance *VpnConnectionProxy) GetPropertyConnectionName() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExceptionPrefix sets the value of ExceptionPrefix for the instance
func (instance *VpnConnectionProxy) SetPropertyExceptionPrefix(value []string) (err error) {
	return instance.SetProperty("ExceptionPrefix", value)
}

// GetExceptionPrefix gets the value of ExceptionPrefix for the instance
func (instance *VpnConnectionProxy) GetPropertyExceptionPrefix() (value []string, err error) {
	retValue, err := instance.GetProperty("ExceptionPrefix")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProxyServer sets the value of ProxyServer for the instance
func (instance *VpnConnectionProxy) SetPropertyProxyServer(value string) (err error) {
	return instance.SetProperty("ProxyServer", value)
}

// GetProxyServer gets the value of ProxyServer for the instance
func (instance *VpnConnectionProxy) GetPropertyProxyServer() (value string, err error) {
	retValue, err := instance.GetProperty("ProxyServer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
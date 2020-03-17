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

// PS_VpnConnectionTriggerTrustedNetwork struct
type PS_VpnConnectionTriggerTrustedNetwork struct {
	cim.WmiInstance
}

//

// <param name="ConnectionName" type="string "></param>
// <param name="DnsSuffix" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnConnectionTriggerTrustedNetwork "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionTriggerTrustedNetwork) Add( /* IN */ ConnectionName string,
	/* IN */ DnsSuffix []string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput VpnConnectionTriggerTrustedNetwork) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", ConnectionName, DnsSuffix, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ConnectionName" type="string "></param>
// <param name="DnsSuffix" type="string []"></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnConnectionTriggerTrustedNetwork "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionTriggerTrustedNetwork) Remove( /* IN */ ConnectionName string,
	/* IN */ DnsSuffix []string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput VpnConnectionTriggerTrustedNetwork) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", ConnectionName, DnsSuffix, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ConnectionName" type="string "></param>
// <param name="DefaultDnsSuffixes" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnConnectionTriggerTrustedNetwork "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionTriggerTrustedNetwork) Set( /* IN */ ConnectionName string,
	/* IN */ DefaultDnsSuffixes bool,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput VpnConnectionTriggerTrustedNetwork) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", ConnectionName, DefaultDnsSuffixes, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
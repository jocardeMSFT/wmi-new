// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetAdapterPacketDirectSettingData struct
type MSFT_NetAdapterPacketDirectSettingData struct {
	MSFT_NetAdapterSettingData

	//
	Capabilities MSFT_NetAdapter_PacketDirectCapabilities

	//
	DiagnosticCode uint32

	//
	DmaAddressWidth uint8

	//
	DomainId uint32

	//
	Enabled bool

	//
	Operational bool
}

// SetCapabilities sets the value of Capabilities for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyCapabilities(value MSFT_NetAdapter_PacketDirectCapabilities) (err error) {
	return instance.SetProperty("Capabilities", value)
}

// GetCapabilities gets the value of Capabilities for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyCapabilities() (value MSFT_NetAdapter_PacketDirectCapabilities, err error) {
	retValue, err := instance.GetProperty("Capabilities")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapter_PacketDirectCapabilities)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDiagnosticCode sets the value of DiagnosticCode for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyDiagnosticCode(value uint32) (err error) {
	return instance.SetProperty("DiagnosticCode", value)
}

// GetDiagnosticCode gets the value of DiagnosticCode for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyDiagnosticCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("DiagnosticCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDmaAddressWidth sets the value of DmaAddressWidth for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyDmaAddressWidth(value uint8) (err error) {
	return instance.SetProperty("DmaAddressWidth", value)
}

// GetDmaAddressWidth gets the value of DmaAddressWidth for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyDmaAddressWidth() (value uint8, err error) {
	retValue, err := instance.GetProperty("DmaAddressWidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomainId sets the value of DomainId for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyDomainId(value uint32) (err error) {
	return instance.SetProperty("DomainId", value)
}

// GetDomainId gets the value of DomainId for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyDomainId() (value uint32, err error) {
	retValue, err := instance.GetProperty("DomainId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperational sets the value of Operational for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyOperational(value bool) (err error) {
	return instance.SetProperty("Operational", value)
}

// GetOperational gets the value of Operational for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyOperational() (value bool, err error) {
	retValue, err := instance.GetProperty("Operational")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="CmdletOutput" type="MSFT_NetAdapterPacketDirectSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterPacketDirectSettingData) Enable( /* OUT */ CmdletOutput MSFT_NetAdapterPacketDirectSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CmdletOutput" type="MSFT_NetAdapterPacketDirectSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterPacketDirectSettingData) Disable( /* OUT */ CmdletOutput MSFT_NetAdapterPacketDirectSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
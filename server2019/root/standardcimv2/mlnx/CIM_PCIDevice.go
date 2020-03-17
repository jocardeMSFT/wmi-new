// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_PCIDevice struct
type CIM_PCIDevice struct {
	CIM_PCIController

	//
	BaseAddress []uint32

	//
	BusNumber uint8

	//
	DeviceNumber uint8

	//
	FunctionNumber uint8

	//
	MaxLatency uint8

	//
	MinGrantTime uint8

	//
	PCIDeviceID uint16

	//
	RevisionID uint8

	//
	SubsystemID uint16

	//
	SubsystemVendorID uint16

	//
	VendorID uint16
}

// SetBaseAddress sets the value of BaseAddress for the instance
func (instance *CIM_PCIDevice) SetPropertyBaseAddress(value []uint32) (err error) {
	return instance.SetProperty("BaseAddress", value)
}

// GetBaseAddress gets the value of BaseAddress for the instance
func (instance *CIM_PCIDevice) GetPropertyBaseAddress() (value []uint32, err error) {
	retValue, err := instance.GetProperty("BaseAddress")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBusNumber sets the value of BusNumber for the instance
func (instance *CIM_PCIDevice) SetPropertyBusNumber(value uint8) (err error) {
	return instance.SetProperty("BusNumber", value)
}

// GetBusNumber gets the value of BusNumber for the instance
func (instance *CIM_PCIDevice) GetPropertyBusNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("BusNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceNumber sets the value of DeviceNumber for the instance
func (instance *CIM_PCIDevice) SetPropertyDeviceNumber(value uint8) (err error) {
	return instance.SetProperty("DeviceNumber", value)
}

// GetDeviceNumber gets the value of DeviceNumber for the instance
func (instance *CIM_PCIDevice) GetPropertyDeviceNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("DeviceNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFunctionNumber sets the value of FunctionNumber for the instance
func (instance *CIM_PCIDevice) SetPropertyFunctionNumber(value uint8) (err error) {
	return instance.SetProperty("FunctionNumber", value)
}

// GetFunctionNumber gets the value of FunctionNumber for the instance
func (instance *CIM_PCIDevice) GetPropertyFunctionNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("FunctionNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxLatency sets the value of MaxLatency for the instance
func (instance *CIM_PCIDevice) SetPropertyMaxLatency(value uint8) (err error) {
	return instance.SetProperty("MaxLatency", value)
}

// GetMaxLatency gets the value of MaxLatency for the instance
func (instance *CIM_PCIDevice) GetPropertyMaxLatency() (value uint8, err error) {
	retValue, err := instance.GetProperty("MaxLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinGrantTime sets the value of MinGrantTime for the instance
func (instance *CIM_PCIDevice) SetPropertyMinGrantTime(value uint8) (err error) {
	return instance.SetProperty("MinGrantTime", value)
}

// GetMinGrantTime gets the value of MinGrantTime for the instance
func (instance *CIM_PCIDevice) GetPropertyMinGrantTime() (value uint8, err error) {
	retValue, err := instance.GetProperty("MinGrantTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPCIDeviceID sets the value of PCIDeviceID for the instance
func (instance *CIM_PCIDevice) SetPropertyPCIDeviceID(value uint16) (err error) {
	return instance.SetProperty("PCIDeviceID", value)
}

// GetPCIDeviceID gets the value of PCIDeviceID for the instance
func (instance *CIM_PCIDevice) GetPropertyPCIDeviceID() (value uint16, err error) {
	retValue, err := instance.GetProperty("PCIDeviceID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRevisionID sets the value of RevisionID for the instance
func (instance *CIM_PCIDevice) SetPropertyRevisionID(value uint8) (err error) {
	return instance.SetProperty("RevisionID", value)
}

// GetRevisionID gets the value of RevisionID for the instance
func (instance *CIM_PCIDevice) GetPropertyRevisionID() (value uint8, err error) {
	retValue, err := instance.GetProperty("RevisionID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubsystemID sets the value of SubsystemID for the instance
func (instance *CIM_PCIDevice) SetPropertySubsystemID(value uint16) (err error) {
	return instance.SetProperty("SubsystemID", value)
}

// GetSubsystemID gets the value of SubsystemID for the instance
func (instance *CIM_PCIDevice) GetPropertySubsystemID() (value uint16, err error) {
	retValue, err := instance.GetProperty("SubsystemID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubsystemVendorID sets the value of SubsystemVendorID for the instance
func (instance *CIM_PCIDevice) SetPropertySubsystemVendorID(value uint16) (err error) {
	return instance.SetProperty("SubsystemVendorID", value)
}

// GetSubsystemVendorID gets the value of SubsystemVendorID for the instance
func (instance *CIM_PCIDevice) GetPropertySubsystemVendorID() (value uint16, err error) {
	retValue, err := instance.GetProperty("SubsystemVendorID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVendorID sets the value of VendorID for the instance
func (instance *CIM_PCIDevice) SetPropertyVendorID(value uint16) (err error) {
	return instance.SetProperty("VendorID", value)
}

// GetVendorID gets the value of VendorID for the instance
func (instance *CIM_PCIDevice) GetPropertyVendorID() (value uint16, err error) {
	retValue, err := instance.GetProperty("VendorID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
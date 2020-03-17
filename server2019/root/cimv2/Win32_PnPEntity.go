// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PnPEntity struct
type Win32_PnPEntity struct {
	CIM_LogicalDevice

	//
	ClassGuid string

	//
	CompatibleID []string

	//
	HardwareID []string

	//
	Manufacturer string

	//
	PNPClass string

	//
	Present bool

	//
	Service string
}

// SetClassGuid sets the value of ClassGuid for the instance
func (instance *Win32_PnPEntity) SetPropertyClassGuid(value string) (err error) {
	return instance.SetProperty("ClassGuid", value)
}

// GetClassGuid gets the value of ClassGuid for the instance
func (instance *Win32_PnPEntity) GetPropertyClassGuid() (value string, err error) {
	retValue, err := instance.GetProperty("ClassGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompatibleID sets the value of CompatibleID for the instance
func (instance *Win32_PnPEntity) SetPropertyCompatibleID(value []string) (err error) {
	return instance.SetProperty("CompatibleID", value)
}

// GetCompatibleID gets the value of CompatibleID for the instance
func (instance *Win32_PnPEntity) GetPropertyCompatibleID() (value []string, err error) {
	retValue, err := instance.GetProperty("CompatibleID")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHardwareID sets the value of HardwareID for the instance
func (instance *Win32_PnPEntity) SetPropertyHardwareID(value []string) (err error) {
	return instance.SetProperty("HardwareID", value)
}

// GetHardwareID gets the value of HardwareID for the instance
func (instance *Win32_PnPEntity) GetPropertyHardwareID() (value []string, err error) {
	retValue, err := instance.GetProperty("HardwareID")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *Win32_PnPEntity) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *Win32_PnPEntity) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPNPClass sets the value of PNPClass for the instance
func (instance *Win32_PnPEntity) SetPropertyPNPClass(value string) (err error) {
	return instance.SetProperty("PNPClass", value)
}

// GetPNPClass gets the value of PNPClass for the instance
func (instance *Win32_PnPEntity) GetPropertyPNPClass() (value string, err error) {
	retValue, err := instance.GetProperty("PNPClass")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPresent sets the value of Present for the instance
func (instance *Win32_PnPEntity) SetPropertyPresent(value bool) (err error) {
	return instance.SetProperty("Present", value)
}

// GetPresent gets the value of Present for the instance
func (instance *Win32_PnPEntity) GetPropertyPresent() (value bool, err error) {
	retValue, err := instance.GetProperty("Present")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetService sets the value of Service for the instance
func (instance *Win32_PnPEntity) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", value)
}

// GetService gets the value of Service for the instance
func (instance *Win32_PnPEntity) GetPropertyService() (value string, err error) {
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

//

// <param name="rebootNeeded" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_PnPEntity) Enable( /* OUT */ rebootNeeded bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="rebootNeeded" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_PnPEntity) Disable( /* OUT */ rebootNeeded bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="devicePropertyKeys" type="string []"></param>

// <param name="deviceProperties" type="Win32_PnPDeviceProperty []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_PnPEntity) GetDeviceProperties( /* OPTIONAL IN */ devicePropertyKeys []string,
	/* OUT */ deviceProperties []Win32_PnPDeviceProperty) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetDeviceProperties", devicePropertyKeys)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
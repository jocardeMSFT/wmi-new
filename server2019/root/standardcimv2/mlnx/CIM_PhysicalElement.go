// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_PhysicalElement struct
type CIM_PhysicalElement struct {
	CIM_ManagedSystemElement

	//
	CanBeFRUed bool

	//
	CreationClassName string

	//
	ManufactureDate string

	//
	Manufacturer string

	//
	Model string

	//
	OtherIdentifyingInfo string

	//
	PartNumber string

	//
	PoweredOn bool

	//
	SerialNumber string

	//
	SKU string

	//
	Tag string

	//
	UserTracking string

	//
	VendorEquipmentType string

	//
	Version string
}

// SetCanBeFRUed sets the value of CanBeFRUed for the instance
func (instance *CIM_PhysicalElement) SetPropertyCanBeFRUed(value bool) (err error) {
	return instance.SetProperty("CanBeFRUed", value)
}

// GetCanBeFRUed gets the value of CanBeFRUed for the instance
func (instance *CIM_PhysicalElement) GetPropertyCanBeFRUed() (value bool, err error) {
	retValue, err := instance.GetProperty("CanBeFRUed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_PhysicalElement) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_PhysicalElement) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufactureDate sets the value of ManufactureDate for the instance
func (instance *CIM_PhysicalElement) SetPropertyManufactureDate(value string) (err error) {
	return instance.SetProperty("ManufactureDate", value)
}

// GetManufactureDate gets the value of ManufactureDate for the instance
func (instance *CIM_PhysicalElement) GetPropertyManufactureDate() (value string, err error) {
	retValue, err := instance.GetProperty("ManufactureDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *CIM_PhysicalElement) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *CIM_PhysicalElement) GetPropertyManufacturer() (value string, err error) {
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

// SetModel sets the value of Model for the instance
func (instance *CIM_PhysicalElement) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", value)
}

// GetModel gets the value of Model for the instance
func (instance *CIM_PhysicalElement) GetPropertyModel() (value string, err error) {
	retValue, err := instance.GetProperty("Model")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherIdentifyingInfo sets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_PhysicalElement) SetPropertyOtherIdentifyingInfo(value string) (err error) {
	return instance.SetProperty("OtherIdentifyingInfo", value)
}

// GetOtherIdentifyingInfo gets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_PhysicalElement) GetPropertyOtherIdentifyingInfo() (value string, err error) {
	retValue, err := instance.GetProperty("OtherIdentifyingInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPartNumber sets the value of PartNumber for the instance
func (instance *CIM_PhysicalElement) SetPropertyPartNumber(value string) (err error) {
	return instance.SetProperty("PartNumber", value)
}

// GetPartNumber gets the value of PartNumber for the instance
func (instance *CIM_PhysicalElement) GetPropertyPartNumber() (value string, err error) {
	retValue, err := instance.GetProperty("PartNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPoweredOn sets the value of PoweredOn for the instance
func (instance *CIM_PhysicalElement) SetPropertyPoweredOn(value bool) (err error) {
	return instance.SetProperty("PoweredOn", value)
}

// GetPoweredOn gets the value of PoweredOn for the instance
func (instance *CIM_PhysicalElement) GetPropertyPoweredOn() (value bool, err error) {
	retValue, err := instance.GetProperty("PoweredOn")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *CIM_PhysicalElement) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", value)
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *CIM_PhysicalElement) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSKU sets the value of SKU for the instance
func (instance *CIM_PhysicalElement) SetPropertySKU(value string) (err error) {
	return instance.SetProperty("SKU", value)
}

// GetSKU gets the value of SKU for the instance
func (instance *CIM_PhysicalElement) GetPropertySKU() (value string, err error) {
	retValue, err := instance.GetProperty("SKU")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTag sets the value of Tag for the instance
func (instance *CIM_PhysicalElement) SetPropertyTag(value string) (err error) {
	return instance.SetProperty("Tag", value)
}

// GetTag gets the value of Tag for the instance
func (instance *CIM_PhysicalElement) GetPropertyTag() (value string, err error) {
	retValue, err := instance.GetProperty("Tag")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserTracking sets the value of UserTracking for the instance
func (instance *CIM_PhysicalElement) SetPropertyUserTracking(value string) (err error) {
	return instance.SetProperty("UserTracking", value)
}

// GetUserTracking gets the value of UserTracking for the instance
func (instance *CIM_PhysicalElement) GetPropertyUserTracking() (value string, err error) {
	retValue, err := instance.GetProperty("UserTracking")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVendorEquipmentType sets the value of VendorEquipmentType for the instance
func (instance *CIM_PhysicalElement) SetPropertyVendorEquipmentType(value string) (err error) {
	return instance.SetProperty("VendorEquipmentType", value)
}

// GetVendorEquipmentType gets the value of VendorEquipmentType for the instance
func (instance *CIM_PhysicalElement) GetPropertyVendorEquipmentType() (value string, err error) {
	retValue, err := instance.GetProperty("VendorEquipmentType")
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
func (instance *CIM_PhysicalElement) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *CIM_PhysicalElement) GetPropertyVersion() (value string, err error) {
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
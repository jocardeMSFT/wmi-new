// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_ProtocolEndpoint struct
type CIM_ProtocolEndpoint struct {
	CIM_ServiceAccessPoint

	// 381
	NameFormat string

	// 646
	OtherTypeDescription string

	// 409
	ProtocolIFType ProtocolEndpoint_ProtocolIFType

	// 382
	ProtocolType ProtocolEndpoint_ProtocolType
}

// SetNameFormat sets the value of NameFormat for the instance
func (instance *CIM_ProtocolEndpoint) SetPropertyNameFormat(value string) (err error) {
	return instance.SetProperty("NameFormat", value)
}

// GetNameFormat gets the value of NameFormat for the instance
func (instance *CIM_ProtocolEndpoint) GetPropertyNameFormat() (value string, err error) {
	retValue, err := instance.GetProperty("NameFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherTypeDescription sets the value of OtherTypeDescription for the instance
func (instance *CIM_ProtocolEndpoint) SetPropertyOtherTypeDescription(value string) (err error) {
	return instance.SetProperty("OtherTypeDescription", value)
}

// GetOtherTypeDescription gets the value of OtherTypeDescription for the instance
func (instance *CIM_ProtocolEndpoint) GetPropertyOtherTypeDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherTypeDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocolIFType sets the value of ProtocolIFType for the instance
func (instance *CIM_ProtocolEndpoint) SetPropertyProtocolIFType(value ProtocolEndpoint_ProtocolIFType) (err error) {
	return instance.SetProperty("ProtocolIFType", value)
}

// GetProtocolIFType gets the value of ProtocolIFType for the instance
func (instance *CIM_ProtocolEndpoint) GetPropertyProtocolIFType() (value ProtocolEndpoint_ProtocolIFType, err error) {
	retValue, err := instance.GetProperty("ProtocolIFType")
	if err != nil {
		return
	}
	value, ok := retValue.(ProtocolEndpoint_ProtocolIFType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocolType sets the value of ProtocolType for the instance
func (instance *CIM_ProtocolEndpoint) SetPropertyProtocolType(value ProtocolEndpoint_ProtocolType) (err error) {
	return instance.SetProperty("ProtocolType", value)
}

// GetProtocolType gets the value of ProtocolType for the instance
func (instance *CIM_ProtocolEndpoint) GetPropertyProtocolType() (value ProtocolEndpoint_ProtocolType, err error) {
	retValue, err := instance.GetProperty("ProtocolType")
	if err != nil {
		return
	}
	value, ok := retValue.(ProtocolEndpoint_ProtocolType)
	if !ok {
		// TODO: Set an error
	}
	return
}
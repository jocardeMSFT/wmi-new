// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_ProtocolEndpoint struct
type CIM_ProtocolEndpoint struct {
	CIM_ServiceAccessPoint

	// NameFormat contains the naming heuristic that is selected to ensure that the value of the Name property is unique. For example, you might choose to prepend the name of the port or interface with the Type of ProtocolEndpoint (for example, IPv4) of this instance followed by an underscore.
	NameFormat string

	// A string that describes the type of ProtocolEndpoint when the Type property of this class (or any of its subclasses) is set to 1 (Other). This property should be set to null when the Type property is any value other than 1.
	OtherTypeDescription string

	// ProtocolIFType is an enumeration that is synchronized with the IANA ifType MIB. The ifType MIB is maintained at the URL, http://www.iana.org/assignments/ianaiftype-mib. Also, additional values defined by the DMTF are included. The property is used to categorize and classify instances of the ProtocolEndpoint class. Note that if the ProtocolIFType is set to 1 (Other), then the type information should be provided in the OtherTypeDescription string property.
	ProtocolIFType ProtocolEndpoint_ProtocolIFType

	// Note: This property is deprecated in lieu of the ProtocolIFType enumeration. This deprecation was done to have better alignment between the IF-MIB of the IETF and this CIM class.
	///Deprecated description: ProtocolType is an enumeration that provides information to categorize and classify different instances of this class. For most instances, information in this enumeration and the definition of the subclass overlap. However, there are several cases where a specific subclass of ProtocolEndpoint is not required (for example, there is no Fibre Channel subclass of ProtocolEndpoint). Therefore, this property is needed to define the type of Endpoint.
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
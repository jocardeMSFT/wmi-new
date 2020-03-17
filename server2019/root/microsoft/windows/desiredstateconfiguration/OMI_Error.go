// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// OMI_Error struct
type OMI_Error struct {
	CIM_Error

	//
	error_Category uint16

	//
	error_Code uint32

	//
	error_Type string
}

// Seterror_Category sets the value of error_Category for the instance
func (instance *OMI_Error) SetPropertyerror_Category(value uint16) (err error) {
	return instance.SetProperty("error_Category", value)
}

// Geterror_Category gets the value of error_Category for the instance
func (instance *OMI_Error) GetPropertyerror_Category() (value uint16, err error) {
	retValue, err := instance.GetProperty("error_Category")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Seterror_Code sets the value of error_Code for the instance
func (instance *OMI_Error) SetPropertyerror_Code(value uint32) (err error) {
	return instance.SetProperty("error_Code", value)
}

// Geterror_Code gets the value of error_Code for the instance
func (instance *OMI_Error) GetPropertyerror_Code() (value uint32, err error) {
	retValue, err := instance.GetProperty("error_Code")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Seterror_Type sets the value of error_Type for the instance
func (instance *OMI_Error) SetPropertyerror_Type(value string) (err error) {
	return instance.SetProperty("error_Type", value)
}

// Geterror_Type gets the value of error_Type for the instance
func (instance *OMI_Error) GetPropertyerror_Type() (value string, err error) {
	retValue, err := instance.GetProperty("error_Type")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_Dependency struct
type CIM_Dependency struct {
	cim.WmiInstance

	//
	Antecedent CIM_ManagedSystemElement

	//
	Dependent CIM_ManagedSystemElement
}

// SetAntecedent sets the value of Antecedent for the instance
func (instance *CIM_Dependency) SetPropertyAntecedent(value CIM_ManagedSystemElement) (err error) {
	return instance.SetProperty("Antecedent", value)
}

// GetAntecedent gets the value of Antecedent for the instance
func (instance *CIM_Dependency) GetPropertyAntecedent() (value CIM_ManagedSystemElement, err error) {
	retValue, err := instance.GetProperty("Antecedent")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedSystemElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDependent sets the value of Dependent for the instance
func (instance *CIM_Dependency) SetPropertyDependent(value CIM_ManagedSystemElement) (err error) {
	return instance.SetProperty("Dependent", value)
}

// GetDependent gets the value of Dependent for the instance
func (instance *CIM_Dependency) GetPropertyDependent() (value CIM_ManagedSystemElement, err error) {
	retValue, err := instance.GetProperty("Dependent")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedSystemElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
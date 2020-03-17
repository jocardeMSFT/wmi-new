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

// CIM_ActionSequence struct
type CIM_ActionSequence struct {
	cim.WmiInstance

	//
	Next CIM_Action

	//
	Prior CIM_Action
}

// SetNext sets the value of Next for the instance
func (instance *CIM_ActionSequence) SetPropertyNext(value CIM_Action) (err error) {
	return instance.SetProperty("Next", value)
}

// GetNext gets the value of Next for the instance
func (instance *CIM_ActionSequence) GetPropertyNext() (value CIM_Action, err error) {
	retValue, err := instance.GetProperty("Next")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Action)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrior sets the value of Prior for the instance
func (instance *CIM_ActionSequence) SetPropertyPrior(value CIM_Action) (err error) {
	return instance.SetProperty("Prior", value)
}

// GetPrior gets the value of Prior for the instance
func (instance *CIM_ActionSequence) GetPropertyPrior() (value CIM_Action, err error) {
	retValue, err := instance.GetProperty("Prior")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Action)
	if !ok {
		// TODO: Set an error
	}
	return
}
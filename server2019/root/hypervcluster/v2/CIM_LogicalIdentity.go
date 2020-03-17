// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_LogicalIdentity struct
type CIM_LogicalIdentity struct {
	cim.WmiInstance

	// SameElement represents an alternate aspect of the ManagedElement.
	SameElement CIM_ManagedElement

	// SystemElement represents one aspect of the Managed Element. The use of 'System' in the role name does not limit the scope of the association. The role name was defined in the original association, where the referenced elements were limited to LogicalElements. Since that time, it has been found valuable to instantiate these types of relationships for ManagedElements, such as Collections. So, the referenced elements of the association were redefined to be ManagedElements. Unfortunately, the role name could not be changed without deprecating the entire association. This was not deemed necessary just to correct the role name.
	SystemElement CIM_ManagedElement
}

// SetSameElement sets the value of SameElement for the instance
func (instance *CIM_LogicalIdentity) SetPropertySameElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("SameElement", value)
}

// GetSameElement gets the value of SameElement for the instance
func (instance *CIM_LogicalIdentity) GetPropertySameElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("SameElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemElement sets the value of SystemElement for the instance
func (instance *CIM_LogicalIdentity) SetPropertySystemElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("SystemElement", value)
}

// GetSystemElement gets the value of SystemElement for the instance
func (instance *CIM_LogicalIdentity) GetPropertySystemElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("SystemElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
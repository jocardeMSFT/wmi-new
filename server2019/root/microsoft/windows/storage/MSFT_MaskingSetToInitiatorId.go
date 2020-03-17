// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MaskingSetToInitiatorId struct
type MSFT_MaskingSetToInitiatorId struct {
	cim.WmiInstance

	//
	InitiatorId MSFT_InitiatorId

	//
	MaskingSet MSFT_MaskingSet
}

// SetInitiatorId sets the value of InitiatorId for the instance
func (instance *MSFT_MaskingSetToInitiatorId) SetPropertyInitiatorId(value MSFT_InitiatorId) (err error) {
	return instance.SetProperty("InitiatorId", value)
}

// GetInitiatorId gets the value of InitiatorId for the instance
func (instance *MSFT_MaskingSetToInitiatorId) GetPropertyInitiatorId() (value MSFT_InitiatorId, err error) {
	retValue, err := instance.GetProperty("InitiatorId")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_InitiatorId)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaskingSet sets the value of MaskingSet for the instance
func (instance *MSFT_MaskingSetToInitiatorId) SetPropertyMaskingSet(value MSFT_MaskingSet) (err error) {
	return instance.SetProperty("MaskingSet", value)
}

// GetMaskingSet gets the value of MaskingSet for the instance
func (instance *MSFT_MaskingSetToInitiatorId) GetPropertyMaskingSet() (value MSFT_MaskingSet, err error) {
	retValue, err := instance.GetProperty("MaskingSet")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_MaskingSet)
	if !ok {
		// TODO: Set an error
	}
	return
}
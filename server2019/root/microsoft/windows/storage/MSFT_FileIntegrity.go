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

// MSFT_FileIntegrity struct
type MSFT_FileIntegrity struct {
	cim.WmiInstance

	//
	Enabled bool

	//
	Enforced bool

	//
	FileName string
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_FileIntegrity) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_FileIntegrity) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnforced sets the value of Enforced for the instance
func (instance *MSFT_FileIntegrity) SetPropertyEnforced(value bool) (err error) {
	return instance.SetProperty("Enforced", value)
}

// GetEnforced gets the value of Enforced for the instance
func (instance *MSFT_FileIntegrity) GetPropertyEnforced() (value bool, err error) {
	retValue, err := instance.GetProperty("Enforced")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileName sets the value of FileName for the instance
func (instance *MSFT_FileIntegrity) SetPropertyFileName(value string) (err error) {
	return instance.SetProperty("FileName", value)
}

// GetFileName gets the value of FileName for the instance
func (instance *MSFT_FileIntegrity) GetPropertyFileName() (value string, err error) {
	retValue, err := instance.GetProperty("FileName")
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

// <param name="FileName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="FileIntegrity" type="MSFT_FileIntegrity "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileIntegrity) Get( /* IN */ FileName string,
	/* OUT */ FileIntegrity MSFT_FileIntegrity,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", FileName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FileName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileIntegrity) Repair( /* IN */ FileName string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Repair", FileName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Enable" type="bool "></param>
// <param name="Enforce" type="bool "></param>
// <param name="FileName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileIntegrity) Set( /* IN */ FileName string,
	/* IN */ Enable bool,
	/* IN */ Enforce bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", FileName, Enable, Enforce)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
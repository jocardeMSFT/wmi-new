// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Appv
//////////////////////////////////////////////
package appv

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// AppvClientApplication struct
type AppvClientApplication struct {
	cim.WmiInstance

	//
	ApplicationId string

	//
	EnabledForUser bool

	//
	EnabledGlobally bool

	//
	Name string

	//
	PackageId string

	//
	PackageVersionId string

	//
	TargetPath string

	//
	Version string
}

// SetApplicationId sets the value of ApplicationId for the instance
func (instance *AppvClientApplication) SetPropertyApplicationId(value string) (err error) {
	return instance.SetProperty("ApplicationId", value)
}

// GetApplicationId gets the value of ApplicationId for the instance
func (instance *AppvClientApplication) GetPropertyApplicationId() (value string, err error) {
	retValue, err := instance.GetProperty("ApplicationId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabledForUser sets the value of EnabledForUser for the instance
func (instance *AppvClientApplication) SetPropertyEnabledForUser(value bool) (err error) {
	return instance.SetProperty("EnabledForUser", value)
}

// GetEnabledForUser gets the value of EnabledForUser for the instance
func (instance *AppvClientApplication) GetPropertyEnabledForUser() (value bool, err error) {
	retValue, err := instance.GetProperty("EnabledForUser")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabledGlobally sets the value of EnabledGlobally for the instance
func (instance *AppvClientApplication) SetPropertyEnabledGlobally(value bool) (err error) {
	return instance.SetProperty("EnabledGlobally", value)
}

// GetEnabledGlobally gets the value of EnabledGlobally for the instance
func (instance *AppvClientApplication) GetPropertyEnabledGlobally() (value bool, err error) {
	retValue, err := instance.GetProperty("EnabledGlobally")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *AppvClientApplication) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *AppvClientApplication) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageId sets the value of PackageId for the instance
func (instance *AppvClientApplication) SetPropertyPackageId(value string) (err error) {
	return instance.SetProperty("PackageId", value)
}

// GetPackageId gets the value of PackageId for the instance
func (instance *AppvClientApplication) GetPropertyPackageId() (value string, err error) {
	retValue, err := instance.GetProperty("PackageId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageVersionId sets the value of PackageVersionId for the instance
func (instance *AppvClientApplication) SetPropertyPackageVersionId(value string) (err error) {
	return instance.SetProperty("PackageVersionId", value)
}

// GetPackageVersionId gets the value of PackageVersionId for the instance
func (instance *AppvClientApplication) GetPropertyPackageVersionId() (value string, err error) {
	retValue, err := instance.GetProperty("PackageVersionId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetPath sets the value of TargetPath for the instance
func (instance *AppvClientApplication) SetPropertyTargetPath(value string) (err error) {
	return instance.SetProperty("TargetPath", value)
}

// GetTargetPath gets the value of TargetPath for the instance
func (instance *AppvClientApplication) GetPropertyTargetPath() (value string, err error) {
	retValue, err := instance.GetProperty("TargetPath")
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
func (instance *AppvClientApplication) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *AppvClientApplication) GetPropertyVersion() (value string, err error) {
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
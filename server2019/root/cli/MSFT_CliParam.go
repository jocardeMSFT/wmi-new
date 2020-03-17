// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Cli
//////////////////////////////////////////////
package cli

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_CliParam struct
type MSFT_CliParam struct {
	cim.WmiInstance

	//
	Default string

	//
	Description string

	//
	Optional bool

	//
	ParaId string

	//
	Qualifiers []MSFT_CliQualifier

	//
	Type string
}

// SetDefault sets the value of Default for the instance
func (instance *MSFT_CliParam) SetPropertyDefault(value string) (err error) {
	return instance.SetProperty("Default", value)
}

// GetDefault gets the value of Default for the instance
func (instance *MSFT_CliParam) GetPropertyDefault() (value string, err error) {
	retValue, err := instance.GetProperty("Default")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_CliParam) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_CliParam) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOptional sets the value of Optional for the instance
func (instance *MSFT_CliParam) SetPropertyOptional(value bool) (err error) {
	return instance.SetProperty("Optional", value)
}

// GetOptional gets the value of Optional for the instance
func (instance *MSFT_CliParam) GetPropertyOptional() (value bool, err error) {
	retValue, err := instance.GetProperty("Optional")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParaId sets the value of ParaId for the instance
func (instance *MSFT_CliParam) SetPropertyParaId(value string) (err error) {
	return instance.SetProperty("ParaId", value)
}

// GetParaId gets the value of ParaId for the instance
func (instance *MSFT_CliParam) GetPropertyParaId() (value string, err error) {
	retValue, err := instance.GetProperty("ParaId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQualifiers sets the value of Qualifiers for the instance
func (instance *MSFT_CliParam) SetPropertyQualifiers(value []MSFT_CliQualifier) (err error) {
	return instance.SetProperty("Qualifiers", value)
}

// GetQualifiers gets the value of Qualifiers for the instance
func (instance *MSFT_CliParam) GetPropertyQualifiers() (value []MSFT_CliQualifier, err error) {
	retValue, err := instance.GetProperty("Qualifiers")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_CliQualifier)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_CliParam) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_CliParam) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
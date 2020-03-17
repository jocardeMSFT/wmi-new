// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_FileResourceManager struct
type MSFT_FileResourceManager struct {
	OMI_ResourceModuleManager

	//
	Credential MSFT_Credential

	//
	SourcePath string
}

// SetCredential sets the value of Credential for the instance
func (instance *MSFT_FileResourceManager) SetPropertyCredential(value MSFT_Credential) (err error) {
	return instance.SetProperty("Credential", value)
}

// GetCredential gets the value of Credential for the instance
func (instance *MSFT_FileResourceManager) GetPropertyCredential() (value MSFT_Credential, err error) {
	retValue, err := instance.GetProperty("Credential")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Credential)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourcePath sets the value of SourcePath for the instance
func (instance *MSFT_FileResourceManager) SetPropertySourcePath(value string) (err error) {
	return instance.SetProperty("SourcePath", value)
}

// GetSourcePath gets the value of SourcePath for the instance
func (instance *MSFT_FileResourceManager) GetPropertySourcePath() (value string, err error) {
	retValue, err := instance.GetProperty("SourcePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
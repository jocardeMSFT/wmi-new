// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_ResourcePoolConfigurationCapabilities struct
type CIM_ResourcePoolConfigurationCapabilities struct {
	CIM_Capabilities

	// This property reflects the methods of the associated service class that are supported that may return a Job.
	AsynchronousMethodsSupported []ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported

	// This property reflects the methods of the associated service class that are supported andblock until completed (no Job is returned.)
	SynchronousMethodsSupported []ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported
}

// SetAsynchronousMethodsSupported sets the value of AsynchronousMethodsSupported for the instance
func (instance *CIM_ResourcePoolConfigurationCapabilities) SetPropertyAsynchronousMethodsSupported(value []ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported) (err error) {
	return instance.SetProperty("AsynchronousMethodsSupported", value)
}

// GetAsynchronousMethodsSupported gets the value of AsynchronousMethodsSupported for the instance
func (instance *CIM_ResourcePoolConfigurationCapabilities) GetPropertyAsynchronousMethodsSupported() (value []ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported, err error) {
	retValue, err := instance.GetProperty("AsynchronousMethodsSupported")
	if err != nil {
		return
	}
	value, ok := retValue.([]ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSynchronousMethodsSupported sets the value of SynchronousMethodsSupported for the instance
func (instance *CIM_ResourcePoolConfigurationCapabilities) SetPropertySynchronousMethodsSupported(value []ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported) (err error) {
	return instance.SetProperty("SynchronousMethodsSupported", value)
}

// GetSynchronousMethodsSupported gets the value of SynchronousMethodsSupported for the instance
func (instance *CIM_ResourcePoolConfigurationCapabilities) GetPropertySynchronousMethodsSupported() (value []ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported, err error) {
	retValue, err := instance.GetProperty("SynchronousMethodsSupported")
	if err != nil {
		return
	}
	value, ok := retValue.([]ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported)
	if !ok {
		// TODO: Set an error
	}
	return
}
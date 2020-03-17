// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_EthernetSwitchHardwareOffloadSettingData struct
type Msvm_EthernetSwitchHardwareOffloadSettingData struct {
	Msvm_EthernetSwitchFeatureSettingData

	//
	DefaultQueueVmmqEnabled bool

	//
	DefaultQueueVmmqQueuePairs uint32

	//
	DefaultQueueVrssEnabled bool

	//
	DefaultQueueVrssExcludePrimaryProcessor bool

	//
	DefaultQueueVrssIndependentHostSpreading bool

	//
	DefaultQueueVrssMinQueuePairs uint32

	//
	DefaultQueueVrssQueueSchedulingMode uint32

	//
	SoftwareRscEnabled bool
}

// SetDefaultQueueVmmqEnabled sets the value of DefaultQueueVmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVmmqEnabled(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVmmqEnabled", value)
}

// GetDefaultQueueVmmqEnabled gets the value of DefaultQueueVmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVmmqEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVmmqEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVmmqQueuePairs sets the value of DefaultQueueVmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVmmqQueuePairs(value uint32) (err error) {
	return instance.SetProperty("DefaultQueueVmmqQueuePairs", value)
}

// GetDefaultQueueVmmqQueuePairs gets the value of DefaultQueueVmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVmmqQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVmmqQueuePairs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVrssEnabled sets the value of DefaultQueueVrssEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVrssEnabled(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVrssEnabled", value)
}

// GetDefaultQueueVrssEnabled gets the value of DefaultQueueVrssEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVrssEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVrssExcludePrimaryProcessor sets the value of DefaultQueueVrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVrssExcludePrimaryProcessor(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVrssExcludePrimaryProcessor", value)
}

// GetDefaultQueueVrssExcludePrimaryProcessor gets the value of DefaultQueueVrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVrssExcludePrimaryProcessor() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssExcludePrimaryProcessor")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVrssIndependentHostSpreading sets the value of DefaultQueueVrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVrssIndependentHostSpreading(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVrssIndependentHostSpreading", value)
}

// GetDefaultQueueVrssIndependentHostSpreading gets the value of DefaultQueueVrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVrssIndependentHostSpreading() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssIndependentHostSpreading")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVrssMinQueuePairs sets the value of DefaultQueueVrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVrssMinQueuePairs(value uint32) (err error) {
	return instance.SetProperty("DefaultQueueVrssMinQueuePairs", value)
}

// GetDefaultQueueVrssMinQueuePairs gets the value of DefaultQueueVrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVrssMinQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssMinQueuePairs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQueueVrssQueueSchedulingMode sets the value of DefaultQueueVrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVrssQueueSchedulingMode(value uint32) (err error) {
	return instance.SetProperty("DefaultQueueVrssQueueSchedulingMode", value)
}

// GetDefaultQueueVrssQueueSchedulingMode gets the value of DefaultQueueVrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVrssQueueSchedulingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssQueueSchedulingMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSoftwareRscEnabled sets the value of SoftwareRscEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertySoftwareRscEnabled(value bool) (err error) {
	return instance.SetProperty("SoftwareRscEnabled", value)
}

// GetSoftwareRscEnabled gets the value of SoftwareRscEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertySoftwareRscEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("SoftwareRscEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
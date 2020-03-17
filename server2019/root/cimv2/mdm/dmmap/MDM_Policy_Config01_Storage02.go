// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Config01_Storage02 struct
type MDM_Policy_Config01_Storage02 struct {
	cim.WmiInstance

	//
	AllowDiskHealthModelUpdates int32

	//
	AllowStorageSenseGlobal int32

	//
	AllowStorageSenseTemporaryFilesCleanup int32

	//
	ConfigStorageSenseCloudContentDehydrationThreshold int32

	//
	ConfigStorageSenseDownloadsCleanupThreshold int32

	//
	ConfigStorageSenseGlobalCadence int32

	//
	ConfigStorageSenseRecycleBinCleanupThreshold int32

	//
	EnhancedStorageDevices string

	//
	InstanceID string

	//
	ParentID string

	//
	RemovableDiskDenyWriteAccess int32
}

// SetAllowDiskHealthModelUpdates sets the value of AllowDiskHealthModelUpdates for the instance
func (instance *MDM_Policy_Config01_Storage02) SetPropertyAllowDiskHealthModelUpdates(value int32) (err error) {
	return instance.SetProperty("AllowDiskHealthModelUpdates", value)
}

// GetAllowDiskHealthModelUpdates gets the value of AllowDiskHealthModelUpdates for the instance
func (instance *MDM_Policy_Config01_Storage02) GetPropertyAllowDiskHealthModelUpdates() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowDiskHealthModelUpdates")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowStorageSenseGlobal sets the value of AllowStorageSenseGlobal for the instance
func (instance *MDM_Policy_Config01_Storage02) SetPropertyAllowStorageSenseGlobal(value int32) (err error) {
	return instance.SetProperty("AllowStorageSenseGlobal", value)
}

// GetAllowStorageSenseGlobal gets the value of AllowStorageSenseGlobal for the instance
func (instance *MDM_Policy_Config01_Storage02) GetPropertyAllowStorageSenseGlobal() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowStorageSenseGlobal")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowStorageSenseTemporaryFilesCleanup sets the value of AllowStorageSenseTemporaryFilesCleanup for the instance
func (instance *MDM_Policy_Config01_Storage02) SetPropertyAllowStorageSenseTemporaryFilesCleanup(value int32) (err error) {
	return instance.SetProperty("AllowStorageSenseTemporaryFilesCleanup", value)
}

// GetAllowStorageSenseTemporaryFilesCleanup gets the value of AllowStorageSenseTemporaryFilesCleanup for the instance
func (instance *MDM_Policy_Config01_Storage02) GetPropertyAllowStorageSenseTemporaryFilesCleanup() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowStorageSenseTemporaryFilesCleanup")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigStorageSenseCloudContentDehydrationThreshold sets the value of ConfigStorageSenseCloudContentDehydrationThreshold for the instance
func (instance *MDM_Policy_Config01_Storage02) SetPropertyConfigStorageSenseCloudContentDehydrationThreshold(value int32) (err error) {
	return instance.SetProperty("ConfigStorageSenseCloudContentDehydrationThreshold", value)
}

// GetConfigStorageSenseCloudContentDehydrationThreshold gets the value of ConfigStorageSenseCloudContentDehydrationThreshold for the instance
func (instance *MDM_Policy_Config01_Storage02) GetPropertyConfigStorageSenseCloudContentDehydrationThreshold() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigStorageSenseCloudContentDehydrationThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigStorageSenseDownloadsCleanupThreshold sets the value of ConfigStorageSenseDownloadsCleanupThreshold for the instance
func (instance *MDM_Policy_Config01_Storage02) SetPropertyConfigStorageSenseDownloadsCleanupThreshold(value int32) (err error) {
	return instance.SetProperty("ConfigStorageSenseDownloadsCleanupThreshold", value)
}

// GetConfigStorageSenseDownloadsCleanupThreshold gets the value of ConfigStorageSenseDownloadsCleanupThreshold for the instance
func (instance *MDM_Policy_Config01_Storage02) GetPropertyConfigStorageSenseDownloadsCleanupThreshold() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigStorageSenseDownloadsCleanupThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigStorageSenseGlobalCadence sets the value of ConfigStorageSenseGlobalCadence for the instance
func (instance *MDM_Policy_Config01_Storage02) SetPropertyConfigStorageSenseGlobalCadence(value int32) (err error) {
	return instance.SetProperty("ConfigStorageSenseGlobalCadence", value)
}

// GetConfigStorageSenseGlobalCadence gets the value of ConfigStorageSenseGlobalCadence for the instance
func (instance *MDM_Policy_Config01_Storage02) GetPropertyConfigStorageSenseGlobalCadence() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigStorageSenseGlobalCadence")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigStorageSenseRecycleBinCleanupThreshold sets the value of ConfigStorageSenseRecycleBinCleanupThreshold for the instance
func (instance *MDM_Policy_Config01_Storage02) SetPropertyConfigStorageSenseRecycleBinCleanupThreshold(value int32) (err error) {
	return instance.SetProperty("ConfigStorageSenseRecycleBinCleanupThreshold", value)
}

// GetConfigStorageSenseRecycleBinCleanupThreshold gets the value of ConfigStorageSenseRecycleBinCleanupThreshold for the instance
func (instance *MDM_Policy_Config01_Storage02) GetPropertyConfigStorageSenseRecycleBinCleanupThreshold() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigStorageSenseRecycleBinCleanupThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnhancedStorageDevices sets the value of EnhancedStorageDevices for the instance
func (instance *MDM_Policy_Config01_Storage02) SetPropertyEnhancedStorageDevices(value string) (err error) {
	return instance.SetProperty("EnhancedStorageDevices", value)
}

// GetEnhancedStorageDevices gets the value of EnhancedStorageDevices for the instance
func (instance *MDM_Policy_Config01_Storage02) GetPropertyEnhancedStorageDevices() (value string, err error) {
	retValue, err := instance.GetProperty("EnhancedStorageDevices")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Storage02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Storage02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Storage02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Storage02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemovableDiskDenyWriteAccess sets the value of RemovableDiskDenyWriteAccess for the instance
func (instance *MDM_Policy_Config01_Storage02) SetPropertyRemovableDiskDenyWriteAccess(value int32) (err error) {
	return instance.SetProperty("RemovableDiskDenyWriteAccess", value)
}

// GetRemovableDiskDenyWriteAccess gets the value of RemovableDiskDenyWriteAccess for the instance
func (instance *MDM_Policy_Config01_Storage02) GetPropertyRemovableDiskDenyWriteAccess() (value int32, err error) {
	retValue, err := instance.GetProperty("RemovableDiskDenyWriteAccess")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_VirtualDiskToReplicaPeer struct
type MSFT_VirtualDiskToReplicaPeer struct {
	MSFT_Synchronized

	//
	ReplicaPeer MSFT_ReplicaPeer

	//
	VirtualDisk MSFT_VirtualDisk
}

// SetReplicaPeer sets the value of ReplicaPeer for the instance
func (instance *MSFT_VirtualDiskToReplicaPeer) SetPropertyReplicaPeer(value MSFT_ReplicaPeer) (err error) {
	return instance.SetProperty("ReplicaPeer", value)
}

// GetReplicaPeer gets the value of ReplicaPeer for the instance
func (instance *MSFT_VirtualDiskToReplicaPeer) GetPropertyReplicaPeer() (value MSFT_ReplicaPeer, err error) {
	retValue, err := instance.GetProperty("ReplicaPeer")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_ReplicaPeer)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualDisk sets the value of VirtualDisk for the instance
func (instance *MSFT_VirtualDiskToReplicaPeer) SetPropertyVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("VirtualDisk", value)
}

// GetVirtualDisk gets the value of VirtualDisk for the instance
func (instance *MSFT_VirtualDiskToReplicaPeer) GetPropertyVirtualDisk() (value MSFT_VirtualDisk, err error) {
	retValue, err := instance.GetProperty("VirtualDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_VirtualDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}
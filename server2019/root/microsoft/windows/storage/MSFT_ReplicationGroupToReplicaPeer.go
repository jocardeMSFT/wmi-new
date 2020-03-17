// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_ReplicationGroupToReplicaPeer struct
type MSFT_ReplicationGroupToReplicaPeer struct {
	MSFT_Synchronized

	//
	ConsistencyState uint16

	//
	ConsistencyType uint16

	//
	ReplicaPeer MSFT_ReplicaPeer

	//
	ReplicationGroup MSFT_ReplicationGroup
}

// SetConsistencyState sets the value of ConsistencyState for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) SetPropertyConsistencyState(value uint16) (err error) {
	return instance.SetProperty("ConsistencyState", value)
}

// GetConsistencyState gets the value of ConsistencyState for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) GetPropertyConsistencyState() (value uint16, err error) {
	retValue, err := instance.GetProperty("ConsistencyState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConsistencyType sets the value of ConsistencyType for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) SetPropertyConsistencyType(value uint16) (err error) {
	return instance.SetProperty("ConsistencyType", value)
}

// GetConsistencyType gets the value of ConsistencyType for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) GetPropertyConsistencyType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ConsistencyType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicaPeer sets the value of ReplicaPeer for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) SetPropertyReplicaPeer(value MSFT_ReplicaPeer) (err error) {
	return instance.SetProperty("ReplicaPeer", value)
}

// GetReplicaPeer gets the value of ReplicaPeer for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) GetPropertyReplicaPeer() (value MSFT_ReplicaPeer, err error) {
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

// SetReplicationGroup sets the value of ReplicationGroup for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) SetPropertyReplicationGroup(value MSFT_ReplicationGroup) (err error) {
	return instance.SetProperty("ReplicationGroup", value)
}

// GetReplicationGroup gets the value of ReplicationGroup for the instance
func (instance *MSFT_ReplicationGroupToReplicaPeer) GetPropertyReplicationGroup() (value MSFT_ReplicationGroup, err error) {
	retValue, err := instance.GetProperty("ReplicationGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_ReplicationGroup)
	if !ok {
		// TODO: Set an error
	}
	return
}
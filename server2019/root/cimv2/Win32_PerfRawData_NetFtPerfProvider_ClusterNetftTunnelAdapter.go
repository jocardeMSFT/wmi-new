// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter struct
type Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter struct {
	Win32_PerfRawData

	//
	TotalReceives uint64

	//
	TotalReceivesAccepted uint64

	//
	TotalReceivesDropped uint64

	//
	TotalSendRequests uint64

	//
	TotalSendRequestsAccepted uint64

	//
	TotalSendRequestsDropped uint64
}

// SetTotalReceives sets the value of TotalReceives for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) SetPropertyTotalReceives(value uint64) (err error) {
	return instance.SetProperty("TotalReceives", value)
}

// GetTotalReceives gets the value of TotalReceives for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) GetPropertyTotalReceives() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalReceives")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalReceivesAccepted sets the value of TotalReceivesAccepted for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) SetPropertyTotalReceivesAccepted(value uint64) (err error) {
	return instance.SetProperty("TotalReceivesAccepted", value)
}

// GetTotalReceivesAccepted gets the value of TotalReceivesAccepted for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) GetPropertyTotalReceivesAccepted() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalReceivesAccepted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalReceivesDropped sets the value of TotalReceivesDropped for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) SetPropertyTotalReceivesDropped(value uint64) (err error) {
	return instance.SetProperty("TotalReceivesDropped", value)
}

// GetTotalReceivesDropped gets the value of TotalReceivesDropped for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) GetPropertyTotalReceivesDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalReceivesDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalSendRequests sets the value of TotalSendRequests for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) SetPropertyTotalSendRequests(value uint64) (err error) {
	return instance.SetProperty("TotalSendRequests", value)
}

// GetTotalSendRequests gets the value of TotalSendRequests for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) GetPropertyTotalSendRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalSendRequests")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalSendRequestsAccepted sets the value of TotalSendRequestsAccepted for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) SetPropertyTotalSendRequestsAccepted(value uint64) (err error) {
	return instance.SetProperty("TotalSendRequestsAccepted", value)
}

// GetTotalSendRequestsAccepted gets the value of TotalSendRequestsAccepted for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) GetPropertyTotalSendRequestsAccepted() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalSendRequestsAccepted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalSendRequestsDropped sets the value of TotalSendRequestsDropped for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) SetPropertyTotalSendRequestsDropped(value uint64) (err error) {
	return instance.SetProperty("TotalSendRequestsDropped", value)
}

// GetTotalSendRequestsDropped gets the value of TotalSendRequestsDropped for the instance
func (instance *Win32_PerfRawData_NetFtPerfProvider_ClusterNetftTunnelAdapter) GetPropertyTotalSendRequestsDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalSendRequestsDropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
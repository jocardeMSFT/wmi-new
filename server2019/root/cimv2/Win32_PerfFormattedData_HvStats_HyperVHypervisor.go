// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_HvStats_HyperVHypervisor struct
type Win32_PerfFormattedData_HvStats_HyperVHypervisor struct {
	Win32_PerfFormattedData

	//
	HypervisorStartupCost uint64

	//
	LogicalProcessors uint64

	//
	ModernStandbyEntries uint64

	//
	MonitoredNotifications uint64

	//
	Partitions uint64

	//
	PlatformIdleTransitions uint64

	//
	TotalPages uint64

	//
	VirtualProcessors uint64
}

// SetHypervisorStartupCost sets the value of HypervisorStartupCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) SetPropertyHypervisorStartupCost(value uint64) (err error) {
	return instance.SetProperty("HypervisorStartupCost", value)
}

// GetHypervisorStartupCost gets the value of HypervisorStartupCost for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) GetPropertyHypervisorStartupCost() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypervisorStartupCost")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogicalProcessors sets the value of LogicalProcessors for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) SetPropertyLogicalProcessors(value uint64) (err error) {
	return instance.SetProperty("LogicalProcessors", value)
}

// GetLogicalProcessors gets the value of LogicalProcessors for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) GetPropertyLogicalProcessors() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogicalProcessors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModernStandbyEntries sets the value of ModernStandbyEntries for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) SetPropertyModernStandbyEntries(value uint64) (err error) {
	return instance.SetProperty("ModernStandbyEntries", value)
}

// GetModernStandbyEntries gets the value of ModernStandbyEntries for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) GetPropertyModernStandbyEntries() (value uint64, err error) {
	retValue, err := instance.GetProperty("ModernStandbyEntries")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonitoredNotifications sets the value of MonitoredNotifications for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) SetPropertyMonitoredNotifications(value uint64) (err error) {
	return instance.SetProperty("MonitoredNotifications", value)
}

// GetMonitoredNotifications gets the value of MonitoredNotifications for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) GetPropertyMonitoredNotifications() (value uint64, err error) {
	retValue, err := instance.GetProperty("MonitoredNotifications")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPartitions sets the value of Partitions for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) SetPropertyPartitions(value uint64) (err error) {
	return instance.SetProperty("Partitions", value)
}

// GetPartitions gets the value of Partitions for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) GetPropertyPartitions() (value uint64, err error) {
	retValue, err := instance.GetProperty("Partitions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlatformIdleTransitions sets the value of PlatformIdleTransitions for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) SetPropertyPlatformIdleTransitions(value uint64) (err error) {
	return instance.SetProperty("PlatformIdleTransitions", value)
}

// GetPlatformIdleTransitions gets the value of PlatformIdleTransitions for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) GetPropertyPlatformIdleTransitions() (value uint64, err error) {
	retValue, err := instance.GetProperty("PlatformIdleTransitions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalPages sets the value of TotalPages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) SetPropertyTotalPages(value uint64) (err error) {
	return instance.SetProperty("TotalPages", value)
}

// GetTotalPages gets the value of TotalPages for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) GetPropertyTotalPages() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalPages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualProcessors sets the value of VirtualProcessors for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) SetPropertyVirtualProcessors(value uint64) (err error) {
	return instance.SetProperty("VirtualProcessors", value)
}

// GetVirtualProcessors gets the value of VirtualProcessors for the instance
func (instance *Win32_PerfFormattedData_HvStats_HyperVHypervisor) GetPropertyVirtualProcessors() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualProcessors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
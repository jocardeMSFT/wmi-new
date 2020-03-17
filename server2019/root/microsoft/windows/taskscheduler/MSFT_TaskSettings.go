// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_TaskSettings struct
type MSFT_TaskSettings struct {
	cim.WmiInstance

	//
	AllowDemandStart bool

	//
	AllowHardTerminate bool

	//
	Compatibility TaskSettings_Compatibility

	//
	DeleteExpiredTaskAfter string

	//
	DisallowStartIfOnBatteries bool

	//
	Enabled bool

	//
	ExecutionTimeLimit string

	//
	Hidden bool

	//
	IdleSettings MSFT_TaskIdleSettings

	//
	MultipleInstances TaskSettings_MultipleInstances

	//
	NetworkSettings MSFT_TaskNetworkSettings

	//
	Priority uint32

	//
	RestartCount uint32

	//
	RestartInterval string

	//
	RunOnlyIfIdle bool

	//
	RunOnlyIfNetworkAvailable bool

	//
	StartWhenAvailable bool

	//
	StopIfGoingOnBatteries bool

	//
	WakeToRun bool
}

// SetAllowDemandStart sets the value of AllowDemandStart for the instance
func (instance *MSFT_TaskSettings) SetPropertyAllowDemandStart(value bool) (err error) {
	return instance.SetProperty("AllowDemandStart", value)
}

// GetAllowDemandStart gets the value of AllowDemandStart for the instance
func (instance *MSFT_TaskSettings) GetPropertyAllowDemandStart() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowDemandStart")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowHardTerminate sets the value of AllowHardTerminate for the instance
func (instance *MSFT_TaskSettings) SetPropertyAllowHardTerminate(value bool) (err error) {
	return instance.SetProperty("AllowHardTerminate", value)
}

// GetAllowHardTerminate gets the value of AllowHardTerminate for the instance
func (instance *MSFT_TaskSettings) GetPropertyAllowHardTerminate() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowHardTerminate")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompatibility sets the value of Compatibility for the instance
func (instance *MSFT_TaskSettings) SetPropertyCompatibility(value TaskSettings_Compatibility) (err error) {
	return instance.SetProperty("Compatibility", value)
}

// GetCompatibility gets the value of Compatibility for the instance
func (instance *MSFT_TaskSettings) GetPropertyCompatibility() (value TaskSettings_Compatibility, err error) {
	retValue, err := instance.GetProperty("Compatibility")
	if err != nil {
		return
	}
	value, ok := retValue.(TaskSettings_Compatibility)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeleteExpiredTaskAfter sets the value of DeleteExpiredTaskAfter for the instance
func (instance *MSFT_TaskSettings) SetPropertyDeleteExpiredTaskAfter(value string) (err error) {
	return instance.SetProperty("DeleteExpiredTaskAfter", value)
}

// GetDeleteExpiredTaskAfter gets the value of DeleteExpiredTaskAfter for the instance
func (instance *MSFT_TaskSettings) GetPropertyDeleteExpiredTaskAfter() (value string, err error) {
	retValue, err := instance.GetProperty("DeleteExpiredTaskAfter")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisallowStartIfOnBatteries sets the value of DisallowStartIfOnBatteries for the instance
func (instance *MSFT_TaskSettings) SetPropertyDisallowStartIfOnBatteries(value bool) (err error) {
	return instance.SetProperty("DisallowStartIfOnBatteries", value)
}

// GetDisallowStartIfOnBatteries gets the value of DisallowStartIfOnBatteries for the instance
func (instance *MSFT_TaskSettings) GetPropertyDisallowStartIfOnBatteries() (value bool, err error) {
	retValue, err := instance.GetProperty("DisallowStartIfOnBatteries")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_TaskSettings) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_TaskSettings) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExecutionTimeLimit sets the value of ExecutionTimeLimit for the instance
func (instance *MSFT_TaskSettings) SetPropertyExecutionTimeLimit(value string) (err error) {
	return instance.SetProperty("ExecutionTimeLimit", value)
}

// GetExecutionTimeLimit gets the value of ExecutionTimeLimit for the instance
func (instance *MSFT_TaskSettings) GetPropertyExecutionTimeLimit() (value string, err error) {
	retValue, err := instance.GetProperty("ExecutionTimeLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHidden sets the value of Hidden for the instance
func (instance *MSFT_TaskSettings) SetPropertyHidden(value bool) (err error) {
	return instance.SetProperty("Hidden", value)
}

// GetHidden gets the value of Hidden for the instance
func (instance *MSFT_TaskSettings) GetPropertyHidden() (value bool, err error) {
	retValue, err := instance.GetProperty("Hidden")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdleSettings sets the value of IdleSettings for the instance
func (instance *MSFT_TaskSettings) SetPropertyIdleSettings(value MSFT_TaskIdleSettings) (err error) {
	return instance.SetProperty("IdleSettings", value)
}

// GetIdleSettings gets the value of IdleSettings for the instance
func (instance *MSFT_TaskSettings) GetPropertyIdleSettings() (value MSFT_TaskIdleSettings, err error) {
	retValue, err := instance.GetProperty("IdleSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_TaskIdleSettings)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMultipleInstances sets the value of MultipleInstances for the instance
func (instance *MSFT_TaskSettings) SetPropertyMultipleInstances(value TaskSettings_MultipleInstances) (err error) {
	return instance.SetProperty("MultipleInstances", value)
}

// GetMultipleInstances gets the value of MultipleInstances for the instance
func (instance *MSFT_TaskSettings) GetPropertyMultipleInstances() (value TaskSettings_MultipleInstances, err error) {
	retValue, err := instance.GetProperty("MultipleInstances")
	if err != nil {
		return
	}
	value, ok := retValue.(TaskSettings_MultipleInstances)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkSettings sets the value of NetworkSettings for the instance
func (instance *MSFT_TaskSettings) SetPropertyNetworkSettings(value MSFT_TaskNetworkSettings) (err error) {
	return instance.SetProperty("NetworkSettings", value)
}

// GetNetworkSettings gets the value of NetworkSettings for the instance
func (instance *MSFT_TaskSettings) GetPropertyNetworkSettings() (value MSFT_TaskNetworkSettings, err error) {
	retValue, err := instance.GetProperty("NetworkSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_TaskNetworkSettings)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriority sets the value of Priority for the instance
func (instance *MSFT_TaskSettings) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *MSFT_TaskSettings) GetPropertyPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestartCount sets the value of RestartCount for the instance
func (instance *MSFT_TaskSettings) SetPropertyRestartCount(value uint32) (err error) {
	return instance.SetProperty("RestartCount", value)
}

// GetRestartCount gets the value of RestartCount for the instance
func (instance *MSFT_TaskSettings) GetPropertyRestartCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("RestartCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestartInterval sets the value of RestartInterval for the instance
func (instance *MSFT_TaskSettings) SetPropertyRestartInterval(value string) (err error) {
	return instance.SetProperty("RestartInterval", value)
}

// GetRestartInterval gets the value of RestartInterval for the instance
func (instance *MSFT_TaskSettings) GetPropertyRestartInterval() (value string, err error) {
	retValue, err := instance.GetProperty("RestartInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunOnlyIfIdle sets the value of RunOnlyIfIdle for the instance
func (instance *MSFT_TaskSettings) SetPropertyRunOnlyIfIdle(value bool) (err error) {
	return instance.SetProperty("RunOnlyIfIdle", value)
}

// GetRunOnlyIfIdle gets the value of RunOnlyIfIdle for the instance
func (instance *MSFT_TaskSettings) GetPropertyRunOnlyIfIdle() (value bool, err error) {
	retValue, err := instance.GetProperty("RunOnlyIfIdle")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunOnlyIfNetworkAvailable sets the value of RunOnlyIfNetworkAvailable for the instance
func (instance *MSFT_TaskSettings) SetPropertyRunOnlyIfNetworkAvailable(value bool) (err error) {
	return instance.SetProperty("RunOnlyIfNetworkAvailable", value)
}

// GetRunOnlyIfNetworkAvailable gets the value of RunOnlyIfNetworkAvailable for the instance
func (instance *MSFT_TaskSettings) GetPropertyRunOnlyIfNetworkAvailable() (value bool, err error) {
	retValue, err := instance.GetProperty("RunOnlyIfNetworkAvailable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartWhenAvailable sets the value of StartWhenAvailable for the instance
func (instance *MSFT_TaskSettings) SetPropertyStartWhenAvailable(value bool) (err error) {
	return instance.SetProperty("StartWhenAvailable", value)
}

// GetStartWhenAvailable gets the value of StartWhenAvailable for the instance
func (instance *MSFT_TaskSettings) GetPropertyStartWhenAvailable() (value bool, err error) {
	retValue, err := instance.GetProperty("StartWhenAvailable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStopIfGoingOnBatteries sets the value of StopIfGoingOnBatteries for the instance
func (instance *MSFT_TaskSettings) SetPropertyStopIfGoingOnBatteries(value bool) (err error) {
	return instance.SetProperty("StopIfGoingOnBatteries", value)
}

// GetStopIfGoingOnBatteries gets the value of StopIfGoingOnBatteries for the instance
func (instance *MSFT_TaskSettings) GetPropertyStopIfGoingOnBatteries() (value bool, err error) {
	retValue, err := instance.GetProperty("StopIfGoingOnBatteries")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWakeToRun sets the value of WakeToRun for the instance
func (instance *MSFT_TaskSettings) SetPropertyWakeToRun(value bool) (err error) {
	return instance.SetProperty("WakeToRun", value)
}

// GetWakeToRun gets the value of WakeToRun for the instance
func (instance *MSFT_TaskSettings) GetPropertyWakeToRun() (value bool, err error) {
	retValue, err := instance.GetProperty("WakeToRun")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
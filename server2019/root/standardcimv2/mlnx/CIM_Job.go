// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_Job struct
type CIM_Job struct {
	CIM_LogicalElement

	//
	DeleteOnCompletion bool

	//
	ElapsedTime string

	//
	ErrorCode uint16

	//
	ErrorDescription string

	//
	JobRunTimes uint32

	//
	JobStatus string

	//
	LocalOrUtcTime Job_LocalOrUtcTime

	//
	Notify string

	//
	OtherRecoveryAction string

	//
	Owner string

	//
	PercentComplete uint16

	//
	Priority uint32

	//
	RecoveryAction Job_RecoveryAction

	//
	RunDay int8

	//
	RunDayOfWeek Job_RunDayOfWeek

	//
	RunMonth Job_RunMonth

	//
	RunStartInterval string

	//
	ScheduledStartTime string

	//
	StartTime string

	//
	TimeSubmitted string

	//
	UntilTime string
}

// SetDeleteOnCompletion sets the value of DeleteOnCompletion for the instance
func (instance *CIM_Job) SetPropertyDeleteOnCompletion(value bool) (err error) {
	return instance.SetProperty("DeleteOnCompletion", value)
}

// GetDeleteOnCompletion gets the value of DeleteOnCompletion for the instance
func (instance *CIM_Job) GetPropertyDeleteOnCompletion() (value bool, err error) {
	retValue, err := instance.GetProperty("DeleteOnCompletion")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElapsedTime sets the value of ElapsedTime for the instance
func (instance *CIM_Job) SetPropertyElapsedTime(value string) (err error) {
	return instance.SetProperty("ElapsedTime", value)
}

// GetElapsedTime gets the value of ElapsedTime for the instance
func (instance *CIM_Job) GetPropertyElapsedTime() (value string, err error) {
	retValue, err := instance.GetProperty("ElapsedTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *CIM_Job) SetPropertyErrorCode(value uint16) (err error) {
	return instance.SetProperty("ErrorCode", value)
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *CIM_Job) GetPropertyErrorCode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorDescription sets the value of ErrorDescription for the instance
func (instance *CIM_Job) SetPropertyErrorDescription(value string) (err error) {
	return instance.SetProperty("ErrorDescription", value)
}

// GetErrorDescription gets the value of ErrorDescription for the instance
func (instance *CIM_Job) GetPropertyErrorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobRunTimes sets the value of JobRunTimes for the instance
func (instance *CIM_Job) SetPropertyJobRunTimes(value uint32) (err error) {
	return instance.SetProperty("JobRunTimes", value)
}

// GetJobRunTimes gets the value of JobRunTimes for the instance
func (instance *CIM_Job) GetPropertyJobRunTimes() (value uint32, err error) {
	retValue, err := instance.GetProperty("JobRunTimes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobStatus sets the value of JobStatus for the instance
func (instance *CIM_Job) SetPropertyJobStatus(value string) (err error) {
	return instance.SetProperty("JobStatus", value)
}

// GetJobStatus gets the value of JobStatus for the instance
func (instance *CIM_Job) GetPropertyJobStatus() (value string, err error) {
	retValue, err := instance.GetProperty("JobStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalOrUtcTime sets the value of LocalOrUtcTime for the instance
func (instance *CIM_Job) SetPropertyLocalOrUtcTime(value Job_LocalOrUtcTime) (err error) {
	return instance.SetProperty("LocalOrUtcTime", value)
}

// GetLocalOrUtcTime gets the value of LocalOrUtcTime for the instance
func (instance *CIM_Job) GetPropertyLocalOrUtcTime() (value Job_LocalOrUtcTime, err error) {
	retValue, err := instance.GetProperty("LocalOrUtcTime")
	if err != nil {
		return
	}
	value, ok := retValue.(Job_LocalOrUtcTime)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotify sets the value of Notify for the instance
func (instance *CIM_Job) SetPropertyNotify(value string) (err error) {
	return instance.SetProperty("Notify", value)
}

// GetNotify gets the value of Notify for the instance
func (instance *CIM_Job) GetPropertyNotify() (value string, err error) {
	retValue, err := instance.GetProperty("Notify")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherRecoveryAction sets the value of OtherRecoveryAction for the instance
func (instance *CIM_Job) SetPropertyOtherRecoveryAction(value string) (err error) {
	return instance.SetProperty("OtherRecoveryAction", value)
}

// GetOtherRecoveryAction gets the value of OtherRecoveryAction for the instance
func (instance *CIM_Job) GetPropertyOtherRecoveryAction() (value string, err error) {
	retValue, err := instance.GetProperty("OtherRecoveryAction")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOwner sets the value of Owner for the instance
func (instance *CIM_Job) SetPropertyOwner(value string) (err error) {
	return instance.SetProperty("Owner", value)
}

// GetOwner gets the value of Owner for the instance
func (instance *CIM_Job) GetPropertyOwner() (value string, err error) {
	retValue, err := instance.GetProperty("Owner")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentComplete sets the value of PercentComplete for the instance
func (instance *CIM_Job) SetPropertyPercentComplete(value uint16) (err error) {
	return instance.SetProperty("PercentComplete", value)
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *CIM_Job) GetPropertyPercentComplete() (value uint16, err error) {
	retValue, err := instance.GetProperty("PercentComplete")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriority sets the value of Priority for the instance
func (instance *CIM_Job) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *CIM_Job) GetPropertyPriority() (value uint32, err error) {
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

// SetRecoveryAction sets the value of RecoveryAction for the instance
func (instance *CIM_Job) SetPropertyRecoveryAction(value Job_RecoveryAction) (err error) {
	return instance.SetProperty("RecoveryAction", value)
}

// GetRecoveryAction gets the value of RecoveryAction for the instance
func (instance *CIM_Job) GetPropertyRecoveryAction() (value Job_RecoveryAction, err error) {
	retValue, err := instance.GetProperty("RecoveryAction")
	if err != nil {
		return
	}
	value, ok := retValue.(Job_RecoveryAction)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunDay sets the value of RunDay for the instance
func (instance *CIM_Job) SetPropertyRunDay(value int8) (err error) {
	return instance.SetProperty("RunDay", value)
}

// GetRunDay gets the value of RunDay for the instance
func (instance *CIM_Job) GetPropertyRunDay() (value int8, err error) {
	retValue, err := instance.GetProperty("RunDay")
	if err != nil {
		return
	}
	value, ok := retValue.(int8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunDayOfWeek sets the value of RunDayOfWeek for the instance
func (instance *CIM_Job) SetPropertyRunDayOfWeek(value Job_RunDayOfWeek) (err error) {
	return instance.SetProperty("RunDayOfWeek", value)
}

// GetRunDayOfWeek gets the value of RunDayOfWeek for the instance
func (instance *CIM_Job) GetPropertyRunDayOfWeek() (value Job_RunDayOfWeek, err error) {
	retValue, err := instance.GetProperty("RunDayOfWeek")
	if err != nil {
		return
	}
	value, ok := retValue.(Job_RunDayOfWeek)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunMonth sets the value of RunMonth for the instance
func (instance *CIM_Job) SetPropertyRunMonth(value Job_RunMonth) (err error) {
	return instance.SetProperty("RunMonth", value)
}

// GetRunMonth gets the value of RunMonth for the instance
func (instance *CIM_Job) GetPropertyRunMonth() (value Job_RunMonth, err error) {
	retValue, err := instance.GetProperty("RunMonth")
	if err != nil {
		return
	}
	value, ok := retValue.(Job_RunMonth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunStartInterval sets the value of RunStartInterval for the instance
func (instance *CIM_Job) SetPropertyRunStartInterval(value string) (err error) {
	return instance.SetProperty("RunStartInterval", value)
}

// GetRunStartInterval gets the value of RunStartInterval for the instance
func (instance *CIM_Job) GetPropertyRunStartInterval() (value string, err error) {
	retValue, err := instance.GetProperty("RunStartInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScheduledStartTime sets the value of ScheduledStartTime for the instance
func (instance *CIM_Job) SetPropertyScheduledStartTime(value string) (err error) {
	return instance.SetProperty("ScheduledStartTime", value)
}

// GetScheduledStartTime gets the value of ScheduledStartTime for the instance
func (instance *CIM_Job) GetPropertyScheduledStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("ScheduledStartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartTime sets the value of StartTime for the instance
func (instance *CIM_Job) SetPropertyStartTime(value string) (err error) {
	return instance.SetProperty("StartTime", value)
}

// GetStartTime gets the value of StartTime for the instance
func (instance *CIM_Job) GetPropertyStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeSubmitted sets the value of TimeSubmitted for the instance
func (instance *CIM_Job) SetPropertyTimeSubmitted(value string) (err error) {
	return instance.SetProperty("TimeSubmitted", value)
}

// GetTimeSubmitted gets the value of TimeSubmitted for the instance
func (instance *CIM_Job) GetPropertyTimeSubmitted() (value string, err error) {
	retValue, err := instance.GetProperty("TimeSubmitted")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUntilTime sets the value of UntilTime for the instance
func (instance *CIM_Job) SetPropertyUntilTime(value string) (err error) {
	return instance.SetProperty("UntilTime", value)
}

// GetUntilTime gets the value of UntilTime for the instance
func (instance *CIM_Job) GetPropertyUntilTime() (value string, err error) {
	retValue, err := instance.GetProperty("UntilTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="DeleteOnKill" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Job) KillJob( /* IN */ DeleteOnKill bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("KillJob", DeleteOnKill)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
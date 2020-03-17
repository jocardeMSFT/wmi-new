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

// MSFT_ScheduledTask struct
type MSFT_ScheduledTask struct {
	cim.WmiInstance

	//
	Actions []MSFT_TaskAction

	//
	Author string

	//
	Date string

	//
	Description string

	//
	Documentation string

	//
	Principal MSFT_TaskPrincipal

	//
	SecurityDescriptor string

	//
	Settings MSFT_TaskSettings

	//
	Source string

	//
	State ScheduledTask_State

	//
	TaskName string

	//
	TaskPath string

	//
	Triggers []MSFT_TaskTrigger

	//
	URI string

	//
	Version string
}

// SetActions sets the value of Actions for the instance
func (instance *MSFT_ScheduledTask) SetPropertyActions(value []MSFT_TaskAction) (err error) {
	return instance.SetProperty("Actions", value)
}

// GetActions gets the value of Actions for the instance
func (instance *MSFT_ScheduledTask) GetPropertyActions() (value []MSFT_TaskAction, err error) {
	retValue, err := instance.GetProperty("Actions")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_TaskAction)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAuthor sets the value of Author for the instance
func (instance *MSFT_ScheduledTask) SetPropertyAuthor(value string) (err error) {
	return instance.SetProperty("Author", value)
}

// GetAuthor gets the value of Author for the instance
func (instance *MSFT_ScheduledTask) GetPropertyAuthor() (value string, err error) {
	retValue, err := instance.GetProperty("Author")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDate sets the value of Date for the instance
func (instance *MSFT_ScheduledTask) SetPropertyDate(value string) (err error) {
	return instance.SetProperty("Date", value)
}

// GetDate gets the value of Date for the instance
func (instance *MSFT_ScheduledTask) GetPropertyDate() (value string, err error) {
	retValue, err := instance.GetProperty("Date")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_ScheduledTask) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_ScheduledTask) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDocumentation sets the value of Documentation for the instance
func (instance *MSFT_ScheduledTask) SetPropertyDocumentation(value string) (err error) {
	return instance.SetProperty("Documentation", value)
}

// GetDocumentation gets the value of Documentation for the instance
func (instance *MSFT_ScheduledTask) GetPropertyDocumentation() (value string, err error) {
	retValue, err := instance.GetProperty("Documentation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrincipal sets the value of Principal for the instance
func (instance *MSFT_ScheduledTask) SetPropertyPrincipal(value MSFT_TaskPrincipal) (err error) {
	return instance.SetProperty("Principal", value)
}

// GetPrincipal gets the value of Principal for the instance
func (instance *MSFT_ScheduledTask) GetPropertyPrincipal() (value MSFT_TaskPrincipal, err error) {
	retValue, err := instance.GetProperty("Principal")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_TaskPrincipal)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecurityDescriptor sets the value of SecurityDescriptor for the instance
func (instance *MSFT_ScheduledTask) SetPropertySecurityDescriptor(value string) (err error) {
	return instance.SetProperty("SecurityDescriptor", value)
}

// GetSecurityDescriptor gets the value of SecurityDescriptor for the instance
func (instance *MSFT_ScheduledTask) GetPropertySecurityDescriptor() (value string, err error) {
	retValue, err := instance.GetProperty("SecurityDescriptor")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettings sets the value of Settings for the instance
func (instance *MSFT_ScheduledTask) SetPropertySettings(value MSFT_TaskSettings) (err error) {
	return instance.SetProperty("Settings", value)
}

// GetSettings gets the value of Settings for the instance
func (instance *MSFT_ScheduledTask) GetPropertySettings() (value MSFT_TaskSettings, err error) {
	retValue, err := instance.GetProperty("Settings")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_TaskSettings)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSource sets the value of Source for the instance
func (instance *MSFT_ScheduledTask) SetPropertySource(value string) (err error) {
	return instance.SetProperty("Source", value)
}

// GetSource gets the value of Source for the instance
func (instance *MSFT_ScheduledTask) GetPropertySource() (value string, err error) {
	retValue, err := instance.GetProperty("Source")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_ScheduledTask) SetPropertyState(value ScheduledTask_State) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_ScheduledTask) GetPropertyState() (value ScheduledTask_State, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(ScheduledTask_State)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTaskName sets the value of TaskName for the instance
func (instance *MSFT_ScheduledTask) SetPropertyTaskName(value string) (err error) {
	return instance.SetProperty("TaskName", value)
}

// GetTaskName gets the value of TaskName for the instance
func (instance *MSFT_ScheduledTask) GetPropertyTaskName() (value string, err error) {
	retValue, err := instance.GetProperty("TaskName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTaskPath sets the value of TaskPath for the instance
func (instance *MSFT_ScheduledTask) SetPropertyTaskPath(value string) (err error) {
	return instance.SetProperty("TaskPath", value)
}

// GetTaskPath gets the value of TaskPath for the instance
func (instance *MSFT_ScheduledTask) GetPropertyTaskPath() (value string, err error) {
	retValue, err := instance.GetProperty("TaskPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTriggers sets the value of Triggers for the instance
func (instance *MSFT_ScheduledTask) SetPropertyTriggers(value []MSFT_TaskTrigger) (err error) {
	return instance.SetProperty("Triggers", value)
}

// GetTriggers gets the value of Triggers for the instance
func (instance *MSFT_ScheduledTask) GetPropertyTriggers() (value []MSFT_TaskTrigger, err error) {
	retValue, err := instance.GetProperty("Triggers")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_TaskTrigger)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetURI sets the value of URI for the instance
func (instance *MSFT_ScheduledTask) SetPropertyURI(value string) (err error) {
	return instance.SetProperty("URI", value)
}

// GetURI gets the value of URI for the instance
func (instance *MSFT_ScheduledTask) GetPropertyURI() (value string, err error) {
	retValue, err := instance.GetProperty("URI")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *MSFT_ScheduledTask) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *MSFT_ScheduledTask) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
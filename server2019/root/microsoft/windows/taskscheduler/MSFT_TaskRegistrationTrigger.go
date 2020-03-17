// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

// MSFT_TaskRegistrationTrigger struct
type MSFT_TaskRegistrationTrigger struct {
	MSFT_TaskTrigger

	//
	Delay string
}

// SetDelay sets the value of Delay for the instance
func (instance *MSFT_TaskRegistrationTrigger) SetPropertyDelay(value string) (err error) {
	return instance.SetProperty("Delay", value)
}

// GetDelay gets the value of Delay for the instance
func (instance *MSFT_TaskRegistrationTrigger) GetPropertyDelay() (value string, err error) {
	retValue, err := instance.GetProperty("Delay")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
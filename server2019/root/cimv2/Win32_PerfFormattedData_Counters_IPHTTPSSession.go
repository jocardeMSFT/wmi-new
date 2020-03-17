// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_IPHTTPSSession struct
type Win32_PerfFormattedData_Counters_IPHTTPSSession struct {
	Win32_PerfFormattedData

	//
	Bytesreceivedonthissession uint64

	//
	Bytessentonthissession uint64

	//
	DurationDurationofthesessionSeconds uint64

	//
	ErrorsReceiveerrorsonthissession uint64

	//
	ErrorsTransmiterrorsonthissession uint64

	//
	Packetsreceivedonthissession uint64

	//
	Packetssentonthissession uint64
}

// SetBytesreceivedonthissession sets the value of Bytesreceivedonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) SetPropertyBytesreceivedonthissession(value uint64) (err error) {
	return instance.SetProperty("Bytesreceivedonthissession", value)
}

// GetBytesreceivedonthissession gets the value of Bytesreceivedonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) GetPropertyBytesreceivedonthissession() (value uint64, err error) {
	retValue, err := instance.GetProperty("Bytesreceivedonthissession")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytessentonthissession sets the value of Bytessentonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) SetPropertyBytessentonthissession(value uint64) (err error) {
	return instance.SetProperty("Bytessentonthissession", value)
}

// GetBytessentonthissession gets the value of Bytessentonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) GetPropertyBytessentonthissession() (value uint64, err error) {
	retValue, err := instance.GetProperty("Bytessentonthissession")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDurationDurationofthesessionSeconds sets the value of DurationDurationofthesessionSeconds for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) SetPropertyDurationDurationofthesessionSeconds(value uint64) (err error) {
	return instance.SetProperty("DurationDurationofthesessionSeconds", value)
}

// GetDurationDurationofthesessionSeconds gets the value of DurationDurationofthesessionSeconds for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) GetPropertyDurationDurationofthesessionSeconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("DurationDurationofthesessionSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorsReceiveerrorsonthissession sets the value of ErrorsReceiveerrorsonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) SetPropertyErrorsReceiveerrorsonthissession(value uint64) (err error) {
	return instance.SetProperty("ErrorsReceiveerrorsonthissession", value)
}

// GetErrorsReceiveerrorsonthissession gets the value of ErrorsReceiveerrorsonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) GetPropertyErrorsReceiveerrorsonthissession() (value uint64, err error) {
	retValue, err := instance.GetProperty("ErrorsReceiveerrorsonthissession")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorsTransmiterrorsonthissession sets the value of ErrorsTransmiterrorsonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) SetPropertyErrorsTransmiterrorsonthissession(value uint64) (err error) {
	return instance.SetProperty("ErrorsTransmiterrorsonthissession", value)
}

// GetErrorsTransmiterrorsonthissession gets the value of ErrorsTransmiterrorsonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) GetPropertyErrorsTransmiterrorsonthissession() (value uint64, err error) {
	retValue, err := instance.GetProperty("ErrorsTransmiterrorsonthissession")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsreceivedonthissession sets the value of Packetsreceivedonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) SetPropertyPacketsreceivedonthissession(value uint64) (err error) {
	return instance.SetProperty("Packetsreceivedonthissession", value)
}

// GetPacketsreceivedonthissession gets the value of Packetsreceivedonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) GetPropertyPacketsreceivedonthissession() (value uint64, err error) {
	retValue, err := instance.GetProperty("Packetsreceivedonthissession")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketssentonthissession sets the value of Packetssentonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) SetPropertyPacketssentonthissession(value uint64) (err error) {
	return instance.SetProperty("Packetssentonthissession", value)
}

// GetPacketssentonthissession gets the value of Packetssentonthissession for the instance
func (instance *Win32_PerfFormattedData_Counters_IPHTTPSSession) GetPropertyPacketssentonthissession() (value uint64, err error) {
	retValue, err := instance.GetProperty("Packetssentonthissession")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetDnsTransitionMonitoring struct
type MSFT_NetDnsTransitionMonitoring struct {
	MSFT_NetSettingData

	//
	NumAAAAQueriesFailed uint32

	//
	NumAAAAQueriesIn6ArpaPtr uint32

	//
	NumAAAAQueriesSucceeded uint32

	//
	NumAAAAQueriesSynthesized uint32

	//
	NumOtherQueriesFailed uint32

	//
	NumOtherQueriesSucceeded uint32
}

// SetNumAAAAQueriesFailed sets the value of NumAAAAQueriesFailed for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) SetPropertyNumAAAAQueriesFailed(value uint32) (err error) {
	return instance.SetProperty("NumAAAAQueriesFailed", value)
}

// GetNumAAAAQueriesFailed gets the value of NumAAAAQueriesFailed for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) GetPropertyNumAAAAQueriesFailed() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumAAAAQueriesFailed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumAAAAQueriesIn6ArpaPtr sets the value of NumAAAAQueriesIn6ArpaPtr for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) SetPropertyNumAAAAQueriesIn6ArpaPtr(value uint32) (err error) {
	return instance.SetProperty("NumAAAAQueriesIn6ArpaPtr", value)
}

// GetNumAAAAQueriesIn6ArpaPtr gets the value of NumAAAAQueriesIn6ArpaPtr for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) GetPropertyNumAAAAQueriesIn6ArpaPtr() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumAAAAQueriesIn6ArpaPtr")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumAAAAQueriesSucceeded sets the value of NumAAAAQueriesSucceeded for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) SetPropertyNumAAAAQueriesSucceeded(value uint32) (err error) {
	return instance.SetProperty("NumAAAAQueriesSucceeded", value)
}

// GetNumAAAAQueriesSucceeded gets the value of NumAAAAQueriesSucceeded for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) GetPropertyNumAAAAQueriesSucceeded() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumAAAAQueriesSucceeded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumAAAAQueriesSynthesized sets the value of NumAAAAQueriesSynthesized for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) SetPropertyNumAAAAQueriesSynthesized(value uint32) (err error) {
	return instance.SetProperty("NumAAAAQueriesSynthesized", value)
}

// GetNumAAAAQueriesSynthesized gets the value of NumAAAAQueriesSynthesized for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) GetPropertyNumAAAAQueriesSynthesized() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumAAAAQueriesSynthesized")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumOtherQueriesFailed sets the value of NumOtherQueriesFailed for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) SetPropertyNumOtherQueriesFailed(value uint32) (err error) {
	return instance.SetProperty("NumOtherQueriesFailed", value)
}

// GetNumOtherQueriesFailed gets the value of NumOtherQueriesFailed for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) GetPropertyNumOtherQueriesFailed() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumOtherQueriesFailed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumOtherQueriesSucceeded sets the value of NumOtherQueriesSucceeded for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) SetPropertyNumOtherQueriesSucceeded(value uint32) (err error) {
	return instance.SetProperty("NumOtherQueriesSucceeded", value)
}

// GetNumOtherQueriesSucceeded gets the value of NumOtherQueriesSucceeded for the instance
func (instance *MSFT_NetDnsTransitionMonitoring) GetPropertyNumOtherQueriesSucceeded() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumOtherQueriesSucceeded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
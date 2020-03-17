// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator struct
type Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator struct {
	Win32_PerfFormattedData

	//
	AbortedTransactions uint32

	//
	AbortedTransactionsPersec uint32

	//
	ActiveTransactions uint32

	//
	ActiveTransactionsMaximum uint32

	//
	CommittedTransactions uint32

	//
	CommittedTransactionsPersec uint32

	//
	ForceAbortedTransactions uint32

	//
	ForceCommittedTransactions uint32

	//
	InDoubtTransactions uint32

	//
	ResponseTimeAverage uint32

	//
	ResponseTimeMaximum uint32

	//
	ResponseTimeMinimum uint32

	//
	TransactionsPersec uint32
}

// SetAbortedTransactions sets the value of AbortedTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyAbortedTransactions(value uint32) (err error) {
	return instance.SetProperty("AbortedTransactions", value)
}

// GetAbortedTransactions gets the value of AbortedTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyAbortedTransactions() (value uint32, err error) {
	retValue, err := instance.GetProperty("AbortedTransactions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAbortedTransactionsPersec sets the value of AbortedTransactionsPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyAbortedTransactionsPersec(value uint32) (err error) {
	return instance.SetProperty("AbortedTransactionsPersec", value)
}

// GetAbortedTransactionsPersec gets the value of AbortedTransactionsPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyAbortedTransactionsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("AbortedTransactionsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActiveTransactions sets the value of ActiveTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyActiveTransactions(value uint32) (err error) {
	return instance.SetProperty("ActiveTransactions", value)
}

// GetActiveTransactions gets the value of ActiveTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyActiveTransactions() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveTransactions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActiveTransactionsMaximum sets the value of ActiveTransactionsMaximum for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyActiveTransactionsMaximum(value uint32) (err error) {
	return instance.SetProperty("ActiveTransactionsMaximum", value)
}

// GetActiveTransactionsMaximum gets the value of ActiveTransactionsMaximum for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyActiveTransactionsMaximum() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveTransactionsMaximum")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCommittedTransactions sets the value of CommittedTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyCommittedTransactions(value uint32) (err error) {
	return instance.SetProperty("CommittedTransactions", value)
}

// GetCommittedTransactions gets the value of CommittedTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyCommittedTransactions() (value uint32, err error) {
	retValue, err := instance.GetProperty("CommittedTransactions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCommittedTransactionsPersec sets the value of CommittedTransactionsPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyCommittedTransactionsPersec(value uint32) (err error) {
	return instance.SetProperty("CommittedTransactionsPersec", value)
}

// GetCommittedTransactionsPersec gets the value of CommittedTransactionsPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyCommittedTransactionsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("CommittedTransactionsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetForceAbortedTransactions sets the value of ForceAbortedTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyForceAbortedTransactions(value uint32) (err error) {
	return instance.SetProperty("ForceAbortedTransactions", value)
}

// GetForceAbortedTransactions gets the value of ForceAbortedTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyForceAbortedTransactions() (value uint32, err error) {
	retValue, err := instance.GetProperty("ForceAbortedTransactions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetForceCommittedTransactions sets the value of ForceCommittedTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyForceCommittedTransactions(value uint32) (err error) {
	return instance.SetProperty("ForceCommittedTransactions", value)
}

// GetForceCommittedTransactions gets the value of ForceCommittedTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyForceCommittedTransactions() (value uint32, err error) {
	retValue, err := instance.GetProperty("ForceCommittedTransactions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInDoubtTransactions sets the value of InDoubtTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyInDoubtTransactions(value uint32) (err error) {
	return instance.SetProperty("InDoubtTransactions", value)
}

// GetInDoubtTransactions gets the value of InDoubtTransactions for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyInDoubtTransactions() (value uint32, err error) {
	retValue, err := instance.GetProperty("InDoubtTransactions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResponseTimeAverage sets the value of ResponseTimeAverage for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyResponseTimeAverage(value uint32) (err error) {
	return instance.SetProperty("ResponseTimeAverage", value)
}

// GetResponseTimeAverage gets the value of ResponseTimeAverage for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyResponseTimeAverage() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResponseTimeAverage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResponseTimeMaximum sets the value of ResponseTimeMaximum for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyResponseTimeMaximum(value uint32) (err error) {
	return instance.SetProperty("ResponseTimeMaximum", value)
}

// GetResponseTimeMaximum gets the value of ResponseTimeMaximum for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyResponseTimeMaximum() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResponseTimeMaximum")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResponseTimeMinimum sets the value of ResponseTimeMinimum for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyResponseTimeMinimum(value uint32) (err error) {
	return instance.SetProperty("ResponseTimeMinimum", value)
}

// GetResponseTimeMinimum gets the value of ResponseTimeMinimum for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyResponseTimeMinimum() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResponseTimeMinimum")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTransactionsPersec sets the value of TransactionsPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) SetPropertyTransactionsPersec(value uint32) (err error) {
	return instance.SetProperty("TransactionsPersec", value)
}

// GetTransactionsPersec gets the value of TransactionsPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTC_DistributedTransactionCoordinator) GetPropertyTransactionsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("TransactionsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
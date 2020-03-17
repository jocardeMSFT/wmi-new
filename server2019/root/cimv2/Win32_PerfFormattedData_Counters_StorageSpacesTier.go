// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_StorageSpacesTier struct
type Win32_PerfFormattedData_Counters_StorageSpacesTier struct {
	Win32_PerfFormattedData

	//
	TierReadBytesAverage uint64

	//
	TierReadBytesPersec uint64

	//
	TierReadLatency uint32

	//
	TierReadsAverage uint64

	//
	TierReadsPersec uint64

	//
	TierTransferBytesAverage uint64

	//
	TierTransferBytesPersec uint64

	//
	TierTransferLatency uint32

	//
	TierTransfersAverage uint64

	//
	TierTransfersCurrent uint32

	//
	TierTransfersPersec uint64

	//
	TierWriteBytesAverage uint64

	//
	TierWriteBytesPersec uint64

	//
	TierWriteLatency uint32

	//
	TierWritesAverage uint64

	//
	TierWritesPersec uint64
}

// SetTierReadBytesAverage sets the value of TierReadBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierReadBytesAverage(value uint64) (err error) {
	return instance.SetProperty("TierReadBytesAverage", value)
}

// GetTierReadBytesAverage gets the value of TierReadBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierReadBytesAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierReadBytesAverage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierReadBytesPersec sets the value of TierReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("TierReadBytesPersec", value)
}

// GetTierReadBytesPersec gets the value of TierReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierReadBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierReadLatency sets the value of TierReadLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierReadLatency(value uint32) (err error) {
	return instance.SetProperty("TierReadLatency", value)
}

// GetTierReadLatency gets the value of TierReadLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierReadLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierReadLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierReadsAverage sets the value of TierReadsAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierReadsAverage(value uint64) (err error) {
	return instance.SetProperty("TierReadsAverage", value)
}

// GetTierReadsAverage gets the value of TierReadsAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierReadsAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierReadsAverage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierReadsPersec sets the value of TierReadsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierReadsPersec(value uint64) (err error) {
	return instance.SetProperty("TierReadsPersec", value)
}

// GetTierReadsPersec gets the value of TierReadsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierReadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierReadsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierTransferBytesAverage sets the value of TierTransferBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierTransferBytesAverage(value uint64) (err error) {
	return instance.SetProperty("TierTransferBytesAverage", value)
}

// GetTierTransferBytesAverage gets the value of TierTransferBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierTransferBytesAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierTransferBytesAverage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierTransferBytesPersec sets the value of TierTransferBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierTransferBytesPersec(value uint64) (err error) {
	return instance.SetProperty("TierTransferBytesPersec", value)
}

// GetTierTransferBytesPersec gets the value of TierTransferBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierTransferBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierTransferBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierTransferLatency sets the value of TierTransferLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierTransferLatency(value uint32) (err error) {
	return instance.SetProperty("TierTransferLatency", value)
}

// GetTierTransferLatency gets the value of TierTransferLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierTransferLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierTransferLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierTransfersAverage sets the value of TierTransfersAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierTransfersAverage(value uint64) (err error) {
	return instance.SetProperty("TierTransfersAverage", value)
}

// GetTierTransfersAverage gets the value of TierTransfersAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierTransfersAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierTransfersAverage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierTransfersCurrent sets the value of TierTransfersCurrent for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierTransfersCurrent(value uint32) (err error) {
	return instance.SetProperty("TierTransfersCurrent", value)
}

// GetTierTransfersCurrent gets the value of TierTransfersCurrent for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierTransfersCurrent() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierTransfersCurrent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierTransfersPersec sets the value of TierTransfersPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierTransfersPersec(value uint64) (err error) {
	return instance.SetProperty("TierTransfersPersec", value)
}

// GetTierTransfersPersec gets the value of TierTransfersPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierTransfersPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierTransfersPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierWriteBytesAverage sets the value of TierWriteBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierWriteBytesAverage(value uint64) (err error) {
	return instance.SetProperty("TierWriteBytesAverage", value)
}

// GetTierWriteBytesAverage gets the value of TierWriteBytesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierWriteBytesAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierWriteBytesAverage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierWriteBytesPersec sets the value of TierWriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("TierWriteBytesPersec", value)
}

// GetTierWriteBytesPersec gets the value of TierWriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierWriteBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierWriteLatency sets the value of TierWriteLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierWriteLatency(value uint32) (err error) {
	return instance.SetProperty("TierWriteLatency", value)
}

// GetTierWriteLatency gets the value of TierWriteLatency for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierWriteLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("TierWriteLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierWritesAverage sets the value of TierWritesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierWritesAverage(value uint64) (err error) {
	return instance.SetProperty("TierWritesAverage", value)
}

// GetTierWritesAverage gets the value of TierWritesAverage for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierWritesAverage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierWritesAverage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTierWritesPersec sets the value of TierWritesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) SetPropertyTierWritesPersec(value uint64) (err error) {
	return instance.SetProperty("TierWritesPersec", value)
}

// GetTierWritesPersec gets the value of TierWritesPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_StorageSpacesTier) GetPropertyTierWritesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TierWritesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
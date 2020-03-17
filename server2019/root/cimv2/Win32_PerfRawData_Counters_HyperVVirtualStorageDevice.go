// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_HyperVVirtualStorageDevice struct
type Win32_PerfRawData_Counters_HyperVVirtualStorageDevice struct {
	Win32_PerfRawData

	//
	AdapterOpenChannelCount uint32

	//
	ByteQuotaReplenishmentRate uint64

	//
	ErrorCount uint32

	//
	FlushCount uint32

	//
	IoQuotaReplenishmentRate uint64

	//
	Latency uint32

	//
	Latency_Base uint32

	//
	LowerLatency uint32

	//
	LowerLatency_Base uint32

	//
	LowerQueueLength uint64

	//
	MaximumAdapterWorkerCount uint32

	//
	MaximumBandwidth uint64

	//
	MaximumIORate uint64

	//
	MinimumIORate uint64

	//
	NormalizedThroughput uint64

	//
	QueueLength uint64

	//
	ReadBytesPersec uint64

	//
	ReadCount uint32

	//
	ReadOperationsPerSec uint32

	//
	Throughput uint32

	//
	WriteBytesPersec uint64

	//
	WriteCount uint32

	//
	WriteOperationsPerSec uint32
}

// SetAdapterOpenChannelCount sets the value of AdapterOpenChannelCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyAdapterOpenChannelCount(value uint32) (err error) {
	return instance.SetProperty("AdapterOpenChannelCount", value)
}

// GetAdapterOpenChannelCount gets the value of AdapterOpenChannelCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyAdapterOpenChannelCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("AdapterOpenChannelCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetByteQuotaReplenishmentRate sets the value of ByteQuotaReplenishmentRate for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyByteQuotaReplenishmentRate(value uint64) (err error) {
	return instance.SetProperty("ByteQuotaReplenishmentRate", value)
}

// GetByteQuotaReplenishmentRate gets the value of ByteQuotaReplenishmentRate for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyByteQuotaReplenishmentRate() (value uint64, err error) {
	retValue, err := instance.GetProperty("ByteQuotaReplenishmentRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorCount sets the value of ErrorCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyErrorCount(value uint32) (err error) {
	return instance.SetProperty("ErrorCount", value)
}

// GetErrorCount gets the value of ErrorCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyErrorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlushCount sets the value of FlushCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyFlushCount(value uint32) (err error) {
	return instance.SetProperty("FlushCount", value)
}

// GetFlushCount gets the value of FlushCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyFlushCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("FlushCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIoQuotaReplenishmentRate sets the value of IoQuotaReplenishmentRate for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyIoQuotaReplenishmentRate(value uint64) (err error) {
	return instance.SetProperty("IoQuotaReplenishmentRate", value)
}

// GetIoQuotaReplenishmentRate gets the value of IoQuotaReplenishmentRate for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyIoQuotaReplenishmentRate() (value uint64, err error) {
	retValue, err := instance.GetProperty("IoQuotaReplenishmentRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLatency sets the value of Latency for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyLatency(value uint32) (err error) {
	return instance.SetProperty("Latency", value)
}

// GetLatency gets the value of Latency for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("Latency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLatency_Base sets the value of Latency_Base for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyLatency_Base(value uint32) (err error) {
	return instance.SetProperty("Latency_Base", value)
}

// GetLatency_Base gets the value of Latency_Base for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyLatency_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("Latency_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLowerLatency sets the value of LowerLatency for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyLowerLatency(value uint32) (err error) {
	return instance.SetProperty("LowerLatency", value)
}

// GetLowerLatency gets the value of LowerLatency for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyLowerLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("LowerLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLowerLatency_Base sets the value of LowerLatency_Base for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyLowerLatency_Base(value uint32) (err error) {
	return instance.SetProperty("LowerLatency_Base", value)
}

// GetLowerLatency_Base gets the value of LowerLatency_Base for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyLowerLatency_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("LowerLatency_Base")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLowerQueueLength sets the value of LowerQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyLowerQueueLength(value uint64) (err error) {
	return instance.SetProperty("LowerQueueLength", value)
}

// GetLowerQueueLength gets the value of LowerQueueLength for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyLowerQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("LowerQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumAdapterWorkerCount sets the value of MaximumAdapterWorkerCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyMaximumAdapterWorkerCount(value uint32) (err error) {
	return instance.SetProperty("MaximumAdapterWorkerCount", value)
}

// GetMaximumAdapterWorkerCount gets the value of MaximumAdapterWorkerCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyMaximumAdapterWorkerCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumAdapterWorkerCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumBandwidth sets the value of MaximumBandwidth for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyMaximumBandwidth(value uint64) (err error) {
	return instance.SetProperty("MaximumBandwidth", value)
}

// GetMaximumBandwidth gets the value of MaximumBandwidth for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyMaximumBandwidth() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaximumBandwidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumIORate sets the value of MaximumIORate for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyMaximumIORate(value uint64) (err error) {
	return instance.SetProperty("MaximumIORate", value)
}

// GetMaximumIORate gets the value of MaximumIORate for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyMaximumIORate() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaximumIORate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinimumIORate sets the value of MinimumIORate for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyMinimumIORate(value uint64) (err error) {
	return instance.SetProperty("MinimumIORate", value)
}

// GetMinimumIORate gets the value of MinimumIORate for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyMinimumIORate() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinimumIORate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNormalizedThroughput sets the value of NormalizedThroughput for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyNormalizedThroughput(value uint64) (err error) {
	return instance.SetProperty("NormalizedThroughput", value)
}

// GetNormalizedThroughput gets the value of NormalizedThroughput for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyNormalizedThroughput() (value uint64, err error) {
	retValue, err := instance.GetProperty("NormalizedThroughput")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQueueLength sets the value of QueueLength for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyQueueLength(value uint64) (err error) {
	return instance.SetProperty("QueueLength", value)
}

// GetQueueLength gets the value of QueueLength for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyQueueLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("QueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadBytesPersec sets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReadBytesPersec", value)
}

// GetReadBytesPersec gets the value of ReadBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadCount sets the value of ReadCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyReadCount(value uint32) (err error) {
	return instance.SetProperty("ReadCount", value)
}

// GetReadCount gets the value of ReadCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyReadCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReadCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadOperationsPerSec sets the value of ReadOperationsPerSec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyReadOperationsPerSec(value uint32) (err error) {
	return instance.SetProperty("ReadOperationsPerSec", value)
}

// GetReadOperationsPerSec gets the value of ReadOperationsPerSec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyReadOperationsPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReadOperationsPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThroughput sets the value of Throughput for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyThroughput(value uint32) (err error) {
	return instance.SetProperty("Throughput", value)
}

// GetThroughput gets the value of Throughput for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyThroughput() (value uint32, err error) {
	retValue, err := instance.GetProperty("Throughput")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteBytesPersec sets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("WriteBytesPersec", value)
}

// GetWriteBytesPersec gets the value of WriteBytesPersec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteCount sets the value of WriteCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyWriteCount(value uint32) (err error) {
	return instance.SetProperty("WriteCount", value)
}

// GetWriteCount gets the value of WriteCount for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyWriteCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("WriteCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteOperationsPerSec sets the value of WriteOperationsPerSec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) SetPropertyWriteOperationsPerSec(value uint32) (err error) {
	return instance.SetProperty("WriteOperationsPerSec", value)
}

// GetWriteOperationsPerSec gets the value of WriteOperationsPerSec for the instance
func (instance *Win32_PerfRawData_Counters_HyperVVirtualStorageDevice) GetPropertyWriteOperationsPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WriteOperationsPerSec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
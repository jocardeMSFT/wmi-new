// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager struct
type Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager struct {
	Win32_PerfFormattedData

	//
	DirectIOFailureRedirection uint64

	//
	DirectIOFailureRedirectionPersec uint64

	//
	IOReadBytes uint64

	//
	IOReadBytesPersec uint64

	//
	IOReadBytesPersecRedirected uint64

	//
	IOReadBytesRedirected uint64

	//
	IOReadPersecRedirected uint64

	//
	IOReads uint64

	//
	IOReadsPersec uint64

	//
	IOReadsRedirected uint64

	//
	IOWriteBytes uint64

	//
	IOWriteBytesPersec uint64

	//
	IOWriteBytesPersecRedirected uint64

	//
	IOWriteBytesRedirected uint64

	//
	IOWrites uint64

	//
	IOWritesPersec uint64

	//
	IOWritesPersecRedirected uint64

	//
	IOWritesRedirected uint64
}

// SetDirectIOFailureRedirection sets the value of DirectIOFailureRedirection for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyDirectIOFailureRedirection(value uint64) (err error) {
	return instance.SetProperty("DirectIOFailureRedirection", value)
}

// GetDirectIOFailureRedirection gets the value of DirectIOFailureRedirection for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyDirectIOFailureRedirection() (value uint64, err error) {
	retValue, err := instance.GetProperty("DirectIOFailureRedirection")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirectIOFailureRedirectionPersec sets the value of DirectIOFailureRedirectionPersec for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyDirectIOFailureRedirectionPersec(value uint64) (err error) {
	return instance.SetProperty("DirectIOFailureRedirectionPersec", value)
}

// GetDirectIOFailureRedirectionPersec gets the value of DirectIOFailureRedirectionPersec for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyDirectIOFailureRedirectionPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DirectIOFailureRedirectionPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadBytes sets the value of IOReadBytes for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOReadBytes(value uint64) (err error) {
	return instance.SetProperty("IOReadBytes", value)
}

// GetIOReadBytes gets the value of IOReadBytes for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOReadBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadBytesPersec sets the value of IOReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOReadBytesPersec(value uint64) (err error) {
	return instance.SetProperty("IOReadBytesPersec", value)
}

// GetIOReadBytesPersec gets the value of IOReadBytesPersec for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOReadBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadBytesPersecRedirected sets the value of IOReadBytesPersecRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOReadBytesPersecRedirected(value uint64) (err error) {
	return instance.SetProperty("IOReadBytesPersecRedirected", value)
}

// GetIOReadBytesPersecRedirected gets the value of IOReadBytesPersecRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOReadBytesPersecRedirected() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadBytesPersecRedirected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadBytesRedirected sets the value of IOReadBytesRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOReadBytesRedirected(value uint64) (err error) {
	return instance.SetProperty("IOReadBytesRedirected", value)
}

// GetIOReadBytesRedirected gets the value of IOReadBytesRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOReadBytesRedirected() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadBytesRedirected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadPersecRedirected sets the value of IOReadPersecRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOReadPersecRedirected(value uint64) (err error) {
	return instance.SetProperty("IOReadPersecRedirected", value)
}

// GetIOReadPersecRedirected gets the value of IOReadPersecRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOReadPersecRedirected() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadPersecRedirected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReads sets the value of IOReads for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOReads(value uint64) (err error) {
	return instance.SetProperty("IOReads", value)
}

// GetIOReads gets the value of IOReads for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOReads() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadsPersec sets the value of IOReadsPersec for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOReadsPersec(value uint64) (err error) {
	return instance.SetProperty("IOReadsPersec", value)
}

// GetIOReadsPersec gets the value of IOReadsPersec for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOReadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOReadsRedirected sets the value of IOReadsRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOReadsRedirected(value uint64) (err error) {
	return instance.SetProperty("IOReadsRedirected", value)
}

// GetIOReadsRedirected gets the value of IOReadsRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOReadsRedirected() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOReadsRedirected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWriteBytes sets the value of IOWriteBytes for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOWriteBytes(value uint64) (err error) {
	return instance.SetProperty("IOWriteBytes", value)
}

// GetIOWriteBytes gets the value of IOWriteBytes for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOWriteBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWriteBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWriteBytesPersec sets the value of IOWriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOWriteBytesPersec(value uint64) (err error) {
	return instance.SetProperty("IOWriteBytesPersec", value)
}

// GetIOWriteBytesPersec gets the value of IOWriteBytesPersec for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOWriteBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWriteBytesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWriteBytesPersecRedirected sets the value of IOWriteBytesPersecRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOWriteBytesPersecRedirected(value uint64) (err error) {
	return instance.SetProperty("IOWriteBytesPersecRedirected", value)
}

// GetIOWriteBytesPersecRedirected gets the value of IOWriteBytesPersecRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOWriteBytesPersecRedirected() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWriteBytesPersecRedirected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWriteBytesRedirected sets the value of IOWriteBytesRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOWriteBytesRedirected(value uint64) (err error) {
	return instance.SetProperty("IOWriteBytesRedirected", value)
}

// GetIOWriteBytesRedirected gets the value of IOWriteBytesRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOWriteBytesRedirected() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWriteBytesRedirected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWrites sets the value of IOWrites for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOWrites(value uint64) (err error) {
	return instance.SetProperty("IOWrites", value)
}

// GetIOWrites gets the value of IOWrites for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOWrites() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWrites")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWritesPersec sets the value of IOWritesPersec for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOWritesPersec(value uint64) (err error) {
	return instance.SetProperty("IOWritesPersec", value)
}

// GetIOWritesPersec gets the value of IOWritesPersec for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOWritesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWritesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWritesPersecRedirected sets the value of IOWritesPersecRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOWritesPersecRedirected(value uint64) (err error) {
	return instance.SetProperty("IOWritesPersecRedirected", value)
}

// GetIOWritesPersecRedirected gets the value of IOWritesPersecRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOWritesPersecRedirected() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWritesPersecRedirected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIOWritesRedirected sets the value of IOWritesRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) SetPropertyIOWritesRedirected(value uint64) (err error) {
	return instance.SetProperty("IOWritesRedirected", value)
}

// GetIOWritesRedirected gets the value of IOWritesRedirected for the instance
func (instance *Win32_PerfFormattedData_CsvVolumePerfProvider_ClusterCSVVolumeManager) GetPropertyIOWritesRedirected() (value uint64, err error) {
	retValue, err := instance.GetProperty("IOWritesRedirected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
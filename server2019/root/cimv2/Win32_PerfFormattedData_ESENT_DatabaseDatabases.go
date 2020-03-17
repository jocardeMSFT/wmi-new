// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_ESENT_DatabaseDatabases struct
type Win32_PerfFormattedData_ESENT_DatabaseDatabases struct {
	Win32_PerfFormattedData

	//
	DatabaseCachePercentHitUnique uint32

	//
	DatabaseCacheRequestsPersecUnique uint32

	//
	DatabaseCacheSizeMB uint64

	//
	IODatabaseReadsAverageLatency uint64

	//
	IODatabaseReadsPersec uint32

	//
	IODatabaseWritesAverageLatency uint64

	//
	IODatabaseWritesPersec uint32
}

// SetDatabaseCachePercentHitUnique sets the value of DatabaseCachePercentHitUnique for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) SetPropertyDatabaseCachePercentHitUnique(value uint32) (err error) {
	return instance.SetProperty("DatabaseCachePercentHitUnique", value)
}

// GetDatabaseCachePercentHitUnique gets the value of DatabaseCachePercentHitUnique for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) GetPropertyDatabaseCachePercentHitUnique() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatabaseCachePercentHitUnique")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatabaseCacheRequestsPersecUnique sets the value of DatabaseCacheRequestsPersecUnique for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) SetPropertyDatabaseCacheRequestsPersecUnique(value uint32) (err error) {
	return instance.SetProperty("DatabaseCacheRequestsPersecUnique", value)
}

// GetDatabaseCacheRequestsPersecUnique gets the value of DatabaseCacheRequestsPersecUnique for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) GetPropertyDatabaseCacheRequestsPersecUnique() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatabaseCacheRequestsPersecUnique")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatabaseCacheSizeMB sets the value of DatabaseCacheSizeMB for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) SetPropertyDatabaseCacheSizeMB(value uint64) (err error) {
	return instance.SetProperty("DatabaseCacheSizeMB", value)
}

// GetDatabaseCacheSizeMB gets the value of DatabaseCacheSizeMB for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) GetPropertyDatabaseCacheSizeMB() (value uint64, err error) {
	retValue, err := instance.GetProperty("DatabaseCacheSizeMB")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIODatabaseReadsAverageLatency sets the value of IODatabaseReadsAverageLatency for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) SetPropertyIODatabaseReadsAverageLatency(value uint64) (err error) {
	return instance.SetProperty("IODatabaseReadsAverageLatency", value)
}

// GetIODatabaseReadsAverageLatency gets the value of IODatabaseReadsAverageLatency for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) GetPropertyIODatabaseReadsAverageLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("IODatabaseReadsAverageLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIODatabaseReadsPersec sets the value of IODatabaseReadsPersec for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) SetPropertyIODatabaseReadsPersec(value uint32) (err error) {
	return instance.SetProperty("IODatabaseReadsPersec", value)
}

// GetIODatabaseReadsPersec gets the value of IODatabaseReadsPersec for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) GetPropertyIODatabaseReadsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IODatabaseReadsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIODatabaseWritesAverageLatency sets the value of IODatabaseWritesAverageLatency for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) SetPropertyIODatabaseWritesAverageLatency(value uint64) (err error) {
	return instance.SetProperty("IODatabaseWritesAverageLatency", value)
}

// GetIODatabaseWritesAverageLatency gets the value of IODatabaseWritesAverageLatency for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) GetPropertyIODatabaseWritesAverageLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("IODatabaseWritesAverageLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIODatabaseWritesPersec sets the value of IODatabaseWritesPersec for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) SetPropertyIODatabaseWritesPersec(value uint32) (err error) {
	return instance.SetProperty("IODatabaseWritesPersec", value)
}

// GetIODatabaseWritesPersec gets the value of IODatabaseWritesPersec for the instance
func (instance *Win32_PerfFormattedData_ESENT_DatabaseDatabases) GetPropertyIODatabaseWritesPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("IODatabaseWritesPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
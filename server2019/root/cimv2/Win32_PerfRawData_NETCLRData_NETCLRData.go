// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_NETCLRData_NETCLRData struct
type Win32_PerfRawData_NETCLRData_NETCLRData struct {
	Win32_PerfRawData

	//
	SqlClientCurrentNumberconnectionpools uint32

	//
	SqlClientCurrentNumberpooledandnonpooledconnections uint32

	//
	SqlClientCurrentNumberpooledconnections uint32

	//
	SqlClientPeakNumberpooledconnections uint32

	//
	SqlClientTotalNumberfailedcommands uint32

	//
	SqlClientTotalNumberfailedconnects uint32
}

// SetSqlClientCurrentNumberconnectionpools sets the value of SqlClientCurrentNumberconnectionpools for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) SetPropertySqlClientCurrentNumberconnectionpools(value uint32) (err error) {
	return instance.SetProperty("SqlClientCurrentNumberconnectionpools", value)
}

// GetSqlClientCurrentNumberconnectionpools gets the value of SqlClientCurrentNumberconnectionpools for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) GetPropertySqlClientCurrentNumberconnectionpools() (value uint32, err error) {
	retValue, err := instance.GetProperty("SqlClientCurrentNumberconnectionpools")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSqlClientCurrentNumberpooledandnonpooledconnections sets the value of SqlClientCurrentNumberpooledandnonpooledconnections for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) SetPropertySqlClientCurrentNumberpooledandnonpooledconnections(value uint32) (err error) {
	return instance.SetProperty("SqlClientCurrentNumberpooledandnonpooledconnections", value)
}

// GetSqlClientCurrentNumberpooledandnonpooledconnections gets the value of SqlClientCurrentNumberpooledandnonpooledconnections for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) GetPropertySqlClientCurrentNumberpooledandnonpooledconnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("SqlClientCurrentNumberpooledandnonpooledconnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSqlClientCurrentNumberpooledconnections sets the value of SqlClientCurrentNumberpooledconnections for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) SetPropertySqlClientCurrentNumberpooledconnections(value uint32) (err error) {
	return instance.SetProperty("SqlClientCurrentNumberpooledconnections", value)
}

// GetSqlClientCurrentNumberpooledconnections gets the value of SqlClientCurrentNumberpooledconnections for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) GetPropertySqlClientCurrentNumberpooledconnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("SqlClientCurrentNumberpooledconnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSqlClientPeakNumberpooledconnections sets the value of SqlClientPeakNumberpooledconnections for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) SetPropertySqlClientPeakNumberpooledconnections(value uint32) (err error) {
	return instance.SetProperty("SqlClientPeakNumberpooledconnections", value)
}

// GetSqlClientPeakNumberpooledconnections gets the value of SqlClientPeakNumberpooledconnections for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) GetPropertySqlClientPeakNumberpooledconnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("SqlClientPeakNumberpooledconnections")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSqlClientTotalNumberfailedcommands sets the value of SqlClientTotalNumberfailedcommands for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) SetPropertySqlClientTotalNumberfailedcommands(value uint32) (err error) {
	return instance.SetProperty("SqlClientTotalNumberfailedcommands", value)
}

// GetSqlClientTotalNumberfailedcommands gets the value of SqlClientTotalNumberfailedcommands for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) GetPropertySqlClientTotalNumberfailedcommands() (value uint32, err error) {
	retValue, err := instance.GetProperty("SqlClientTotalNumberfailedcommands")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSqlClientTotalNumberfailedconnects sets the value of SqlClientTotalNumberfailedconnects for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) SetPropertySqlClientTotalNumberfailedconnects(value uint32) (err error) {
	return instance.SetProperty("SqlClientTotalNumberfailedconnects", value)
}

// GetSqlClientTotalNumberfailedconnects gets the value of SqlClientTotalNumberfailedconnects for the instance
func (instance *Win32_PerfRawData_NETCLRData_NETCLRData) GetPropertySqlClientTotalNumberfailedconnects() (value uint32, err error) {
	retValue, err := instance.GetProperty("SqlClientTotalNumberfailedconnects")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
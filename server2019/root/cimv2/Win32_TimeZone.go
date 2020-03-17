// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_TimeZone struct
type Win32_TimeZone struct {
	CIM_Setting

	//
	Bias int32

	//
	DaylightBias int32

	//
	DaylightDay uint32

	//
	DaylightDayOfWeek uint8

	//
	DaylightHour uint32

	//
	DaylightMillisecond uint32

	//
	DaylightMinute uint32

	//
	DaylightMonth uint32

	//
	DaylightName string

	//
	DaylightSecond uint32

	//
	DaylightYear uint32

	//
	StandardBias uint32

	//
	StandardDay uint32

	//
	StandardDayOfWeek uint8

	//
	StandardHour uint32

	//
	StandardMillisecond uint32

	//
	StandardMinute uint32

	//
	StandardMonth uint32

	//
	StandardName string

	//
	StandardSecond uint32

	//
	StandardYear uint32
}

// SetBias sets the value of Bias for the instance
func (instance *Win32_TimeZone) SetPropertyBias(value int32) (err error) {
	return instance.SetProperty("Bias", value)
}

// GetBias gets the value of Bias for the instance
func (instance *Win32_TimeZone) GetPropertyBias() (value int32, err error) {
	retValue, err := instance.GetProperty("Bias")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaylightBias sets the value of DaylightBias for the instance
func (instance *Win32_TimeZone) SetPropertyDaylightBias(value int32) (err error) {
	return instance.SetProperty("DaylightBias", value)
}

// GetDaylightBias gets the value of DaylightBias for the instance
func (instance *Win32_TimeZone) GetPropertyDaylightBias() (value int32, err error) {
	retValue, err := instance.GetProperty("DaylightBias")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaylightDay sets the value of DaylightDay for the instance
func (instance *Win32_TimeZone) SetPropertyDaylightDay(value uint32) (err error) {
	return instance.SetProperty("DaylightDay", value)
}

// GetDaylightDay gets the value of DaylightDay for the instance
func (instance *Win32_TimeZone) GetPropertyDaylightDay() (value uint32, err error) {
	retValue, err := instance.GetProperty("DaylightDay")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaylightDayOfWeek sets the value of DaylightDayOfWeek for the instance
func (instance *Win32_TimeZone) SetPropertyDaylightDayOfWeek(value uint8) (err error) {
	return instance.SetProperty("DaylightDayOfWeek", value)
}

// GetDaylightDayOfWeek gets the value of DaylightDayOfWeek for the instance
func (instance *Win32_TimeZone) GetPropertyDaylightDayOfWeek() (value uint8, err error) {
	retValue, err := instance.GetProperty("DaylightDayOfWeek")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaylightHour sets the value of DaylightHour for the instance
func (instance *Win32_TimeZone) SetPropertyDaylightHour(value uint32) (err error) {
	return instance.SetProperty("DaylightHour", value)
}

// GetDaylightHour gets the value of DaylightHour for the instance
func (instance *Win32_TimeZone) GetPropertyDaylightHour() (value uint32, err error) {
	retValue, err := instance.GetProperty("DaylightHour")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaylightMillisecond sets the value of DaylightMillisecond for the instance
func (instance *Win32_TimeZone) SetPropertyDaylightMillisecond(value uint32) (err error) {
	return instance.SetProperty("DaylightMillisecond", value)
}

// GetDaylightMillisecond gets the value of DaylightMillisecond for the instance
func (instance *Win32_TimeZone) GetPropertyDaylightMillisecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("DaylightMillisecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaylightMinute sets the value of DaylightMinute for the instance
func (instance *Win32_TimeZone) SetPropertyDaylightMinute(value uint32) (err error) {
	return instance.SetProperty("DaylightMinute", value)
}

// GetDaylightMinute gets the value of DaylightMinute for the instance
func (instance *Win32_TimeZone) GetPropertyDaylightMinute() (value uint32, err error) {
	retValue, err := instance.GetProperty("DaylightMinute")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaylightMonth sets the value of DaylightMonth for the instance
func (instance *Win32_TimeZone) SetPropertyDaylightMonth(value uint32) (err error) {
	return instance.SetProperty("DaylightMonth", value)
}

// GetDaylightMonth gets the value of DaylightMonth for the instance
func (instance *Win32_TimeZone) GetPropertyDaylightMonth() (value uint32, err error) {
	retValue, err := instance.GetProperty("DaylightMonth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaylightName sets the value of DaylightName for the instance
func (instance *Win32_TimeZone) SetPropertyDaylightName(value string) (err error) {
	return instance.SetProperty("DaylightName", value)
}

// GetDaylightName gets the value of DaylightName for the instance
func (instance *Win32_TimeZone) GetPropertyDaylightName() (value string, err error) {
	retValue, err := instance.GetProperty("DaylightName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaylightSecond sets the value of DaylightSecond for the instance
func (instance *Win32_TimeZone) SetPropertyDaylightSecond(value uint32) (err error) {
	return instance.SetProperty("DaylightSecond", value)
}

// GetDaylightSecond gets the value of DaylightSecond for the instance
func (instance *Win32_TimeZone) GetPropertyDaylightSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("DaylightSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDaylightYear sets the value of DaylightYear for the instance
func (instance *Win32_TimeZone) SetPropertyDaylightYear(value uint32) (err error) {
	return instance.SetProperty("DaylightYear", value)
}

// GetDaylightYear gets the value of DaylightYear for the instance
func (instance *Win32_TimeZone) GetPropertyDaylightYear() (value uint32, err error) {
	retValue, err := instance.GetProperty("DaylightYear")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardBias sets the value of StandardBias for the instance
func (instance *Win32_TimeZone) SetPropertyStandardBias(value uint32) (err error) {
	return instance.SetProperty("StandardBias", value)
}

// GetStandardBias gets the value of StandardBias for the instance
func (instance *Win32_TimeZone) GetPropertyStandardBias() (value uint32, err error) {
	retValue, err := instance.GetProperty("StandardBias")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardDay sets the value of StandardDay for the instance
func (instance *Win32_TimeZone) SetPropertyStandardDay(value uint32) (err error) {
	return instance.SetProperty("StandardDay", value)
}

// GetStandardDay gets the value of StandardDay for the instance
func (instance *Win32_TimeZone) GetPropertyStandardDay() (value uint32, err error) {
	retValue, err := instance.GetProperty("StandardDay")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardDayOfWeek sets the value of StandardDayOfWeek for the instance
func (instance *Win32_TimeZone) SetPropertyStandardDayOfWeek(value uint8) (err error) {
	return instance.SetProperty("StandardDayOfWeek", value)
}

// GetStandardDayOfWeek gets the value of StandardDayOfWeek for the instance
func (instance *Win32_TimeZone) GetPropertyStandardDayOfWeek() (value uint8, err error) {
	retValue, err := instance.GetProperty("StandardDayOfWeek")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardHour sets the value of StandardHour for the instance
func (instance *Win32_TimeZone) SetPropertyStandardHour(value uint32) (err error) {
	return instance.SetProperty("StandardHour", value)
}

// GetStandardHour gets the value of StandardHour for the instance
func (instance *Win32_TimeZone) GetPropertyStandardHour() (value uint32, err error) {
	retValue, err := instance.GetProperty("StandardHour")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardMillisecond sets the value of StandardMillisecond for the instance
func (instance *Win32_TimeZone) SetPropertyStandardMillisecond(value uint32) (err error) {
	return instance.SetProperty("StandardMillisecond", value)
}

// GetStandardMillisecond gets the value of StandardMillisecond for the instance
func (instance *Win32_TimeZone) GetPropertyStandardMillisecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("StandardMillisecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardMinute sets the value of StandardMinute for the instance
func (instance *Win32_TimeZone) SetPropertyStandardMinute(value uint32) (err error) {
	return instance.SetProperty("StandardMinute", value)
}

// GetStandardMinute gets the value of StandardMinute for the instance
func (instance *Win32_TimeZone) GetPropertyStandardMinute() (value uint32, err error) {
	retValue, err := instance.GetProperty("StandardMinute")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardMonth sets the value of StandardMonth for the instance
func (instance *Win32_TimeZone) SetPropertyStandardMonth(value uint32) (err error) {
	return instance.SetProperty("StandardMonth", value)
}

// GetStandardMonth gets the value of StandardMonth for the instance
func (instance *Win32_TimeZone) GetPropertyStandardMonth() (value uint32, err error) {
	retValue, err := instance.GetProperty("StandardMonth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardName sets the value of StandardName for the instance
func (instance *Win32_TimeZone) SetPropertyStandardName(value string) (err error) {
	return instance.SetProperty("StandardName", value)
}

// GetStandardName gets the value of StandardName for the instance
func (instance *Win32_TimeZone) GetPropertyStandardName() (value string, err error) {
	retValue, err := instance.GetProperty("StandardName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardSecond sets the value of StandardSecond for the instance
func (instance *Win32_TimeZone) SetPropertyStandardSecond(value uint32) (err error) {
	return instance.SetProperty("StandardSecond", value)
}

// GetStandardSecond gets the value of StandardSecond for the instance
func (instance *Win32_TimeZone) GetPropertyStandardSecond() (value uint32, err error) {
	retValue, err := instance.GetProperty("StandardSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStandardYear sets the value of StandardYear for the instance
func (instance *Win32_TimeZone) SetPropertyStandardYear(value uint32) (err error) {
	return instance.SetProperty("StandardYear", value)
}

// GetStandardYear gets the value of StandardYear for the instance
func (instance *Win32_TimeZone) GetPropertyStandardYear() (value uint32, err error) {
	retValue, err := instance.GetProperty("StandardYear")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
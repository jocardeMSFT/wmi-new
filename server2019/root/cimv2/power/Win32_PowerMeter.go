// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

// Win32_PowerMeter struct
type Win32_PowerMeter struct {
	CIM_NumericSensor

	//
	AveragingInterval uint32

	//
	BudgetEnabled bool

	//
	BudgetWriteable bool

	//
	ConfiguredBudget uint32

	//
	MaximumAveragingInterval uint32

	//
	MaxOperatingBudget uint32

	//
	MeterType uint32

	//
	MinimumAveragingInterval uint32

	//
	MinOperatingBudget uint32

	//
	SamplingPeriod uint32

	//
	SupportCapabilities uint32
}

// SetAveragingInterval sets the value of AveragingInterval for the instance
func (instance *Win32_PowerMeter) SetPropertyAveragingInterval(value uint32) (err error) {
	return instance.SetProperty("AveragingInterval", value)
}

// GetAveragingInterval gets the value of AveragingInterval for the instance
func (instance *Win32_PowerMeter) GetPropertyAveragingInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("AveragingInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBudgetEnabled sets the value of BudgetEnabled for the instance
func (instance *Win32_PowerMeter) SetPropertyBudgetEnabled(value bool) (err error) {
	return instance.SetProperty("BudgetEnabled", value)
}

// GetBudgetEnabled gets the value of BudgetEnabled for the instance
func (instance *Win32_PowerMeter) GetPropertyBudgetEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("BudgetEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBudgetWriteable sets the value of BudgetWriteable for the instance
func (instance *Win32_PowerMeter) SetPropertyBudgetWriteable(value bool) (err error) {
	return instance.SetProperty("BudgetWriteable", value)
}

// GetBudgetWriteable gets the value of BudgetWriteable for the instance
func (instance *Win32_PowerMeter) GetPropertyBudgetWriteable() (value bool, err error) {
	retValue, err := instance.GetProperty("BudgetWriteable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfiguredBudget sets the value of ConfiguredBudget for the instance
func (instance *Win32_PowerMeter) SetPropertyConfiguredBudget(value uint32) (err error) {
	return instance.SetProperty("ConfiguredBudget", value)
}

// GetConfiguredBudget gets the value of ConfiguredBudget for the instance
func (instance *Win32_PowerMeter) GetPropertyConfiguredBudget() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConfiguredBudget")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumAveragingInterval sets the value of MaximumAveragingInterval for the instance
func (instance *Win32_PowerMeter) SetPropertyMaximumAveragingInterval(value uint32) (err error) {
	return instance.SetProperty("MaximumAveragingInterval", value)
}

// GetMaximumAveragingInterval gets the value of MaximumAveragingInterval for the instance
func (instance *Win32_PowerMeter) GetPropertyMaximumAveragingInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumAveragingInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxOperatingBudget sets the value of MaxOperatingBudget for the instance
func (instance *Win32_PowerMeter) SetPropertyMaxOperatingBudget(value uint32) (err error) {
	return instance.SetProperty("MaxOperatingBudget", value)
}

// GetMaxOperatingBudget gets the value of MaxOperatingBudget for the instance
func (instance *Win32_PowerMeter) GetPropertyMaxOperatingBudget() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxOperatingBudget")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMeterType sets the value of MeterType for the instance
func (instance *Win32_PowerMeter) SetPropertyMeterType(value uint32) (err error) {
	return instance.SetProperty("MeterType", value)
}

// GetMeterType gets the value of MeterType for the instance
func (instance *Win32_PowerMeter) GetPropertyMeterType() (value uint32, err error) {
	retValue, err := instance.GetProperty("MeterType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinimumAveragingInterval sets the value of MinimumAveragingInterval for the instance
func (instance *Win32_PowerMeter) SetPropertyMinimumAveragingInterval(value uint32) (err error) {
	return instance.SetProperty("MinimumAveragingInterval", value)
}

// GetMinimumAveragingInterval gets the value of MinimumAveragingInterval for the instance
func (instance *Win32_PowerMeter) GetPropertyMinimumAveragingInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumAveragingInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinOperatingBudget sets the value of MinOperatingBudget for the instance
func (instance *Win32_PowerMeter) SetPropertyMinOperatingBudget(value uint32) (err error) {
	return instance.SetProperty("MinOperatingBudget", value)
}

// GetMinOperatingBudget gets the value of MinOperatingBudget for the instance
func (instance *Win32_PowerMeter) GetPropertyMinOperatingBudget() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinOperatingBudget")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSamplingPeriod sets the value of SamplingPeriod for the instance
func (instance *Win32_PowerMeter) SetPropertySamplingPeriod(value uint32) (err error) {
	return instance.SetProperty("SamplingPeriod", value)
}

// GetSamplingPeriod gets the value of SamplingPeriod for the instance
func (instance *Win32_PowerMeter) GetPropertySamplingPeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("SamplingPeriod")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportCapabilities sets the value of SupportCapabilities for the instance
func (instance *Win32_PowerMeter) SetPropertySupportCapabilities(value uint32) (err error) {
	return instance.SetProperty("SupportCapabilities", value)
}

// GetSupportCapabilities gets the value of SupportCapabilities for the instance
func (instance *Win32_PowerMeter) GetPropertySupportCapabilities() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportCapabilities")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
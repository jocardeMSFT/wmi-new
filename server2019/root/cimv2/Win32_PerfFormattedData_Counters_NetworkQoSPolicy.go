// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Counters_NetworkQoSPolicy struct
type Win32_PerfFormattedData_Counters_NetworkQoSPolicy struct {
	Win32_PerfFormattedData

	//
	Bytestransmitted uint64

	//
	BytestransmittedPersec uint64

	//
	Packetsdropped uint32

	//
	PacketsdroppedPersec uint32

	//
	Packetstransmitted uint32

	//
	PacketstransmittedPersec uint32
}

// SetBytestransmitted sets the value of Bytestransmitted for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) SetPropertyBytestransmitted(value uint64) (err error) {
	return instance.SetProperty("Bytestransmitted", value)
}

// GetBytestransmitted gets the value of Bytestransmitted for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) GetPropertyBytestransmitted() (value uint64, err error) {
	retValue, err := instance.GetProperty("Bytestransmitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytestransmittedPersec sets the value of BytestransmittedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) SetPropertyBytestransmittedPersec(value uint64) (err error) {
	return instance.SetProperty("BytestransmittedPersec", value)
}

// GetBytestransmittedPersec gets the value of BytestransmittedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) GetPropertyBytestransmittedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytestransmittedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsdropped sets the value of Packetsdropped for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) SetPropertyPacketsdropped(value uint32) (err error) {
	return instance.SetProperty("Packetsdropped", value)
}

// GetPacketsdropped gets the value of Packetsdropped for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) GetPropertyPacketsdropped() (value uint32, err error) {
	retValue, err := instance.GetProperty("Packetsdropped")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketsdroppedPersec sets the value of PacketsdroppedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) SetPropertyPacketsdroppedPersec(value uint32) (err error) {
	return instance.SetProperty("PacketsdroppedPersec", value)
}

// GetPacketsdroppedPersec gets the value of PacketsdroppedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) GetPropertyPacketsdroppedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketsdroppedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketstransmitted sets the value of Packetstransmitted for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) SetPropertyPacketstransmitted(value uint32) (err error) {
	return instance.SetProperty("Packetstransmitted", value)
}

// GetPacketstransmitted gets the value of Packetstransmitted for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) GetPropertyPacketstransmitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("Packetstransmitted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPacketstransmittedPersec sets the value of PacketstransmittedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) SetPropertyPacketstransmittedPersec(value uint32) (err error) {
	return instance.SetProperty("PacketstransmittedPersec", value)
}

// GetPacketstransmittedPersec gets the value of PacketstransmittedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_NetworkQoSPolicy) GetPropertyPacketstransmittedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PacketstransmittedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_PushPrinterConnectionsPolicySetting struct
type RSOP_PushPrinterConnectionsPolicySetting struct {
	RSOP_PolicySetting

	// Whether the the connection is applied per machine, per user1 = User, 2 = Machine.
	ConnectionType PushPrinterConnectionsPolicySetting_ConnectionType

	// Indicates whether the print connectionhas been deleted.
	deleted bool

	// Printer name,
	printerName string

	// The final result of pushed printer connection. 0 indicate success
	PushResult uint32

	// Short server name
	serverName string

	// Short server name
	uncName string
}

// SetConnectionType sets the value of ConnectionType for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertyConnectionType(value PushPrinterConnectionsPolicySetting_ConnectionType) (err error) {
	return instance.SetProperty("ConnectionType", value)
}

// GetConnectionType gets the value of ConnectionType for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertyConnectionType() (value PushPrinterConnectionsPolicySetting_ConnectionType, err error) {
	retValue, err := instance.GetProperty("ConnectionType")
	if err != nil {
		return
	}
	value, ok := retValue.(PushPrinterConnectionsPolicySetting_ConnectionType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setdeleted sets the value of deleted for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertydeleted(value bool) (err error) {
	return instance.SetProperty("deleted", value)
}

// Getdeleted gets the value of deleted for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertydeleted() (value bool, err error) {
	retValue, err := instance.GetProperty("deleted")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetprinterName sets the value of printerName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertyprinterName(value string) (err error) {
	return instance.SetProperty("printerName", value)
}

// GetprinterName gets the value of printerName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertyprinterName() (value string, err error) {
	retValue, err := instance.GetProperty("printerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPushResult sets the value of PushResult for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertyPushResult(value uint32) (err error) {
	return instance.SetProperty("PushResult", value)
}

// GetPushResult gets the value of PushResult for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertyPushResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("PushResult")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetserverName sets the value of serverName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertyserverName(value string) (err error) {
	return instance.SetProperty("serverName", value)
}

// GetserverName gets the value of serverName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertyserverName() (value string, err error) {
	retValue, err := instance.GetProperty("serverName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetuncName sets the value of uncName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertyuncName(value string) (err error) {
	return instance.SetProperty("uncName", value)
}

// GetuncName gets the value of uncName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertyuncName() (value string, err error) {
	retValue, err := instance.GetProperty("uncName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
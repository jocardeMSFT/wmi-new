// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.subscription
//////////////////////////////////////////////
package subscription

// NTEventLogEventConsumer struct
type NTEventLogEventConsumer struct {
	__EventConsumer

	//
	Category uint16

	//
	EventID uint32

	//
	EventType NTEventLogEventConsumer_EventType

	//
	InsertionStringTemplates []string

	//
	Name string

	//
	NameOfRawDataProperty string

	//
	NameOfUserSIDProperty string

	//
	NumberOfInsertionStrings uint32

	//
	SourceName string

	//
	UNCServerName string
}

// SetCategory sets the value of Category for the instance
func (instance *NTEventLogEventConsumer) SetPropertyCategory(value uint16) (err error) {
	return instance.SetProperty("Category", value)
}

// GetCategory gets the value of Category for the instance
func (instance *NTEventLogEventConsumer) GetPropertyCategory() (value uint16, err error) {
	retValue, err := instance.GetProperty("Category")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventID sets the value of EventID for the instance
func (instance *NTEventLogEventConsumer) SetPropertyEventID(value uint32) (err error) {
	return instance.SetProperty("EventID", value)
}

// GetEventID gets the value of EventID for the instance
func (instance *NTEventLogEventConsumer) GetPropertyEventID() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventType sets the value of EventType for the instance
func (instance *NTEventLogEventConsumer) SetPropertyEventType(value NTEventLogEventConsumer_EventType) (err error) {
	return instance.SetProperty("EventType", value)
}

// GetEventType gets the value of EventType for the instance
func (instance *NTEventLogEventConsumer) GetPropertyEventType() (value NTEventLogEventConsumer_EventType, err error) {
	retValue, err := instance.GetProperty("EventType")
	if err != nil {
		return
	}
	value, ok := retValue.(NTEventLogEventConsumer_EventType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInsertionStringTemplates sets the value of InsertionStringTemplates for the instance
func (instance *NTEventLogEventConsumer) SetPropertyInsertionStringTemplates(value []string) (err error) {
	return instance.SetProperty("InsertionStringTemplates", value)
}

// GetInsertionStringTemplates gets the value of InsertionStringTemplates for the instance
func (instance *NTEventLogEventConsumer) GetPropertyInsertionStringTemplates() (value []string, err error) {
	retValue, err := instance.GetProperty("InsertionStringTemplates")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *NTEventLogEventConsumer) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *NTEventLogEventConsumer) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNameOfRawDataProperty sets the value of NameOfRawDataProperty for the instance
func (instance *NTEventLogEventConsumer) SetPropertyNameOfRawDataProperty(value string) (err error) {
	return instance.SetProperty("NameOfRawDataProperty", value)
}

// GetNameOfRawDataProperty gets the value of NameOfRawDataProperty for the instance
func (instance *NTEventLogEventConsumer) GetPropertyNameOfRawDataProperty() (value string, err error) {
	retValue, err := instance.GetProperty("NameOfRawDataProperty")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNameOfUserSIDProperty sets the value of NameOfUserSIDProperty for the instance
func (instance *NTEventLogEventConsumer) SetPropertyNameOfUserSIDProperty(value string) (err error) {
	return instance.SetProperty("NameOfUserSIDProperty", value)
}

// GetNameOfUserSIDProperty gets the value of NameOfUserSIDProperty for the instance
func (instance *NTEventLogEventConsumer) GetPropertyNameOfUserSIDProperty() (value string, err error) {
	retValue, err := instance.GetProperty("NameOfUserSIDProperty")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfInsertionStrings sets the value of NumberOfInsertionStrings for the instance
func (instance *NTEventLogEventConsumer) SetPropertyNumberOfInsertionStrings(value uint32) (err error) {
	return instance.SetProperty("NumberOfInsertionStrings", value)
}

// GetNumberOfInsertionStrings gets the value of NumberOfInsertionStrings for the instance
func (instance *NTEventLogEventConsumer) GetPropertyNumberOfInsertionStrings() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfInsertionStrings")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceName sets the value of SourceName for the instance
func (instance *NTEventLogEventConsumer) SetPropertySourceName(value string) (err error) {
	return instance.SetProperty("SourceName", value)
}

// GetSourceName gets the value of SourceName for the instance
func (instance *NTEventLogEventConsumer) GetPropertySourceName() (value string, err error) {
	retValue, err := instance.GetProperty("SourceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUNCServerName sets the value of UNCServerName for the instance
func (instance *NTEventLogEventConsumer) SetPropertyUNCServerName(value string) (err error) {
	return instance.SetProperty("UNCServerName", value)
}

// GetUNCServerName gets the value of UNCServerName for the instance
func (instance *NTEventLogEventConsumer) GetPropertyUNCServerName() (value string, err error) {
	retValue, err := instance.GetProperty("UNCServerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
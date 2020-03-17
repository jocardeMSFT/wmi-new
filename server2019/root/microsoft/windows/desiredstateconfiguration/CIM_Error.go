// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_Error struct
type CIM_Error struct {
	cim.WmiInstance

	//
	CIMStatusCode uint32

	//
	CIMStatusCodeDescription string

	//
	ErrorSource string

	//
	ErrorSourceFormat uint16

	//
	ErrorType uint16

	//
	Message string

	//
	MessageArguments []string

	//
	MessageID string

	//
	OtherErrorSourceFormat string

	//
	OtherErrorType string

	//
	OwningEntity string

	//
	PerceivedSeverity uint16

	//
	ProbableCause uint16

	//
	ProbableCauseDescription string

	//
	RecommendedActions []string
}

// SetCIMStatusCode sets the value of CIMStatusCode for the instance
func (instance *CIM_Error) SetPropertyCIMStatusCode(value uint32) (err error) {
	return instance.SetProperty("CIMStatusCode", value)
}

// GetCIMStatusCode gets the value of CIMStatusCode for the instance
func (instance *CIM_Error) GetPropertyCIMStatusCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("CIMStatusCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCIMStatusCodeDescription sets the value of CIMStatusCodeDescription for the instance
func (instance *CIM_Error) SetPropertyCIMStatusCodeDescription(value string) (err error) {
	return instance.SetProperty("CIMStatusCodeDescription", value)
}

// GetCIMStatusCodeDescription gets the value of CIMStatusCodeDescription for the instance
func (instance *CIM_Error) GetPropertyCIMStatusCodeDescription() (value string, err error) {
	retValue, err := instance.GetProperty("CIMStatusCodeDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorSource sets the value of ErrorSource for the instance
func (instance *CIM_Error) SetPropertyErrorSource(value string) (err error) {
	return instance.SetProperty("ErrorSource", value)
}

// GetErrorSource gets the value of ErrorSource for the instance
func (instance *CIM_Error) GetPropertyErrorSource() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorSource")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorSourceFormat sets the value of ErrorSourceFormat for the instance
func (instance *CIM_Error) SetPropertyErrorSourceFormat(value uint16) (err error) {
	return instance.SetProperty("ErrorSourceFormat", value)
}

// GetErrorSourceFormat gets the value of ErrorSourceFormat for the instance
func (instance *CIM_Error) GetPropertyErrorSourceFormat() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorSourceFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorType sets the value of ErrorType for the instance
func (instance *CIM_Error) SetPropertyErrorType(value uint16) (err error) {
	return instance.SetProperty("ErrorType", value)
}

// GetErrorType gets the value of ErrorType for the instance
func (instance *CIM_Error) GetPropertyErrorType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMessage sets the value of Message for the instance
func (instance *CIM_Error) SetPropertyMessage(value string) (err error) {
	return instance.SetProperty("Message", value)
}

// GetMessage gets the value of Message for the instance
func (instance *CIM_Error) GetPropertyMessage() (value string, err error) {
	retValue, err := instance.GetProperty("Message")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMessageArguments sets the value of MessageArguments for the instance
func (instance *CIM_Error) SetPropertyMessageArguments(value []string) (err error) {
	return instance.SetProperty("MessageArguments", value)
}

// GetMessageArguments gets the value of MessageArguments for the instance
func (instance *CIM_Error) GetPropertyMessageArguments() (value []string, err error) {
	retValue, err := instance.GetProperty("MessageArguments")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMessageID sets the value of MessageID for the instance
func (instance *CIM_Error) SetPropertyMessageID(value string) (err error) {
	return instance.SetProperty("MessageID", value)
}

// GetMessageID gets the value of MessageID for the instance
func (instance *CIM_Error) GetPropertyMessageID() (value string, err error) {
	retValue, err := instance.GetProperty("MessageID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherErrorSourceFormat sets the value of OtherErrorSourceFormat for the instance
func (instance *CIM_Error) SetPropertyOtherErrorSourceFormat(value string) (err error) {
	return instance.SetProperty("OtherErrorSourceFormat", value)
}

// GetOtherErrorSourceFormat gets the value of OtherErrorSourceFormat for the instance
func (instance *CIM_Error) GetPropertyOtherErrorSourceFormat() (value string, err error) {
	retValue, err := instance.GetProperty("OtherErrorSourceFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherErrorType sets the value of OtherErrorType for the instance
func (instance *CIM_Error) SetPropertyOtherErrorType(value string) (err error) {
	return instance.SetProperty("OtherErrorType", value)
}

// GetOtherErrorType gets the value of OtherErrorType for the instance
func (instance *CIM_Error) GetPropertyOtherErrorType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherErrorType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOwningEntity sets the value of OwningEntity for the instance
func (instance *CIM_Error) SetPropertyOwningEntity(value string) (err error) {
	return instance.SetProperty("OwningEntity", value)
}

// GetOwningEntity gets the value of OwningEntity for the instance
func (instance *CIM_Error) GetPropertyOwningEntity() (value string, err error) {
	retValue, err := instance.GetProperty("OwningEntity")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPerceivedSeverity sets the value of PerceivedSeverity for the instance
func (instance *CIM_Error) SetPropertyPerceivedSeverity(value uint16) (err error) {
	return instance.SetProperty("PerceivedSeverity", value)
}

// GetPerceivedSeverity gets the value of PerceivedSeverity for the instance
func (instance *CIM_Error) GetPropertyPerceivedSeverity() (value uint16, err error) {
	retValue, err := instance.GetProperty("PerceivedSeverity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProbableCause sets the value of ProbableCause for the instance
func (instance *CIM_Error) SetPropertyProbableCause(value uint16) (err error) {
	return instance.SetProperty("ProbableCause", value)
}

// GetProbableCause gets the value of ProbableCause for the instance
func (instance *CIM_Error) GetPropertyProbableCause() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProbableCause")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProbableCauseDescription sets the value of ProbableCauseDescription for the instance
func (instance *CIM_Error) SetPropertyProbableCauseDescription(value string) (err error) {
	return instance.SetProperty("ProbableCauseDescription", value)
}

// GetProbableCauseDescription gets the value of ProbableCauseDescription for the instance
func (instance *CIM_Error) GetPropertyProbableCauseDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ProbableCauseDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecommendedActions sets the value of RecommendedActions for the instance
func (instance *CIM_Error) SetPropertyRecommendedActions(value []string) (err error) {
	return instance.SetProperty("RecommendedActions", value)
}

// GetRecommendedActions gets the value of RecommendedActions for the instance
func (instance *CIM_Error) GetPropertyRecommendedActions() (value []string, err error) {
	retValue, err := instance.GetProperty("RecommendedActions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
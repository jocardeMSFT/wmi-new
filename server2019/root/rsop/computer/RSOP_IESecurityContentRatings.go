// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IESecurityContentRatings struct
type RSOP_IESecurityContentRatings struct {
	cim.WmiInstance

	//
	alwaysViewableSites []string

	//
	neverViewableSites []string

	//
	passwordOverrideEnabled bool

	//
	ratingSystemFileNames []string

	//
	ratingSystems []string

	//
	rsopID string

	//
	rsopPrecedence int32

	//
	selectedRatingsBureau string

	//
	viewUnknownRatedSites bool
}

// SetalwaysViewableSites sets the value of alwaysViewableSites for the instance
func (instance *RSOP_IESecurityContentRatings) SetPropertyalwaysViewableSites(value []string) (err error) {
	return instance.SetProperty("alwaysViewableSites", value)
}

// GetalwaysViewableSites gets the value of alwaysViewableSites for the instance
func (instance *RSOP_IESecurityContentRatings) GetPropertyalwaysViewableSites() (value []string, err error) {
	retValue, err := instance.GetProperty("alwaysViewableSites")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetneverViewableSites sets the value of neverViewableSites for the instance
func (instance *RSOP_IESecurityContentRatings) SetPropertyneverViewableSites(value []string) (err error) {
	return instance.SetProperty("neverViewableSites", value)
}

// GetneverViewableSites gets the value of neverViewableSites for the instance
func (instance *RSOP_IESecurityContentRatings) GetPropertyneverViewableSites() (value []string, err error) {
	retValue, err := instance.GetProperty("neverViewableSites")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetpasswordOverrideEnabled sets the value of passwordOverrideEnabled for the instance
func (instance *RSOP_IESecurityContentRatings) SetPropertypasswordOverrideEnabled(value bool) (err error) {
	return instance.SetProperty("passwordOverrideEnabled", value)
}

// GetpasswordOverrideEnabled gets the value of passwordOverrideEnabled for the instance
func (instance *RSOP_IESecurityContentRatings) GetPropertypasswordOverrideEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("passwordOverrideEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetratingSystemFileNames sets the value of ratingSystemFileNames for the instance
func (instance *RSOP_IESecurityContentRatings) SetPropertyratingSystemFileNames(value []string) (err error) {
	return instance.SetProperty("ratingSystemFileNames", value)
}

// GetratingSystemFileNames gets the value of ratingSystemFileNames for the instance
func (instance *RSOP_IESecurityContentRatings) GetPropertyratingSystemFileNames() (value []string, err error) {
	retValue, err := instance.GetProperty("ratingSystemFileNames")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetratingSystems sets the value of ratingSystems for the instance
func (instance *RSOP_IESecurityContentRatings) SetPropertyratingSystems(value []string) (err error) {
	return instance.SetProperty("ratingSystems", value)
}

// GetratingSystems gets the value of ratingSystems for the instance
func (instance *RSOP_IESecurityContentRatings) GetPropertyratingSystems() (value []string, err error) {
	retValue, err := instance.GetProperty("ratingSystems")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IESecurityContentRatings) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", value)
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IESecurityContentRatings) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IESecurityContentRatings) SetPropertyrsopPrecedence(value int32) (err error) {
	return instance.SetProperty("rsopPrecedence", value)
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IESecurityContentRatings) GetPropertyrsopPrecedence() (value int32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetselectedRatingsBureau sets the value of selectedRatingsBureau for the instance
func (instance *RSOP_IESecurityContentRatings) SetPropertyselectedRatingsBureau(value string) (err error) {
	return instance.SetProperty("selectedRatingsBureau", value)
}

// GetselectedRatingsBureau gets the value of selectedRatingsBureau for the instance
func (instance *RSOP_IESecurityContentRatings) GetPropertyselectedRatingsBureau() (value string, err error) {
	retValue, err := instance.GetProperty("selectedRatingsBureau")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetviewUnknownRatedSites sets the value of viewUnknownRatedSites for the instance
func (instance *RSOP_IESecurityContentRatings) SetPropertyviewUnknownRatedSites(value bool) (err error) {
	return instance.SetProperty("viewUnknownRatedSites", value)
}

// GetviewUnknownRatedSites gets the value of viewUnknownRatedSites for the instance
func (instance *RSOP_IESecurityContentRatings) GetPropertyviewUnknownRatedSites() (value bool, err error) {
	retValue, err := instance.GetProperty("viewUnknownRatedSites")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
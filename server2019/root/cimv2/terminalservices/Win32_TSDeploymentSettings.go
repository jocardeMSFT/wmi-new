// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TSDeploymentSettings struct
type Win32_TSDeploymentSettings struct {
	CIM_LogicalElement

	// Allow Font Smoothing
	AllowFontSmoothing bool

	// Certificate Expires On, stored as a 64bit FILETIME format
	CertificateExpiresOn string

	// Certificate used to sign RDP files
	CertificateHash []uint8

	// Certificate Issued By
	CertificateIssuedBy string

	// Certificate Issued To
	CertificateIssuedTo string

	// Color Bit Depth
	ColorBitDepth TSDeploymentSettings_ColorBitDepth

	// Contents of the RDP file corresponding to the Custom RDP Settings
	CustomRDPSettings string

	// Contents of the RDP file corresponding to the Deployment Settings, if this is set the corresponding Redirection settings and other Deployment settings are ignored and this RDP file is used.
	DeploymentRDPSettings string

	// Farm Name
	FarmName string

	// Gateway Authentication Mode
	GatewayAuthMode TSDeploymentSettings_GatewayAuthMode

	// Gateway Name
	GatewayName string

	// How Gateway is Used
	GatewayUsage TSDeploymentSettings_GatewayUsage

	// Use the same user credentials for TS Gateway and TS Server when possible
	GatewayUseCachedCreds bool

	// Use a Certificate to Sign the RDP Files
	HasCertificate bool

	// RDP Port
	Port int32

	// Redirection Options is configured by adding the following flags  None(0), Drives(1), Printers(2), Clipboard(4), Plug and Play(8), Smart Card(16)
	RedirectionOptions int32

	// Require Server Authentication (DEPRECATED)
	RequireServerAuth bool

	// Enable Multi-Monitor for desktop (not RAIL)
	UseMultimon bool
}

// SetAllowFontSmoothing sets the value of AllowFontSmoothing for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyAllowFontSmoothing(value bool) (err error) {
	return instance.SetProperty("AllowFontSmoothing", value)
}

// GetAllowFontSmoothing gets the value of AllowFontSmoothing for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyAllowFontSmoothing() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowFontSmoothing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCertificateExpiresOn sets the value of CertificateExpiresOn for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyCertificateExpiresOn(value string) (err error) {
	return instance.SetProperty("CertificateExpiresOn", value)
}

// GetCertificateExpiresOn gets the value of CertificateExpiresOn for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyCertificateExpiresOn() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateExpiresOn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCertificateHash sets the value of CertificateHash for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyCertificateHash(value []uint8) (err error) {
	return instance.SetProperty("CertificateHash", value)
}

// GetCertificateHash gets the value of CertificateHash for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyCertificateHash() (value []uint8, err error) {
	retValue, err := instance.GetProperty("CertificateHash")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCertificateIssuedBy sets the value of CertificateIssuedBy for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyCertificateIssuedBy(value string) (err error) {
	return instance.SetProperty("CertificateIssuedBy", value)
}

// GetCertificateIssuedBy gets the value of CertificateIssuedBy for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyCertificateIssuedBy() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateIssuedBy")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCertificateIssuedTo sets the value of CertificateIssuedTo for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyCertificateIssuedTo(value string) (err error) {
	return instance.SetProperty("CertificateIssuedTo", value)
}

// GetCertificateIssuedTo gets the value of CertificateIssuedTo for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyCertificateIssuedTo() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateIssuedTo")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetColorBitDepth sets the value of ColorBitDepth for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyColorBitDepth(value TSDeploymentSettings_ColorBitDepth) (err error) {
	return instance.SetProperty("ColorBitDepth", value)
}

// GetColorBitDepth gets the value of ColorBitDepth for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyColorBitDepth() (value TSDeploymentSettings_ColorBitDepth, err error) {
	retValue, err := instance.GetProperty("ColorBitDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(TSDeploymentSettings_ColorBitDepth)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCustomRDPSettings sets the value of CustomRDPSettings for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyCustomRDPSettings(value string) (err error) {
	return instance.SetProperty("CustomRDPSettings", value)
}

// GetCustomRDPSettings gets the value of CustomRDPSettings for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyCustomRDPSettings() (value string, err error) {
	retValue, err := instance.GetProperty("CustomRDPSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeploymentRDPSettings sets the value of DeploymentRDPSettings for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyDeploymentRDPSettings(value string) (err error) {
	return instance.SetProperty("DeploymentRDPSettings", value)
}

// GetDeploymentRDPSettings gets the value of DeploymentRDPSettings for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyDeploymentRDPSettings() (value string, err error) {
	retValue, err := instance.GetProperty("DeploymentRDPSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFarmName sets the value of FarmName for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyFarmName(value string) (err error) {
	return instance.SetProperty("FarmName", value)
}

// GetFarmName gets the value of FarmName for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyFarmName() (value string, err error) {
	retValue, err := instance.GetProperty("FarmName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGatewayAuthMode sets the value of GatewayAuthMode for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyGatewayAuthMode(value TSDeploymentSettings_GatewayAuthMode) (err error) {
	return instance.SetProperty("GatewayAuthMode", value)
}

// GetGatewayAuthMode gets the value of GatewayAuthMode for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyGatewayAuthMode() (value TSDeploymentSettings_GatewayAuthMode, err error) {
	retValue, err := instance.GetProperty("GatewayAuthMode")
	if err != nil {
		return
	}
	value, ok := retValue.(TSDeploymentSettings_GatewayAuthMode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGatewayName sets the value of GatewayName for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyGatewayName(value string) (err error) {
	return instance.SetProperty("GatewayName", value)
}

// GetGatewayName gets the value of GatewayName for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyGatewayName() (value string, err error) {
	retValue, err := instance.GetProperty("GatewayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGatewayUsage sets the value of GatewayUsage for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyGatewayUsage(value TSDeploymentSettings_GatewayUsage) (err error) {
	return instance.SetProperty("GatewayUsage", value)
}

// GetGatewayUsage gets the value of GatewayUsage for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyGatewayUsage() (value TSDeploymentSettings_GatewayUsage, err error) {
	retValue, err := instance.GetProperty("GatewayUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(TSDeploymentSettings_GatewayUsage)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGatewayUseCachedCreds sets the value of GatewayUseCachedCreds for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyGatewayUseCachedCreds(value bool) (err error) {
	return instance.SetProperty("GatewayUseCachedCreds", value)
}

// GetGatewayUseCachedCreds gets the value of GatewayUseCachedCreds for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyGatewayUseCachedCreds() (value bool, err error) {
	retValue, err := instance.GetProperty("GatewayUseCachedCreds")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHasCertificate sets the value of HasCertificate for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyHasCertificate(value bool) (err error) {
	return instance.SetProperty("HasCertificate", value)
}

// GetHasCertificate gets the value of HasCertificate for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyHasCertificate() (value bool, err error) {
	retValue, err := instance.GetProperty("HasCertificate")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPort sets the value of Port for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyPort(value int32) (err error) {
	return instance.SetProperty("Port", value)
}

// GetPort gets the value of Port for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyPort() (value int32, err error) {
	retValue, err := instance.GetProperty("Port")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRedirectionOptions sets the value of RedirectionOptions for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyRedirectionOptions(value int32) (err error) {
	return instance.SetProperty("RedirectionOptions", value)
}

// GetRedirectionOptions gets the value of RedirectionOptions for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyRedirectionOptions() (value int32, err error) {
	retValue, err := instance.GetProperty("RedirectionOptions")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireServerAuth sets the value of RequireServerAuth for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyRequireServerAuth(value bool) (err error) {
	return instance.SetProperty("RequireServerAuth", value)
}

// GetRequireServerAuth gets the value of RequireServerAuth for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyRequireServerAuth() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireServerAuth")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseMultimon sets the value of UseMultimon for the instance
func (instance *Win32_TSDeploymentSettings) SetPropertyUseMultimon(value bool) (err error) {
	return instance.SetProperty("UseMultimon", value)
}

// GetUseMultimon gets the value of UseMultimon for the instance
func (instance *Win32_TSDeploymentSettings) GetPropertyUseMultimon() (value bool, err error) {
	retValue, err := instance.GetProperty("UseMultimon")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
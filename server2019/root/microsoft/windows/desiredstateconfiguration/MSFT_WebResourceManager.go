// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_WebResourceManager struct
type MSFT_WebResourceManager struct {
	OMI_ResourceModuleManager

	//
	AllowUnsecureConnection bool

	//
	CertificateID string

	//
	ProxyCredential MSFT_Credential

	//
	ProxyURL string

	//
	RegistrationKey string

	//
	ServerURL string
}

// SetAllowUnsecureConnection sets the value of AllowUnsecureConnection for the instance
func (instance *MSFT_WebResourceManager) SetPropertyAllowUnsecureConnection(value bool) (err error) {
	return instance.SetProperty("AllowUnsecureConnection", value)
}

// GetAllowUnsecureConnection gets the value of AllowUnsecureConnection for the instance
func (instance *MSFT_WebResourceManager) GetPropertyAllowUnsecureConnection() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowUnsecureConnection")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCertificateID sets the value of CertificateID for the instance
func (instance *MSFT_WebResourceManager) SetPropertyCertificateID(value string) (err error) {
	return instance.SetProperty("CertificateID", value)
}

// GetCertificateID gets the value of CertificateID for the instance
func (instance *MSFT_WebResourceManager) GetPropertyCertificateID() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProxyCredential sets the value of ProxyCredential for the instance
func (instance *MSFT_WebResourceManager) SetPropertyProxyCredential(value MSFT_Credential) (err error) {
	return instance.SetProperty("ProxyCredential", value)
}

// GetProxyCredential gets the value of ProxyCredential for the instance
func (instance *MSFT_WebResourceManager) GetPropertyProxyCredential() (value MSFT_Credential, err error) {
	retValue, err := instance.GetProperty("ProxyCredential")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_Credential)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProxyURL sets the value of ProxyURL for the instance
func (instance *MSFT_WebResourceManager) SetPropertyProxyURL(value string) (err error) {
	return instance.SetProperty("ProxyURL", value)
}

// GetProxyURL gets the value of ProxyURL for the instance
func (instance *MSFT_WebResourceManager) GetPropertyProxyURL() (value string, err error) {
	retValue, err := instance.GetProperty("ProxyURL")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistrationKey sets the value of RegistrationKey for the instance
func (instance *MSFT_WebResourceManager) SetPropertyRegistrationKey(value string) (err error) {
	return instance.SetProperty("RegistrationKey", value)
}

// GetRegistrationKey gets the value of RegistrationKey for the instance
func (instance *MSFT_WebResourceManager) GetPropertyRegistrationKey() (value string, err error) {
	retValue, err := instance.GetProperty("RegistrationKey")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerURL sets the value of ServerURL for the instance
func (instance *MSFT_WebResourceManager) SetPropertyServerURL(value string) (err error) {
	return instance.SetProperty("ServerURL", value)
}

// GetServerURL gets the value of ServerURL for the instance
func (instance *MSFT_WebResourceManager) GetPropertyServerURL() (value string, err error) {
	retValue, err := instance.GetProperty("ServerURL")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
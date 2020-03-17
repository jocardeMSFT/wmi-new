// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_ClientCertificateInstall_Install03 struct
type MDM_ClientCertificateInstall_Install03 struct {
	cim.WmiInstance

	//
	AADKeyIdentifierList string

	//
	CAThumbprint string

	//
	Challenge string

	//
	ContainerName string

	//
	CustomTextToShowInPrompt string

	//
	EKUMapping string

	//
	HashAlgorithm string

	//
	InstanceID string

	//
	KeyLength int32

	//
	KeyProtection int32

	//
	KeyUsage int32

	//
	ParentID string

	//
	RetryCount int32

	//
	RetryDelay int32

	//
	ServerURL string

	//
	SubjectAlternativeNames string

	//
	SubjectName string

	//
	TemplateName string

	//
	ValidPeriod string

	//
	ValidPeriodUnits int32
}

// SetAADKeyIdentifierList sets the value of AADKeyIdentifierList for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyAADKeyIdentifierList(value string) (err error) {
	return instance.SetProperty("AADKeyIdentifierList", value)
}

// GetAADKeyIdentifierList gets the value of AADKeyIdentifierList for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyAADKeyIdentifierList() (value string, err error) {
	retValue, err := instance.GetProperty("AADKeyIdentifierList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCAThumbprint sets the value of CAThumbprint for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyCAThumbprint(value string) (err error) {
	return instance.SetProperty("CAThumbprint", value)
}

// GetCAThumbprint gets the value of CAThumbprint for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyCAThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("CAThumbprint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetChallenge sets the value of Challenge for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyChallenge(value string) (err error) {
	return instance.SetProperty("Challenge", value)
}

// GetChallenge gets the value of Challenge for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyChallenge() (value string, err error) {
	retValue, err := instance.GetProperty("Challenge")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContainerName sets the value of ContainerName for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyContainerName(value string) (err error) {
	return instance.SetProperty("ContainerName", value)
}

// GetContainerName gets the value of ContainerName for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyContainerName() (value string, err error) {
	retValue, err := instance.GetProperty("ContainerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCustomTextToShowInPrompt sets the value of CustomTextToShowInPrompt for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyCustomTextToShowInPrompt(value string) (err error) {
	return instance.SetProperty("CustomTextToShowInPrompt", value)
}

// GetCustomTextToShowInPrompt gets the value of CustomTextToShowInPrompt for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyCustomTextToShowInPrompt() (value string, err error) {
	retValue, err := instance.GetProperty("CustomTextToShowInPrompt")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEKUMapping sets the value of EKUMapping for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyEKUMapping(value string) (err error) {
	return instance.SetProperty("EKUMapping", value)
}

// GetEKUMapping gets the value of EKUMapping for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyEKUMapping() (value string, err error) {
	retValue, err := instance.GetProperty("EKUMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHashAlgorithm sets the value of HashAlgorithm for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyHashAlgorithm(value string) (err error) {
	return instance.SetProperty("HashAlgorithm", value)
}

// GetHashAlgorithm gets the value of HashAlgorithm for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyHashAlgorithm() (value string, err error) {
	retValue, err := instance.GetProperty("HashAlgorithm")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKeyLength sets the value of KeyLength for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyKeyLength(value int32) (err error) {
	return instance.SetProperty("KeyLength", value)
}

// GetKeyLength gets the value of KeyLength for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyKeyLength() (value int32, err error) {
	retValue, err := instance.GetProperty("KeyLength")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKeyProtection sets the value of KeyProtection for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyKeyProtection(value int32) (err error) {
	return instance.SetProperty("KeyProtection", value)
}

// GetKeyProtection gets the value of KeyProtection for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyKeyProtection() (value int32, err error) {
	retValue, err := instance.GetProperty("KeyProtection")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKeyUsage sets the value of KeyUsage for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyKeyUsage(value int32) (err error) {
	return instance.SetProperty("KeyUsage", value)
}

// GetKeyUsage gets the value of KeyUsage for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyKeyUsage() (value int32, err error) {
	retValue, err := instance.GetProperty("KeyUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRetryCount sets the value of RetryCount for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyRetryCount(value int32) (err error) {
	return instance.SetProperty("RetryCount", value)
}

// GetRetryCount gets the value of RetryCount for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyRetryCount() (value int32, err error) {
	retValue, err := instance.GetProperty("RetryCount")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRetryDelay sets the value of RetryDelay for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyRetryDelay(value int32) (err error) {
	return instance.SetProperty("RetryDelay", value)
}

// GetRetryDelay gets the value of RetryDelay for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyRetryDelay() (value int32, err error) {
	retValue, err := instance.GetProperty("RetryDelay")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerURL sets the value of ServerURL for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyServerURL(value string) (err error) {
	return instance.SetProperty("ServerURL", value)
}

// GetServerURL gets the value of ServerURL for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyServerURL() (value string, err error) {
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

// SetSubjectAlternativeNames sets the value of SubjectAlternativeNames for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertySubjectAlternativeNames(value string) (err error) {
	return instance.SetProperty("SubjectAlternativeNames", value)
}

// GetSubjectAlternativeNames gets the value of SubjectAlternativeNames for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertySubjectAlternativeNames() (value string, err error) {
	retValue, err := instance.GetProperty("SubjectAlternativeNames")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubjectName sets the value of SubjectName for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertySubjectName(value string) (err error) {
	return instance.SetProperty("SubjectName", value)
}

// GetSubjectName gets the value of SubjectName for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertySubjectName() (value string, err error) {
	retValue, err := instance.GetProperty("SubjectName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemplateName sets the value of TemplateName for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyTemplateName(value string) (err error) {
	return instance.SetProperty("TemplateName", value)
}

// GetTemplateName gets the value of TemplateName for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyTemplateName() (value string, err error) {
	retValue, err := instance.GetProperty("TemplateName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidPeriod sets the value of ValidPeriod for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyValidPeriod(value string) (err error) {
	return instance.SetProperty("ValidPeriod", value)
}

// GetValidPeriod gets the value of ValidPeriod for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyValidPeriod() (value string, err error) {
	retValue, err := instance.GetProperty("ValidPeriod")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidPeriodUnits sets the value of ValidPeriodUnits for the instance
func (instance *MDM_ClientCertificateInstall_Install03) SetPropertyValidPeriodUnits(value int32) (err error) {
	return instance.SetProperty("ValidPeriodUnits", value)
}

// GetValidPeriodUnits gets the value of ValidPeriodUnits for the instance
func (instance *MDM_ClientCertificateInstall_Install03) GetPropertyValidPeriodUnits() (value int32, err error) {
	retValue, err := instance.GetProperty("ValidPeriodUnits")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_ClientCertificateInstall_Install03) EnrollMethod() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EnrollMethod")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
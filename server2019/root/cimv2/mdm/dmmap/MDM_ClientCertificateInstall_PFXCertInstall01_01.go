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

// MDM_ClientCertificateInstall_PFXCertInstall01_01 struct
type MDM_ClientCertificateInstall_PFXCertInstall01_01 struct {
	cim.WmiInstance

	//
	ContainerName string

	//
	InstanceID string

	//
	KeyLocation int32

	//
	ParentID string

	//
	PFXCertBlob string

	//
	PFXCertPassword string

	//
	PFXCertPasswordEncryptionStore string

	//
	PFXCertPasswordEncryptionType int32

	//
	PFXKeyExportable bool

	//
	Status int32

	//
	Thumbprint string
}

// SetContainerName sets the value of ContainerName for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) SetPropertyContainerName(value string) (err error) {
	return instance.SetProperty("ContainerName", value)
}

// GetContainerName gets the value of ContainerName for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) GetPropertyContainerName() (value string, err error) {
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) GetPropertyInstanceID() (value string, err error) {
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

// SetKeyLocation sets the value of KeyLocation for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) SetPropertyKeyLocation(value int32) (err error) {
	return instance.SetProperty("KeyLocation", value)
}

// GetKeyLocation gets the value of KeyLocation for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) GetPropertyKeyLocation() (value int32, err error) {
	retValue, err := instance.GetProperty("KeyLocation")
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
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) GetPropertyParentID() (value string, err error) {
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

// SetPFXCertBlob sets the value of PFXCertBlob for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) SetPropertyPFXCertBlob(value string) (err error) {
	return instance.SetProperty("PFXCertBlob", value)
}

// GetPFXCertBlob gets the value of PFXCertBlob for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) GetPropertyPFXCertBlob() (value string, err error) {
	retValue, err := instance.GetProperty("PFXCertBlob")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPFXCertPassword sets the value of PFXCertPassword for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) SetPropertyPFXCertPassword(value string) (err error) {
	return instance.SetProperty("PFXCertPassword", value)
}

// GetPFXCertPassword gets the value of PFXCertPassword for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) GetPropertyPFXCertPassword() (value string, err error) {
	retValue, err := instance.GetProperty("PFXCertPassword")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPFXCertPasswordEncryptionStore sets the value of PFXCertPasswordEncryptionStore for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) SetPropertyPFXCertPasswordEncryptionStore(value string) (err error) {
	return instance.SetProperty("PFXCertPasswordEncryptionStore", value)
}

// GetPFXCertPasswordEncryptionStore gets the value of PFXCertPasswordEncryptionStore for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) GetPropertyPFXCertPasswordEncryptionStore() (value string, err error) {
	retValue, err := instance.GetProperty("PFXCertPasswordEncryptionStore")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPFXCertPasswordEncryptionType sets the value of PFXCertPasswordEncryptionType for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) SetPropertyPFXCertPasswordEncryptionType(value int32) (err error) {
	return instance.SetProperty("PFXCertPasswordEncryptionType", value)
}

// GetPFXCertPasswordEncryptionType gets the value of PFXCertPasswordEncryptionType for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) GetPropertyPFXCertPasswordEncryptionType() (value int32, err error) {
	retValue, err := instance.GetProperty("PFXCertPasswordEncryptionType")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPFXKeyExportable sets the value of PFXKeyExportable for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) SetPropertyPFXKeyExportable(value bool) (err error) {
	return instance.SetProperty("PFXKeyExportable", value)
}

// GetPFXKeyExportable gets the value of PFXKeyExportable for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) GetPropertyPFXKeyExportable() (value bool, err error) {
	retValue, err := instance.GetProperty("PFXKeyExportable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThumbprint sets the value of Thumbprint for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) SetPropertyThumbprint(value string) (err error) {
	return instance.SetProperty("Thumbprint", value)
}

// GetThumbprint gets the value of Thumbprint for the instance
func (instance *MDM_ClientCertificateInstall_PFXCertInstall01_01) GetPropertyThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("Thumbprint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
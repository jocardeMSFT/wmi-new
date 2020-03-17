// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_StorageProvider struct
type MSFT_StorageProvider struct {
	MSFT_StorageObject

	//
	CimServerName string

	//
	Manufacturer string

	//
	Name string

	//
	RemoteSubsystemCacheMode uint16

	//
	SupportedRemoteSubsystemCacheModes []uint16

	//
	SupportsSubsystemRegistration bool

	//
	Type uint16

	//
	URI string

	//
	URI_IP string

	//
	Version string
}

// SetCimServerName sets the value of CimServerName for the instance
func (instance *MSFT_StorageProvider) SetPropertyCimServerName(value string) (err error) {
	return instance.SetProperty("CimServerName", value)
}

// GetCimServerName gets the value of CimServerName for the instance
func (instance *MSFT_StorageProvider) GetPropertyCimServerName() (value string, err error) {
	retValue, err := instance.GetProperty("CimServerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *MSFT_StorageProvider) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MSFT_StorageProvider) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_StorageProvider) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_StorageProvider) GetPropertyName() (value string, err error) {
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

// SetRemoteSubsystemCacheMode sets the value of RemoteSubsystemCacheMode for the instance
func (instance *MSFT_StorageProvider) SetPropertyRemoteSubsystemCacheMode(value uint16) (err error) {
	return instance.SetProperty("RemoteSubsystemCacheMode", value)
}

// GetRemoteSubsystemCacheMode gets the value of RemoteSubsystemCacheMode for the instance
func (instance *MSFT_StorageProvider) GetPropertyRemoteSubsystemCacheMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("RemoteSubsystemCacheMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedRemoteSubsystemCacheModes sets the value of SupportedRemoteSubsystemCacheModes for the instance
func (instance *MSFT_StorageProvider) SetPropertySupportedRemoteSubsystemCacheModes(value []uint16) (err error) {
	return instance.SetProperty("SupportedRemoteSubsystemCacheModes", value)
}

// GetSupportedRemoteSubsystemCacheModes gets the value of SupportedRemoteSubsystemCacheModes for the instance
func (instance *MSFT_StorageProvider) GetPropertySupportedRemoteSubsystemCacheModes() (value []uint16, err error) {
	retValue, err := instance.GetProperty("SupportedRemoteSubsystemCacheModes")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsSubsystemRegistration sets the value of SupportsSubsystemRegistration for the instance
func (instance *MSFT_StorageProvider) SetPropertySupportsSubsystemRegistration(value bool) (err error) {
	return instance.SetProperty("SupportsSubsystemRegistration", value)
}

// GetSupportsSubsystemRegistration gets the value of SupportsSubsystemRegistration for the instance
func (instance *MSFT_StorageProvider) GetPropertySupportsSubsystemRegistration() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsSubsystemRegistration")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_StorageProvider) SetPropertyType(value uint16) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_StorageProvider) GetPropertyType() (value uint16, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetURI sets the value of URI for the instance
func (instance *MSFT_StorageProvider) SetPropertyURI(value string) (err error) {
	return instance.SetProperty("URI", value)
}

// GetURI gets the value of URI for the instance
func (instance *MSFT_StorageProvider) GetPropertyURI() (value string, err error) {
	retValue, err := instance.GetProperty("URI")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetURI_IP sets the value of URI_IP for the instance
func (instance *MSFT_StorageProvider) SetPropertyURI_IP(value string) (err error) {
	return instance.SetProperty("URI_IP", value)
}

// GetURI_IP gets the value of URI_IP for the instance
func (instance *MSFT_StorageProvider) GetPropertyURI_IP() (value string, err error) {
	retValue, err := instance.GetProperty("URI_IP")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *MSFT_StorageProvider) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *MSFT_StorageProvider) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="DiscoveryLevel" type="uint16 "></param>
// <param name="RootObject" type="MSFT_StorageObject "></param>
// <param name="RunAsJob" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageProvider) Discover( /* IN */ DiscoveryLevel uint16,
	/* IN */ RootObject MSFT_StorageObject,
	/* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Discover", DiscoveryLevel, RootObject, RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SecurityDescriptor" type="string "></param>
func (instance *MSFT_StorageProvider) GetSecurityDescriptor( /* OUT */ SecurityDescriptor string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSecurityDescriptor")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SecurityDescriptor" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageProvider) SetSecurityDescriptor( /* IN */ SecurityDescriptor string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetSecurityDescriptor", SecurityDescriptor)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="Credential" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="RegisteredSubsystem" type="MSFT_StorageSubSystem "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageProvider) RegisterSubsystem( /* IN */ ComputerName string,
	/* IN */ Credential string,
	/* OUT */ RegisteredSubsystem MSFT_StorageSubSystem,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RegisterSubsystem", ComputerName, Credential)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="StorageSubSystemUniqueId" type="string "></param>
// <param name="Subsystem" type="MSFT_StorageSubSystem "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageProvider) UnregisterSubsystem( /* IN */ Subsystem MSFT_StorageSubSystem,
	/* IN */ StorageSubSystemUniqueId string,
	/* IN */ Force bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("UnregisterSubsystem", Subsystem, StorageSubSystemUniqueId, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RemoteSubsystemCacheMode" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageProvider) SetAttributes( /* IN */ RemoteSubsystemCacheMode uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes", RemoteSubsystemCacheMode)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_Disk struct
type MSFT_Disk struct {
	MSFT_StorageObject

	//
	AdapterSerialNumber string

	//
	AllocatedSize uint64

	//
	BootFromDisk bool

	//
	BusType uint16

	//
	FirmwareVersion string

	//
	FriendlyName string

	//
	Guid string

	//
	HealthStatus uint16

	//
	IsBoot bool

	//
	IsClustered bool

	//
	IsHighlyAvailable bool

	//
	IsOffline bool

	//
	IsReadOnly bool

	//
	IsScaleOut bool

	//
	IsSystem bool

	//
	LargestFreeExtent uint64

	//
	Location string

	//
	LogicalSectorSize uint32

	//
	Manufacturer string

	//
	Model string

	//
	Number uint32

	//
	NumberOfPartitions uint32

	//
	OfflineReason uint16

	//
	OperationalStatus []uint16

	//
	PartitionStyle uint16

	//
	Path string

	//
	PhysicalSectorSize uint32

	//
	ProvisioningType uint16

	//
	SerialNumber string

	//
	Signature uint32

	//
	Size uint64

	//
	UniqueIdFormat uint16
}

// SetAdapterSerialNumber sets the value of AdapterSerialNumber for the instance
func (instance *MSFT_Disk) SetPropertyAdapterSerialNumber(value string) (err error) {
	return instance.SetProperty("AdapterSerialNumber", value)
}

// GetAdapterSerialNumber gets the value of AdapterSerialNumber for the instance
func (instance *MSFT_Disk) GetPropertyAdapterSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("AdapterSerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllocatedSize sets the value of AllocatedSize for the instance
func (instance *MSFT_Disk) SetPropertyAllocatedSize(value uint64) (err error) {
	return instance.SetProperty("AllocatedSize", value)
}

// GetAllocatedSize gets the value of AllocatedSize for the instance
func (instance *MSFT_Disk) GetPropertyAllocatedSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("AllocatedSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBootFromDisk sets the value of BootFromDisk for the instance
func (instance *MSFT_Disk) SetPropertyBootFromDisk(value bool) (err error) {
	return instance.SetProperty("BootFromDisk", value)
}

// GetBootFromDisk gets the value of BootFromDisk for the instance
func (instance *MSFT_Disk) GetPropertyBootFromDisk() (value bool, err error) {
	retValue, err := instance.GetProperty("BootFromDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBusType sets the value of BusType for the instance
func (instance *MSFT_Disk) SetPropertyBusType(value uint16) (err error) {
	return instance.SetProperty("BusType", value)
}

// GetBusType gets the value of BusType for the instance
func (instance *MSFT_Disk) GetPropertyBusType() (value uint16, err error) {
	retValue, err := instance.GetProperty("BusType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFirmwareVersion sets the value of FirmwareVersion for the instance
func (instance *MSFT_Disk) SetPropertyFirmwareVersion(value string) (err error) {
	return instance.SetProperty("FirmwareVersion", value)
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *MSFT_Disk) GetPropertyFirmwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_Disk) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_Disk) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuid sets the value of Guid for the instance
func (instance *MSFT_Disk) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", value)
}

// GetGuid gets the value of Guid for the instance
func (instance *MSFT_Disk) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_Disk) SetPropertyHealthStatus(value uint16) (err error) {
	return instance.SetProperty("HealthStatus", value)
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_Disk) GetPropertyHealthStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsBoot sets the value of IsBoot for the instance
func (instance *MSFT_Disk) SetPropertyIsBoot(value bool) (err error) {
	return instance.SetProperty("IsBoot", value)
}

// GetIsBoot gets the value of IsBoot for the instance
func (instance *MSFT_Disk) GetPropertyIsBoot() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBoot")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsClustered sets the value of IsClustered for the instance
func (instance *MSFT_Disk) SetPropertyIsClustered(value bool) (err error) {
	return instance.SetProperty("IsClustered", value)
}

// GetIsClustered gets the value of IsClustered for the instance
func (instance *MSFT_Disk) GetPropertyIsClustered() (value bool, err error) {
	retValue, err := instance.GetProperty("IsClustered")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsHighlyAvailable sets the value of IsHighlyAvailable for the instance
func (instance *MSFT_Disk) SetPropertyIsHighlyAvailable(value bool) (err error) {
	return instance.SetProperty("IsHighlyAvailable", value)
}

// GetIsHighlyAvailable gets the value of IsHighlyAvailable for the instance
func (instance *MSFT_Disk) GetPropertyIsHighlyAvailable() (value bool, err error) {
	retValue, err := instance.GetProperty("IsHighlyAvailable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsOffline sets the value of IsOffline for the instance
func (instance *MSFT_Disk) SetPropertyIsOffline(value bool) (err error) {
	return instance.SetProperty("IsOffline", value)
}

// GetIsOffline gets the value of IsOffline for the instance
func (instance *MSFT_Disk) GetPropertyIsOffline() (value bool, err error) {
	retValue, err := instance.GetProperty("IsOffline")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsReadOnly sets the value of IsReadOnly for the instance
func (instance *MSFT_Disk) SetPropertyIsReadOnly(value bool) (err error) {
	return instance.SetProperty("IsReadOnly", value)
}

// GetIsReadOnly gets the value of IsReadOnly for the instance
func (instance *MSFT_Disk) GetPropertyIsReadOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("IsReadOnly")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsScaleOut sets the value of IsScaleOut for the instance
func (instance *MSFT_Disk) SetPropertyIsScaleOut(value bool) (err error) {
	return instance.SetProperty("IsScaleOut", value)
}

// GetIsScaleOut gets the value of IsScaleOut for the instance
func (instance *MSFT_Disk) GetPropertyIsScaleOut() (value bool, err error) {
	retValue, err := instance.GetProperty("IsScaleOut")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsSystem sets the value of IsSystem for the instance
func (instance *MSFT_Disk) SetPropertyIsSystem(value bool) (err error) {
	return instance.SetProperty("IsSystem", value)
}

// GetIsSystem gets the value of IsSystem for the instance
func (instance *MSFT_Disk) GetPropertyIsSystem() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSystem")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLargestFreeExtent sets the value of LargestFreeExtent for the instance
func (instance *MSFT_Disk) SetPropertyLargestFreeExtent(value uint64) (err error) {
	return instance.SetProperty("LargestFreeExtent", value)
}

// GetLargestFreeExtent gets the value of LargestFreeExtent for the instance
func (instance *MSFT_Disk) GetPropertyLargestFreeExtent() (value uint64, err error) {
	retValue, err := instance.GetProperty("LargestFreeExtent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocation sets the value of Location for the instance
func (instance *MSFT_Disk) SetPropertyLocation(value string) (err error) {
	return instance.SetProperty("Location", value)
}

// GetLocation gets the value of Location for the instance
func (instance *MSFT_Disk) GetPropertyLocation() (value string, err error) {
	retValue, err := instance.GetProperty("Location")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogicalSectorSize sets the value of LogicalSectorSize for the instance
func (instance *MSFT_Disk) SetPropertyLogicalSectorSize(value uint32) (err error) {
	return instance.SetProperty("LogicalSectorSize", value)
}

// GetLogicalSectorSize gets the value of LogicalSectorSize for the instance
func (instance *MSFT_Disk) GetPropertyLogicalSectorSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogicalSectorSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *MSFT_Disk) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MSFT_Disk) GetPropertyManufacturer() (value string, err error) {
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

// SetModel sets the value of Model for the instance
func (instance *MSFT_Disk) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", value)
}

// GetModel gets the value of Model for the instance
func (instance *MSFT_Disk) GetPropertyModel() (value string, err error) {
	retValue, err := instance.GetProperty("Model")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumber sets the value of Number for the instance
func (instance *MSFT_Disk) SetPropertyNumber(value uint32) (err error) {
	return instance.SetProperty("Number", value)
}

// GetNumber gets the value of Number for the instance
func (instance *MSFT_Disk) GetPropertyNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("Number")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfPartitions sets the value of NumberOfPartitions for the instance
func (instance *MSFT_Disk) SetPropertyNumberOfPartitions(value uint32) (err error) {
	return instance.SetProperty("NumberOfPartitions", value)
}

// GetNumberOfPartitions gets the value of NumberOfPartitions for the instance
func (instance *MSFT_Disk) GetPropertyNumberOfPartitions() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfPartitions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOfflineReason sets the value of OfflineReason for the instance
func (instance *MSFT_Disk) SetPropertyOfflineReason(value uint16) (err error) {
	return instance.SetProperty("OfflineReason", value)
}

// GetOfflineReason gets the value of OfflineReason for the instance
func (instance *MSFT_Disk) GetPropertyOfflineReason() (value uint16, err error) {
	retValue, err := instance.GetProperty("OfflineReason")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_Disk) SetPropertyOperationalStatus(value []uint16) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_Disk) GetPropertyOperationalStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPartitionStyle sets the value of PartitionStyle for the instance
func (instance *MSFT_Disk) SetPropertyPartitionStyle(value uint16) (err error) {
	return instance.SetProperty("PartitionStyle", value)
}

// GetPartitionStyle gets the value of PartitionStyle for the instance
func (instance *MSFT_Disk) GetPropertyPartitionStyle() (value uint16, err error) {
	retValue, err := instance.GetProperty("PartitionStyle")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPath sets the value of Path for the instance
func (instance *MSFT_Disk) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *MSFT_Disk) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalSectorSize sets the value of PhysicalSectorSize for the instance
func (instance *MSFT_Disk) SetPropertyPhysicalSectorSize(value uint32) (err error) {
	return instance.SetProperty("PhysicalSectorSize", value)
}

// GetPhysicalSectorSize gets the value of PhysicalSectorSize for the instance
func (instance *MSFT_Disk) GetPropertyPhysicalSectorSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PhysicalSectorSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProvisioningType sets the value of ProvisioningType for the instance
func (instance *MSFT_Disk) SetPropertyProvisioningType(value uint16) (err error) {
	return instance.SetProperty("ProvisioningType", value)
}

// GetProvisioningType gets the value of ProvisioningType for the instance
func (instance *MSFT_Disk) GetPropertyProvisioningType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProvisioningType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MSFT_Disk) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", value)
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MSFT_Disk) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSignature sets the value of Signature for the instance
func (instance *MSFT_Disk) SetPropertySignature(value uint32) (err error) {
	return instance.SetProperty("Signature", value)
}

// GetSignature gets the value of Signature for the instance
func (instance *MSFT_Disk) GetPropertySignature() (value uint32, err error) {
	retValue, err := instance.GetProperty("Signature")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSize sets the value of Size for the instance
func (instance *MSFT_Disk) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", value)
}

// GetSize gets the value of Size for the instance
func (instance *MSFT_Disk) GetPropertySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Size")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniqueIdFormat sets the value of UniqueIdFormat for the instance
func (instance *MSFT_Disk) SetPropertyUniqueIdFormat(value uint16) (err error) {
	return instance.SetProperty("UniqueIdFormat", value)
}

// GetUniqueIdFormat gets the value of UniqueIdFormat for the instance
func (instance *MSFT_Disk) GetPropertyUniqueIdFormat() (value uint16, err error) {
	retValue, err := instance.GetProperty("UniqueIdFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Alignment" type="uint32 "></param>
// <param name="AssignDriveLetter" type="bool "></param>
// <param name="DriveLetter" type="byte "></param>
// <param name="GptType" type="string "></param>
// <param name="IsActive" type="bool "></param>
// <param name="IsHidden" type="bool "></param>
// <param name="MbrType" type="uint16 "></param>
// <param name="Offset" type="uint64 "></param>
// <param name="Size" type="uint64 "></param>
// <param name="UseMaximumSize" type="bool "></param>

// <param name="CreatedPartition" type="MSFT_Partition "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) CreatePartition( /* IN */ Size uint64,
	/* IN */ UseMaximumSize bool,
	/* IN */ Offset uint64,
	/* IN */ Alignment uint32,
	/* IN */ DriveLetter byte,
	/* IN */ AssignDriveLetter bool,
	/* IN */ MbrType uint16,
	/* IN */ GptType string,
	/* IN */ IsHidden bool,
	/* IN */ IsActive bool,
	/* OUT */ CreatedPartition MSFT_Partition,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreatePartition", Size, UseMaximumSize, Offset, Alignment, DriveLetter, AssignDriveLetter, MbrType, GptType, IsHidden, IsActive)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PartitionStyle" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) Initialize( /* IN */ PartitionStyle uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Initialize", PartitionStyle)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RemoveData" type="bool "></param>
// <param name="RemoveOEM" type="bool "></param>
// <param name="RunAsJob" type="bool "></param>
// <param name="Sanitize" type="bool "></param>
// <param name="ZeroOutEntireDisk" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) Clear( /* IN */ RemoveData bool,
	/* IN */ RemoveOEM bool,
	/* IN */ ZeroOutEntireDisk bool,
	/* IN */ Sanitize bool,
	/* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Clear", RemoveData, RemoveOEM, ZeroOutEntireDisk, Sanitize, RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PartitionStyle" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) ConvertStyle( /* IN */ PartitionStyle uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ConvertStyle", PartitionStyle)
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
func (instance *MSFT_Disk) Offline( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Offline")
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
func (instance *MSFT_Disk) Online( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Online")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Guid" type="string "></param>
// <param name="IsReadOnly" type="bool "></param>
// <param name="Signature" type="uint32 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) SetAttributes( /* IN */ IsReadOnly bool,
	/* IN */ Signature uint32,
	/* IN */ Guid string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes", IsReadOnly, Signature, Guid)
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
func (instance *MSFT_Disk) Refresh( /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Refresh")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccessPath" type="string "></param>
// <param name="AllocationUnitSize" type="uint32 "></param>
// <param name="FileSystem" type="uint16 "></param>
// <param name="FriendlyName" type="string "></param>
// <param name="RunAsJob" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="CreatedVolume" type="MSFT_Volume "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) CreateVolume( /* IN */ FriendlyName string,
	/* IN */ FileSystem uint16,
	/* IN */ AccessPath string,
	/* IN */ AllocationUnitSize uint32,
	/* OUT */ CreatedVolume MSFT_Volume,
	/* OPTIONAL IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateVolume", FriendlyName, FileSystem, AccessPath, AllocationUnitSize, RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RunAsJob" type="bool "></param>
// <param name="ScaleOut" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) EnableHighAvailability( /* IN */ ScaleOut bool,
	/* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnableHighAvailability", ScaleOut, RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RunAsJob" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Disk) DisableHighAvailability( /* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisableHighAvailability", RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
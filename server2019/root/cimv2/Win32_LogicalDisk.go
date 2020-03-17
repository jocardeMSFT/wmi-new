// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_LogicalDisk struct
type Win32_LogicalDisk struct {
	CIM_LogicalDisk

	//
	Compressed bool

	//
	DriveType uint32

	//
	FileSystem string

	//
	MaximumComponentLength uint32

	//
	MediaType uint32

	//
	ProviderName string

	//
	QuotasDisabled bool

	//
	QuotasIncomplete bool

	//
	QuotasRebuilding bool

	//
	SupportsDiskQuotas bool

	//
	SupportsFileBasedCompression bool

	//
	VolumeDirty bool

	//
	VolumeName string

	//
	VolumeSerialNumber string
}

// SetCompressed sets the value of Compressed for the instance
func (instance *Win32_LogicalDisk) SetPropertyCompressed(value bool) (err error) {
	return instance.SetProperty("Compressed", value)
}

// GetCompressed gets the value of Compressed for the instance
func (instance *Win32_LogicalDisk) GetPropertyCompressed() (value bool, err error) {
	retValue, err := instance.GetProperty("Compressed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriveType sets the value of DriveType for the instance
func (instance *Win32_LogicalDisk) SetPropertyDriveType(value uint32) (err error) {
	return instance.SetProperty("DriveType", value)
}

// GetDriveType gets the value of DriveType for the instance
func (instance *Win32_LogicalDisk) GetPropertyDriveType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DriveType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileSystem sets the value of FileSystem for the instance
func (instance *Win32_LogicalDisk) SetPropertyFileSystem(value string) (err error) {
	return instance.SetProperty("FileSystem", value)
}

// GetFileSystem gets the value of FileSystem for the instance
func (instance *Win32_LogicalDisk) GetPropertyFileSystem() (value string, err error) {
	retValue, err := instance.GetProperty("FileSystem")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumComponentLength sets the value of MaximumComponentLength for the instance
func (instance *Win32_LogicalDisk) SetPropertyMaximumComponentLength(value uint32) (err error) {
	return instance.SetProperty("MaximumComponentLength", value)
}

// GetMaximumComponentLength gets the value of MaximumComponentLength for the instance
func (instance *Win32_LogicalDisk) GetPropertyMaximumComponentLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumComponentLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMediaType sets the value of MediaType for the instance
func (instance *Win32_LogicalDisk) SetPropertyMediaType(value uint32) (err error) {
	return instance.SetProperty("MediaType", value)
}

// GetMediaType gets the value of MediaType for the instance
func (instance *Win32_LogicalDisk) GetPropertyMediaType() (value uint32, err error) {
	retValue, err := instance.GetProperty("MediaType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProviderName sets the value of ProviderName for the instance
func (instance *Win32_LogicalDisk) SetPropertyProviderName(value string) (err error) {
	return instance.SetProperty("ProviderName", value)
}

// GetProviderName gets the value of ProviderName for the instance
func (instance *Win32_LogicalDisk) GetPropertyProviderName() (value string, err error) {
	retValue, err := instance.GetProperty("ProviderName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuotasDisabled sets the value of QuotasDisabled for the instance
func (instance *Win32_LogicalDisk) SetPropertyQuotasDisabled(value bool) (err error) {
	return instance.SetProperty("QuotasDisabled", value)
}

// GetQuotasDisabled gets the value of QuotasDisabled for the instance
func (instance *Win32_LogicalDisk) GetPropertyQuotasDisabled() (value bool, err error) {
	retValue, err := instance.GetProperty("QuotasDisabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuotasIncomplete sets the value of QuotasIncomplete for the instance
func (instance *Win32_LogicalDisk) SetPropertyQuotasIncomplete(value bool) (err error) {
	return instance.SetProperty("QuotasIncomplete", value)
}

// GetQuotasIncomplete gets the value of QuotasIncomplete for the instance
func (instance *Win32_LogicalDisk) GetPropertyQuotasIncomplete() (value bool, err error) {
	retValue, err := instance.GetProperty("QuotasIncomplete")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuotasRebuilding sets the value of QuotasRebuilding for the instance
func (instance *Win32_LogicalDisk) SetPropertyQuotasRebuilding(value bool) (err error) {
	return instance.SetProperty("QuotasRebuilding", value)
}

// GetQuotasRebuilding gets the value of QuotasRebuilding for the instance
func (instance *Win32_LogicalDisk) GetPropertyQuotasRebuilding() (value bool, err error) {
	retValue, err := instance.GetProperty("QuotasRebuilding")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsDiskQuotas sets the value of SupportsDiskQuotas for the instance
func (instance *Win32_LogicalDisk) SetPropertySupportsDiskQuotas(value bool) (err error) {
	return instance.SetProperty("SupportsDiskQuotas", value)
}

// GetSupportsDiskQuotas gets the value of SupportsDiskQuotas for the instance
func (instance *Win32_LogicalDisk) GetPropertySupportsDiskQuotas() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsDiskQuotas")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsFileBasedCompression sets the value of SupportsFileBasedCompression for the instance
func (instance *Win32_LogicalDisk) SetPropertySupportsFileBasedCompression(value bool) (err error) {
	return instance.SetProperty("SupportsFileBasedCompression", value)
}

// GetSupportsFileBasedCompression gets the value of SupportsFileBasedCompression for the instance
func (instance *Win32_LogicalDisk) GetPropertySupportsFileBasedCompression() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsFileBasedCompression")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolumeDirty sets the value of VolumeDirty for the instance
func (instance *Win32_LogicalDisk) SetPropertyVolumeDirty(value bool) (err error) {
	return instance.SetProperty("VolumeDirty", value)
}

// GetVolumeDirty gets the value of VolumeDirty for the instance
func (instance *Win32_LogicalDisk) GetPropertyVolumeDirty() (value bool, err error) {
	retValue, err := instance.GetProperty("VolumeDirty")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolumeName sets the value of VolumeName for the instance
func (instance *Win32_LogicalDisk) SetPropertyVolumeName(value string) (err error) {
	return instance.SetProperty("VolumeName", value)
}

// GetVolumeName gets the value of VolumeName for the instance
func (instance *Win32_LogicalDisk) GetPropertyVolumeName() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolumeSerialNumber sets the value of VolumeSerialNumber for the instance
func (instance *Win32_LogicalDisk) SetPropertyVolumeSerialNumber(value string) (err error) {
	return instance.SetProperty("VolumeSerialNumber", value)
}

// GetVolumeSerialNumber gets the value of VolumeSerialNumber for the instance
func (instance *Win32_LogicalDisk) GetPropertyVolumeSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeSerialNumber")
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

// <param name="FixErrors" type="bool "></param>
// <param name="ForceDismount" type="bool "></param>
// <param name="OkToRunAtBootUp" type="bool "></param>
// <param name="RecoverBadSectors" type="bool "></param>
// <param name="SkipFolderCycle" type="bool "></param>
// <param name="VigorousIndexCheck" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_LogicalDisk) Chkdsk( /* IN */ FixErrors bool,
	/* IN */ VigorousIndexCheck bool,
	/* IN */ SkipFolderCycle bool,
	/* IN */ ForceDismount bool,
	/* IN */ RecoverBadSectors bool,
	/* IN */ OkToRunAtBootUp bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Chkdsk", FixErrors, VigorousIndexCheck, SkipFolderCycle, ForceDismount, RecoverBadSectors, OkToRunAtBootUp)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="LogicalDisk" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_LogicalDisk) ScheduleAutoChk( /* IN */ LogicalDisk []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ScheduleAutoChk", LogicalDisk)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="LogicalDisk" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_LogicalDisk) ExcludeFromAutochk( /* IN */ LogicalDisk []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ExcludeFromAutochk", LogicalDisk)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
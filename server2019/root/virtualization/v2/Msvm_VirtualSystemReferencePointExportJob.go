// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_VirtualSystemReferencePointExportJob struct
type Msvm_VirtualSystemReferencePointExportJob struct {
	CIM_ConcreteJob

	//
	BaseReferencePointId string

	//
	Cancellable bool

	//
	ErrorSummaryDescription string

	//
	ExportDirectory string

	//
	ExportedConfigFilePath string

	//
	ExportedDisks []string

	//
	ExportedGuestStateFilePath string

	//
	ExportedLogFilePaths []string

	//
	ExportedRuntimeFilePath string

	//
	ReferencePointId string

	//
	VirtualMachineId string
}

// SetBaseReferencePointId sets the value of BaseReferencePointId for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) SetPropertyBaseReferencePointId(value string) (err error) {
	return instance.SetProperty("BaseReferencePointId", value)
}

// GetBaseReferencePointId gets the value of BaseReferencePointId for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetPropertyBaseReferencePointId() (value string, err error) {
	retValue, err := instance.GetProperty("BaseReferencePointId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCancellable sets the value of Cancellable for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) SetPropertyCancellable(value bool) (err error) {
	return instance.SetProperty("Cancellable", value)
}

// GetCancellable gets the value of Cancellable for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetPropertyCancellable() (value bool, err error) {
	retValue, err := instance.GetProperty("Cancellable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorSummaryDescription sets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) SetPropertyErrorSummaryDescription(value string) (err error) {
	return instance.SetProperty("ErrorSummaryDescription", value)
}

// GetErrorSummaryDescription gets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetPropertyErrorSummaryDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorSummaryDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExportDirectory sets the value of ExportDirectory for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) SetPropertyExportDirectory(value string) (err error) {
	return instance.SetProperty("ExportDirectory", value)
}

// GetExportDirectory gets the value of ExportDirectory for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetPropertyExportDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("ExportDirectory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExportedConfigFilePath sets the value of ExportedConfigFilePath for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) SetPropertyExportedConfigFilePath(value string) (err error) {
	return instance.SetProperty("ExportedConfigFilePath", value)
}

// GetExportedConfigFilePath gets the value of ExportedConfigFilePath for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetPropertyExportedConfigFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("ExportedConfigFilePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExportedDisks sets the value of ExportedDisks for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) SetPropertyExportedDisks(value []string) (err error) {
	return instance.SetProperty("ExportedDisks", value)
}

// GetExportedDisks gets the value of ExportedDisks for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetPropertyExportedDisks() (value []string, err error) {
	retValue, err := instance.GetProperty("ExportedDisks")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExportedGuestStateFilePath sets the value of ExportedGuestStateFilePath for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) SetPropertyExportedGuestStateFilePath(value string) (err error) {
	return instance.SetProperty("ExportedGuestStateFilePath", value)
}

// GetExportedGuestStateFilePath gets the value of ExportedGuestStateFilePath for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetPropertyExportedGuestStateFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("ExportedGuestStateFilePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExportedLogFilePaths sets the value of ExportedLogFilePaths for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) SetPropertyExportedLogFilePaths(value []string) (err error) {
	return instance.SetProperty("ExportedLogFilePaths", value)
}

// GetExportedLogFilePaths gets the value of ExportedLogFilePaths for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetPropertyExportedLogFilePaths() (value []string, err error) {
	retValue, err := instance.GetProperty("ExportedLogFilePaths")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExportedRuntimeFilePath sets the value of ExportedRuntimeFilePath for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) SetPropertyExportedRuntimeFilePath(value string) (err error) {
	return instance.SetProperty("ExportedRuntimeFilePath", value)
}

// GetExportedRuntimeFilePath gets the value of ExportedRuntimeFilePath for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetPropertyExportedRuntimeFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("ExportedRuntimeFilePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReferencePointId sets the value of ReferencePointId for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) SetPropertyReferencePointId(value string) (err error) {
	return instance.SetProperty("ReferencePointId", value)
}

// GetReferencePointId gets the value of ReferencePointId for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetPropertyReferencePointId() (value string, err error) {
	retValue, err := instance.GetProperty("ReferencePointId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualMachineId sets the value of VirtualMachineId for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) SetPropertyVirtualMachineId(value string) (err error) {
	return instance.SetProperty("VirtualMachineId", value)
}

// GetVirtualMachineId gets the value of VirtualMachineId for the instance
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetPropertyVirtualMachineId() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualMachineId")
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

// <param name="Errors" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VirtualSystemReferencePointExportJob) GetErrorEx( /* OUT */ Errors []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetErrorEx")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
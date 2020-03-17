// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetRoute struct
type MSFT_NetRoute struct {
	CIM_NextHopRoute

	//
	AddressFamily uint16

	//
	CompartmentId uint32

	//
	DestinationPrefix string

	//
	InterfaceAlias string

	//
	InterfaceIndex uint32

	//
	InterfaceMetric uint32

	//
	NextHop string

	//
	PreferredLifetime string

	//
	Protocol uint16

	//
	Publish uint8

	//
	State uint8

	//
	Store uint8

	//
	ValidLifetime string
}

// SetAddressFamily sets the value of AddressFamily for the instance
func (instance *MSFT_NetRoute) SetPropertyAddressFamily(value uint16) (err error) {
	return instance.SetProperty("AddressFamily", value)
}

// GetAddressFamily gets the value of AddressFamily for the instance
func (instance *MSFT_NetRoute) GetPropertyAddressFamily() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressFamily")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompartmentId sets the value of CompartmentId for the instance
func (instance *MSFT_NetRoute) SetPropertyCompartmentId(value uint32) (err error) {
	return instance.SetProperty("CompartmentId", value)
}

// GetCompartmentId gets the value of CompartmentId for the instance
func (instance *MSFT_NetRoute) GetPropertyCompartmentId() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompartmentId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDestinationPrefix sets the value of DestinationPrefix for the instance
func (instance *MSFT_NetRoute) SetPropertyDestinationPrefix(value string) (err error) {
	return instance.SetProperty("DestinationPrefix", value)
}

// GetDestinationPrefix gets the value of DestinationPrefix for the instance
func (instance *MSFT_NetRoute) GetPropertyDestinationPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationPrefix")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetRoute) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", value)
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetRoute) GetPropertyInterfaceAlias() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_NetRoute) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", value)
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetRoute) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceMetric sets the value of InterfaceMetric for the instance
func (instance *MSFT_NetRoute) SetPropertyInterfaceMetric(value uint32) (err error) {
	return instance.SetProperty("InterfaceMetric", value)
}

// GetInterfaceMetric gets the value of InterfaceMetric for the instance
func (instance *MSFT_NetRoute) GetPropertyInterfaceMetric() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceMetric")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNextHop sets the value of NextHop for the instance
func (instance *MSFT_NetRoute) SetPropertyNextHop(value string) (err error) {
	return instance.SetProperty("NextHop", value)
}

// GetNextHop gets the value of NextHop for the instance
func (instance *MSFT_NetRoute) GetPropertyNextHop() (value string, err error) {
	retValue, err := instance.GetProperty("NextHop")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreferredLifetime sets the value of PreferredLifetime for the instance
func (instance *MSFT_NetRoute) SetPropertyPreferredLifetime(value string) (err error) {
	return instance.SetProperty("PreferredLifetime", value)
}

// GetPreferredLifetime gets the value of PreferredLifetime for the instance
func (instance *MSFT_NetRoute) GetPropertyPreferredLifetime() (value string, err error) {
	retValue, err := instance.GetProperty("PreferredLifetime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocol sets the value of Protocol for the instance
func (instance *MSFT_NetRoute) SetPropertyProtocol(value uint16) (err error) {
	return instance.SetProperty("Protocol", value)
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MSFT_NetRoute) GetPropertyProtocol() (value uint16, err error) {
	retValue, err := instance.GetProperty("Protocol")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPublish sets the value of Publish for the instance
func (instance *MSFT_NetRoute) SetPropertyPublish(value uint8) (err error) {
	return instance.SetProperty("Publish", value)
}

// GetPublish gets the value of Publish for the instance
func (instance *MSFT_NetRoute) GetPropertyPublish() (value uint8, err error) {
	retValue, err := instance.GetProperty("Publish")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_NetRoute) SetPropertyState(value uint8) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_NetRoute) GetPropertyState() (value uint8, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStore sets the value of Store for the instance
func (instance *MSFT_NetRoute) SetPropertyStore(value uint8) (err error) {
	return instance.SetProperty("Store", value)
}

// GetStore gets the value of Store for the instance
func (instance *MSFT_NetRoute) GetPropertyStore() (value uint8, err error) {
	retValue, err := instance.GetProperty("Store")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidLifetime sets the value of ValidLifetime for the instance
func (instance *MSFT_NetRoute) SetPropertyValidLifetime(value string) (err error) {
	return instance.SetProperty("ValidLifetime", value)
}

// GetValidLifetime gets the value of ValidLifetime for the instance
func (instance *MSFT_NetRoute) GetPropertyValidLifetime() (value string, err error) {
	retValue, err := instance.GetProperty("ValidLifetime")
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

// <param name="AddressFamily" type="uint16 "></param>
// <param name="CompartmentId" type="uint32 "></param>
// <param name="DestinationPrefix" type="string "></param>
// <param name="InterfaceAlias" type="string "></param>
// <param name="InterfaceIndex" type="uint32 "></param>
// <param name="NextHop" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PolicyStore" type="string "></param>
// <param name="PreferredLifetime" type="string "></param>
// <param name="Protocol" type="uint16 "></param>
// <param name="Publish" type="uint8 "></param>
// <param name="RouteMetric" type="uint16 "></param>
// <param name="ValidLifetime" type="string "></param>

// <param name="CmdletOutput" type="MSFT_NetRoute []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetRoute) Create( /* IN */ InterfaceIndex uint32,
	/* IN */ InterfaceAlias string,
	/* IN */ DestinationPrefix string,
	/* IN */ NextHop string,
	/* IN */ Publish uint8,
	/* IN */ RouteMetric uint16,
	/* IN */ Protocol uint16,
	/* IN */ CompartmentId uint32,
	/* IN */ ValidLifetime string,
	/* IN */ PreferredLifetime string,
	/* IN */ PolicyStore string,
	/* IN */ AddressFamily uint16,
	/* IN */ PassThru bool,
	/* OUT */ CmdletOutput []MSFT_NetRoute) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Create", InterfaceIndex, InterfaceAlias, DestinationPrefix, NextHop, Publish, RouteMetric, Protocol, CompartmentId, ValidLifetime, PreferredLifetime, PolicyStore, AddressFamily, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InterfaceIndex" type="uint32 "></param>
// <param name="LocalIPAddress" type="string "></param>
// <param name="RemoteIPAddress" type="string "></param>

// <param name="CmdletOutput" type="CIM_ManagedElement []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetRoute) Find( /* IN */ InterfaceIndex uint32,
	/* IN */ LocalIPAddress string,
	/* IN */ RemoteIPAddress string,
	/* OUT */ CmdletOutput []CIM_ManagedElement) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Find", InterfaceIndex, LocalIPAddress, RemoteIPAddress)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
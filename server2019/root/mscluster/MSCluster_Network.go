// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_Network struct
type MSCluster_Network struct {
	MSCluster_LogicalElement

	//
	Address string

	//
	AddressMask string

	//
	AutoMetric bool

	//
	Id string

	//
	IPv4Addresses []string

	//
	IPv4PrefixLengths []string

	//
	IPv6Addresses []string

	//
	IPv6PrefixLengths []string

	//
	Metric uint32

	//
	PrivateProperties MSCluster_Property

	//
	Role uint32

	//
	State uint32
}

// SetAddress sets the value of Address for the instance
func (instance *MSCluster_Network) SetPropertyAddress(value string) (err error) {
	return instance.SetProperty("Address", value)
}

// GetAddress gets the value of Address for the instance
func (instance *MSCluster_Network) GetPropertyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAddressMask sets the value of AddressMask for the instance
func (instance *MSCluster_Network) SetPropertyAddressMask(value string) (err error) {
	return instance.SetProperty("AddressMask", value)
}

// GetAddressMask gets the value of AddressMask for the instance
func (instance *MSCluster_Network) GetPropertyAddressMask() (value string, err error) {
	retValue, err := instance.GetProperty("AddressMask")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoMetric sets the value of AutoMetric for the instance
func (instance *MSCluster_Network) SetPropertyAutoMetric(value bool) (err error) {
	return instance.SetProperty("AutoMetric", value)
}

// GetAutoMetric gets the value of AutoMetric for the instance
func (instance *MSCluster_Network) GetPropertyAutoMetric() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoMetric")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSCluster_Network) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSCluster_Network) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4Addresses sets the value of IPv4Addresses for the instance
func (instance *MSCluster_Network) SetPropertyIPv4Addresses(value []string) (err error) {
	return instance.SetProperty("IPv4Addresses", value)
}

// GetIPv4Addresses gets the value of IPv4Addresses for the instance
func (instance *MSCluster_Network) GetPropertyIPv4Addresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv4Addresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4PrefixLengths sets the value of IPv4PrefixLengths for the instance
func (instance *MSCluster_Network) SetPropertyIPv4PrefixLengths(value []string) (err error) {
	return instance.SetProperty("IPv4PrefixLengths", value)
}

// GetIPv4PrefixLengths gets the value of IPv4PrefixLengths for the instance
func (instance *MSCluster_Network) GetPropertyIPv4PrefixLengths() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv4PrefixLengths")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6Addresses sets the value of IPv6Addresses for the instance
func (instance *MSCluster_Network) SetPropertyIPv6Addresses(value []string) (err error) {
	return instance.SetProperty("IPv6Addresses", value)
}

// GetIPv6Addresses gets the value of IPv6Addresses for the instance
func (instance *MSCluster_Network) GetPropertyIPv6Addresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv6Addresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6PrefixLengths sets the value of IPv6PrefixLengths for the instance
func (instance *MSCluster_Network) SetPropertyIPv6PrefixLengths(value []string) (err error) {
	return instance.SetProperty("IPv6PrefixLengths", value)
}

// GetIPv6PrefixLengths gets the value of IPv6PrefixLengths for the instance
func (instance *MSCluster_Network) GetPropertyIPv6PrefixLengths() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv6PrefixLengths")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMetric sets the value of Metric for the instance
func (instance *MSCluster_Network) SetPropertyMetric(value uint32) (err error) {
	return instance.SetProperty("Metric", value)
}

// GetMetric gets the value of Metric for the instance
func (instance *MSCluster_Network) GetPropertyMetric() (value uint32, err error) {
	retValue, err := instance.GetProperty("Metric")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrivateProperties sets the value of PrivateProperties for the instance
func (instance *MSCluster_Network) SetPropertyPrivateProperties(value MSCluster_Property) (err error) {
	return instance.SetProperty("PrivateProperties", value)
}

// GetPrivateProperties gets the value of PrivateProperties for the instance
func (instance *MSCluster_Network) GetPropertyPrivateProperties() (value MSCluster_Property, err error) {
	retValue, err := instance.GetProperty("PrivateProperties")
	if err != nil {
		return
	}
	value, ok := retValue.(MSCluster_Property)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRole sets the value of Role for the instance
func (instance *MSCluster_Network) SetPropertyRole(value uint32) (err error) {
	return instance.SetProperty("Role", value)
}

// GetRole gets the value of Role for the instance
func (instance *MSCluster_Network) GetPropertyRole() (value uint32, err error) {
	retValue, err := instance.GetProperty("Role")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSCluster_Network) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSCluster_Network) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="newName" type="string "></param>
func (instance *MSCluster_Network) Rename( /* IN */ newName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Rename", newName)
	if err != nil {
		return
	}
	return

}

//

// <param name="ControlCode" type="int32 "></param>
// <param name="InputBuffer" type="uint8 []"></param>

// <param name="OutputBuffer" type="uint8 []"></param>
// <param name="OutputBufferSize" type="int32 "></param>
func (instance *MSCluster_Network) ExecuteNetworkControl( /* IN */ ControlCode int32,
	/* IN */ InputBuffer []uint8,
	/* OUT */ OutputBuffer []uint8,
	/* OUT */ OutputBufferSize int32) (err error) {
	_, err = instance.InvokeMethod("ExecuteNetworkControl", ControlCode, InputBuffer)
	if err != nil {
		return
	}
	return

}
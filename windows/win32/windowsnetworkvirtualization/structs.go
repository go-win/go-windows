// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsnetworkvirtualization implements the Windows.Win32.WindowsNetworkVirtualization namespace.
package windowsnetworkvirtualization

type WNV_OBJECT_HEADER struct {
	MajorVersion int
	MinorVersion int
	Size         int
}

type WNV_NOTIFICATION_PARAM struct {
	Header               int
	NotificationType     int
	PendingNotifications int
	Buffer               int
}

type WNV_IP_ADDRESS struct {
	IP int
}

type WNV_POLICY_MISMATCH_PARAM struct {
	CAFamily        int
	PAFamily        int
	VirtualSubnetId int
	CA              int
	PA              int
}

type WNV_PROVIDER_ADDRESS_CHANGE_PARAM struct {
	PAFamily     int
	PA           int
	AddressState int
}

type WNV_CUSTOMER_ADDRESS_CHANGE_PARAM struct {
	MACAddress         int
	CAFamily           int
	CA                 int
	VirtualSubnetId    int
	PAFamily           int
	PA                 int
	NotificationReason int
}

type WNV_OBJECT_CHANGE_PARAM struct {
	ObjectType  int
	ObjectParam int
}

type WNV_REDIRECT_PARAM struct {
	CAFamily        int
	PAFamily        int
	NewPAFamily     int
	VirtualSubnetId int
	CA              int
	PA              int
	NewPA           int
}

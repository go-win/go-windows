// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package networkdrivers implements the Windows.Win32.NetworkDrivers namespace.
package networkdrivers

type NET_PHYSICAL_LOCATION_LH struct {
	BusNumber      int
	SlotNumber     int
	FunctionNumber int
}

type IF_COUNTED_STRING_LH struct {
	Length int
	String int
}

type NDIS_INTERFACE_INFORMATION struct {
	IFOPERSTATUS               int
	IFOPERSTATUSFLAGS          int
	MediaConnectState          int
	MediaDuplexState           int
	IFMTU                      int
	IFPROMISCUOUSMODE          int
	IFDEVICEWAKEUPENABLE       int
	XmitLinkSpeed              int
	RcvLinkSpeed               int
	IFLASTCHANGE               int
	IFCOUNTERDISCONTINUITYTIME int
	IFINUNKNOWNPROTOS          int
	IFINDISCARDS               int
	IFINERRORS                 int
	IFHCINOCTETS               int
	IFHCINUCASTPKTS            int
	IFHCINMULTICASTPKTS        int
	IFHCINBROADCASTPKTS        int
	IFHCOUTOCTETS              int
	IFHCOUTUCASTPKTS           int
	IFHCOUTMULTICASTPKTS       int
	IFHCOUTBROADCASTPKTS       int
	IFOUTERRORS                int
	IFOUTDISCARDS              int
	IFHCINUCASTOCTETS          int
	IFHCINMULTICASTOCTETS      int
	IFHCINBROADCASTOCTETS      int
	IFHCOUTUCASTOCTETS         int
	IFHCOUTMULTICASTOCTETS     int
	IFHCOUTBROADCASTOCTETS     int
	CompartmentId              int
	SupportedStatistics        int
}

type SOCKET_ADDRESS_LIST struct {
	IADDRESSCOUNT int
	Address       int
}

type SOCKADDR_STORAGE_LH struct {
	SS_FAMILY int
	SS_PAD1   int
	SS_ALIGN  int
	SS_PAD2   int
}

type SOCKADDR_IN6_LH struct {
	SIN6_FAMILY   int
	SIN6_PORT     int
	SIN6_FLOWINFO int
	SIN6_ADDR     int
	Anonymous     int
}

type L2_NOTIFICATION_DATA struct {
	NotificationSource int
	NotificationCode   int
	InterfaceGuid      int
	DWDATASIZE         int
	PDATA              int
}

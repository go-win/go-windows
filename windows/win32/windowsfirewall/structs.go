// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsfirewall implements the Windows.Win32.WindowsFirewall namespace.
package windowsfirewall

type UPnPNAT struct {
}

type NetSharingManager struct {
}

type NETCON_PROPERTIES struct {
	GUIDID          int
	PSZWNAME        int
	PSZWDEVICENAME  int
	Status          int
	MediaType       int
	DWCHARACTER     int
	CLSIDTHISOBJECT int
	CLSIDUIOBJECT   int
}

type NetFwRule struct {
}

type NetFwOpenPort struct {
}

type NetFwAuthorizedApplication struct {
}

type NetFwPolicy2 struct {
}

type NetFwProduct struct {
}

type NetFwProducts struct {
}

type NetFwMgr struct {
}

type INET_FIREWALL_AC_CAPABILITIES struct {
	COUNT        int
	CAPABILITIES int
}

type INET_FIREWALL_AC_BINARIES struct {
	COUNT    int
	BINARIES int
}

type INET_FIREWALL_AC_CHANGE struct {
	CHANGETYPE      int
	CREATETYPE      int
	APPCONTAINERSID int
	USERSID         int
	DISPLAYNAME     int
	Anonymous       int
}

type INET_FIREWALL_APP_CONTAINER struct {
	APPCONTAINERSID  int
	USERSID          int
	APPCONTAINERNAME int
	DISPLAYNAME      int
	DESCRIPTION      int
	CAPABILITIES     int
	BINARIES         int
	WORKINGDIRECTORY int
	PACKAGEFULLNAME  int
}

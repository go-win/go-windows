// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package routingandremoteaccessservice implements the Windows.Win32.RoutingAndRemoteAccessService namespace.
package routingandremoteaccessservice

type RASAPIVERSION int32

const (
	RASAPIVERSION_500 = 1
	RASAPIVERSION_501 = 2
	RASAPIVERSION_600 = 3
	RASAPIVERSION_601 = 4
)

type tagRASCONNSTATE int32

const (
	RASCS_OpenPort = 0
	RASCS_PortOpened = 1
	RASCS_ConnectDevice = 2
	RASCS_DeviceConnected = 3
	RASCS_AllDevicesConnected = 4
	RASCS_Authenticate = 5
	RASCS_AuthNotify = 6
	RASCS_AuthRetry = 7
	RASCS_AuthCallback = 8
	RASCS_AuthChangePassword = 9
	RASCS_AuthProject = 10
	RASCS_AuthLinkSpeed = 11
	RASCS_AuthAck = 12
	RASCS_ReAuthenticate = 13
	RASCS_Authenticated = 14
	RASCS_PrepareForCallback = 15
	RASCS_WaitForModemReset = 16
	RASCS_WaitForCallback = 17
	RASCS_Projected = 18
	RASCS_StartAuthentication = 19
	RASCS_CallbackComplete = 20
	RASCS_LogonNetwork = 21
	RASCS_SubEntryConnected = 22
	RASCS_SubEntryDisconnected = 23
	RASCS_ApplySettings = 24
	RASCS_Interactive = 4096
	RASCS_RetryAuthentication = 4097
	RASCS_CallbackSetByCaller = 4098
	RASCS_PasswordExpired = 4099
	RASCS_InvokeEapUI = 4100
	RASCS_Connected = 8192
	RASCS_Disconnected = 8193
)

type tagRASCONNSUBSTATE int32

const (
	RASCSS_None = 0
	RASCSS_Dormant = 1
	RASCSS_Reconnecting = 2
	RASCSS_Reconnected = 8192
)

type tagRASPROJECTION int32

const (
	RASP_Amb = 65536
	RASP_PppNbf = 32831
	RASP_PppIpx = 32811
	RASP_PppIp = 32801
	RASP_PppCcp = 33021
	RASP_PppLcp = 49185
	RASP_PppIpv6 = 32855
)

type RASPROJECTION_INFO_TYPE int32

const (
	PROJECTION_INFO_TYPE_PPP = 1
	PROJECTION_INFO_TYPE_IKEv2 = 2
)

type IKEV2_ID_PAYLOAD_TYPE int32

const (
	IKEV2_ID_PAYLOAD_TYPE_INVALID = 0
	IKEV2_ID_PAYLOAD_TYPE_IPV4_ADDR = 1
	IKEV2_ID_PAYLOAD_TYPE_FQDN = 2
	IKEV2_ID_PAYLOAD_TYPE_RFC822_ADDR = 3
	IKEV2_ID_PAYLOAD_TYPE_RESERVED1 = 4
	IKEV2_ID_PAYLOAD_TYPE_ID_IPV6_ADDR = 5
	IKEV2_ID_PAYLOAD_TYPE_RESERVED2 = 6
	IKEV2_ID_PAYLOAD_TYPE_RESERVED3 = 7
	IKEV2_ID_PAYLOAD_TYPE_RESERVED4 = 8
	IKEV2_ID_PAYLOAD_TYPE_DER_ASN1_DN = 9
	IKEV2_ID_PAYLOAD_TYPE_DER_ASN1_GN = 10
	IKEV2_ID_PAYLOAD_TYPE_KEY_ID = 11
	IKEV2_ID_PAYLOAD_TYPE_MAX = 12
)

type ROUTER_INTERFACE_TYPE int32

const (
	ROUTER_IF_TYPE_CLIENT = 0
	ROUTER_IF_TYPE_HOME_ROUTER = 1
	ROUTER_IF_TYPE_FULL_ROUTER = 2
	ROUTER_IF_TYPE_DEDICATED = 3
	ROUTER_IF_TYPE_INTERNAL = 4
	ROUTER_IF_TYPE_LOOPBACK = 5
	ROUTER_IF_TYPE_TUNNEL1 = 6
	ROUTER_IF_TYPE_DIALOUT = 7
	ROUTER_IF_TYPE_MAX = 8
)

type ROUTER_CONNECTION_STATE int32

const (
	ROUTER_IF_STATE_UNREACHABLE = 0
	ROUTER_IF_STATE_DISCONNECTED = 1
	ROUTER_IF_STATE_CONNECTING = 2
	ROUTER_IF_STATE_CONNECTED = 3
)

type RAS_PORT_CONDITION int32

const (
	RAS_PORT_NON_OPERATIONAL = 0
	RAS_PORT_DISCONNECTED = 1
	RAS_PORT_CALLING_BACK = 2
	RAS_PORT_LISTENING = 3
	RAS_PORT_AUTHENTICATING = 4
	RAS_PORT_AUTHENTICATED = 5
	RAS_PORT_INITIALIZING = 6
)

type RAS_HARDWARE_CONDITION int32

const (
	RAS_HARDWARE_OPERATIONAL = 0
	RAS_HARDWARE_FAILURE = 1
)

type RAS_QUARANTINE_STATE int32

const (
	RAS_QUAR_STATE_NORMAL = 0
	RAS_QUAR_STATE_QUARANTINE = 1
	RAS_QUAR_STATE_PROBATION = 2
	RAS_QUAR_STATE_NOT_CAPABLE = 3
)

type MPRAPI_OBJECT_TYPE int32

const (
	MPRAPI_OBJECT_TYPE_RAS_CONNECTION_OBJECT = 1
	MPRAPI_OBJECT_TYPE_MPR_SERVER_OBJECT = 2
	MPRAPI_OBJECT_TYPE_MPR_SERVER_SET_CONFIG_OBJECT = 3
	MPRAPI_OBJECT_TYPE_AUTH_VALIDATION_OBJECT = 4
	MPRAPI_OBJECT_TYPE_UPDATE_CONNECTION_OBJECT = 5
	MPRAPI_OBJECT_TYPE_IF_CUSTOM_CONFIG_OBJECT = 6
)

type MPR_VPN_TS_TYPE int32

const (
	MPR_VPN_TS_IPv4_ADDR_RANGE = 7
	MPR_VPN_TS_IPv6_ADDR_RANGE = 8
)

type MGM_ENUM_TYPES int32

const (
	ANY_SOURCE = 0
	ALL_SOURCES = 1
)

type RTM_EVENT_TYPE int32

const (
	RTM_ENTITY_REGISTERED = 0
	RTM_ENTITY_DEREGISTERED = 1
	RTM_ROUTE_EXPIRED = 2
	RTM_CHANGE_NOTIFICATION = 3
)

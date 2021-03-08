// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package winsock implements the Windows.Win32.WinSock namespace.
package winsock

type IPPROTO int32

const (
	IPPROTO_HOPOPTS               IPPROTO = 0
	IPPROTO_ICMP                  IPPROTO = 1
	IPPROTO_IGMP                  IPPROTO = 2
	IPPROTO_GGP                   IPPROTO = 3
	IPPROTO_IPV4                  IPPROTO = 4
	IPPROTO_ST                    IPPROTO = 5
	IPPROTO_TCP                   IPPROTO = 6
	IPPROTO_CBT                   IPPROTO = 7
	IPPROTO_EGP                   IPPROTO = 8
	IPPROTO_IGP                   IPPROTO = 9
	IPPROTO_PUP                   IPPROTO = 12
	IPPROTO_UDP                   IPPROTO = 17
	IPPROTO_IDP                   IPPROTO = 22
	IPPROTO_RDP                   IPPROTO = 27
	IPPROTO_IPV6                  IPPROTO = 41
	IPPROTO_ROUTING               IPPROTO = 43
	IPPROTO_FRAGMENT              IPPROTO = 44
	IPPROTO_ESP                   IPPROTO = 50
	IPPROTO_AH                    IPPROTO = 51
	IPPROTO_ICMPV6                IPPROTO = 58
	IPPROTO_NONE                  IPPROTO = 59
	IPPROTO_DSTOPTS               IPPROTO = 60
	IPPROTO_ND                    IPPROTO = 77
	IPPROTO_ICLFXBM               IPPROTO = 78
	IPPROTO_PIM                   IPPROTO = 103
	IPPROTO_PGM                   IPPROTO = 113
	IPPROTO_L2TP                  IPPROTO = 115
	IPPROTO_SCTP                  IPPROTO = 132
	IPPROTO_RAW                   IPPROTO = 255
	IPPROTO_MAX                   IPPROTO = 256
	IPPROTO_RESERVED_RAW          IPPROTO = 257
	IPPROTO_RESERVED_IPSEC        IPPROTO = 258
	IPPROTO_RESERVED_IPSECOFFLOAD IPPROTO = 259
	IPPROTO_RESERVED_WNV          IPPROTO = 260
	IPPROTO_RESERVED_MAX          IPPROTO = 261
)

type WSACOMPLETIONTYPE int32

const (
	NSP_NOTIFY_IMMEDIATELY WSACOMPLETIONTYPE = 0
	NSP_NOTIFY_HWND        WSACOMPLETIONTYPE = 1
	NSP_NOTIFY_EVENT       WSACOMPLETIONTYPE = 2
	NSP_NOTIFY_PORT        WSACOMPLETIONTYPE = 3
	NSP_NOTIFY_APC         WSACOMPLETIONTYPE = 4
)

type WSAECOMPARATOR int32

const (
	COMP_EQUAL   WSAECOMPARATOR = 0
	COMP_NOTLESS WSAECOMPARATOR = 1
)

type WSAESETSERVICEOP int32

const (
	RNRSERVICE_REGISTER   WSAESETSERVICEOP = 0
	RNRSERVICE_DEREGISTER WSAESETSERVICEOP = 1
	RNRSERVICE_DELETE     WSAESETSERVICEOP = 2
)

type PMTUD_STATE int32

const (
	IP_PMTUDISC_NOT_SET PMTUD_STATE = 0
	IP_PMTUDISC_DO      PMTUD_STATE = 1
	IP_PMTUDISC_DONT    PMTUD_STATE = 2
	IP_PMTUDISC_PROBE   PMTUD_STATE = 3
	IP_PMTUDISC_MAX     PMTUD_STATE = 4
)

type MULTICAST_MODE_TYPE int32

const (
	MCAST_INCLUDE MULTICAST_MODE_TYPE = 0
	MCAST_EXCLUDE MULTICAST_MODE_TYPE = 1
)

type eWINDOW_ADVANCE_METHOD int32

const (
	E_WINDOW_ADVANCE_BY_TIME   eWINDOW_ADVANCE_METHOD = 1
	E_WINDOW_USE_AS_DATA_CACHE eWINDOW_ADVANCE_METHOD = 2
)

type NL_BANDWIDTH_FLAG int32

const (
	NlbwDisabled  NL_BANDWIDTH_FLAG = 0
	NlbwEnabled   NL_BANDWIDTH_FLAG = 1
	NlbwUnchanged NL_BANDWIDTH_FLAG = -1
)

type NL_NETWORK_CATEGORY int32

const (
	NetworkCategoryPublic              NL_NETWORK_CATEGORY = 0
	NetworkCategoryPrivate             NL_NETWORK_CATEGORY = 1
	NetworkCategoryDomainAuthenticated NL_NETWORK_CATEGORY = 2
	NetworkCategoryUnchanged           NL_NETWORK_CATEGORY = -1
	NetworkCategoryUnknown             NL_NETWORK_CATEGORY = -1
)

type NL_INTERFACE_NETWORK_CATEGORY_STATE int32

const (
	NlincCategoryUnknown     NL_INTERFACE_NETWORK_CATEGORY_STATE = 0
	NlincPublic              NL_INTERFACE_NETWORK_CATEGORY_STATE = 1
	NlincPrivate             NL_INTERFACE_NETWORK_CATEGORY_STATE = 2
	NlincDomainAuthenticated NL_INTERFACE_NETWORK_CATEGORY_STATE = 3
	NlincCategoryStateMax    NL_INTERFACE_NETWORK_CATEGORY_STATE = 4
)

type TCPSTATE int32

const (
	TCPSTATE_CLOSED      TCPSTATE = 0
	TCPSTATE_LISTEN      TCPSTATE = 1
	TCPSTATE_SYN_SENT    TCPSTATE = 2
	TCPSTATE_SYN_RCVD    TCPSTATE = 3
	TCPSTATE_ESTABLISHED TCPSTATE = 4
	TCPSTATE_FIN_WAIT_1  TCPSTATE = 5
	TCPSTATE_FIN_WAIT_2  TCPSTATE = 6
	TCPSTATE_CLOSE_WAIT  TCPSTATE = 7
	TCPSTATE_CLOSING     TCPSTATE = 8
	TCPSTATE_LAST_ACK    TCPSTATE = 9
	TCPSTATE_TIME_WAIT   TCPSTATE = 10
	TCPSTATE_MAX         TCPSTATE = 11
)

type CONTROL_CHANNEL_TRIGGER_STATUS int32

const (
	CONTROL_CHANNEL_TRIGGER_STATUS_INVALID                 CONTROL_CHANNEL_TRIGGER_STATUS = 0
	CONTROL_CHANNEL_TRIGGER_STATUS_SOFTWARE_SLOT_ALLOCATED CONTROL_CHANNEL_TRIGGER_STATUS = 1
	CONTROL_CHANNEL_TRIGGER_STATUS_HARDWARE_SLOT_ALLOCATED CONTROL_CHANNEL_TRIGGER_STATUS = 2
	CONTROL_CHANNEL_TRIGGER_STATUS_POLICY_ERROR            CONTROL_CHANNEL_TRIGGER_STATUS = 3
	CONTROL_CHANNEL_TRIGGER_STATUS_SYSTEM_ERROR            CONTROL_CHANNEL_TRIGGER_STATUS = 4
	CONTROL_CHANNEL_TRIGGER_STATUS_TRANSPORT_DISCONNECTED  CONTROL_CHANNEL_TRIGGER_STATUS = 5
	CONTROL_CHANNEL_TRIGGER_STATUS_SERVICE_UNAVAILABLE     CONTROL_CHANNEL_TRIGGER_STATUS = 6
)

type RCVALL_VALUE int32

const (
	RCVALL_OFF             RCVALL_VALUE = 0
	RCVALL_ON              RCVALL_VALUE = 1
	RCVALL_SOCKETLEVELONLY RCVALL_VALUE = 2
	RCVALL_IPLEVEL         RCVALL_VALUE = 3
)

type TCP_ICW_LEVEL int32

const (
	TCP_ICW_LEVEL_DEFAULT      TCP_ICW_LEVEL = 0
	TCP_ICW_LEVEL_HIGH         TCP_ICW_LEVEL = 1
	TCP_ICW_LEVEL_VERY_HIGH    TCP_ICW_LEVEL = 2
	TCP_ICW_LEVEL_AGGRESSIVE   TCP_ICW_LEVEL = 3
	TCP_ICW_LEVEL_EXPERIMENTAL TCP_ICW_LEVEL = 4
	TCP_ICW_LEVEL_COMPAT       TCP_ICW_LEVEL = 254
	TCP_ICW_LEVEL_MAX          TCP_ICW_LEVEL = 255
)

type SOCKET_USAGE_TYPE int32

const (
	SYSTEM_CRITICAL_SOCKET SOCKET_USAGE_TYPE = 1
)

type SOCKET_SECURITY_PROTOCOL int32

const (
	SOCKET_SECURITY_PROTOCOL_DEFAULT SOCKET_SECURITY_PROTOCOL = 0
	SOCKET_SECURITY_PROTOCOL_IPSEC   SOCKET_SECURITY_PROTOCOL = 1
	SOCKET_SECURITY_PROTOCOL_IPSEC2  SOCKET_SECURITY_PROTOCOL = 2
	SOCKET_SECURITY_PROTOCOL_INVALID SOCKET_SECURITY_PROTOCOL = 3
)

type WSA_COMPATIBILITY_BEHAVIOR_ID int32

const (
	WsaBehaviorAll              WSA_COMPATIBILITY_BEHAVIOR_ID = 0
	WsaBehaviorReceiveBuffering WSA_COMPATIBILITY_BEHAVIOR_ID = 1
	WsaBehaviorAutoTuning       WSA_COMPATIBILITY_BEHAVIOR_ID = 2
)

type Q2931_IE_TYPE int32

const (
	IE_AALParameters             Q2931_IE_TYPE = 0
	IE_TrafficDescriptor         Q2931_IE_TYPE = 1
	IE_BroadbandBearerCapability Q2931_IE_TYPE = 2
	IE_BHLI                      Q2931_IE_TYPE = 3
	IE_BLLI                      Q2931_IE_TYPE = 4
	IE_CalledPartyNumber         Q2931_IE_TYPE = 5
	IE_CalledPartySubaddress     Q2931_IE_TYPE = 6
	IE_CallingPartyNumber        Q2931_IE_TYPE = 7
	IE_CallingPartySubaddress    Q2931_IE_TYPE = 8
	IE_Cause                     Q2931_IE_TYPE = 9
	IE_QOSClass                  Q2931_IE_TYPE = 10
	IE_TransitNetworkSelection   Q2931_IE_TYPE = 11
)

type AAL_TYPE int32

const (
	AALTYPE_5    AAL_TYPE = 5
	AALTYPE_USER AAL_TYPE = 16
)

type NAPI_PROVIDER_TYPE int32

const (
	ProviderType_Application NAPI_PROVIDER_TYPE = 1
	ProviderType_Service     NAPI_PROVIDER_TYPE = 2
)

type NAPI_PROVIDER_LEVEL int32

const (
	ProviderLevel_None      NAPI_PROVIDER_LEVEL = 0
	ProviderLevel_Secondary NAPI_PROVIDER_LEVEL = 1
	ProviderLevel_Primary   NAPI_PROVIDER_LEVEL = 2
)

type NLA_BLOB_DATA_TYPE int32

const (
	NLA_RAW_DATA        NLA_BLOB_DATA_TYPE = 0
	NLA_INTERFACE       NLA_BLOB_DATA_TYPE = 1
	NLA_802_1X_LOCATION NLA_BLOB_DATA_TYPE = 2
	NLA_CONNECTIVITY    NLA_BLOB_DATA_TYPE = 3
	NLA_ICS             NLA_BLOB_DATA_TYPE = 4
)

type NLA_CONNECTIVITY_TYPE int32

const (
	NLA_NETWORK_AD_HOC    NLA_CONNECTIVITY_TYPE = 0
	NLA_NETWORK_MANAGED   NLA_CONNECTIVITY_TYPE = 1
	NLA_NETWORK_UNMANAGED NLA_CONNECTIVITY_TYPE = 2
	NLA_NETWORK_UNKNOWN   NLA_CONNECTIVITY_TYPE = 3
)

type NLA_INTERNET int32

const (
	NLA_INTERNET_UNKNOWN NLA_INTERNET = 0
	NLA_INTERNET_NO      NLA_INTERNET = 1
	NLA_INTERNET_YES     NLA_INTERNET = 2
)

type RIO_NOTIFICATION_COMPLETION_TYPE int32

const (
	RIO_EVENT_COMPLETION RIO_NOTIFICATION_COMPLETION_TYPE = 1
	RIO_IOCP_COMPLETION  RIO_NOTIFICATION_COMPLETION_TYPE = 2
)

type WSC_PROVIDER_INFO_TYPE int32

const (
	ProviderInfoLspCategories WSC_PROVIDER_INFO_TYPE = 0
	ProviderInfoAudit         WSC_PROVIDER_INFO_TYPE = 1
)

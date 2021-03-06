// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package networkdrivers implements the Windows.Win32.NetworkDrivers namespace.
package networkdrivers

type NET_IF_OPER_STATUS int32

const (
	NET_IF_OPER_STATUS_UP               NET_IF_OPER_STATUS = 1
	NET_IF_OPER_STATUS_DOWN             NET_IF_OPER_STATUS = 2
	NET_IF_OPER_STATUS_TESTING          NET_IF_OPER_STATUS = 3
	NET_IF_OPER_STATUS_UNKNOWN          NET_IF_OPER_STATUS = 4
	NET_IF_OPER_STATUS_DORMANT          NET_IF_OPER_STATUS = 5
	NET_IF_OPER_STATUS_NOT_PRESENT      NET_IF_OPER_STATUS = 6
	NET_IF_OPER_STATUS_LOWER_LAYER_DOWN NET_IF_OPER_STATUS = 7
)

type NET_IF_ADMIN_STATUS int32

const (
	NET_IF_ADMIN_STATUS_UP      NET_IF_ADMIN_STATUS = 1
	NET_IF_ADMIN_STATUS_DOWN    NET_IF_ADMIN_STATUS = 2
	NET_IF_ADMIN_STATUS_TESTING NET_IF_ADMIN_STATUS = 3
)

type NET_IF_CONNECTION_TYPE int32

const (
	NET_IF_CONNECTION_DEDICATED NET_IF_CONNECTION_TYPE = 1
	NET_IF_CONNECTION_PASSIVE   NET_IF_CONNECTION_TYPE = 2
	NET_IF_CONNECTION_DEMAND    NET_IF_CONNECTION_TYPE = 3
	NET_IF_CONNECTION_MAXIMUM   NET_IF_CONNECTION_TYPE = 4
)

type TUNNEL_TYPE int32

const (
	TUNNEL_TYPE_NONE    TUNNEL_TYPE = 0
	TUNNEL_TYPE_OTHER   TUNNEL_TYPE = 1
	TUNNEL_TYPE_DIRECT  TUNNEL_TYPE = 2
	TUNNEL_TYPE_6TO4    TUNNEL_TYPE = 11
	TUNNEL_TYPE_ISATAP  TUNNEL_TYPE = 13
	TUNNEL_TYPE_TEREDO  TUNNEL_TYPE = 14
	TUNNEL_TYPE_IPHTTPS TUNNEL_TYPE = 15
)

type NET_IF_ACCESS_TYPE int32

const (
	NET_IF_ACCESS_LOOPBACK             NET_IF_ACCESS_TYPE = 1
	NET_IF_ACCESS_BROADCAST            NET_IF_ACCESS_TYPE = 2
	NET_IF_ACCESS_POINT_TO_POINT       NET_IF_ACCESS_TYPE = 3
	NET_IF_ACCESS_POINT_TO_MULTI_POINT NET_IF_ACCESS_TYPE = 4
	NET_IF_ACCESS_MAXIMUM              NET_IF_ACCESS_TYPE = 5
)

type NET_IF_DIRECTION_TYPE int32

const (
	NET_IF_DIRECTION_SENDRECEIVE NET_IF_DIRECTION_TYPE = 0
	NET_IF_DIRECTION_SENDONLY    NET_IF_DIRECTION_TYPE = 1
	NET_IF_DIRECTION_RECEIVEONLY NET_IF_DIRECTION_TYPE = 2
	NET_IF_DIRECTION_MAXIMUM     NET_IF_DIRECTION_TYPE = 3
)

type NET_IF_MEDIA_CONNECT_STATE int32

const (
	MediaConnectStateUnknown      NET_IF_MEDIA_CONNECT_STATE = 0
	MediaConnectStateConnected    NET_IF_MEDIA_CONNECT_STATE = 1
	MediaConnectStateDisconnected NET_IF_MEDIA_CONNECT_STATE = 2
)

type NET_IF_MEDIA_DUPLEX_STATE int32

const (
	MediaDuplexStateUnknown NET_IF_MEDIA_DUPLEX_STATE = 0
	MediaDuplexStateHalf    NET_IF_MEDIA_DUPLEX_STATE = 1
	MediaDuplexStateFull    NET_IF_MEDIA_DUPLEX_STATE = 2
)

type MIB_IF_TABLE_LEVEL int32

const (
	MibIfTableNormal                  MIB_IF_TABLE_LEVEL = 0
	MibIfTableRaw                     MIB_IF_TABLE_LEVEL = 1
	MibIfTableNormalWithoutStatistics MIB_IF_TABLE_LEVEL = 2
)

type NL_ROUTE_PROTOCOL int32

const (
	RouteProtocolOther            NL_ROUTE_PROTOCOL = 1
	RouteProtocolLocal            NL_ROUTE_PROTOCOL = 2
	RouteProtocolNetMgmt          NL_ROUTE_PROTOCOL = 3
	RouteProtocolIcmp             NL_ROUTE_PROTOCOL = 4
	RouteProtocolEgp              NL_ROUTE_PROTOCOL = 5
	RouteProtocolGgp              NL_ROUTE_PROTOCOL = 6
	RouteProtocolHello            NL_ROUTE_PROTOCOL = 7
	RouteProtocolRip              NL_ROUTE_PROTOCOL = 8
	RouteProtocolIsIs             NL_ROUTE_PROTOCOL = 9
	RouteProtocolEsIs             NL_ROUTE_PROTOCOL = 10
	RouteProtocolCisco            NL_ROUTE_PROTOCOL = 11
	RouteProtocolBbn              NL_ROUTE_PROTOCOL = 12
	RouteProtocolOspf             NL_ROUTE_PROTOCOL = 13
	RouteProtocolBgp              NL_ROUTE_PROTOCOL = 14
	RouteProtocolIdpr             NL_ROUTE_PROTOCOL = 15
	RouteProtocolEigrp            NL_ROUTE_PROTOCOL = 16
	RouteProtocolDvmrp            NL_ROUTE_PROTOCOL = 17
	RouteProtocolRpl              NL_ROUTE_PROTOCOL = 18
	RouteProtocolDhcp             NL_ROUTE_PROTOCOL = 19
	MIB_IPPROTO_OTHER             NL_ROUTE_PROTOCOL = 1
	PROTO_IP_OTHER                NL_ROUTE_PROTOCOL = 1
	MIB_IPPROTO_LOCAL             NL_ROUTE_PROTOCOL = 2
	PROTO_IP_LOCAL                NL_ROUTE_PROTOCOL = 2
	MIB_IPPROTO_NETMGMT           NL_ROUTE_PROTOCOL = 3
	PROTO_IP_NETMGMT              NL_ROUTE_PROTOCOL = 3
	MIB_IPPROTO_ICMP              NL_ROUTE_PROTOCOL = 4
	PROTO_IP_ICMP                 NL_ROUTE_PROTOCOL = 4
	MIB_IPPROTO_EGP               NL_ROUTE_PROTOCOL = 5
	PROTO_IP_EGP                  NL_ROUTE_PROTOCOL = 5
	MIB_IPPROTO_GGP               NL_ROUTE_PROTOCOL = 6
	PROTO_IP_GGP                  NL_ROUTE_PROTOCOL = 6
	MIB_IPPROTO_HELLO             NL_ROUTE_PROTOCOL = 7
	PROTO_IP_HELLO                NL_ROUTE_PROTOCOL = 7
	MIB_IPPROTO_RIP               NL_ROUTE_PROTOCOL = 8
	PROTO_IP_RIP                  NL_ROUTE_PROTOCOL = 8
	MIB_IPPROTO_IS_IS             NL_ROUTE_PROTOCOL = 9
	PROTO_IP_IS_IS                NL_ROUTE_PROTOCOL = 9
	MIB_IPPROTO_ES_IS             NL_ROUTE_PROTOCOL = 10
	PROTO_IP_ES_IS                NL_ROUTE_PROTOCOL = 10
	MIB_IPPROTO_CISCO             NL_ROUTE_PROTOCOL = 11
	PROTO_IP_CISCO                NL_ROUTE_PROTOCOL = 11
	MIB_IPPROTO_BBN               NL_ROUTE_PROTOCOL = 12
	PROTO_IP_BBN                  NL_ROUTE_PROTOCOL = 12
	MIB_IPPROTO_OSPF              NL_ROUTE_PROTOCOL = 13
	PROTO_IP_OSPF                 NL_ROUTE_PROTOCOL = 13
	MIB_IPPROTO_BGP               NL_ROUTE_PROTOCOL = 14
	PROTO_IP_BGP                  NL_ROUTE_PROTOCOL = 14
	MIB_IPPROTO_IDPR              NL_ROUTE_PROTOCOL = 15
	PROTO_IP_IDPR                 NL_ROUTE_PROTOCOL = 15
	MIB_IPPROTO_EIGRP             NL_ROUTE_PROTOCOL = 16
	PROTO_IP_EIGRP                NL_ROUTE_PROTOCOL = 16
	MIB_IPPROTO_DVMRP             NL_ROUTE_PROTOCOL = 17
	PROTO_IP_DVMRP                NL_ROUTE_PROTOCOL = 17
	MIB_IPPROTO_RPL               NL_ROUTE_PROTOCOL = 18
	PROTO_IP_RPL                  NL_ROUTE_PROTOCOL = 18
	MIB_IPPROTO_DHCP              NL_ROUTE_PROTOCOL = 19
	PROTO_IP_DHCP                 NL_ROUTE_PROTOCOL = 19
	MIB_IPPROTO_NT_AUTOSTATIC     NL_ROUTE_PROTOCOL = 10002
	PROTO_IP_NT_AUTOSTATIC        NL_ROUTE_PROTOCOL = 10002
	MIB_IPPROTO_NT_STATIC         NL_ROUTE_PROTOCOL = 10006
	PROTO_IP_NT_STATIC            NL_ROUTE_PROTOCOL = 10006
	MIB_IPPROTO_NT_STATIC_NON_DOD NL_ROUTE_PROTOCOL = 10007
	PROTO_IP_NT_STATIC_NON_DOD    NL_ROUTE_PROTOCOL = 10007
)

type NL_ADDRESS_TYPE int32

const (
	NlatUnspecified NL_ADDRESS_TYPE = 0
	NlatUnicast     NL_ADDRESS_TYPE = 1
	NlatAnycast     NL_ADDRESS_TYPE = 2
	NlatMulticast   NL_ADDRESS_TYPE = 3
	NlatBroadcast   NL_ADDRESS_TYPE = 4
	NlatInvalid     NL_ADDRESS_TYPE = 5
)

type NL_ROUTE_ORIGIN int32

const (
	NlroManual              NL_ROUTE_ORIGIN = 0
	NlroWellKnown           NL_ROUTE_ORIGIN = 1
	NlroDHCP                NL_ROUTE_ORIGIN = 2
	NlroRouterAdvertisement NL_ROUTE_ORIGIN = 3
	Nlro6to4                NL_ROUTE_ORIGIN = 4
)

type NL_NEIGHBOR_STATE int32

const (
	NlnsUnreachable NL_NEIGHBOR_STATE = 0
	NlnsIncomplete  NL_NEIGHBOR_STATE = 1
	NlnsProbe       NL_NEIGHBOR_STATE = 2
	NlnsDelay       NL_NEIGHBOR_STATE = 3
	NlnsStale       NL_NEIGHBOR_STATE = 4
	NlnsReachable   NL_NEIGHBOR_STATE = 5
	NlnsPermanent   NL_NEIGHBOR_STATE = 6
	NlnsMaximum     NL_NEIGHBOR_STATE = 7
)

type NL_LINK_LOCAL_ADDRESS_BEHAVIOR int32

const (
	LinkLocalAlwaysOff NL_LINK_LOCAL_ADDRESS_BEHAVIOR = 0
	LinkLocalDelayed   NL_LINK_LOCAL_ADDRESS_BEHAVIOR = 1
	LinkLocalAlwaysOn  NL_LINK_LOCAL_ADDRESS_BEHAVIOR = 2
	LinkLocalUnchanged NL_LINK_LOCAL_ADDRESS_BEHAVIOR = -1
)

type NL_ROUTER_DISCOVERY_BEHAVIOR int32

const (
	RouterDiscoveryDisabled  NL_ROUTER_DISCOVERY_BEHAVIOR = 0
	RouterDiscoveryEnabled   NL_ROUTER_DISCOVERY_BEHAVIOR = 1
	RouterDiscoveryDhcp      NL_ROUTER_DISCOVERY_BEHAVIOR = 2
	RouterDiscoveryUnchanged NL_ROUTER_DISCOVERY_BEHAVIOR = -1
)

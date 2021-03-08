// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package mib implements the Windows.Win32.Mib namespace.
package mib

type MIB_NOTIFICATION_TYPE int32

const (
	MibParameterNotification = 0
	MibAddInstance = 1
	MibDeleteInstance = 2
	MibInitialNotification = 3
)

type ICMP6_TYPE int32

const (
	ICMP6_DST_UNREACH = 1
	ICMP6_PACKET_TOO_BIG = 2
	ICMP6_TIME_EXCEEDED = 3
	ICMP6_PARAM_PROB = 4
	ICMP6_ECHO_REQUEST = 128
	ICMP6_ECHO_REPLY = 129
	ICMP6_MEMBERSHIP_QUERY = 130
	ICMP6_MEMBERSHIP_REPORT = 131
	ICMP6_MEMBERSHIP_REDUCTION = 132
	ND_ROUTER_SOLICIT = 133
	ND_ROUTER_ADVERT = 134
	ND_NEIGHBOR_SOLICIT = 135
	ND_NEIGHBOR_ADVERT = 136
	ND_REDIRECT = 137
	ICMP6_V2_MEMBERSHIP_REPORT = 143
)

type ICMP4_TYPE int32

const (
	ICMP4_ECHO_REPLY = 0
	ICMP4_DST_UNREACH = 3
	ICMP4_SOURCE_QUENCH = 4
	ICMP4_REDIRECT = 5
	ICMP4_ECHO_REQUEST = 8
	ICMP4_ROUTER_ADVERT = 9
	ICMP4_ROUTER_SOLICIT = 10
	ICMP4_TIME_EXCEEDED = 11
	ICMP4_PARAM_PROB = 12
	ICMP4_TIMESTAMP_REQUEST = 13
	ICMP4_TIMESTAMP_REPLY = 14
	ICMP4_MASK_REQUEST = 17
	ICMP4_MASK_REPLY = 18
)

type TCP_CONNECTION_OFFLOAD_STATE int32

const (
	TcpConnectionOffloadStateInHost = 0
	TcpConnectionOffloadStateOffloading = 1
	TcpConnectionOffloadStateOffloaded = 2
	TcpConnectionOffloadStateUploading = 3
	TcpConnectionOffloadStateMax = 4
)

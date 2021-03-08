// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package websock implements the Windows.Win32.WebSock namespace.
package websock

type WEB_SOCKET_CLOSE_STATUS int32

const (
	WEB_SOCKET_SUCCESS_CLOSE_STATUS = 1000
	WEB_SOCKET_ENDPOINT_UNAVAILABLE_CLOSE_STATUS = 1001
	WEB_SOCKET_PROTOCOL_ERROR_CLOSE_STATUS = 1002
	WEB_SOCKET_INVALID_DATA_TYPE_CLOSE_STATUS = 1003
	WEB_SOCKET_EMPTY_CLOSE_STATUS = 1005
	WEB_SOCKET_ABORTED_CLOSE_STATUS = 1006
	WEB_SOCKET_INVALID_PAYLOAD_CLOSE_STATUS = 1007
	WEB_SOCKET_POLICY_VIOLATION_CLOSE_STATUS = 1008
	WEB_SOCKET_MESSAGE_TOO_BIG_CLOSE_STATUS = 1009
	WEB_SOCKET_UNSUPPORTED_EXTENSIONS_CLOSE_STATUS = 1010
	WEB_SOCKET_SERVER_ERROR_CLOSE_STATUS = 1011
	WEB_SOCKET_SECURE_HANDSHAKE_ERROR_CLOSE_STATUS = 1015
)

type WEB_SOCKET_PROPERTY_TYPE int32

const (
	WEB_SOCKET_RECEIVE_BUFFER_SIZE_PROPERTY_TYPE = 0
	WEB_SOCKET_SEND_BUFFER_SIZE_PROPERTY_TYPE = 1
	WEB_SOCKET_DISABLE_MASKING_PROPERTY_TYPE = 2
	WEB_SOCKET_ALLOCATED_BUFFER_PROPERTY_TYPE = 3
	WEB_SOCKET_DISABLE_UTF8_VERIFICATION_PROPERTY_TYPE = 4
	WEB_SOCKET_KEEPALIVE_INTERVAL_PROPERTY_TYPE = 5
	WEB_SOCKET_SUPPORTED_VERSIONS_PROPERTY_TYPE = 6
)

type WEB_SOCKET_ACTION_QUEUE int32

const (
	WEB_SOCKET_SEND_ACTION_QUEUE = 1
	WEB_SOCKET_RECEIVE_ACTION_QUEUE = 2
	WEB_SOCKET_ALL_ACTION_QUEUE = 3
)

type WEB_SOCKET_BUFFER_TYPE int32

const (
	WEB_SOCKET_UTF8_MESSAGE_BUFFER_TYPE = -2147483648
	WEB_SOCKET_UTF8_FRAGMENT_BUFFER_TYPE = -2147483647
	WEB_SOCKET_BINARY_MESSAGE_BUFFER_TYPE = -2147483646
	WEB_SOCKET_BINARY_FRAGMENT_BUFFER_TYPE = -2147483645
	WEB_SOCKET_CLOSE_BUFFER_TYPE = -2147483644
	WEB_SOCKET_PING_PONG_BUFFER_TYPE = -2147483643
	WEB_SOCKET_UNSOLICITED_PONG_BUFFER_TYPE = -2147483642
)

type WEB_SOCKET_ACTION int32

const (
	WEB_SOCKET_NO_ACTION = 0
	WEB_SOCKET_SEND_TO_NETWORK_ACTION = 1
	WEB_SOCKET_INDICATE_SEND_COMPLETE_ACTION = 2
	WEB_SOCKET_RECEIVE_FROM_NETWORK_ACTION = 3
	WEB_SOCKET_INDICATE_RECEIVE_COMPLETE_ACTION = 4
)

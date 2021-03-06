// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsremotemanagement implements the Windows.Win32.WindowsRemoteManagement namespace.
package windowsremotemanagement

type WSManDataType int32

const (
	WSMAN_DATA_NONE        WSManDataType = 0
	WSMAN_DATA_TYPE_TEXT   WSManDataType = 1
	WSMAN_DATA_TYPE_BINARY WSManDataType = 2
	WSMAN_DATA_TYPE_DWORD  WSManDataType = 4
)

type WSManAuthenticationFlags int32

const (
	WSMAN_FLAG_DEFAULT_AUTHENTICATION  WSManAuthenticationFlags = 0
	WSMAN_FLAG_NO_AUTHENTICATION       WSManAuthenticationFlags = 1
	WSMAN_FLAG_AUTH_DIGEST             WSManAuthenticationFlags = 2
	WSMAN_FLAG_AUTH_NEGOTIATE          WSManAuthenticationFlags = 4
	WSMAN_FLAG_AUTH_BASIC              WSManAuthenticationFlags = 8
	WSMAN_FLAG_AUTH_KERBEROS           WSManAuthenticationFlags = 16
	WSMAN_FLAG_AUTH_CREDSSP            WSManAuthenticationFlags = 128
	WSMAN_FLAG_AUTH_CLIENT_CERTIFICATE WSManAuthenticationFlags = 32
)

type WSManProxyAccessType int32

const (
	WSMAN_OPTION_PROXY_IE_PROXY_CONFIG      WSManProxyAccessType = 1
	WSMAN_OPTION_PROXY_WINHTTP_PROXY_CONFIG WSManProxyAccessType = 2
	WSMAN_OPTION_PROXY_AUTO_DETECT          WSManProxyAccessType = 4
	WSMAN_OPTION_PROXY_NO_PROXY_SERVER      WSManProxyAccessType = 8
)

type WSManSessionOption int32

const (
	WSMAN_OPTION_DEFAULT_OPERATION_TIMEOUTMS          WSManSessionOption = 1
	WSMAN_OPTION_MAX_RETRY_TIME                       WSManSessionOption = 11
	WSMAN_OPTION_TIMEOUTMS_CREATE_SHELL               WSManSessionOption = 12
	WSMAN_OPTION_TIMEOUTMS_RUN_SHELL_COMMAND          WSManSessionOption = 13
	WSMAN_OPTION_TIMEOUTMS_RECEIVE_SHELL_OUTPUT       WSManSessionOption = 14
	WSMAN_OPTION_TIMEOUTMS_SEND_SHELL_INPUT           WSManSessionOption = 15
	WSMAN_OPTION_TIMEOUTMS_SIGNAL_SHELL               WSManSessionOption = 16
	WSMAN_OPTION_TIMEOUTMS_CLOSE_SHELL                WSManSessionOption = 17
	WSMAN_OPTION_SKIP_CA_CHECK                        WSManSessionOption = 18
	WSMAN_OPTION_SKIP_CN_CHECK                        WSManSessionOption = 19
	WSMAN_OPTION_UNENCRYPTED_MESSAGES                 WSManSessionOption = 20
	WSMAN_OPTION_UTF16                                WSManSessionOption = 21
	WSMAN_OPTION_ENABLE_SPN_SERVER_PORT               WSManSessionOption = 22
	WSMAN_OPTION_MACHINE_ID                           WSManSessionOption = 23
	WSMAN_OPTION_LOCALE                               WSManSessionOption = 25
	WSMAN_OPTION_UI_LANGUAGE                          WSManSessionOption = 26
	WSMAN_OPTION_MAX_ENVELOPE_SIZE_KB                 WSManSessionOption = 28
	WSMAN_OPTION_SHELL_MAX_DATA_SIZE_PER_MESSAGE_KB   WSManSessionOption = 29
	WSMAN_OPTION_REDIRECT_LOCATION                    WSManSessionOption = 30
	WSMAN_OPTION_SKIP_REVOCATION_CHECK                WSManSessionOption = 31
	WSMAN_OPTION_ALLOW_NEGOTIATE_IMPLICIT_CREDENTIALS WSManSessionOption = 32
	WSMAN_OPTION_USE_SSL                              WSManSessionOption = 33
	WSMAN_OPTION_USE_INTEARACTIVE_TOKEN               WSManSessionOption = 34
)

type WSManCallbackFlags int32

const (
	WSMAN_FLAG_CALLBACK_END_OF_OPERATION                       WSManCallbackFlags = 1
	WSMAN_FLAG_CALLBACK_END_OF_STREAM                          WSManCallbackFlags = 8
	WSMAN_FLAG_CALLBACK_SHELL_SUPPORTS_DISCONNECT              WSManCallbackFlags = 32
	WSMAN_FLAG_CALLBACK_SHELL_AUTODISCONNECTED                 WSManCallbackFlags = 64
	WSMAN_FLAG_CALLBACK_NETWORK_FAILURE_DETECTED               WSManCallbackFlags = 256
	WSMAN_FLAG_CALLBACK_RETRYING_AFTER_NETWORK_FAILURE         WSManCallbackFlags = 512
	WSMAN_FLAG_CALLBACK_RECONNECTED_AFTER_NETWORK_FAILURE      WSManCallbackFlags = 1024
	WSMAN_FLAG_CALLBACK_SHELL_AUTODISCONNECTING                WSManCallbackFlags = 2048
	WSMAN_FLAG_CALLBACK_RETRY_ABORTED_DUE_TO_INTERNAL_ERROR    WSManCallbackFlags = 4096
	WSMAN_FLAG_CALLBACK_RECEIVE_DELAY_STREAM_REQUEST_PROCESSED WSManCallbackFlags = 8192
)

type WSManShellFlag int32

const (
	WSMAN_FLAG_NO_COMPRESSION              WSManShellFlag = 1
	WSMAN_FLAG_DELETE_SERVER_SESSION       WSManShellFlag = 2
	WSMAN_FLAG_SERVER_BUFFERING_MODE_DROP  WSManShellFlag = 4
	WSMAN_FLAG_SERVER_BUFFERING_MODE_BLOCK WSManShellFlag = 8
	WSMAN_FLAG_RECEIVE_DELAY_OUTPUT_STREAM WSManShellFlag = 16
)

type WSManSessionFlags int32

const (
	WSManFlagUTF8                              WSManSessionFlags = 1
	WSManFlagCredUsernamePassword              WSManSessionFlags = 4096
	WSManFlagSkipCACheck                       WSManSessionFlags = 8192
	WSManFlagSkipCNCheck                       WSManSessionFlags = 16384
	WSManFlagUseNoAuthentication               WSManSessionFlags = 32768
	WSManFlagUseDigest                         WSManSessionFlags = 65536
	WSManFlagUseNegotiate                      WSManSessionFlags = 131072
	WSManFlagUseBasic                          WSManSessionFlags = 262144
	WSManFlagUseKerberos                       WSManSessionFlags = 524288
	WSManFlagNoEncryption                      WSManSessionFlags = 1048576
	WSManFlagUseClientCertificate              WSManSessionFlags = 2097152
	WSManFlagEnableSPNServerPort               WSManSessionFlags = 4194304
	WSManFlagUTF16                             WSManSessionFlags = 8388608
	WSManFlagUseCredSsp                        WSManSessionFlags = 16777216
	WSManFlagSkipRevocationCheck               WSManSessionFlags = 33554432
	WSManFlagAllowNegotiateImplicitCredentials WSManSessionFlags = 67108864
	WSManFlagUseSsl                            WSManSessionFlags = 134217728
)

type WSManEnumFlags int32

const (
	WSManFlagNonXmlText                 WSManEnumFlags = 1
	WSManFlagReturnObject               WSManEnumFlags = 0
	WSManFlagReturnEPR                  WSManEnumFlags = 2
	WSManFlagReturnObjectAndEPR         WSManEnumFlags = 4
	WSManFlagHierarchyDeep              WSManEnumFlags = 0
	WSManFlagHierarchyShallow           WSManEnumFlags = 32
	WSManFlagHierarchyDeepBasePropsOnly WSManEnumFlags = 64
	WSManFlagAssociatedInstance         WSManEnumFlags = 0
	WSManFlagAssociationInstance        WSManEnumFlags = 128
)

type WSManProxyAccessTypeFlags int32

const (
	WSManProxyIEConfig      WSManProxyAccessTypeFlags = 1
	WSManProxyWinHttpConfig WSManProxyAccessTypeFlags = 2
	WSManProxyAutoDetect    WSManProxyAccessTypeFlags = 4
	WSManProxyNoProxyServer WSManProxyAccessTypeFlags = 8
)

type WSManProxyAuthenticationFlags int32

const (
	WSManFlagProxyAuthenticationUseNegotiate WSManProxyAuthenticationFlags = 1
	WSManFlagProxyAuthenticationUseBasic     WSManProxyAuthenticationFlags = 2
	WSManFlagProxyAuthenticationUseDigest    WSManProxyAuthenticationFlags = 4
)

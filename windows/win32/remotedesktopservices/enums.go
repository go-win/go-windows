// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package remotedesktopservices implements the Windows.Win32.RemoteDesktopServices namespace.
package remotedesktopservices

type AAAuthSchemes int32

const (
	AA_AUTH_MIN                 AAAuthSchemes = 0
	AA_AUTH_BASIC               AAAuthSchemes = 1
	AA_AUTH_NTLM                AAAuthSchemes = 2
	AA_AUTH_SC                  AAAuthSchemes = 3
	AA_AUTH_LOGGEDONCREDENTIALS AAAuthSchemes = 4
	AA_AUTH_NEGOTIATE           AAAuthSchemes = 5
	AA_AUTH_ANY                 AAAuthSchemes = 6
	AA_AUTH_COOKIE              AAAuthSchemes = 7
	AA_AUTH_DIGEST              AAAuthSchemes = 8
	AA_AUTH_ORGID               AAAuthSchemes = 9
	AA_AUTH_CONID               AAAuthSchemes = 10
	AA_AUTH_SSPI_NTLM           AAAuthSchemes = 11
	AA_AUTH_MAX                 AAAuthSchemes = 12
)

type AAAccountingDataType int32

const (
	AA_MAIN_SESSION_CREATION AAAccountingDataType = 0
	AA_SUB_SESSION_CREATION  AAAccountingDataType = 1
	AA_SUB_SESSION_CLOSED    AAAccountingDataType = 2
	AA_MAIN_SESSION_CLOSED   AAAccountingDataType = 3
)

type MIDL___MIDL_itf_tsgpolicyengine_0000_0000_0004 int32

const (
	SESSION_TIMEOUT_ACTION_DISCONNECT    MIDL___MIDL_itf_tsgpolicyengine_0000_0000_0004 = 0
	SESSION_TIMEOUT_ACTION_SILENT_REAUTH MIDL___MIDL_itf_tsgpolicyengine_0000_0000_0004 = 1
)

type PolicyAttributeType int32

const (
	EnableAllRedirections        PolicyAttributeType = 0
	DisableAllRedirections       PolicyAttributeType = 1
	DriveRedirectionDisabled     PolicyAttributeType = 2
	PrinterRedirectionDisabled   PolicyAttributeType = 3
	PortRedirectionDisabled      PolicyAttributeType = 4
	ClipboardRedirectionDisabled PolicyAttributeType = 5
	PnpRedirectionDisabled       PolicyAttributeType = 6
	AllowOnlySDRServers          PolicyAttributeType = 7
)

type MIDL___MIDL_itf_tsgpolicyengine_0000_0000_0006 int32

const (
	AA_UNTRUSTED                   MIDL___MIDL_itf_tsgpolicyengine_0000_0000_0006 = 0
	AA_TRUSTEDUSER_UNTRUSTEDCLIENT MIDL___MIDL_itf_tsgpolicyengine_0000_0000_0006 = 1
	AA_TRUSTEDUSER_TRUSTEDCLIENT   MIDL___MIDL_itf_tsgpolicyengine_0000_0000_0006 = 2
)

type WTS_CONNECTSTATE_CLASS int32

const (
	WTSActive       WTS_CONNECTSTATE_CLASS = 0
	WTSConnected    WTS_CONNECTSTATE_CLASS = 1
	WTSConnectQuery WTS_CONNECTSTATE_CLASS = 2
	WTSShadow       WTS_CONNECTSTATE_CLASS = 3
	WTSDisconnected WTS_CONNECTSTATE_CLASS = 4
	WTSIdle         WTS_CONNECTSTATE_CLASS = 5
	WTSListen       WTS_CONNECTSTATE_CLASS = 6
	WTSReset        WTS_CONNECTSTATE_CLASS = 7
	WTSDown         WTS_CONNECTSTATE_CLASS = 8
	WTSInit         WTS_CONNECTSTATE_CLASS = 9
)

type WTS_INFO_CLASS int32

const (
	WTSInitialProgram     WTS_INFO_CLASS = 0
	WTSApplicationName    WTS_INFO_CLASS = 1
	WTSWorkingDirectory   WTS_INFO_CLASS = 2
	WTSOEMId              WTS_INFO_CLASS = 3
	WTSSessionId          WTS_INFO_CLASS = 4
	WTSUserName           WTS_INFO_CLASS = 5
	WTSWinStationName     WTS_INFO_CLASS = 6
	WTSDomainName         WTS_INFO_CLASS = 7
	WTSConnectState       WTS_INFO_CLASS = 8
	WTSClientBuildNumber  WTS_INFO_CLASS = 9
	WTSClientName         WTS_INFO_CLASS = 10
	WTSClientDirectory    WTS_INFO_CLASS = 11
	WTSClientProductId    WTS_INFO_CLASS = 12
	WTSClientHardwareId   WTS_INFO_CLASS = 13
	WTSClientAddress      WTS_INFO_CLASS = 14
	WTSClientDisplay      WTS_INFO_CLASS = 15
	WTSClientProtocolType WTS_INFO_CLASS = 16
	WTSIdleTime           WTS_INFO_CLASS = 17
	WTSLogonTime          WTS_INFO_CLASS = 18
	WTSIncomingBytes      WTS_INFO_CLASS = 19
	WTSOutgoingBytes      WTS_INFO_CLASS = 20
	WTSIncomingFrames     WTS_INFO_CLASS = 21
	WTSOutgoingFrames     WTS_INFO_CLASS = 22
	WTSClientInfo         WTS_INFO_CLASS = 23
	WTSSessionInfo        WTS_INFO_CLASS = 24
	WTSSessionInfoEx      WTS_INFO_CLASS = 25
	WTSConfigInfo         WTS_INFO_CLASS = 26
	WTSValidationInfo     WTS_INFO_CLASS = 27
	WTSSessionAddressV4   WTS_INFO_CLASS = 28
	WTSIsRemoteSession    WTS_INFO_CLASS = 29
)

type WTS_CONFIG_CLASS int32

const (
	WTSUserConfigInitialProgram                WTS_CONFIG_CLASS = 0
	WTSUserConfigWorkingDirectory              WTS_CONFIG_CLASS = 1
	WTSUserConfigfInheritInitialProgram        WTS_CONFIG_CLASS = 2
	WTSUserConfigfAllowLogonTerminalServer     WTS_CONFIG_CLASS = 3
	WTSUserConfigTimeoutSettingsConnections    WTS_CONFIG_CLASS = 4
	WTSUserConfigTimeoutSettingsDisconnections WTS_CONFIG_CLASS = 5
	WTSUserConfigTimeoutSettingsIdle           WTS_CONFIG_CLASS = 6
	WTSUserConfigfDeviceClientDrives           WTS_CONFIG_CLASS = 7
	WTSUserConfigfDeviceClientPrinters         WTS_CONFIG_CLASS = 8
	WTSUserConfigfDeviceClientDefaultPrinter   WTS_CONFIG_CLASS = 9
	WTSUserConfigBrokenTimeoutSettings         WTS_CONFIG_CLASS = 10
	WTSUserConfigReconnectSettings             WTS_CONFIG_CLASS = 11
	WTSUserConfigModemCallbackSettings         WTS_CONFIG_CLASS = 12
	WTSUserConfigModemCallbackPhoneNumber      WTS_CONFIG_CLASS = 13
	WTSUserConfigShadowingSettings             WTS_CONFIG_CLASS = 14
	WTSUserConfigTerminalServerProfilePath     WTS_CONFIG_CLASS = 15
	WTSUserConfigTerminalServerHomeDir         WTS_CONFIG_CLASS = 16
	WTSUserConfigTerminalServerHomeDirDrive    WTS_CONFIG_CLASS = 17
	WTSUserConfigfTerminalServerRemoteHomeDir  WTS_CONFIG_CLASS = 18
	WTSUserConfigUser                          WTS_CONFIG_CLASS = 19
)

type WTS_CONFIG_SOURCE int32

const (
	WTSUserConfigSourceSAM WTS_CONFIG_SOURCE = 0
)

type WTS_VIRTUAL_CLASS int32

const (
	WTSVirtualClientData WTS_VIRTUAL_CLASS = 0
	WTSVirtualFileHandle WTS_VIRTUAL_CLASS = 1
)

type WTS_TYPE_CLASS int32

const (
	WTSTypeProcessInfoLevel0 WTS_TYPE_CLASS = 0
	WTSTypeProcessInfoLevel1 WTS_TYPE_CLASS = 1
	WTSTypeSessionInfoLevel1 WTS_TYPE_CLASS = 2
)

type WTSSBX_MACHINE_DRAIN int32

const (
	WTSSBX_MACHINE_DRAIN_UNSPEC WTSSBX_MACHINE_DRAIN = 0
	WTSSBX_MACHINE_DRAIN_OFF    WTSSBX_MACHINE_DRAIN = 1
	WTSSBX_MACHINE_DRAIN_ON     WTSSBX_MACHINE_DRAIN = 2
)

type WTSSBX_MACHINE_SESSION_MODE int32

const (
	WTSSBX_MACHINE_SESSION_MODE_UNSPEC   WTSSBX_MACHINE_SESSION_MODE = 0
	WTSSBX_MACHINE_SESSION_MODE_SINGLE   WTSSBX_MACHINE_SESSION_MODE = 1
	WTSSBX_MACHINE_SESSION_MODE_MULTIPLE WTSSBX_MACHINE_SESSION_MODE = 2
)

type WTSSBX_ADDRESS_FAMILY int32

const (
	WTSSBX_ADDRESS_FAMILY_AF_UNSPEC  WTSSBX_ADDRESS_FAMILY = 0
	WTSSBX_ADDRESS_FAMILY_AF_INET    WTSSBX_ADDRESS_FAMILY = 1
	WTSSBX_ADDRESS_FAMILY_AF_INET6   WTSSBX_ADDRESS_FAMILY = 2
	WTSSBX_ADDRESS_FAMILY_AF_IPX     WTSSBX_ADDRESS_FAMILY = 3
	WTSSBX_ADDRESS_FAMILY_AF_NETBIOS WTSSBX_ADDRESS_FAMILY = 4
)

type WTSSBX_MACHINE_STATE int32

const (
	WTSSBX_MACHINE_STATE_UNSPEC        WTSSBX_MACHINE_STATE = 0
	WTSSBX_MACHINE_STATE_READY         WTSSBX_MACHINE_STATE = 1
	WTSSBX_MACHINE_STATE_SYNCHRONIZING WTSSBX_MACHINE_STATE = 2
)

type WTSSBX_SESSION_STATE int32

const (
	WTSSBX_SESSION_STATE_UNSPEC       WTSSBX_SESSION_STATE = 0
	WTSSBX_SESSION_STATE_ACTIVE       WTSSBX_SESSION_STATE = 1
	WTSSBX_SESSION_STATE_DISCONNECTED WTSSBX_SESSION_STATE = 2
)

type WTSSBX_NOTIFICATION_TYPE int32

const (
	WTSSBX_NOTIFICATION_REMOVED WTSSBX_NOTIFICATION_TYPE = 1
	WTSSBX_NOTIFICATION_CHANGED WTSSBX_NOTIFICATION_TYPE = 2
	WTSSBX_NOTIFICATION_ADDED   WTSSBX_NOTIFICATION_TYPE = 4
	WTSSBX_NOTIFICATION_RESYNC  WTSSBX_NOTIFICATION_TYPE = 8
)

type TSSD_AddrV46Type int32

const (
	TSSD_ADDR_UNDEFINED TSSD_AddrV46Type = 0
	TSSD_ADDR_IPv4      TSSD_AddrV46Type = 4
	TSSD_ADDR_IPv6      TSSD_AddrV46Type = 6
)

type TSSB_NOTIFICATION_TYPE int32

const (
	TSSB_NOTIFY_INVALID                   TSSB_NOTIFICATION_TYPE = 0
	TSSB_NOTIFY_TARGET_CHANGE             TSSB_NOTIFICATION_TYPE = 1
	TSSB_NOTIFY_SESSION_CHANGE            TSSB_NOTIFICATION_TYPE = 2
	TSSB_NOTIFY_CONNECTION_REQUEST_CHANGE TSSB_NOTIFICATION_TYPE = 4
)

type TARGET_STATE int32

const (
	TARGET_UNKNOWN      TARGET_STATE = 1
	TARGET_INITIALIZING TARGET_STATE = 2
	TARGET_RUNNING      TARGET_STATE = 3
	TARGET_DOWN         TARGET_STATE = 4
	TARGET_HIBERNATED   TARGET_STATE = 5
	TARGET_CHECKED_OUT  TARGET_STATE = 6
	TARGET_STOPPED      TARGET_STATE = 7
	TARGET_INVALID      TARGET_STATE = 8
	TARGET_STARTING     TARGET_STATE = 9
	TARGET_STOPPING     TARGET_STATE = 10
	TARGET_MAXSTATE     TARGET_STATE = 11
)

type TARGET_CHANGE_TYPE int32

const (
	TARGET_CHANGE_UNSPEC           TARGET_CHANGE_TYPE = 1
	TARGET_EXTERNALIP_CHANGED      TARGET_CHANGE_TYPE = 2
	TARGET_INTERNALIP_CHANGED      TARGET_CHANGE_TYPE = 4
	TARGET_JOINED                  TARGET_CHANGE_TYPE = 8
	TARGET_REMOVED                 TARGET_CHANGE_TYPE = 16
	TARGET_STATE_CHANGED           TARGET_CHANGE_TYPE = 32
	TARGET_IDLE                    TARGET_CHANGE_TYPE = 64
	TARGET_PENDING                 TARGET_CHANGE_TYPE = 128
	TARGET_INUSE                   TARGET_CHANGE_TYPE = 256
	TARGET_PATCH_STATE_CHANGED     TARGET_CHANGE_TYPE = 512
	TARGET_FARM_MEMBERSHIP_CHANGED TARGET_CHANGE_TYPE = 1024
)

type TARGET_TYPE int32

const (
	UNKNOWN TARGET_TYPE = 0
	FARM    TARGET_TYPE = 1
	NONFARM TARGET_TYPE = 2
)

type TARGET_PATCH_STATE int32

const (
	TARGET_PATCH_UNKNOWN     TARGET_PATCH_STATE = 0
	TARGET_PATCH_NOT_STARTED TARGET_PATCH_STATE = 1
	TARGET_PATCH_IN_PROGRESS TARGET_PATCH_STATE = 2
	TARGET_PATCH_COMPLETED   TARGET_PATCH_STATE = 3
	TARGET_PATCH_FAILED      TARGET_PATCH_STATE = 4
)

type CLIENT_MESSAGE_TYPE int32

const (
	CLIENT_MESSAGE_CONNECTION_INVALID CLIENT_MESSAGE_TYPE = 0
	CLIENT_MESSAGE_CONNECTION_STATUS  CLIENT_MESSAGE_TYPE = 1
	CLIENT_MESSAGE_CONNECTION_ERROR   CLIENT_MESSAGE_TYPE = 2
)

type CONNECTION_CHANGE_NOTIFICATION int32

const (
	CONNECTION_REQUEST_INVALID            CONNECTION_CHANGE_NOTIFICATION = 0
	CONNECTION_REQUEST_PENDING            CONNECTION_CHANGE_NOTIFICATION = 1
	CONNECTION_REQUEST_FAILED             CONNECTION_CHANGE_NOTIFICATION = 2
	CONNECTION_REQUEST_TIMEDOUT           CONNECTION_CHANGE_NOTIFICATION = 3
	CONNECTION_REQUEST_SUCCEEDED          CONNECTION_CHANGE_NOTIFICATION = 4
	CONNECTION_REQUEST_CANCELLED          CONNECTION_CHANGE_NOTIFICATION = 5
	CONNECTION_REQUEST_LB_COMPLETED       CONNECTION_CHANGE_NOTIFICATION = 6
	CONNECTION_REQUEST_QUERY_PL_COMPLETED CONNECTION_CHANGE_NOTIFICATION = 7
	CONNECTION_REQUEST_ORCH_COMPLETED     CONNECTION_CHANGE_NOTIFICATION = 8
)

type RD_FARM_TYPE int32

const (
	RD_FARM_RDSH                 RD_FARM_TYPE = 0
	RD_FARM_TEMP_VM              RD_FARM_TYPE = 1
	RD_FARM_MANUAL_PERSONAL_VM   RD_FARM_TYPE = 2
	RD_FARM_AUTO_PERSONAL_VM     RD_FARM_TYPE = 3
	RD_FARM_MANUAL_PERSONAL_RDSH RD_FARM_TYPE = 4
	RD_FARM_AUTO_PERSONAL_RDSH   RD_FARM_TYPE = 5
	RD_FARM_TYPE_UNKNOWN         RD_FARM_TYPE = -1
)

type PLUGIN_TYPE int32

const (
	UNKNOWN_PLUGIN        PLUGIN_TYPE = 0
	POLICY_PLUGIN         PLUGIN_TYPE = 1
	RESOURCE_PLUGIN       PLUGIN_TYPE = 2
	LOAD_BALANCING_PLUGIN PLUGIN_TYPE = 4
	PLACEMENT_PLUGIN      PLUGIN_TYPE = 8
	ORCHESTRATION_PLUGIN  PLUGIN_TYPE = 16
	PROVISIONING_PLUGIN   PLUGIN_TYPE = 32
	TASK_PLUGIN           PLUGIN_TYPE = 64
)

type TSSESSION_STATE int32

const (
	STATE_INVALID      TSSESSION_STATE = -1
	STATE_ACTIVE       TSSESSION_STATE = 0
	STATE_CONNECTED    TSSESSION_STATE = 1
	STATE_CONNECTQUERY TSSESSION_STATE = 2
	STATE_SHADOW       TSSESSION_STATE = 3
	STATE_DISCONNECTED TSSESSION_STATE = 4
	STATE_IDLE         TSSESSION_STATE = 5
	STATE_LISTEN       TSSESSION_STATE = 6
	STATE_RESET        TSSESSION_STATE = 7
	STATE_DOWN         TSSESSION_STATE = 8
	STATE_INIT         TSSESSION_STATE = 9
	STATE_MAX          TSSESSION_STATE = 10
)

type TARGET_OWNER int32

const (
	OWNER_UNKNOWN      TARGET_OWNER = 0
	OWNER_MS_TS_PLUGIN TARGET_OWNER = 1
	OWNER_MS_VM_PLUGIN TARGET_OWNER = 2
)

type VM_NOTIFY_STATUS int32

const (
	VM_NOTIFY_STATUS_PENDING     VM_NOTIFY_STATUS = 0
	VM_NOTIFY_STATUS_IN_PROGRESS VM_NOTIFY_STATUS = 1
	VM_NOTIFY_STATUS_COMPLETE    VM_NOTIFY_STATUS = 2
	VM_NOTIFY_STATUS_FAILED      VM_NOTIFY_STATUS = 3
	VM_NOTIFY_STATUS_CANCELED    VM_NOTIFY_STATUS = 4
)

type VM_HOST_NOTIFY_STATUS int32

const (
	VM_HOST_STATUS_INIT_PENDING     VM_HOST_NOTIFY_STATUS = 0
	VM_HOST_STATUS_INIT_IN_PROGRESS VM_HOST_NOTIFY_STATUS = 1
	VM_HOST_STATUS_INIT_COMPLETE    VM_HOST_NOTIFY_STATUS = 2
	VM_HOST_STATUS_INIT_FAILED      VM_HOST_NOTIFY_STATUS = 3
)

type RDV_TASK_STATUS int32

const (
	RDV_TASK_STATUS_UNKNOWN     RDV_TASK_STATUS = 0
	RDV_TASK_STATUS_SEARCHING   RDV_TASK_STATUS = 1
	RDV_TASK_STATUS_DOWNLOADING RDV_TASK_STATUS = 2
	RDV_TASK_STATUS_APPLYING    RDV_TASK_STATUS = 3
	RDV_TASK_STATUS_REBOOTING   RDV_TASK_STATUS = 4
	RDV_TASK_STATUS_REBOOTED    RDV_TASK_STATUS = 5
	RDV_TASK_STATUS_SUCCESS     RDV_TASK_STATUS = 6
	RDV_TASK_STATUS_FAILED      RDV_TASK_STATUS = 7
	RDV_TASK_STATUS_TIMEOUT     RDV_TASK_STATUS = 8
)

type TS_SB_SORT_BY int32

const (
	TS_SB_SORT_BY_NONE TS_SB_SORT_BY = 0
	TS_SB_SORT_BY_NAME TS_SB_SORT_BY = 1
	TS_SB_SORT_BY_PROP TS_SB_SORT_BY = 2
)

type TSPUB_PLUGIN_PD_RESOLUTION_TYPE int32

const (
	TSPUB_PLUGIN_PD_QUERY_OR_CREATE TSPUB_PLUGIN_PD_RESOLUTION_TYPE = 0
	TSPUB_PLUGIN_PD_QUERY_EXISTING  TSPUB_PLUGIN_PD_RESOLUTION_TYPE = 1
)

type TSPUB_PLUGIN_PD_ASSIGNMENT_TYPE int32

const (
	TSPUB_PLUGIN_PD_ASSIGNMENT_NEW      TSPUB_PLUGIN_PD_ASSIGNMENT_TYPE = 0
	TSPUB_PLUGIN_PD_ASSIGNMENT_EXISTING TSPUB_PLUGIN_PD_ASSIGNMENT_TYPE = 1
)

type WRdsGraphicsChannelType int32

const (
	WRdsGraphicsChannelType_GuaranteedDelivery WRdsGraphicsChannelType = 0
	WRdsGraphicsChannelType_BestEffortDelivery WRdsGraphicsChannelType = 1
)

type WTS_RCM_SERVICE_STATE int32

const (
	WTS_SERVICE_NONE  WTS_RCM_SERVICE_STATE = 0
	WTS_SERVICE_START WTS_RCM_SERVICE_STATE = 1
	WTS_SERVICE_STOP  WTS_RCM_SERVICE_STATE = 2
)

type WTS_RCM_DRAIN_STATE int32

const (
	WTS_DRAIN_STATE_NONE   WTS_RCM_DRAIN_STATE = 0
	WTS_DRAIN_IN_DRAIN     WTS_RCM_DRAIN_STATE = 1
	WTS_DRAIN_NOT_IN_DRAIN WTS_RCM_DRAIN_STATE = 2
)

type WTS_LOGON_ERROR_REDIRECTOR_RESPONSE int32

const (
	WTS_LOGON_ERR_INVALID                      WTS_LOGON_ERROR_REDIRECTOR_RESPONSE = 0
	WTS_LOGON_ERR_NOT_HANDLED                  WTS_LOGON_ERROR_REDIRECTOR_RESPONSE = 1
	WTS_LOGON_ERR_HANDLED_SHOW                 WTS_LOGON_ERROR_REDIRECTOR_RESPONSE = 2
	WTS_LOGON_ERR_HANDLED_DONT_SHOW            WTS_LOGON_ERROR_REDIRECTOR_RESPONSE = 3
	WTS_LOGON_ERR_HANDLED_DONT_SHOW_START_OVER WTS_LOGON_ERROR_REDIRECTOR_RESPONSE = 4
)

type WTS_CERT_TYPE int32

const (
	WTS_CERT_TYPE_INVALID     WTS_CERT_TYPE = 0
	WTS_CERT_TYPE_PROPRIETORY WTS_CERT_TYPE = 1
	WTS_CERT_TYPE_X509        WTS_CERT_TYPE = 2
)

type WRDS_CONNECTION_SETTING_LEVEL int32

const (
	WRDS_CONNECTION_SETTING_LEVEL_INVALID WRDS_CONNECTION_SETTING_LEVEL = 0
	WRDS_CONNECTION_SETTING_LEVEL_1       WRDS_CONNECTION_SETTING_LEVEL = 1
)

type WRDS_LISTENER_SETTING_LEVEL int32

const (
	WRDS_LISTENER_SETTING_LEVEL_INVALID WRDS_LISTENER_SETTING_LEVEL = 0
	WRDS_LISTENER_SETTING_LEVEL_1       WRDS_LISTENER_SETTING_LEVEL = 1
)

type WRDS_SETTING_TYPE int32

const (
	WRDS_SETTING_TYPE_INVALID WRDS_SETTING_TYPE = 0
	WRDS_SETTING_TYPE_MACHINE WRDS_SETTING_TYPE = 1
	WRDS_SETTING_TYPE_USER    WRDS_SETTING_TYPE = 2
	WRDS_SETTING_TYPE_SAM     WRDS_SETTING_TYPE = 3
)

type WRDS_SETTING_STATUS int32

const (
	WRDS_SETTING_STATUS_NOTAPPLICABLE WRDS_SETTING_STATUS = -1
	WRDS_SETTING_STATUS_DISABLED      WRDS_SETTING_STATUS = 0
	WRDS_SETTING_STATUS_ENABLED       WRDS_SETTING_STATUS = 1
	WRDS_SETTING_STATUS_NOTCONFIGURED WRDS_SETTING_STATUS = 2
)

type WRDS_SETTING_LEVEL int32

const (
	WRDS_SETTING_LEVEL_INVALID WRDS_SETTING_LEVEL = 0
	WRDS_SETTING_LEVEL_1       WRDS_SETTING_LEVEL = 1
)

type MIDL_IRemoteDesktopClientSettings_0001 int32

const (
	PasswordEncodingUTF8    MIDL_IRemoteDesktopClientSettings_0001 = 0
	PasswordEncodingUTF16LE MIDL_IRemoteDesktopClientSettings_0001 = 1
	PasswordEncodingUTF16BE MIDL_IRemoteDesktopClientSettings_0001 = 2
)

type RemoteActionType int32

const (
	RemoteActionCharms      RemoteActionType = 0
	RemoteActionAppbar      RemoteActionType = 1
	RemoteActionSnap        RemoteActionType = 2
	RemoteActionStartScreen RemoteActionType = 3
	RemoteActionAppSwitch   RemoteActionType = 4
)

type SnapshotEncodingType int32

const (
	SnapshotEncodingDataUri SnapshotEncodingType = 0
)

type SnapshotFormatType int32

const (
	SnapshotFormatPng  SnapshotFormatType = 0
	SnapshotFormatJpeg SnapshotFormatType = 1
	SnapshotFormatBmp  SnapshotFormatType = 2
)

type MIDL_IRemoteDesktopClient_0001 int32

const (
	KeyCombinationHome   MIDL_IRemoteDesktopClient_0001 = 0
	KeyCombinationLeft   MIDL_IRemoteDesktopClient_0001 = 1
	KeyCombinationUp     MIDL_IRemoteDesktopClient_0001 = 2
	KeyCombinationRight  MIDL_IRemoteDesktopClient_0001 = 3
	KeyCombinationDown   MIDL_IRemoteDesktopClient_0001 = 4
	KeyCombinationScroll MIDL_IRemoteDesktopClient_0001 = 5
)

type APO_BUFFER_FLAGS int32

const (
	BUFFER_INVALID APO_BUFFER_FLAGS = 0
	BUFFER_VALID   APO_BUFFER_FLAGS = 1
	BUFFER_SILENT  APO_BUFFER_FLAGS = 2
)

type AE_POSITION_FLAGS int32

const (
	POSITION_INVALID       AE_POSITION_FLAGS = 0
	POSITION_DISCONTINUOUS AE_POSITION_FLAGS = 1
	POSITION_CONTINUOUS    AE_POSITION_FLAGS = 2
	POSITION_QPC_ERROR     AE_POSITION_FLAGS = 4
)

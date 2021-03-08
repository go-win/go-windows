// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsfirewall implements the Windows.Win32.WindowsFirewall namespace.
package windowsfirewall

type NETCON_CHARACTERISTIC_FLAGS int32

const (
	NCCF_NONE              NETCON_CHARACTERISTIC_FLAGS = 0
	NCCF_ALL_USERS         NETCON_CHARACTERISTIC_FLAGS = 1
	NCCF_ALLOW_DUPLICATION NETCON_CHARACTERISTIC_FLAGS = 2
	NCCF_ALLOW_REMOVAL     NETCON_CHARACTERISTIC_FLAGS = 4
	NCCF_ALLOW_RENAME      NETCON_CHARACTERISTIC_FLAGS = 8
	NCCF_INCOMING_ONLY     NETCON_CHARACTERISTIC_FLAGS = 32
	NCCF_OUTGOING_ONLY     NETCON_CHARACTERISTIC_FLAGS = 64
	NCCF_BRANDED           NETCON_CHARACTERISTIC_FLAGS = 128
	NCCF_SHARED            NETCON_CHARACTERISTIC_FLAGS = 256
	NCCF_BRIDGED           NETCON_CHARACTERISTIC_FLAGS = 512
	NCCF_FIREWALLED        NETCON_CHARACTERISTIC_FLAGS = 1024
	NCCF_DEFAULT           NETCON_CHARACTERISTIC_FLAGS = 2048
	NCCF_HOMENET_CAPABLE   NETCON_CHARACTERISTIC_FLAGS = 4096
	NCCF_SHARED_PRIVATE    NETCON_CHARACTERISTIC_FLAGS = 8192
	NCCF_QUARANTINED       NETCON_CHARACTERISTIC_FLAGS = 16384
	NCCF_RESERVED          NETCON_CHARACTERISTIC_FLAGS = 32768
	NCCF_HOSTED_NETWORK    NETCON_CHARACTERISTIC_FLAGS = 65536
	NCCF_VIRTUAL_STATION   NETCON_CHARACTERISTIC_FLAGS = 131072
	NCCF_WIFI_DIRECT       NETCON_CHARACTERISTIC_FLAGS = 262144
	NCCF_BLUETOOTH_MASK    NETCON_CHARACTERISTIC_FLAGS = 983040
	NCCF_LAN_MASK          NETCON_CHARACTERISTIC_FLAGS = 15728640
)

type NETCON_STATUS int32

const (
	NCS_DISCONNECTED             NETCON_STATUS = 0
	NCS_CONNECTING               NETCON_STATUS = 1
	NCS_CONNECTED                NETCON_STATUS = 2
	NCS_DISCONNECTING            NETCON_STATUS = 3
	NCS_HARDWARE_NOT_PRESENT     NETCON_STATUS = 4
	NCS_HARDWARE_DISABLED        NETCON_STATUS = 5
	NCS_HARDWARE_MALFUNCTION     NETCON_STATUS = 6
	NCS_MEDIA_DISCONNECTED       NETCON_STATUS = 7
	NCS_AUTHENTICATING           NETCON_STATUS = 8
	NCS_AUTHENTICATION_SUCCEEDED NETCON_STATUS = 9
	NCS_AUTHENTICATION_FAILED    NETCON_STATUS = 10
	NCS_INVALID_ADDRESS          NETCON_STATUS = 11
	NCS_CREDENTIALS_REQUIRED     NETCON_STATUS = 12
	NCS_ACTION_REQUIRED          NETCON_STATUS = 13
	NCS_ACTION_REQUIRED_RETRY    NETCON_STATUS = 14
	NCS_CONNECT_FAILED           NETCON_STATUS = 15
)

type NETCON_TYPE int32

const (
	NCT_DIRECT_CONNECT NETCON_TYPE = 0
	NCT_INBOUND        NETCON_TYPE = 1
	NCT_INTERNET       NETCON_TYPE = 2
	NCT_LAN            NETCON_TYPE = 3
	NCT_PHONE          NETCON_TYPE = 4
	NCT_TUNNEL         NETCON_TYPE = 5
	NCT_BRIDGE         NETCON_TYPE = 6
)

type NETCON_MEDIATYPE int32

const (
	NCM_NONE                 NETCON_MEDIATYPE = 0
	NCM_DIRECT               NETCON_MEDIATYPE = 1
	NCM_ISDN                 NETCON_MEDIATYPE = 2
	NCM_LAN                  NETCON_MEDIATYPE = 3
	NCM_PHONE                NETCON_MEDIATYPE = 4
	NCM_TUNNEL               NETCON_MEDIATYPE = 5
	NCM_PPPOE                NETCON_MEDIATYPE = 6
	NCM_BRIDGE               NETCON_MEDIATYPE = 7
	NCM_SHAREDACCESSHOST_LAN NETCON_MEDIATYPE = 8
	NCM_SHAREDACCESSHOST_RAS NETCON_MEDIATYPE = 9
)

type NETCONMGR_ENUM_FLAGS int32

const (
	NCME_DEFAULT NETCONMGR_ENUM_FLAGS = 0
	NCME_HIDDEN  NETCONMGR_ENUM_FLAGS = 1
)

type NETCONUI_CONNECT_FLAGS int32

const (
	NCUC_DEFAULT        NETCONUI_CONNECT_FLAGS = 0
	NCUC_NO_UI          NETCONUI_CONNECT_FLAGS = 1
	NCUC_ENABLE_DISABLE NETCONUI_CONNECT_FLAGS = 2
)

type SHARINGCONNECTIONTYPE int32

const (
	ICSSHARINGTYPE_PUBLIC  SHARINGCONNECTIONTYPE = 0
	ICSSHARINGTYPE_PRIVATE SHARINGCONNECTIONTYPE = 1
)

type SHARINGCONNECTION_ENUM_FLAGS int32

const (
	ICSSC_DEFAULT SHARINGCONNECTION_ENUM_FLAGS = 0
	ICSSC_ENABLED SHARINGCONNECTION_ENUM_FLAGS = 1
)

type ICS_TARGETTYPE int32

const (
	ICSTT_NAME      ICS_TARGETTYPE = 0
	ICSTT_IPADDRESS ICS_TARGETTYPE = 1
)

type NET_FW_POLICY_TYPE int32

const (
	NET_FW_POLICY_GROUP     NET_FW_POLICY_TYPE = 0
	NET_FW_POLICY_LOCAL     NET_FW_POLICY_TYPE = 1
	NET_FW_POLICY_EFFECTIVE NET_FW_POLICY_TYPE = 2
	NET_FW_POLICY_TYPE_MAX  NET_FW_POLICY_TYPE = 3
)

type NET_FW_PROFILE_TYPE int32

const (
	NET_FW_PROFILE_DOMAIN   NET_FW_PROFILE_TYPE = 0
	NET_FW_PROFILE_STANDARD NET_FW_PROFILE_TYPE = 1
	NET_FW_PROFILE_CURRENT  NET_FW_PROFILE_TYPE = 2
	NET_FW_PROFILE_TYPE_MAX NET_FW_PROFILE_TYPE = 3
)

type NET_FW_PROFILE_TYPE2 int32

const (
	NET_FW_PROFILE2_DOMAIN  NET_FW_PROFILE_TYPE2 = 1
	NET_FW_PROFILE2_PRIVATE NET_FW_PROFILE_TYPE2 = 2
	NET_FW_PROFILE2_PUBLIC  NET_FW_PROFILE_TYPE2 = 4
	NET_FW_PROFILE2_ALL     NET_FW_PROFILE_TYPE2 = 2147483647
)

type NET_FW_IP_VERSION int32

const (
	NET_FW_IP_VERSION_V4  NET_FW_IP_VERSION = 0
	NET_FW_IP_VERSION_V6  NET_FW_IP_VERSION = 1
	NET_FW_IP_VERSION_ANY NET_FW_IP_VERSION = 2
	NET_FW_IP_VERSION_MAX NET_FW_IP_VERSION = 3
)

type NET_FW_SCOPE int32

const (
	NET_FW_SCOPE_ALL          NET_FW_SCOPE = 0
	NET_FW_SCOPE_LOCAL_SUBNET NET_FW_SCOPE = 1
	NET_FW_SCOPE_CUSTOM       NET_FW_SCOPE = 2
	NET_FW_SCOPE_MAX          NET_FW_SCOPE = 3
)

type NET_FW_IP_PROTOCOL int32

const (
	NET_FW_IP_PROTOCOL_TCP NET_FW_IP_PROTOCOL = 6
	NET_FW_IP_PROTOCOL_UDP NET_FW_IP_PROTOCOL = 17
	NET_FW_IP_PROTOCOL_ANY NET_FW_IP_PROTOCOL = 256
)

type NET_FW_SERVICE_TYPE int32

const (
	NET_FW_SERVICE_FILE_AND_PRINT NET_FW_SERVICE_TYPE = 0
	NET_FW_SERVICE_UPNP           NET_FW_SERVICE_TYPE = 1
	NET_FW_SERVICE_REMOTE_DESKTOP NET_FW_SERVICE_TYPE = 2
	NET_FW_SERVICE_NONE           NET_FW_SERVICE_TYPE = 3
	NET_FW_SERVICE_TYPE_MAX       NET_FW_SERVICE_TYPE = 4
)

type NET_FW_RULE_DIRECTION int32

const (
	NET_FW_RULE_DIR_IN  NET_FW_RULE_DIRECTION = 1
	NET_FW_RULE_DIR_OUT NET_FW_RULE_DIRECTION = 2
	NET_FW_RULE_DIR_MAX NET_FW_RULE_DIRECTION = 3
)

type NET_FW_ACTION int32

const (
	NET_FW_ACTION_BLOCK NET_FW_ACTION = 0
	NET_FW_ACTION_ALLOW NET_FW_ACTION = 1
	NET_FW_ACTION_MAX   NET_FW_ACTION = 2
)

type NET_FW_MODIFY_STATE int32

const (
	NET_FW_MODIFY_STATE_OK              NET_FW_MODIFY_STATE = 0
	NET_FW_MODIFY_STATE_GP_OVERRIDE     NET_FW_MODIFY_STATE = 1
	NET_FW_MODIFY_STATE_INBOUND_BLOCKED NET_FW_MODIFY_STATE = 2
)

type NET_FW_RULE_CATEGORY int32

const (
	NET_FW_RULE_CATEGORY_BOOT     NET_FW_RULE_CATEGORY = 0
	NET_FW_RULE_CATEGORY_STEALTH  NET_FW_RULE_CATEGORY = 1
	NET_FW_RULE_CATEGORY_FIREWALL NET_FW_RULE_CATEGORY = 2
	NET_FW_RULE_CATEGORY_CONSEC   NET_FW_RULE_CATEGORY = 3
	NET_FW_RULE_CATEGORY_MAX      NET_FW_RULE_CATEGORY = 4
)

type NET_FW_EDGE_TRAVERSAL_TYPE int32

const (
	NET_FW_EDGE_TRAVERSAL_TYPE_DENY          NET_FW_EDGE_TRAVERSAL_TYPE = 0
	NET_FW_EDGE_TRAVERSAL_TYPE_ALLOW         NET_FW_EDGE_TRAVERSAL_TYPE = 1
	NET_FW_EDGE_TRAVERSAL_TYPE_DEFER_TO_APP  NET_FW_EDGE_TRAVERSAL_TYPE = 2
	NET_FW_EDGE_TRAVERSAL_TYPE_DEFER_TO_USER NET_FW_EDGE_TRAVERSAL_TYPE = 3
)

type NET_FW_AUTHENTICATE_TYPE int32

const (
	NET_FW_AUTHENTICATE_NONE                     NET_FW_AUTHENTICATE_TYPE = 0
	NET_FW_AUTHENTICATE_NO_ENCAPSULATION         NET_FW_AUTHENTICATE_TYPE = 1
	NET_FW_AUTHENTICATE_WITH_INTEGRITY           NET_FW_AUTHENTICATE_TYPE = 2
	NET_FW_AUTHENTICATE_AND_NEGOTIATE_ENCRYPTION NET_FW_AUTHENTICATE_TYPE = 3
	NET_FW_AUTHENTICATE_AND_ENCRYPT              NET_FW_AUTHENTICATE_TYPE = 4
)

type NETISO_FLAG int32

const (
	NETISO_FLAG_FORCE_COMPUTE_BINARIES NETISO_FLAG = 1
	NETISO_FLAG_MAX                    NETISO_FLAG = 2
)

type INET_FIREWALL_AC_CREATION_TYPE int32

const (
	INET_FIREWALL_AC_NONE            INET_FIREWALL_AC_CREATION_TYPE = 0
	INET_FIREWALL_AC_PACKAGE_ID_ONLY INET_FIREWALL_AC_CREATION_TYPE = 1
	INET_FIREWALL_AC_BINARY          INET_FIREWALL_AC_CREATION_TYPE = 2
	INET_FIREWALL_AC_MAX             INET_FIREWALL_AC_CREATION_TYPE = 4
)

type INET_FIREWALL_AC_CHANGE_TYPE int32

const (
	INET_FIREWALL_AC_CHANGE_INVALID INET_FIREWALL_AC_CHANGE_TYPE = 0
	INET_FIREWALL_AC_CHANGE_CREATE  INET_FIREWALL_AC_CHANGE_TYPE = 1
	INET_FIREWALL_AC_CHANGE_DELETE  INET_FIREWALL_AC_CHANGE_TYPE = 2
	INET_FIREWALL_AC_CHANGE_MAX     INET_FIREWALL_AC_CHANGE_TYPE = 3
)

type NETISO_ERROR_TYPE int32

const (
	NETISO_ERROR_TYPE_NONE                   NETISO_ERROR_TYPE = 0
	NETISO_ERROR_TYPE_PRIVATE_NETWORK        NETISO_ERROR_TYPE = 1
	NETISO_ERROR_TYPE_INTERNET_CLIENT        NETISO_ERROR_TYPE = 2
	NETISO_ERROR_TYPE_INTERNET_CLIENT_SERVER NETISO_ERROR_TYPE = 3
	NETISO_ERROR_TYPE_MAX                    NETISO_ERROR_TYPE = 4
)

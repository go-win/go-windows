// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package nativewifi implements the Windows.Win32.NativeWiFi namespace.
package nativewifi

type DOT11_BSS_TYPE int32

const (
	dot11_BSS_type_infrastructure = 1
	dot11_BSS_type_independent = 2
	dot11_BSS_type_any = 3
)

type DOT11_AUTH_ALGORITHM int32

const (
	DOT11_AUTH_ALGO_80211_OPEN = 1
	DOT11_AUTH_ALGO_80211_SHARED_KEY = 2
	DOT11_AUTH_ALGO_WPA = 3
	DOT11_AUTH_ALGO_WPA_PSK = 4
	DOT11_AUTH_ALGO_WPA_NONE = 5
	DOT11_AUTH_ALGO_RSNA = 6
	DOT11_AUTH_ALGO_RSNA_PSK = 7
	DOT11_AUTH_ALGO_WPA3 = 8
	DOT11_AUTH_ALGO_WPA3_SAE = 9
	DOT11_AUTH_ALGO_OWE = 10
	DOT11_AUTH_ALGO_IHV_START = -2147483648
	DOT11_AUTH_ALGO_IHV_END = -1
)

type DOT11_CIPHER_ALGORITHM int32

const (
	DOT11_CIPHER_ALGO_NONE = 0
	DOT11_CIPHER_ALGO_WEP40 = 1
	DOT11_CIPHER_ALGO_TKIP = 2
	DOT11_CIPHER_ALGO_CCMP = 4
	DOT11_CIPHER_ALGO_WEP104 = 5
	DOT11_CIPHER_ALGO_BIP = 6
	DOT11_CIPHER_ALGO_GCMP = 8
	DOT11_CIPHER_ALGO_GCMP_256 = 9
	DOT11_CIPHER_ALGO_CCMP_256 = 10
	DOT11_CIPHER_ALGO_BIP_GMAC_128 = 11
	DOT11_CIPHER_ALGO_BIP_GMAC_256 = 12
	DOT11_CIPHER_ALGO_BIP_CMAC_256 = 13
	DOT11_CIPHER_ALGO_WPA_USE_GROUP = 256
	DOT11_CIPHER_ALGO_RSN_USE_GROUP = 256
	DOT11_CIPHER_ALGO_WEP = 257
	DOT11_CIPHER_ALGO_IHV_START = -2147483648
	DOT11_CIPHER_ALGO_IHV_END = -1
)

type NDIS_REQUEST_TYPE int32

const (
	NdisRequestQueryInformation = 0
	NdisRequestSetInformation = 1
	NdisRequestQueryStatistics = 2
	NdisRequestOpen = 3
	NdisRequestClose = 4
	NdisRequestSend = 5
	NdisRequestTransferData = 6
	NdisRequestReset = 7
	NdisRequestGeneric1 = 8
	NdisRequestGeneric2 = 9
	NdisRequestGeneric3 = 10
	NdisRequestGeneric4 = 11
)

type NDIS_INTERRUPT_MODERATION int32

const (
	NdisInterruptModerationUnknown = 0
	NdisInterruptModerationNotSupported = 1
	NdisInterruptModerationEnabled = 2
	NdisInterruptModerationDisabled = 3
)

type NDIS_802_11_STATUS_TYPE int32

const (
	Ndis802_11StatusType_Authentication = 0
	Ndis802_11StatusType_MediaStreamMode = 1
	Ndis802_11StatusType_PMKID_CandidateList = 2
	Ndis802_11StatusTypeMax = 3
)

type NDIS_802_11_NETWORK_TYPE int32

const (
	Ndis802_11FH = 0
	Ndis802_11DS = 1
	Ndis802_11OFDM5 = 2
	Ndis802_11OFDM24 = 3
	Ndis802_11Automode = 4
	Ndis802_11NetworkTypeMax = 5
)

type NDIS_802_11_POWER_MODE int32

const (
	Ndis802_11PowerModeCAM = 0
	Ndis802_11PowerModeMAX_PSP = 1
	Ndis802_11PowerModeFast_PSP = 2
	Ndis802_11PowerModeMax = 3
)

type NDIS_802_11_NETWORK_INFRASTRUCTURE int32

const (
	Ndis802_11IBSS = 0
	Ndis802_11Infrastructure = 1
	Ndis802_11AutoUnknown = 2
	Ndis802_11InfrastructureMax = 3
)

type NDIS_802_11_AUTHENTICATION_MODE int32

const (
	Ndis802_11AuthModeOpen = 0
	Ndis802_11AuthModeShared = 1
	Ndis802_11AuthModeAutoSwitch = 2
	Ndis802_11AuthModeWPA = 3
	Ndis802_11AuthModeWPAPSK = 4
	Ndis802_11AuthModeWPANone = 5
	Ndis802_11AuthModeWPA2 = 6
	Ndis802_11AuthModeWPA2PSK = 7
	Ndis802_11AuthModeWPA3 = 8
	Ndis802_11AuthModeWPA3SAE = 9
	Ndis802_11AuthModeMax = 10
)

type NDIS_802_11_PRIVACY_FILTER int32

const (
	Ndis802_11PrivFilterAcceptAll = 0
	Ndis802_11PrivFilter8021xWEP = 1
)

type NDIS_802_11_WEP_STATUS int32

const (
	Ndis802_11WEPEnabled = 0
	Ndis802_11Encryption1Enabled = 0
	Ndis802_11WEPDisabled = 1
	Ndis802_11EncryptionDisabled = 1
	Ndis802_11WEPKeyAbsent = 2
	Ndis802_11Encryption1KeyAbsent = 2
	Ndis802_11WEPNotSupported = 3
	Ndis802_11EncryptionNotSupported = 3
	Ndis802_11Encryption2Enabled = 4
	Ndis802_11Encryption2KeyAbsent = 5
	Ndis802_11Encryption3Enabled = 6
	Ndis802_11Encryption3KeyAbsent = 7
)

type NDIS_802_11_RELOAD_DEFAULTS int32

const (
	Ndis802_11ReloadWEPKeys = 0
)

type NDIS_802_11_MEDIA_STREAM_MODE int32

const (
	Ndis802_11MediaStreamOff = 0
	Ndis802_11MediaStreamOn = 1
)

type NDIS_802_11_RADIO_STATUS int32

const (
	Ndis802_11RadioStatusOn = 0
	Ndis802_11RadioStatusHardwareOff = 1
	Ndis802_11RadioStatusSoftwareOff = 2
	Ndis802_11RadioStatusHardwareSoftwareOff = 3
	Ndis802_11RadioStatusMax = 4
)

type OFFLOAD_OPERATION_E int32

const (
	AUTHENTICATE = 1
	ENCRYPT = 2
)

type OFFLOAD_CONF_ALGO int32

const (
	OFFLOAD_IPSEC_CONF_NONE = 0
	OFFLOAD_IPSEC_CONF_DES = 1
	OFFLOAD_IPSEC_CONF_RESERVED = 2
	OFFLOAD_IPSEC_CONF_3_DES = 3
	OFFLOAD_IPSEC_CONF_MAX = 4
)

type OFFLOAD_INTEGRITY_ALGO int32

const (
	OFFLOAD_IPSEC_INTEGRITY_NONE = 0
	OFFLOAD_IPSEC_INTEGRITY_MD5 = 1
	OFFLOAD_IPSEC_INTEGRITY_SHA = 2
	OFFLOAD_IPSEC_INTEGRITY_MAX = 3
)

type UDP_ENCAP_TYPE int32

const (
	OFFLOAD_IPSEC_UDPESP_ENCAPTYPE_IKE = 0
	OFFLOAD_IPSEC_UDPESP_ENCAPTYPE_OTHER = 1
)

type NDIS_MEDIUM int32

const (
	NdisMedium802_3 = 0
	NdisMedium802_5 = 1
	NdisMediumFddi = 2
	NdisMediumWan = 3
	NdisMediumLocalTalk = 4
	NdisMediumDix = 5
	NdisMediumArcnetRaw = 6
	NdisMediumArcnet878_2 = 7
	NdisMediumAtm = 8
	NdisMediumWirelessWan = 9
	NdisMediumIrda = 10
	NdisMediumBpc = 11
	NdisMediumCoWan = 12
	NdisMedium1394 = 13
	NdisMediumInfiniBand = 14
	NdisMediumTunnel = 15
	NdisMediumNative802_11 = 16
	NdisMediumLoopback = 17
	NdisMediumWiMAX = 18
	NdisMediumIP = 19
	NdisMediumMax = 20
)

type NDIS_PHYSICAL_MEDIUM int32

const (
	NdisPhysicalMediumUnspecified = 0
	NdisPhysicalMediumWirelessLan = 1
	NdisPhysicalMediumCableModem = 2
	NdisPhysicalMediumPhoneLine = 3
	NdisPhysicalMediumPowerLine = 4
	NdisPhysicalMediumDSL = 5
	NdisPhysicalMediumFibreChannel = 6
	NdisPhysicalMedium1394 = 7
	NdisPhysicalMediumWirelessWan = 8
	NdisPhysicalMediumNative802_11 = 9
	NdisPhysicalMediumBluetooth = 10
	NdisPhysicalMediumInfiniband = 11
	NdisPhysicalMediumWiMax = 12
	NdisPhysicalMediumUWB = 13
	NdisPhysicalMedium802_3 = 14
	NdisPhysicalMedium802_5 = 15
	NdisPhysicalMediumIrda = 16
	NdisPhysicalMediumWiredWAN = 17
	NdisPhysicalMediumWiredCoWan = 18
	NdisPhysicalMediumOther = 19
	NdisPhysicalMediumNative802_15_4 = 20
	NdisPhysicalMediumMax = 21
)

type NDIS_HARDWARE_STATUS int32

const (
	NdisHardwareStatusReady = 0
	NdisHardwareStatusInitializing = 1
	NdisHardwareStatusReset = 2
	NdisHardwareStatusClosing = 3
	NdisHardwareStatusNotReady = 4
)

type NDIS_DEVICE_POWER_STATE int32

const (
	NdisDeviceStateUnspecified = 0
	NdisDeviceStateD0 = 1
	NdisDeviceStateD1 = 2
	NdisDeviceStateD2 = 3
	NdisDeviceStateD3 = 4
	NdisDeviceStateMaximum = 5
)

type NDIS_FDDI_ATTACHMENT_TYPE int32

const (
	NdisFddiTypeIsolated = 1
	NdisFddiTypeLocalA = 2
	NdisFddiTypeLocalB = 3
	NdisFddiTypeLocalAB = 4
	NdisFddiTypeLocalS = 5
	NdisFddiTypeWrapA = 6
	NdisFddiTypeWrapB = 7
	NdisFddiTypeWrapAB = 8
	NdisFddiTypeWrapS = 9
	NdisFddiTypeCWrapA = 10
	NdisFddiTypeCWrapB = 11
	NdisFddiTypeCWrapS = 12
	NdisFddiTypeThrough = 13
)

type NDIS_FDDI_RING_MGT_STATE int32

const (
	NdisFddiRingIsolated = 1
	NdisFddiRingNonOperational = 2
	NdisFddiRingOperational = 3
	NdisFddiRingDetect = 4
	NdisFddiRingNonOperationalDup = 5
	NdisFddiRingOperationalDup = 6
	NdisFddiRingDirected = 7
	NdisFddiRingTrace = 8
)

type NDIS_FDDI_LCONNECTION_STATE int32

const (
	NdisFddiStateOff = 1
	NdisFddiStateBreak = 2
	NdisFddiStateTrace = 3
	NdisFddiStateConnect = 4
	NdisFddiStateNext = 5
	NdisFddiStateSignal = 6
	NdisFddiStateJoin = 7
	NdisFddiStateVerify = 8
	NdisFddiStateActive = 9
	NdisFddiStateMaintenance = 10
)

type NDIS_WAN_MEDIUM_SUBTYPE int32

const (
	NdisWanMediumHub = 0
	NdisWanMediumX_25 = 1
	NdisWanMediumIsdn = 2
	NdisWanMediumSerial = 3
	NdisWanMediumFrameRelay = 4
	NdisWanMediumAtm = 5
	NdisWanMediumSonet = 6
	NdisWanMediumSW56K = 7
	NdisWanMediumPPTP = 8
	NdisWanMediumL2TP = 9
	NdisWanMediumIrda = 10
	NdisWanMediumParallel = 11
	NdisWanMediumPppoe = 12
	NdisWanMediumSSTP = 13
	NdisWanMediumAgileVPN = 14
	NdisWanMediumGre = 15
	NdisWanMediumSubTypeMax = 16
)

type NDIS_WAN_HEADER_FORMAT int32

const (
	NdisWanHeaderNative = 0
	NdisWanHeaderEthernet = 1
)

type NDIS_WAN_QUALITY int32

const (
	NdisWanRaw = 0
	NdisWanErrorControl = 1
	NdisWanReliable = 2
)

type NDIS_802_5_RING_STATE int32

const (
	NdisRingStateOpened = 1
	NdisRingStateClosed = 2
	NdisRingStateOpening = 3
	NdisRingStateClosing = 4
	NdisRingStateOpenFailure = 5
	NdisRingStateRingFailure = 6
)

type NDIS_MEDIA_STATE int32

const (
	NdisMediaStateConnected = 0
	NdisMediaStateDisconnected = 1
)

type NDIS_SUPPORTED_PAUSE_FUNCTIONS int32

const (
	NdisPauseFunctionsUnsupported = 0
	NdisPauseFunctionsSendOnly = 1
	NdisPauseFunctionsReceiveOnly = 2
	NdisPauseFunctionsSendAndReceive = 3
	NdisPauseFunctionsUnknown = 4
)

type NDIS_PORT_TYPE int32

const (
	NdisPortTypeUndefined = 0
	NdisPortTypeBridge = 1
	NdisPortTypeRasConnection = 2
	NdisPortType8021xSupplicant = 3
	NdisPortTypeMax = 4
)

type NDIS_PORT_AUTHORIZATION_STATE int32

const (
	NdisPortAuthorizationUnknown = 0
	NdisPortAuthorized = 1
	NdisPortUnauthorized = 2
	NdisPortReauthorizing = 3
)

type NDIS_PORT_CONTROL_STATE int32

const (
	NdisPortControlStateUnknown = 0
	NdisPortControlStateControlled = 1
	NdisPortControlStateUncontrolled = 2
)

type NDIS_NETWORK_CHANGE_TYPE int32

const (
	NdisPossibleNetworkChange = 1
	NdisDefinitelyNetworkChange = 2
	NdisNetworkChangeFromMediaConnect = 3
	NdisNetworkChangeMax = 4
)

type NDIS_PROCESSOR_VENDOR int32

const (
	NdisProcessorVendorUnknown = 0
	NdisProcessorVendorGenuinIntel = 1
	NdisProcessorVendorGenuineIntel = 1
	NdisProcessorVendorAuthenticAMD = 2
)

type DOT11_PHY_TYPE int32

const (
	dot11_phy_type_unknown = 0
	dot11_phy_type_any = 0
	dot11_phy_type_fhss = 1
	dot11_phy_type_dsss = 2
	dot11_phy_type_irbaseband = 3
	dot11_phy_type_ofdm = 4
	dot11_phy_type_hrdsss = 5
	dot11_phy_type_erp = 6
	dot11_phy_type_ht = 7
	dot11_phy_type_vht = 8
	dot11_phy_type_dmg = 9
	dot11_phy_type_he = 10
	dot11_phy_type_IHV_start = -2147483648
	dot11_phy_type_IHV_end = -1
)

type DOT11_OFFLOAD_TYPE int32

const (
	dot11_offload_type_wep = 1
	dot11_offload_type_auth = 2
)

type DOT11_KEY_DIRECTION int32

const (
	dot11_key_direction_both = 1
	dot11_key_direction_inbound = 2
	dot11_key_direction_outbound = 3
)

type DOT11_SCAN_TYPE int32

const (
	dot11_scan_type_active = 1
	dot11_scan_type_passive = 2
	dot11_scan_type_auto = 3
	dot11_scan_type_forced = -2147483648
)

type CH_DESCRIPTION_TYPE int32

const (
	ch_description_type_logical = 1
	ch_description_type_center_frequency = 2
	ch_description_type_phy_specific = 3
)

type DOT11_UPDATE_IE_OP int32

const (
	dot11_update_ie_op_create_replace = 1
	dot11_update_ie_op_delete = 2
)

type DOT11_RESET_TYPE int32

const (
	dot11_reset_type_phy = 1
	dot11_reset_type_mac = 2
	dot11_reset_type_phy_and_mac = 3
)

type DOT11_POWER_MODE int32

const (
	dot11_power_mode_unknown = 0
	dot11_power_mode_active = 1
	dot11_power_mode_powersave = 2
)

type DOT11_TEMP_TYPE int32

const (
	dot11_temp_type_unknown = 0
	dot11_temp_type_1 = 1
	dot11_temp_type_2 = 2
)

type DOT11_DIVERSITY_SUPPORT int32

const (
	dot11_diversity_support_unknown = 0
	dot11_diversity_support_fixedlist = 1
	dot11_diversity_support_notsupported = 2
	dot11_diversity_support_dynamic = 3
)

type DOT11_HOP_ALGO_ADOPTED int32

const (
	dot11_hop_algo_current = 0
	dot11_hop_algo_hop_index = 1
	dot11_hop_algo_hcc = 2
)

type DOT11_AC_PARAM int32

const (
	dot11_AC_param_BE = 0
	dot11_AC_param_BK = 1
	dot11_AC_param_VI = 2
	dot11_AC_param_VO = 3
	dot11_AC_param_max = 4
)

type DOT11_DIRECTION int32

const (
	DOT11_DIR_INBOUND = 1
	DOT11_DIR_OUTBOUND = 2
	DOT11_DIR_BOTH = 3
)

type DOT11_ASSOCIATION_STATE int32

const (
	dot11_assoc_state_zero = 0
	dot11_assoc_state_unauth_unassoc = 1
	dot11_assoc_state_auth_unassoc = 2
	dot11_assoc_state_auth_assoc = 3
)

type DOT11_DS_INFO int32

const (
	DOT11_DS_CHANGED = 0
	DOT11_DS_UNCHANGED = 1
	DOT11_DS_UNKNOWN = 2
)

type DOT11_WPS_CONFIG_METHOD int32

const (
	DOT11_WPS_CONFIG_METHOD_NULL = 0
	DOT11_WPS_CONFIG_METHOD_DISPLAY = 8
	DOT11_WPS_CONFIG_METHOD_NFC_TAG = 32
	DOT11_WPS_CONFIG_METHOD_NFC_INTERFACE = 64
	DOT11_WPS_CONFIG_METHOD_PUSHBUTTON = 128
	DOT11_WPS_CONFIG_METHOD_KEYPAD = 256
	DOT11_WPS_CONFIG_METHOD_WFDS_DEFAULT = 4096
)

type DOT11_WPS_DEVICE_PASSWORD_ID int32

const (
	DOT11_WPS_PASSWORD_ID_DEFAULT = 0
	DOT11_WPS_PASSWORD_ID_USER_SPECIFIED = 1
	DOT11_WPS_PASSWORD_ID_MACHINE_SPECIFIED = 2
	DOT11_WPS_PASSWORD_ID_REKEY = 3
	DOT11_WPS_PASSWORD_ID_PUSHBUTTON = 4
	DOT11_WPS_PASSWORD_ID_REGISTRAR_SPECIFIED = 5
	DOT11_WPS_PASSWORD_ID_NFC_CONNECTION_HANDOVER = 7
	DOT11_WPS_PASSWORD_ID_WFD_SERVICES = 8
	DOT11_WPS_PASSWORD_ID_OOB_RANGE_MIN = 16
	DOT11_WPS_PASSWORD_ID_OOB_RANGE_MAX = 65535
)

type DOT11_ANQP_QUERY_RESULT int32

const (
	dot11_ANQP_query_result_success = 0
	dot11_ANQP_query_result_failure = 1
	dot11_ANQP_query_result_timed_out = 2
	dot11_ANQP_query_result_resources = 3
	dot11_ANQP_query_result_advertisement_protocol_not_supported_on_remote = 4
	dot11_ANQP_query_result_gas_protocol_failure = 5
	dot11_ANQP_query_result_advertisement_server_not_responding = 6
	dot11_ANQP_query_result_access_issues = 7
)

type DOT11_WFD_DISCOVER_TYPE int32

const (
	dot11_wfd_discover_type_scan_only = 1
	dot11_wfd_discover_type_find_only = 2
	dot11_wfd_discover_type_auto = 3
	dot11_wfd_discover_type_scan_social_channels = 4
	dot11_wfd_discover_type_forced = -2147483648
)

type DOT11_WFD_SCAN_TYPE int32

const (
	dot11_wfd_scan_type_active = 1
	dot11_wfd_scan_type_passive = 2
	dot11_wfd_scan_type_auto = 3
)

type DOT11_POWER_MODE_REASON int32

const (
	dot11_power_mode_reason_no_change = 0
	dot11_power_mode_reason_noncompliant_AP = 1
	dot11_power_mode_reason_legacy_WFD_device = 2
	dot11_power_mode_reason_compliant_AP = 3
	dot11_power_mode_reason_compliant_WFD_device = 4
	dot11_power_mode_reason_others = 5
)

type DOT11_MANUFACTURING_TEST_TYPE int32

const (
	dot11_manufacturing_test_unknown = 0
	dot11_manufacturing_test_self_start = 1
	dot11_manufacturing_test_self_query_result = 2
	dot11_manufacturing_test_rx = 3
	dot11_manufacturing_test_tx = 4
	dot11_manufacturing_test_query_adc = 5
	dot11_manufacturing_test_set_data = 6
	dot11_manufacturing_test_query_data = 7
	dot11_manufacturing_test_sleep = 8
	dot11_manufacturing_test_awake = 9
	dot11_manufacturing_test_IHV_start = -2147483648
	dot11_manufacturing_test_IHV_end = -1
)

type DOT11_MANUFACTURING_SELF_TEST_TYPE int32

const (
	DOT11_MANUFACTURING_SELF_TEST_TYPE_INTERFACE = 1
	DOT11_MANUFACTURING_SELF_TEST_TYPE_RF_INTERFACE = 2
	DOT11_MANUFACTURING_SELF_TEST_TYPE_BT_COEXISTENCE = 3
)

type DOT11_BAND int32

const (
	dot11_band_2p4g = 1
	dot11_band_4p9g = 2
	dot11_band_5g = 3
)

type DOT11_MANUFACTURING_CALLBACK_TYPE int32

const (
	dot11_manufacturing_callback_unknown = 0
	dot11_manufacturing_callback_self_test_complete = 1
	dot11_manufacturing_callback_sleep_complete = 2
	dot11_manufacturing_callback_IHV_start = -2147483648
	dot11_manufacturing_callback_IHV_end = -1
)

type WLAN_CONNECTION_MODE int32

const (
	wlan_connection_mode_profile = 0
	wlan_connection_mode_temporary_profile = 1
	wlan_connection_mode_discovery_secure = 2
	wlan_connection_mode_discovery_unsecure = 3
	wlan_connection_mode_auto = 4
	wlan_connection_mode_invalid = 5
)

type WLAN_INTERFACE_STATE int32

const (
	wlan_interface_state_not_ready = 0
	wlan_interface_state_connected = 1
	wlan_interface_state_ad_hoc_network_formed = 2
	wlan_interface_state_disconnecting = 3
	wlan_interface_state_disconnected = 4
	wlan_interface_state_associating = 5
	wlan_interface_state_discovering = 6
	wlan_interface_state_authenticating = 7
)

type WLAN_ADHOC_NETWORK_STATE int32

const (
	wlan_adhoc_network_state_formed = 0
	wlan_adhoc_network_state_connected = 1
)

type DOT11_RADIO_STATE int32

const (
	dot11_radio_state_unknown = 0
	dot11_radio_state_on = 1
	dot11_radio_state_off = 2
)

type WLAN_OPERATIONAL_STATE int32

const (
	wlan_operational_state_unknown = 0
	wlan_operational_state_off = 1
	wlan_operational_state_on = 2
	wlan_operational_state_going_off = 3
	wlan_operational_state_going_on = 4
)

type WLAN_INTERFACE_TYPE int32

const (
	wlan_interface_type_emulated_802_11 = 0
	wlan_interface_type_native_802_11 = 1
	wlan_interface_type_invalid = 2
)

type WLAN_POWER_SETTING int32

const (
	wlan_power_setting_no_saving = 0
	wlan_power_setting_low_saving = 1
	wlan_power_setting_medium_saving = 2
	wlan_power_setting_maximum_saving = 3
	wlan_power_setting_invalid = 4
)

type WLAN_NOTIFICATION_ACM int32

const (
	wlan_notification_acm_start = 0
	wlan_notification_acm_autoconf_enabled = 1
	wlan_notification_acm_autoconf_disabled = 2
	wlan_notification_acm_background_scan_enabled = 3
	wlan_notification_acm_background_scan_disabled = 4
	wlan_notification_acm_bss_type_change = 5
	wlan_notification_acm_power_setting_change = 6
	wlan_notification_acm_scan_complete = 7
	wlan_notification_acm_scan_fail = 8
	wlan_notification_acm_connection_start = 9
	wlan_notification_acm_connection_complete = 10
	wlan_notification_acm_connection_attempt_fail = 11
	wlan_notification_acm_filter_list_change = 12
	wlan_notification_acm_interface_arrival = 13
	wlan_notification_acm_interface_removal = 14
	wlan_notification_acm_profile_change = 15
	wlan_notification_acm_profile_name_change = 16
	wlan_notification_acm_profiles_exhausted = 17
	wlan_notification_acm_network_not_available = 18
	wlan_notification_acm_network_available = 19
	wlan_notification_acm_disconnecting = 20
	wlan_notification_acm_disconnected = 21
	wlan_notification_acm_adhoc_network_state_change = 22
	wlan_notification_acm_profile_unblocked = 23
	wlan_notification_acm_screen_power_change = 24
	wlan_notification_acm_profile_blocked = 25
	wlan_notification_acm_scan_list_refresh = 26
	wlan_notification_acm_operational_state_change = 27
	wlan_notification_acm_end = 28
)

type WLAN_NOTIFICATION_MSM int32

const (
	wlan_notification_msm_start = 0
	wlan_notification_msm_associating = 1
	wlan_notification_msm_associated = 2
	wlan_notification_msm_authenticating = 3
	wlan_notification_msm_connected = 4
	wlan_notification_msm_roaming_start = 5
	wlan_notification_msm_roaming_end = 6
	wlan_notification_msm_radio_state_change = 7
	wlan_notification_msm_signal_quality_change = 8
	wlan_notification_msm_disassociating = 9
	wlan_notification_msm_disconnected = 10
	wlan_notification_msm_peer_join = 11
	wlan_notification_msm_peer_leave = 12
	wlan_notification_msm_adapter_removal = 13
	wlan_notification_msm_adapter_operation_mode_change = 14
	wlan_notification_msm_link_degraded = 15
	wlan_notification_msm_link_improved = 16
	wlan_notification_msm_end = 17
)

type WLAN_NOTIFICATION_SECURITY int32

const (
	wlan_notification_security_start = 0
	wlan_notification_security_end = 1
)

type WLAN_OPCODE_VALUE_TYPE int32

const (
	wlan_opcode_value_type_query_only = 0
	wlan_opcode_value_type_set_by_group_policy = 1
	wlan_opcode_value_type_set_by_user = 2
	wlan_opcode_value_type_invalid = 3
)

type WLAN_INTF_OPCODE int32

const (
	wlan_intf_opcode_autoconf_start = 0
	wlan_intf_opcode_autoconf_enabled = 1
	wlan_intf_opcode_background_scan_enabled = 2
	wlan_intf_opcode_media_streaming_mode = 3
	wlan_intf_opcode_radio_state = 4
	wlan_intf_opcode_bss_type = 5
	wlan_intf_opcode_interface_state = 6
	wlan_intf_opcode_current_connection = 7
	wlan_intf_opcode_channel_number = 8
	wlan_intf_opcode_supported_infrastructure_auth_cipher_pairs = 9
	wlan_intf_opcode_supported_adhoc_auth_cipher_pairs = 10
	wlan_intf_opcode_supported_country_or_region_string_list = 11
	wlan_intf_opcode_current_operation_mode = 12
	wlan_intf_opcode_supported_safe_mode = 13
	wlan_intf_opcode_certified_safe_mode = 14
	wlan_intf_opcode_hosted_network_capable = 15
	wlan_intf_opcode_management_frame_protection_capable = 16
	wlan_intf_opcode_autoconf_end = 268435455
	wlan_intf_opcode_msm_start = 268435712
	wlan_intf_opcode_statistics = 268435713
	wlan_intf_opcode_rssi = 268435714
	wlan_intf_opcode_msm_end = 536870911
	wlan_intf_opcode_security_start = 536936448
	wlan_intf_opcode_security_end = 805306367
	wlan_intf_opcode_ihv_start = 805306368
	wlan_intf_opcode_ihv_end = 1073741823
)

type WLAN_AUTOCONF_OPCODE int32

const (
	wlan_autoconf_opcode_start = 0
	wlan_autoconf_opcode_show_denied_networks = 1
	wlan_autoconf_opcode_power_setting = 2
	wlan_autoconf_opcode_only_use_gp_profiles_for_allowed_networks = 3
	wlan_autoconf_opcode_allow_explicit_creds = 4
	wlan_autoconf_opcode_block_period = 5
	wlan_autoconf_opcode_allow_virtual_station_extensibility = 6
	wlan_autoconf_opcode_end = 7
)

type WLAN_IHV_CONTROL_TYPE int32

const (
	wlan_ihv_control_type_service = 0
	wlan_ihv_control_type_driver = 1
)

type WLAN_FILTER_LIST_TYPE int32

const (
	wlan_filter_list_type_gp_permit = 0
	wlan_filter_list_type_gp_deny = 1
	wlan_filter_list_type_user_permit = 2
	wlan_filter_list_type_user_deny = 3
)

type WLAN_SECURABLE_OBJECT int32

const (
	wlan_secure_permit_list = 0
	wlan_secure_deny_list = 1
	wlan_secure_ac_enabled = 2
	wlan_secure_bc_scan_enabled = 3
	wlan_secure_bss_type = 4
	wlan_secure_show_denied = 5
	wlan_secure_interface_properties = 6
	wlan_secure_ihv_control = 7
	wlan_secure_all_user_profiles_order = 8
	wlan_secure_add_new_all_user_profiles = 9
	wlan_secure_add_new_per_user_profiles = 10
	wlan_secure_media_streaming_mode_enabled = 11
	wlan_secure_current_operation_mode = 12
	wlan_secure_get_plaintext_key = 13
	wlan_secure_hosted_network_elevated_access = 14
	wlan_secure_virtual_station_extensibility = 15
	wlan_secure_wfd_elevated_access = 16
	WLAN_SECURABLE_OBJECT_COUNT = 17
)

type WFD_ROLE_TYPE int32

const (
	WFD_ROLE_TYPE_NONE = 0
	WFD_ROLE_TYPE_DEVICE = 1
	WFD_ROLE_TYPE_GROUP_OWNER = 2
	WFD_ROLE_TYPE_CLIENT = 4
	WFD_ROLE_TYPE_MAX = 5
)

type WL_DISPLAY_PAGES int32

const (
	WLConnectionPage = 0
	WLSecurityPage = 1
	WLAdvPage = 2
)

type WLAN_HOSTED_NETWORK_STATE int32

const (
	wlan_hosted_network_unavailable = 0
	wlan_hosted_network_idle = 1
	wlan_hosted_network_active = 2
)

type WLAN_HOSTED_NETWORK_REASON int32

const (
	wlan_hosted_network_reason_success = 0
	wlan_hosted_network_reason_unspecified = 1
	wlan_hosted_network_reason_bad_parameters = 2
	wlan_hosted_network_reason_service_shutting_down = 3
	wlan_hosted_network_reason_insufficient_resources = 4
	wlan_hosted_network_reason_elevation_required = 5
	wlan_hosted_network_reason_read_only = 6
	wlan_hosted_network_reason_persistence_failed = 7
	wlan_hosted_network_reason_crypt_error = 8
	wlan_hosted_network_reason_impersonation = 9
	wlan_hosted_network_reason_stop_before_start = 10
	wlan_hosted_network_reason_interface_available = 11
	wlan_hosted_network_reason_interface_unavailable = 12
	wlan_hosted_network_reason_miniport_stopped = 13
	wlan_hosted_network_reason_miniport_started = 14
	wlan_hosted_network_reason_incompatible_connection_started = 15
	wlan_hosted_network_reason_incompatible_connection_stopped = 16
	wlan_hosted_network_reason_user_action = 17
	wlan_hosted_network_reason_client_abort = 18
	wlan_hosted_network_reason_ap_start_failed = 19
	wlan_hosted_network_reason_peer_arrived = 20
	wlan_hosted_network_reason_peer_departed = 21
	wlan_hosted_network_reason_peer_timeout = 22
	wlan_hosted_network_reason_gp_denied = 23
	wlan_hosted_network_reason_service_unavailable = 24
	wlan_hosted_network_reason_device_change = 25
	wlan_hosted_network_reason_properties_change = 26
	wlan_hosted_network_reason_virtual_station_blocking_use = 27
	wlan_hosted_network_reason_service_available_on_virtual_station = 28
)

type WLAN_HOSTED_NETWORK_PEER_AUTH_STATE int32

const (
	wlan_hosted_network_peer_state_invalid = 0
	wlan_hosted_network_peer_state_authenticated = 1
)

type WLAN_HOSTED_NETWORK_NOTIFICATION_CODE int32

const (
	wlan_hosted_network_state_change = 4096
	wlan_hosted_network_peer_state_change = 4097
	wlan_hosted_network_radio_state_change = 4098
)

type WLAN_HOSTED_NETWORK_OPCODE int32

const (
	wlan_hosted_network_opcode_connection_settings = 0
	wlan_hosted_network_opcode_security_settings = 1
	wlan_hosted_network_opcode_station_profile = 2
	wlan_hosted_network_opcode_enable = 3
)

type ONEX_AUTH_IDENTITY int32

const (
	OneXAuthIdentityNone = 0
	OneXAuthIdentityMachine = 1
	OneXAuthIdentityUser = 2
	OneXAuthIdentityExplicitUser = 3
	OneXAuthIdentityGuest = 4
	OneXAuthIdentityInvalid = 5
)

type ONEX_AUTH_STATUS int32

const (
	OneXAuthNotStarted = 0
	OneXAuthInProgress = 1
	OneXAuthNoAuthenticatorFound = 2
	OneXAuthSuccess = 3
	OneXAuthFailure = 4
	OneXAuthInvalid = 5
)

type ONEX_REASON_CODE int32

const (
	ONEX_REASON_CODE_SUCCESS = 0
	ONEX_REASON_START = 327680
	ONEX_UNABLE_TO_IDENTIFY_USER = 327681
	ONEX_IDENTITY_NOT_FOUND = 327682
	ONEX_UI_DISABLED = 327683
	ONEX_UI_FAILURE = 327684
	ONEX_EAP_FAILURE_RECEIVED = 327685
	ONEX_AUTHENTICATOR_NO_LONGER_PRESENT = 327686
	ONEX_NO_RESPONSE_TO_IDENTITY = 327687
	ONEX_PROFILE_VERSION_NOT_SUPPORTED = 327688
	ONEX_PROFILE_INVALID_LENGTH = 327689
	ONEX_PROFILE_DISALLOWED_EAP_TYPE = 327690
	ONEX_PROFILE_INVALID_EAP_TYPE_OR_FLAG = 327691
	ONEX_PROFILE_INVALID_ONEX_FLAGS = 327692
	ONEX_PROFILE_INVALID_TIMER_VALUE = 327693
	ONEX_PROFILE_INVALID_SUPPLICANT_MODE = 327694
	ONEX_PROFILE_INVALID_AUTH_MODE = 327695
	ONEX_PROFILE_INVALID_EAP_CONNECTION_PROPERTIES = 327696
	ONEX_UI_CANCELLED = 327697
	ONEX_PROFILE_INVALID_EXPLICIT_CREDENTIALS = 327698
	ONEX_PROFILE_EXPIRED_EXPLICIT_CREDENTIALS = 327699
	ONEX_UI_NOT_PERMITTED = 327700
)

type ONEX_NOTIFICATION_TYPE int32

const (
	OneXPublicNotificationBase = 0
	OneXNotificationTypeResultUpdate = 1
	OneXNotificationTypeAuthRestarted = 2
	OneXNotificationTypeEventInvalid = 3
	OneXNumNotifications = 3
)

type ONEX_AUTH_RESTART_REASON int32

const (
	OneXRestartReasonPeerInitiated = 0
	OneXRestartReasonMsmInitiated = 1
	OneXRestartReasonOneXHeldStateTimeout = 2
	OneXRestartReasonOneXAuthTimeout = 3
	OneXRestartReasonOneXConfigurationChanged = 4
	OneXRestartReasonOneXUserChanged = 5
	OneXRestartReasonQuarantineStateChanged = 6
	OneXRestartReasonAltCredsTrial = 7
	OneXRestartReasonInvalid = 8
)

type ONEX_EAP_METHOD_BACKEND_SUPPORT int32

const (
	OneXEapMethodBackendSupportUnknown = 0
	OneXEapMethodBackendSupported = 1
	OneXEapMethodBackendUnsupported = 2
)

type DOT11_ADHOC_CIPHER_ALGORITHM int32

const (
	DOT11_ADHOC_CIPHER_ALGO_INVALID = -1
	DOT11_ADHOC_CIPHER_ALGO_NONE = 0
	DOT11_ADHOC_CIPHER_ALGO_CCMP = 4
	DOT11_ADHOC_CIPHER_ALGO_WEP = 257
)

type DOT11_ADHOC_AUTH_ALGORITHM int32

const (
	DOT11_ADHOC_AUTH_ALGO_INVALID = -1
	DOT11_ADHOC_AUTH_ALGO_80211_OPEN = 1
	DOT11_ADHOC_AUTH_ALGO_RSNA_PSK = 7
)

type DOT11_ADHOC_NETWORK_CONNECTION_STATUS int32

const (
	DOT11_ADHOC_NETWORK_CONNECTION_STATUS_INVALID = 0
	DOT11_ADHOC_NETWORK_CONNECTION_STATUS_DISCONNECTED = 11
	DOT11_ADHOC_NETWORK_CONNECTION_STATUS_CONNECTING = 12
	DOT11_ADHOC_NETWORK_CONNECTION_STATUS_CONNECTED = 13
	DOT11_ADHOC_NETWORK_CONNECTION_STATUS_FORMED = 14
)

type DOT11_ADHOC_CONNECT_FAIL_REASON int32

const (
	DOT11_ADHOC_CONNECT_FAIL_DOMAIN_MISMATCH = 0
	DOT11_ADHOC_CONNECT_FAIL_PASSPHRASE_MISMATCH = 1
	DOT11_ADHOC_CONNECT_FAIL_OTHER = 2
)

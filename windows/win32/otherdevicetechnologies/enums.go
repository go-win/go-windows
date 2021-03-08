// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package otherdevicetechnologies implements the Windows.Win32.OtherDeviceTechnologies namespace.
package otherdevicetechnologies

type PropertyConstraint int32

const (
	QC_EQUALS             PropertyConstraint = 0
	QC_NOTEQUAL           PropertyConstraint = 1
	QC_LESSTHAN           PropertyConstraint = 2
	QC_LESSTHANOREQUAL    PropertyConstraint = 3
	QC_GREATERTHAN        PropertyConstraint = 4
	QC_GREATERTHANOREQUAL PropertyConstraint = 5
	QC_STARTSWITH         PropertyConstraint = 6
	QC_EXISTS             PropertyConstraint = 7
	QC_DOESNOTEXIST       PropertyConstraint = 8
	QC_CONTAINS           PropertyConstraint = 9
)

type SystemVisibilityFlags int32

const (
	SVF_SYSTEM SystemVisibilityFlags = 0
	SVF_USER   SystemVisibilityFlags = 1
)

type QueryUpdateAction int32

const (
	QUA_ADD    QueryUpdateAction = 0
	QUA_REMOVE QueryUpdateAction = 1
	QUA_CHANGE QueryUpdateAction = 2
)

type QueryCategoryType int32

const (
	QCT_PROVIDER QueryCategoryType = 0
	QCT_LAYERED  QueryCategoryType = 1
)

type __MIDL___MIDL_itf_wsdxml_0000_0000_0001 int32

const (
	OpNone                  __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 0
	OpEndOfTable            __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 1
	OpBeginElement_         __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 2
	OpBeginAnyElement       __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 3
	OpEndElement            __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 4
	OpElement_              __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 5
	OpAnyElement            __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 6
	OpAnyElements           __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 7
	OpAnyText               __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 8
	OpAttribute_            __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 9
	OpBeginChoice           __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 10
	OpEndChoice             __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 11
	OpBeginSequence         __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 12
	OpEndSequence           __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 13
	OpBeginAll              __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 14
	OpEndAll                __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 15
	OpAnything              __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 16
	OpAnyNumber             __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 17
	OpOneOrMore             __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 18
	OpOptional              __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 19
	OpFormatBool_           __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 20
	OpFormatInt8_           __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 21
	OpFormatInt16_          __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 22
	OpFormatInt32_          __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 23
	OpFormatInt64_          __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 24
	OpFormatUInt8_          __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 25
	OpFormatUInt16_         __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 26
	OpFormatUInt32_         __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 27
	OpFormatUInt64_         __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 28
	OpFormatUnicodeString_  __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 29
	OpFormatDom_            __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 30
	OpFormatStruct_         __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 31
	OpFormatUri_            __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 32
	OpFormatUuidUri_        __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 33
	OpFormatName_           __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 34
	OpFormatListInsertTail_ __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 35
	OpFormatType_           __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 36
	OpFormatDynamicType_    __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 37
	OpFormatLookupType_     __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 38
	OpFormatDuration_       __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 39
	OpFormatDateTime_       __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 40
	OpFormatFloat_          __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 41
	OpFormatDouble_         __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 42
	OpProcess_              __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 43
	OpQualifiedAttribute_   __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 44
	OpFormatXMLDeclaration_ __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 45
	OpFormatMax             __MIDL___MIDL_itf_wsdxml_0000_0000_0001 = 46
)

type WSD_CONFIG_PARAM_TYPE int32

const (
	WSD_CONFIG_MAX_INBOUND_MESSAGE_SIZE                  WSD_CONFIG_PARAM_TYPE = 1
	WSD_CONFIG_MAX_OUTBOUND_MESSAGE_SIZE                 WSD_CONFIG_PARAM_TYPE = 2
	WSD_SECURITY_SSL_CERT_FOR_CLIENT_AUTH                WSD_CONFIG_PARAM_TYPE = 3
	WSD_SECURITY_SSL_SERVER_CERT_VALIDATION              WSD_CONFIG_PARAM_TYPE = 4
	WSD_SECURITY_SSL_CLIENT_CERT_VALIDATION              WSD_CONFIG_PARAM_TYPE = 5
	WSD_SECURITY_SSL_NEGOTIATE_CLIENT_CERT               WSD_CONFIG_PARAM_TYPE = 6
	WSD_SECURITY_COMPACTSIG_SIGNING_CERT                 WSD_CONFIG_PARAM_TYPE = 7
	WSD_SECURITY_COMPACTSIG_VALIDATION                   WSD_CONFIG_PARAM_TYPE = 8
	WSD_CONFIG_HOSTING_ADDRESSES                         WSD_CONFIG_PARAM_TYPE = 9
	WSD_CONFIG_DEVICE_ADDRESSES                          WSD_CONFIG_PARAM_TYPE = 10
	WSD_SECURITY_REQUIRE_HTTP_CLIENT_AUTH                WSD_CONFIG_PARAM_TYPE = 11
	WSD_SECURITY_REQUIRE_CLIENT_CERT_OR_HTTP_CLIENT_AUTH WSD_CONFIG_PARAM_TYPE = 12
	WSD_SECURITY_USE_HTTP_CLIENT_AUTH                    WSD_CONFIG_PARAM_TYPE = 13
)

type WSDUdpMessageType int32

const (
	ONE_WAY WSDUdpMessageType = 0
	TWO_WAY WSDUdpMessageType = 1
)

type DeviceDiscoveryMechanism int32

const (
	MulticastDiscovery      DeviceDiscoveryMechanism = 0
	DirectedDiscovery       DeviceDiscoveryMechanism = 1
	SecureDirectedDiscovery DeviceDiscoveryMechanism = 2
)

type WSD_PROTOCOL_TYPE int32

const (
	WSD_PT_NONE  WSD_PROTOCOL_TYPE = 0
	WSD_PT_UDP   WSD_PROTOCOL_TYPE = 1
	WSD_PT_HTTP  WSD_PROTOCOL_TYPE = 2
	WSD_PT_HTTPS WSD_PROTOCOL_TYPE = 4
	WSD_PT_ALL   WSD_PROTOCOL_TYPE = 255
)

type WSDEventType int32

const (
	WSDET_NONE                 WSDEventType = 0
	WSDET_INCOMING_MESSAGE     WSDEventType = 1
	WSDET_INCOMING_FAULT       WSDEventType = 2
	WSDET_TRANSMISSION_FAILURE WSDEventType = 3
	WSDET_RESPONSE_TIMEOUT     WSDEventType = 4
)

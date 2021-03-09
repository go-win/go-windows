// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package etw implements the Windows.Win32.Etw namespace.
package etw

type WMIDPREQUESTCODE int32

const (
	WMI_GET_ALL_DATA        WMIDPREQUESTCODE = 0
	WMI_GET_SINGLE_INSTANCE WMIDPREQUESTCODE = 1
	WMI_SET_SINGLE_INSTANCE WMIDPREQUESTCODE = 2
	WMI_SET_SINGLE_ITEM     WMIDPREQUESTCODE = 3
	WMI_ENABLE_EVENTS       WMIDPREQUESTCODE = 4
	WMI_DISABLE_EVENTS      WMIDPREQUESTCODE = 5
	WMI_ENABLE_COLLECTION   WMIDPREQUESTCODE = 6
	WMI_DISABLE_COLLECTION  WMIDPREQUESTCODE = 7
	WMI_REGINFO             WMIDPREQUESTCODE = 8
	WMI_EXECUTE_METHOD      WMIDPREQUESTCODE = 9
	WMI_CAPTURE_STATE       WMIDPREQUESTCODE = 10
)

type ETW_COMPRESSION_RESUMPTION_MODE int32

const (
	EtwCompressionModeRestart   ETW_COMPRESSION_RESUMPTION_MODE = 0
	EtwCompressionModeNoDisable ETW_COMPRESSION_RESUMPTION_MODE = 1
	EtwCompressionModeNoRestart ETW_COMPRESSION_RESUMPTION_MODE = 2
)

type TRACE_QUERY_INFO_CLASS int32

const (
	TraceGuidQueryList                TRACE_QUERY_INFO_CLASS = 0
	TraceGuidQueryInfo                TRACE_QUERY_INFO_CLASS = 1
	TraceGuidQueryProcess             TRACE_QUERY_INFO_CLASS = 2
	TraceStackTracingInfo             TRACE_QUERY_INFO_CLASS = 3
	TraceSystemTraceEnableFlagsInfo   TRACE_QUERY_INFO_CLASS = 4
	TraceSampledProfileIntervalInfo   TRACE_QUERY_INFO_CLASS = 5
	TraceProfileSourceConfigInfo      TRACE_QUERY_INFO_CLASS = 6
	TraceProfileSourceListInfo        TRACE_QUERY_INFO_CLASS = 7
	TracePmcEventListInfo             TRACE_QUERY_INFO_CLASS = 8
	TracePmcCounterListInfo           TRACE_QUERY_INFO_CLASS = 9
	TraceSetDisallowList              TRACE_QUERY_INFO_CLASS = 10
	TraceVersionInfo                  TRACE_QUERY_INFO_CLASS = 11
	TraceGroupQueryList               TRACE_QUERY_INFO_CLASS = 12
	TraceGroupQueryInfo               TRACE_QUERY_INFO_CLASS = 13
	TraceDisallowListQuery            TRACE_QUERY_INFO_CLASS = 14
	TraceInfoReserved15               TRACE_QUERY_INFO_CLASS = 15
	TracePeriodicCaptureStateListInfo TRACE_QUERY_INFO_CLASS = 16
	TracePeriodicCaptureStateInfo     TRACE_QUERY_INFO_CLASS = 17
	TraceProviderBinaryTracking       TRACE_QUERY_INFO_CLASS = 18
	TraceMaxLoggersQuery              TRACE_QUERY_INFO_CLASS = 19
	TraceLbrConfigurationInfo         TRACE_QUERY_INFO_CLASS = 20
	TraceLbrEventListInfo             TRACE_QUERY_INFO_CLASS = 21
	TraceMaxPmcCounterQuery           TRACE_QUERY_INFO_CLASS = 22
	MaxTraceSetInfoClass              TRACE_QUERY_INFO_CLASS = 23
)

type ETW_PROCESS_HANDLE_INFO_TYPE int32

const (
	EtwQueryPartitionInformation   ETW_PROCESS_HANDLE_INFO_TYPE = 1
	EtwQueryPartitionInformationV2 ETW_PROCESS_HANDLE_INFO_TYPE = 2
	EtwQueryLastDroppedTimes       ETW_PROCESS_HANDLE_INFO_TYPE = 3
	EtwQueryProcessHandleInfoMax   ETW_PROCESS_HANDLE_INFO_TYPE = 4
)

type EVENT_INFO_CLASS int32

const (
	EventProviderBinaryTrackInfo   EVENT_INFO_CLASS = 0
	EventProviderSetReserved1      EVENT_INFO_CLASS = 1
	EventProviderSetTraits         EVENT_INFO_CLASS = 2
	EventProviderUseDescriptorType EVENT_INFO_CLASS = 3
	MaxEventInfo                   EVENT_INFO_CLASS = 4
)

type ETW_PROVIDER_TRAIT_TYPE int32

const (
	EtwProviderTraitTypeGroup  ETW_PROVIDER_TRAIT_TYPE = 1
	EtwProviderTraitDecodeGuid ETW_PROVIDER_TRAIT_TYPE = 2
	EtwProviderTraitTypeMax    ETW_PROVIDER_TRAIT_TYPE = 3
)

type EVENTSECURITYOPERATION int32

const (
	EventSecuritySetDACL EVENTSECURITYOPERATION = 0
	EventSecuritySetSACL EVENTSECURITYOPERATION = 1
	EventSecurityAddDACL EVENTSECURITYOPERATION = 2
	EventSecurityAddSACL EVENTSECURITYOPERATION = 3
	EventSecurityMax     EVENTSECURITYOPERATION = 4
)

type MAP_FLAGS int32

const (
	EVENTMAP_INFO_FLAG_MANIFEST_VALUEMAP   MAP_FLAGS = 1
	EVENTMAP_INFO_FLAG_MANIFEST_BITMAP     MAP_FLAGS = 2
	EVENTMAP_INFO_FLAG_MANIFEST_PATTERNMAP MAP_FLAGS = 4
	EVENTMAP_INFO_FLAG_WBEM_VALUEMAP       MAP_FLAGS = 8
	EVENTMAP_INFO_FLAG_WBEM_BITMAP         MAP_FLAGS = 16
	EVENTMAP_INFO_FLAG_WBEM_FLAG           MAP_FLAGS = 32
	EVENTMAP_INFO_FLAG_WBEM_NO_MAP         MAP_FLAGS = 64
)

type MAP_VALUETYPE int32

const (
	EVENTMAP_ENTRY_VALUETYPE_ULONG  MAP_VALUETYPE = 0
	EVENTMAP_ENTRY_VALUETYPE_STRING MAP_VALUETYPE = 1
)

type TDH_IN_TYPE int32

const (
	TDH_INTYPE_NULL                        TDH_IN_TYPE = 0
	TDH_INTYPE_UNICODESTRING               TDH_IN_TYPE = 1
	TDH_INTYPE_ANSISTRING                  TDH_IN_TYPE = 2
	TDH_INTYPE_INT8                        TDH_IN_TYPE = 3
	TDH_INTYPE_UINT8                       TDH_IN_TYPE = 4
	TDH_INTYPE_INT16                       TDH_IN_TYPE = 5
	TDH_INTYPE_UINT16                      TDH_IN_TYPE = 6
	TDH_INTYPE_INT32                       TDH_IN_TYPE = 7
	TDH_INTYPE_UINT32                      TDH_IN_TYPE = 8
	TDH_INTYPE_INT64                       TDH_IN_TYPE = 9
	TDH_INTYPE_UINT64                      TDH_IN_TYPE = 10
	TDH_INTYPE_FLOAT                       TDH_IN_TYPE = 11
	TDH_INTYPE_DOUBLE                      TDH_IN_TYPE = 12
	TDH_INTYPE_BOOLEAN                     TDH_IN_TYPE = 13
	TDH_INTYPE_BINARY                      TDH_IN_TYPE = 14
	TDH_INTYPE_GUID                        TDH_IN_TYPE = 15
	TDH_INTYPE_POINTER                     TDH_IN_TYPE = 16
	TDH_INTYPE_FILETIME                    TDH_IN_TYPE = 17
	TDH_INTYPE_SYSTEMTIME                  TDH_IN_TYPE = 18
	TDH_INTYPE_SID                         TDH_IN_TYPE = 19
	TDH_INTYPE_HEXINT32                    TDH_IN_TYPE = 20
	TDH_INTYPE_HEXINT64                    TDH_IN_TYPE = 21
	TDH_INTYPE_MANIFEST_COUNTEDSTRING      TDH_IN_TYPE = 22
	TDH_INTYPE_MANIFEST_COUNTEDANSISTRING  TDH_IN_TYPE = 23
	TDH_INTYPE_RESERVED24                  TDH_IN_TYPE = 24
	TDH_INTYPE_MANIFEST_COUNTEDBINARY      TDH_IN_TYPE = 25
	TDH_INTYPE_COUNTEDSTRING               TDH_IN_TYPE = 300
	TDH_INTYPE_COUNTEDANSISTRING           TDH_IN_TYPE = 301
	TDH_INTYPE_REVERSEDCOUNTEDSTRING       TDH_IN_TYPE = 302
	TDH_INTYPE_REVERSEDCOUNTEDANSISTRING   TDH_IN_TYPE = 303
	TDH_INTYPE_NONNULLTERMINATEDSTRING     TDH_IN_TYPE = 304
	TDH_INTYPE_NONNULLTERMINATEDANSISTRING TDH_IN_TYPE = 305
	TDH_INTYPE_UNICODECHAR                 TDH_IN_TYPE = 306
	TDH_INTYPE_ANSICHAR                    TDH_IN_TYPE = 307
	TDH_INTYPE_SIZET                       TDH_IN_TYPE = 308
	TDH_INTYPE_HEXDUMP                     TDH_IN_TYPE = 309
	TDH_INTYPE_WBEMSID                     TDH_IN_TYPE = 310
)

type TDH_OUT_TYPE int32

const (
	TDH_OUTTYPE_NULL                         TDH_OUT_TYPE = 0
	TDH_OUTTYPE_STRING                       TDH_OUT_TYPE = 1
	TDH_OUTTYPE_DATETIME                     TDH_OUT_TYPE = 2
	TDH_OUTTYPE_BYTE                         TDH_OUT_TYPE = 3
	TDH_OUTTYPE_UNSIGNEDBYTE                 TDH_OUT_TYPE = 4
	TDH_OUTTYPE_SHORT                        TDH_OUT_TYPE = 5
	TDH_OUTTYPE_UNSIGNEDSHORT                TDH_OUT_TYPE = 6
	TDH_OUTTYPE_INT                          TDH_OUT_TYPE = 7
	TDH_OUTTYPE_UNSIGNEDINT                  TDH_OUT_TYPE = 8
	TDH_OUTTYPE_LONG                         TDH_OUT_TYPE = 9
	TDH_OUTTYPE_UNSIGNEDLONG                 TDH_OUT_TYPE = 10
	TDH_OUTTYPE_FLOAT                        TDH_OUT_TYPE = 11
	TDH_OUTTYPE_DOUBLE                       TDH_OUT_TYPE = 12
	TDH_OUTTYPE_BOOLEAN                      TDH_OUT_TYPE = 13
	TDH_OUTTYPE_GUID                         TDH_OUT_TYPE = 14
	TDH_OUTTYPE_HEXBINARY                    TDH_OUT_TYPE = 15
	TDH_OUTTYPE_HEXINT8                      TDH_OUT_TYPE = 16
	TDH_OUTTYPE_HEXINT16                     TDH_OUT_TYPE = 17
	TDH_OUTTYPE_HEXINT32                     TDH_OUT_TYPE = 18
	TDH_OUTTYPE_HEXINT64                     TDH_OUT_TYPE = 19
	TDH_OUTTYPE_PID                          TDH_OUT_TYPE = 20
	TDH_OUTTYPE_TID                          TDH_OUT_TYPE = 21
	TDH_OUTTYPE_PORT                         TDH_OUT_TYPE = 22
	TDH_OUTTYPE_IPV4                         TDH_OUT_TYPE = 23
	TDH_OUTTYPE_IPV6                         TDH_OUT_TYPE = 24
	TDH_OUTTYPE_SOCKETADDRESS                TDH_OUT_TYPE = 25
	TDH_OUTTYPE_CIMDATETIME                  TDH_OUT_TYPE = 26
	TDH_OUTTYPE_ETWTIME                      TDH_OUT_TYPE = 27
	TDH_OUTTYPE_XML                          TDH_OUT_TYPE = 28
	TDH_OUTTYPE_ERRORCODE                    TDH_OUT_TYPE = 29
	TDH_OUTTYPE_WIN32ERROR                   TDH_OUT_TYPE = 30
	TDH_OUTTYPE_NTSTATUS                     TDH_OUT_TYPE = 31
	TDH_OUTTYPE_HRESULT                      TDH_OUT_TYPE = 32
	TDH_OUTTYPE_CULTURE_INSENSITIVE_DATETIME TDH_OUT_TYPE = 33
	TDH_OUTTYPE_JSON                         TDH_OUT_TYPE = 34
	TDH_OUTTYPE_UTF8                         TDH_OUT_TYPE = 35
	TDH_OUTTYPE_PKCS7_WITH_TYPE_INFO         TDH_OUT_TYPE = 36
	TDH_OUTTYPE_CODE_POINTER                 TDH_OUT_TYPE = 37
	TDH_OUTTYPE_DATETIME_UTC                 TDH_OUT_TYPE = 38
	TDH_OUTTYPE_REDUCEDSTRING                TDH_OUT_TYPE = 300
	TDH_OUTTYPE_NOPRINT                      TDH_OUT_TYPE = 301
)

type PROPERTY_FLAGS int32

const (
	PropertyStruct           PROPERTY_FLAGS = 1
	PropertyParamLength      PROPERTY_FLAGS = 2
	PropertyParamCount       PROPERTY_FLAGS = 4
	PropertyWBEMXmlFragment  PROPERTY_FLAGS = 8
	PropertyParamFixedLength PROPERTY_FLAGS = 16
	PropertyParamFixedCount  PROPERTY_FLAGS = 32
	PropertyHasTags          PROPERTY_FLAGS = 64
	PropertyHasCustomSchema  PROPERTY_FLAGS = 128
)

type DECODING_SOURCE int32

const (
	DecodingSourceXMLFile DECODING_SOURCE = 0
	DecodingSourceWbem    DECODING_SOURCE = 1
	DecodingSourceWPP     DECODING_SOURCE = 2
	DecodingSourceTlg     DECODING_SOURCE = 3
	DecodingSourceMax     DECODING_SOURCE = 4
)

type TEMPLATE_FLAGS int32

const (
	TEMPLATE_EVENT_DATA   TEMPLATE_FLAGS = 1
	TEMPLATE_USER_DATA    TEMPLATE_FLAGS = 2
	TEMPLATE_CONTROL_GUID TEMPLATE_FLAGS = 4
)

type PAYLOAD_OPERATOR int32

const (
	PAYLOADFIELD_EQ            PAYLOAD_OPERATOR = 0
	PAYLOADFIELD_NE            PAYLOAD_OPERATOR = 1
	PAYLOADFIELD_LE            PAYLOAD_OPERATOR = 2
	PAYLOADFIELD_GT            PAYLOAD_OPERATOR = 3
	PAYLOADFIELD_LT            PAYLOAD_OPERATOR = 4
	PAYLOADFIELD_GE            PAYLOAD_OPERATOR = 5
	PAYLOADFIELD_BETWEEN       PAYLOAD_OPERATOR = 6
	PAYLOADFIELD_NOTBETWEEN    PAYLOAD_OPERATOR = 7
	PAYLOADFIELD_MODULO        PAYLOAD_OPERATOR = 8
	PAYLOADFIELD_CONTAINS      PAYLOAD_OPERATOR = 20
	PAYLOADFIELD_DOESNTCONTAIN PAYLOAD_OPERATOR = 21
	PAYLOADFIELD_IS            PAYLOAD_OPERATOR = 30
	PAYLOADFIELD_ISNOT         PAYLOAD_OPERATOR = 31
	PAYLOADFIELD_INVALID       PAYLOAD_OPERATOR = 32
)

type EVENT_FIELD_TYPE int32

const (
	EventKeywordInformation EVENT_FIELD_TYPE = 0
	EventLevelInformation   EVENT_FIELD_TYPE = 1
	EventChannelInformation EVENT_FIELD_TYPE = 2
	EventTaskInformation    EVENT_FIELD_TYPE = 3
	EventOpcodeInformation  EVENT_FIELD_TYPE = 4
	EventInformationMax     EVENT_FIELD_TYPE = 5
)

type TDH_CONTEXT_TYPE int32

const (
	TDH_CONTEXT_WPP_TMFFILE       TDH_CONTEXT_TYPE = 0
	TDH_CONTEXT_WPP_TMFSEARCHPATH TDH_CONTEXT_TYPE = 1
	TDH_CONTEXT_WPP_GMT           TDH_CONTEXT_TYPE = 2
	TDH_CONTEXT_POINTERSIZE       TDH_CONTEXT_TYPE = 3
	TDH_CONTEXT_PDB_PATH          TDH_CONTEXT_TYPE = 4
	TDH_CONTEXT_MAXIMUM           TDH_CONTEXT_TYPE = 5
)

// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package alljoyn implements the Windows.Win32.AllJoyn namespace.
package alljoyn

type alljoyn_about_announceflag int32

const (
	UNANNOUNCED = 0
	ANNOUNCED = 1
)

type QStatus int32

const (
	ER_OK = 0
	ER_FAIL = 1
	ER_UTF_CONVERSION_FAILED = 2
	ER_BUFFER_TOO_SMALL = 3
	ER_OS_ERROR = 4
	ER_OUT_OF_MEMORY = 5
	ER_SOCKET_BIND_ERROR = 6
	ER_INIT_FAILED = 7
	ER_WOULDBLOCK = 8
	ER_NOT_IMPLEMENTED = 9
	ER_TIMEOUT = 10
	ER_SOCK_OTHER_END_CLOSED = 11
	ER_BAD_ARG_1 = 12
	ER_BAD_ARG_2 = 13
	ER_BAD_ARG_3 = 14
	ER_BAD_ARG_4 = 15
	ER_BAD_ARG_5 = 16
	ER_BAD_ARG_6 = 17
	ER_BAD_ARG_7 = 18
	ER_BAD_ARG_8 = 19
	ER_INVALID_ADDRESS = 20
	ER_INVALID_DATA = 21
	ER_READ_ERROR = 22
	ER_WRITE_ERROR = 23
	ER_OPEN_FAILED = 24
	ER_PARSE_ERROR = 25
	ER_END_OF_DATA = 26
	ER_CONN_REFUSED = 27
	ER_BAD_ARG_COUNT = 28
	ER_WARNING = 29
	ER_EOF = 30
	ER_DEADLOCK = 31
	ER_COMMON_ERRORS = 4096
	ER_STOPPING_THREAD = 4097
	ER_ALERTED_THREAD = 4098
	ER_XML_MALFORMED = 4099
	ER_AUTH_FAIL = 4100
	ER_AUTH_USER_REJECT = 4101
	ER_NO_SUCH_ALARM = 4102
	ER_TIMER_FALLBEHIND = 4103
	ER_SSL_ERRORS = 4104
	ER_SSL_INIT = 4105
	ER_SSL_CONNECT = 4106
	ER_SSL_VERIFY = 4107
	ER_EXTERNAL_THREAD = 4108
	ER_CRYPTO_ERROR = 4109
	ER_CRYPTO_TRUNCATED = 4110
	ER_CRYPTO_KEY_UNAVAILABLE = 4111
	ER_BAD_HOSTNAME = 4112
	ER_CRYPTO_KEY_UNUSABLE = 4113
	ER_EMPTY_KEY_BLOB = 4114
	ER_CORRUPT_KEYBLOB = 4115
	ER_INVALID_KEY_ENCODING = 4116
	ER_DEAD_THREAD = 4117
	ER_THREAD_RUNNING = 4118
	ER_THREAD_STOPPING = 4119
	ER_BAD_STRING_ENCODING = 4120
	ER_CRYPTO_INSUFFICIENT_SECURITY = 4121
	ER_CRYPTO_ILLEGAL_PARAMETERS = 4122
	ER_CRYPTO_HASH_UNINITIALIZED = 4123
	ER_THREAD_NO_WAIT = 4124
	ER_TIMER_EXITING = 4125
	ER_INVALID_GUID = 4126
	ER_THREADPOOL_EXHAUSTED = 4127
	ER_THREADPOOL_STOPPING = 4128
	ER_INVALID_STREAM = 4129
	ER_TIMER_FULL = 4130
	ER_IODISPATCH_STOPPING = 4131
	ER_SLAP_INVALID_PACKET_LEN = 4132
	ER_SLAP_HDR_CHECKSUM_ERROR = 4133
	ER_SLAP_INVALID_PACKET_TYPE = 4134
	ER_SLAP_LEN_MISMATCH = 4135
	ER_SLAP_PACKET_TYPE_MISMATCH = 4136
	ER_SLAP_CRC_ERROR = 4137
	ER_SLAP_ERROR = 4138
	ER_SLAP_OTHER_END_CLOSED = 4139
	ER_TIMER_NOT_ALLOWED = 4140
	ER_NOT_CONN = 4141
	ER_XML_CONVERTER_ERROR = 8192
	ER_XML_INVALID_RULES_COUNT = 8193
	ER_XML_INTERFACE_MEMBERS_MISSING = 8194
	ER_XML_INVALID_MEMBER_TYPE = 8195
	ER_XML_INVALID_MEMBER_ACTION = 8196
	ER_XML_MEMBER_DENY_ACTION_WITH_OTHER = 8197
	ER_XML_INVALID_ANNOTATIONS_COUNT = 8198
	ER_XML_INVALID_ELEMENT_NAME = 8199
	ER_XML_INVALID_ATTRIBUTE_VALUE = 8200
	ER_XML_INVALID_SECURITY_LEVEL_ANNOTATION_VALUE = 8201
	ER_XML_INVALID_ELEMENT_CHILDREN_COUNT = 8202
	ER_XML_INVALID_POLICY_VERSION = 8203
	ER_XML_INVALID_POLICY_SERIAL_NUMBER = 8204
	ER_XML_INVALID_ACL_PEER_TYPE = 8205
	ER_XML_INVALID_ACL_PEER_CHILDREN_COUNT = 8206
	ER_XML_ACL_ALL_TYPE_PEER_WITH_OTHERS = 8207
	ER_XML_INVALID_ACL_PEER_PUBLIC_KEY = 8208
	ER_XML_ACL_PEER_NOT_UNIQUE = 8209
	ER_XML_ACL_PEER_PUBLIC_KEY_SET = 8210
	ER_XML_ACLS_MISSING = 8211
	ER_XML_ACL_PEERS_MISSING = 8212
	ER_XML_INVALID_OBJECT_PATH = 8213
	ER_XML_INVALID_INTERFACE_NAME = 8214
	ER_XML_INVALID_MEMBER_NAME = 8215
	ER_XML_INVALID_MANIFEST_VERSION = 8216
	ER_XML_INVALID_OID = 8217
	ER_XML_INVALID_BASE64 = 8218
	ER_XML_INTERFACE_NAME_NOT_UNIQUE = 8219
	ER_XML_MEMBER_NAME_NOT_UNIQUE = 8220
	ER_XML_OBJECT_PATH_NOT_UNIQUE = 8221
	ER_XML_ANNOTATION_NOT_UNIQUE = 8222
	ER_NONE = 65535
	ER_BUS_ERRORS = 36864
	ER_BUS_READ_ERROR = 36865
	ER_BUS_WRITE_ERROR = 36866
	ER_BUS_BAD_VALUE_TYPE = 36867
	ER_BUS_BAD_HEADER_FIELD = 36868
	ER_BUS_BAD_SIGNATURE = 36869
	ER_BUS_BAD_OBJ_PATH = 36870
	ER_BUS_BAD_MEMBER_NAME = 36871
	ER_BUS_BAD_INTERFACE_NAME = 36872
	ER_BUS_BAD_ERROR_NAME = 36873
	ER_BUS_BAD_BUS_NAME = 36874
	ER_BUS_NAME_TOO_LONG = 36875
	ER_BUS_BAD_LENGTH = 36876
	ER_BUS_BAD_VALUE = 36877
	ER_BUS_BAD_HDR_FLAGS = 36878
	ER_BUS_BAD_BODY_LEN = 36879
	ER_BUS_BAD_HEADER_LEN = 36880
	ER_BUS_UNKNOWN_SERIAL = 36881
	ER_BUS_UNKNOWN_PATH = 36882
	ER_BUS_UNKNOWN_INTERFACE = 36883
	ER_BUS_ESTABLISH_FAILED = 36884
	ER_BUS_UNEXPECTED_SIGNATURE = 36885
	ER_BUS_INTERFACE_MISSING = 36886
	ER_BUS_PATH_MISSING = 36887
	ER_BUS_MEMBER_MISSING = 36888
	ER_BUS_REPLY_SERIAL_MISSING = 36889
	ER_BUS_ERROR_NAME_MISSING = 36890
	ER_BUS_INTERFACE_NO_SUCH_MEMBER = 36891
	ER_BUS_NO_SUCH_OBJECT = 36892
	ER_BUS_OBJECT_NO_SUCH_MEMBER = 36893
	ER_BUS_OBJECT_NO_SUCH_INTERFACE = 36894
	ER_BUS_NO_SUCH_INTERFACE = 36895
	ER_BUS_MEMBER_NO_SUCH_SIGNATURE = 36896
	ER_BUS_NOT_NUL_TERMINATED = 36897
	ER_BUS_NO_SUCH_PROPERTY = 36898
	ER_BUS_SET_WRONG_SIGNATURE = 36899
	ER_BUS_PROPERTY_VALUE_NOT_SET = 36900
	ER_BUS_PROPERTY_ACCESS_DENIED = 36901
	ER_BUS_NO_TRANSPORTS = 36902
	ER_BUS_BAD_TRANSPORT_ARGS = 36903
	ER_BUS_NO_ROUTE = 36904
	ER_BUS_NO_ENDPOINT = 36905
	ER_BUS_BAD_SEND_PARAMETER = 36906
	ER_BUS_UNMATCHED_REPLY_SERIAL = 36907
	ER_BUS_BAD_SENDER_ID = 36908
	ER_BUS_TRANSPORT_NOT_STARTED = 36909
	ER_BUS_EMPTY_MESSAGE = 36910
	ER_BUS_NOT_OWNER = 36911
	ER_BUS_SET_PROPERTY_REJECTED = 36912
	ER_BUS_CONNECT_FAILED = 36913
	ER_BUS_REPLY_IS_ERROR_MESSAGE = 36914
	ER_BUS_NOT_AUTHENTICATING = 36915
	ER_BUS_NO_LISTENER = 36916
	ER_BUS_NOT_ALLOWED = 36918
	ER_BUS_WRITE_QUEUE_FULL = 36919
	ER_BUS_ENDPOINT_CLOSING = 36920
	ER_BUS_INTERFACE_MISMATCH = 36921
	ER_BUS_MEMBER_ALREADY_EXISTS = 36922
	ER_BUS_PROPERTY_ALREADY_EXISTS = 36923
	ER_BUS_IFACE_ALREADY_EXISTS = 36924
	ER_BUS_ERROR_RESPONSE = 36925
	ER_BUS_BAD_XML = 36926
	ER_BUS_BAD_CHILD_PATH = 36927
	ER_BUS_OBJ_ALREADY_EXISTS = 36928
	ER_BUS_OBJ_NOT_FOUND = 36929
	ER_BUS_CANNOT_EXPAND_MESSAGE = 36930
	ER_BUS_NOT_COMPRESSED = 36931
	ER_BUS_ALREADY_CONNECTED = 36932
	ER_BUS_NOT_CONNECTED = 36933
	ER_BUS_ALREADY_LISTENING = 36934
	ER_BUS_KEY_UNAVAILABLE = 36935
	ER_BUS_TRUNCATED = 36936
	ER_BUS_KEY_STORE_NOT_LOADED = 36937
	ER_BUS_NO_AUTHENTICATION_MECHANISM = 36938
	ER_BUS_BUS_ALREADY_STARTED = 36939
	ER_BUS_BUS_NOT_STARTED = 36940
	ER_BUS_KEYBLOB_OP_INVALID = 36941
	ER_BUS_INVALID_HEADER_CHECKSUM = 36942
	ER_BUS_MESSAGE_NOT_ENCRYPTED = 36943
	ER_BUS_INVALID_HEADER_SERIAL = 36944
	ER_BUS_TIME_TO_LIVE_EXPIRED = 36945
	ER_BUS_HDR_EXPANSION_INVALID = 36946
	ER_BUS_MISSING_COMPRESSION_TOKEN = 36947
	ER_BUS_NO_PEER_GUID = 36948
	ER_BUS_MESSAGE_DECRYPTION_FAILED = 36949
	ER_BUS_SECURITY_FATAL = 36950
	ER_BUS_KEY_EXPIRED = 36951
	ER_BUS_CORRUPT_KEYSTORE = 36952
	ER_BUS_NO_CALL_FOR_REPLY = 36953
	ER_BUS_NOT_A_COMPLETE_TYPE = 36954
	ER_BUS_POLICY_VIOLATION = 36955
	ER_BUS_NO_SUCH_SERVICE = 36956
	ER_BUS_TRANSPORT_NOT_AVAILABLE = 36957
	ER_BUS_INVALID_AUTH_MECHANISM = 36958
	ER_BUS_KEYSTORE_VERSION_MISMATCH = 36959
	ER_BUS_BLOCKING_CALL_NOT_ALLOWED = 36960
	ER_BUS_SIGNATURE_MISMATCH = 36961
	ER_BUS_STOPPING = 36962
	ER_BUS_METHOD_CALL_ABORTED = 36963
	ER_BUS_CANNOT_ADD_INTERFACE = 36964
	ER_BUS_CANNOT_ADD_HANDLER = 36965
	ER_BUS_KEYSTORE_NOT_LOADED = 36966
	ER_BUS_NO_SUCH_HANDLE = 36971
	ER_BUS_HANDLES_NOT_ENABLED = 36972
	ER_BUS_HANDLES_MISMATCH = 36973
	ER_BUS_NO_SESSION = 36975
	ER_BUS_ELEMENT_NOT_FOUND = 36976
	ER_BUS_NOT_A_DICTIONARY = 36977
	ER_BUS_WAIT_FAILED = 36978
	ER_BUS_BAD_SESSION_OPTS = 36980
	ER_BUS_CONNECTION_REJECTED = 36981
	ER_DBUS_REQUEST_NAME_REPLY_PRIMARY_OWNER = 36982
	ER_DBUS_REQUEST_NAME_REPLY_IN_QUEUE = 36983
	ER_DBUS_REQUEST_NAME_REPLY_EXISTS = 36984
	ER_DBUS_REQUEST_NAME_REPLY_ALREADY_OWNER = 36985
	ER_DBUS_RELEASE_NAME_REPLY_RELEASED = 36986
	ER_DBUS_RELEASE_NAME_REPLY_NON_EXISTENT = 36987
	ER_DBUS_RELEASE_NAME_REPLY_NOT_OWNER = 36988
	ER_DBUS_START_REPLY_ALREADY_RUNNING = 36990
	ER_ALLJOYN_BINDSESSIONPORT_REPLY_ALREADY_EXISTS = 36992
	ER_ALLJOYN_BINDSESSIONPORT_REPLY_FAILED = 36993
	ER_ALLJOYN_JOINSESSION_REPLY_NO_SESSION = 36995
	ER_ALLJOYN_JOINSESSION_REPLY_UNREACHABLE = 36996
	ER_ALLJOYN_JOINSESSION_REPLY_CONNECT_FAILED = 36997
	ER_ALLJOYN_JOINSESSION_REPLY_REJECTED = 36998
	ER_ALLJOYN_JOINSESSION_REPLY_BAD_SESSION_OPTS = 36999
	ER_ALLJOYN_JOINSESSION_REPLY_FAILED = 37000
	ER_ALLJOYN_LEAVESESSION_REPLY_NO_SESSION = 37002
	ER_ALLJOYN_LEAVESESSION_REPLY_FAILED = 37003
	ER_ALLJOYN_ADVERTISENAME_REPLY_TRANSPORT_NOT_AVAILABLE = 37004
	ER_ALLJOYN_ADVERTISENAME_REPLY_ALREADY_ADVERTISING = 37005
	ER_ALLJOYN_ADVERTISENAME_REPLY_FAILED = 37006
	ER_ALLJOYN_CANCELADVERTISENAME_REPLY_FAILED = 37008
	ER_ALLJOYN_FINDADVERTISEDNAME_REPLY_TRANSPORT_NOT_AVAILABLE = 37009
	ER_ALLJOYN_FINDADVERTISEDNAME_REPLY_ALREADY_DISCOVERING = 37010
	ER_ALLJOYN_FINDADVERTISEDNAME_REPLY_FAILED = 37011
	ER_ALLJOYN_CANCELFINDADVERTISEDNAME_REPLY_FAILED = 37013
	ER_BUS_UNEXPECTED_DISPOSITION = 37014
	ER_BUS_INTERFACE_ACTIVATED = 37015
	ER_ALLJOYN_UNBINDSESSIONPORT_REPLY_BAD_PORT = 37016
	ER_ALLJOYN_UNBINDSESSIONPORT_REPLY_FAILED = 37017
	ER_ALLJOYN_BINDSESSIONPORT_REPLY_INVALID_OPTS = 37018
	ER_ALLJOYN_JOINSESSION_REPLY_ALREADY_JOINED = 37019
	ER_BUS_SELF_CONNECT = 37020
	ER_BUS_SECURITY_NOT_ENABLED = 37021
	ER_BUS_LISTENER_ALREADY_SET = 37022
	ER_BUS_PEER_AUTH_VERSION_MISMATCH = 37023
	ER_ALLJOYN_SETLINKTIMEOUT_REPLY_NOT_SUPPORTED = 37024
	ER_ALLJOYN_SETLINKTIMEOUT_REPLY_NO_DEST_SUPPORT = 37025
	ER_ALLJOYN_SETLINKTIMEOUT_REPLY_FAILED = 37026
	ER_ALLJOYN_ACCESS_PERMISSION_WARNING = 37027
	ER_ALLJOYN_ACCESS_PERMISSION_ERROR = 37028
	ER_BUS_DESTINATION_NOT_AUTHENTICATED = 37029
	ER_BUS_ENDPOINT_REDIRECTED = 37030
	ER_BUS_AUTHENTICATION_PENDING = 37031
	ER_BUS_NOT_AUTHORIZED = 37032
	ER_PACKET_BUS_NO_SUCH_CHANNEL = 37033
	ER_PACKET_BAD_FORMAT = 37034
	ER_PACKET_CONNECT_TIMEOUT = 37035
	ER_PACKET_CHANNEL_FAIL = 37036
	ER_PACKET_TOO_LARGE = 37037
	ER_PACKET_BAD_PARAMETER = 37038
	ER_PACKET_BAD_CRC = 37039
	ER_RENDEZVOUS_SERVER_DEACTIVATED_USER = 37067
	ER_RENDEZVOUS_SERVER_UNKNOWN_USER = 37068
	ER_UNABLE_TO_CONNECT_TO_RENDEZVOUS_SERVER = 37069
	ER_NOT_CONNECTED_TO_RENDEZVOUS_SERVER = 37070
	ER_UNABLE_TO_SEND_MESSAGE_TO_RENDEZVOUS_SERVER = 37071
	ER_INVALID_RENDEZVOUS_SERVER_INTERFACE_MESSAGE = 37072
	ER_INVALID_PERSISTENT_CONNECTION_MESSAGE_RESPONSE = 37073
	ER_INVALID_ON_DEMAND_CONNECTION_MESSAGE_RESPONSE = 37074
	ER_INVALID_HTTP_METHOD_USED_FOR_RENDEZVOUS_SERVER_INTERFACE_MESSAGE = 37075
	ER_RENDEZVOUS_SERVER_ERR500_INTERNAL_ERROR = 37076
	ER_RENDEZVOUS_SERVER_ERR503_STATUS_UNAVAILABLE = 37077
	ER_RENDEZVOUS_SERVER_ERR401_UNAUTHORIZED_REQUEST = 37078
	ER_RENDEZVOUS_SERVER_UNRECOVERABLE_ERROR = 37079
	ER_RENDEZVOUS_SERVER_ROOT_CERTIFICATE_UNINITIALIZED = 37080
	ER_BUS_NO_SUCH_ANNOTATION = 37081
	ER_BUS_ANNOTATION_ALREADY_EXISTS = 37082
	ER_SOCK_CLOSING = 37083
	ER_NO_SUCH_DEVICE = 37084
	ER_P2P = 37085
	ER_P2P_TIMEOUT = 37086
	ER_P2P_NOT_CONNECTED = 37087
	ER_BAD_TRANSPORT_MASK = 37088
	ER_PROXIMITY_CONNECTION_ESTABLISH_FAIL = 37089
	ER_PROXIMITY_NO_PEERS_FOUND = 37090
	ER_BUS_OBJECT_NOT_REGISTERED = 37091
	ER_P2P_DISABLED = 37092
	ER_P2P_BUSY = 37093
	ER_BUS_INCOMPATIBLE_DAEMON = 37094
	ER_P2P_NO_GO = 37095
	ER_P2P_NO_STA = 37096
	ER_P2P_FORBIDDEN = 37097
	ER_ALLJOYN_ONAPPSUSPEND_REPLY_FAILED = 37098
	ER_ALLJOYN_ONAPPSUSPEND_REPLY_UNSUPPORTED = 37099
	ER_ALLJOYN_ONAPPRESUME_REPLY_FAILED = 37100
	ER_ALLJOYN_ONAPPRESUME_REPLY_UNSUPPORTED = 37101
	ER_BUS_NO_SUCH_MESSAGE = 37102
	ER_ALLJOYN_REMOVESESSIONMEMBER_REPLY_NO_SESSION = 37103
	ER_ALLJOYN_REMOVESESSIONMEMBER_NOT_BINDER = 37104
	ER_ALLJOYN_REMOVESESSIONMEMBER_NOT_MULTIPOINT = 37105
	ER_ALLJOYN_REMOVESESSIONMEMBER_NOT_FOUND = 37106
	ER_ALLJOYN_REMOVESESSIONMEMBER_INCOMPATIBLE_REMOTE_DAEMON = 37107
	ER_ALLJOYN_REMOVESESSIONMEMBER_REPLY_FAILED = 37108
	ER_BUS_REMOVED_BY_BINDER = 37109
	ER_BUS_MATCH_RULE_NOT_FOUND = 37110
	ER_ALLJOYN_PING_FAILED = 37111
	ER_ALLJOYN_PING_REPLY_UNREACHABLE = 37112
	ER_UDP_MSG_TOO_LONG = 37113
	ER_UDP_DEMUX_NO_ENDPOINT = 37114
	ER_UDP_NO_NETWORK = 37115
	ER_UDP_UNEXPECTED_LENGTH = 37116
	ER_UDP_UNEXPECTED_FLOW = 37117
	ER_UDP_DISCONNECT = 37118
	ER_UDP_NOT_IMPLEMENTED = 37119
	ER_UDP_NO_LISTENER = 37120
	ER_UDP_STOPPING = 37121
	ER_ARDP_BACKPRESSURE = 37122
	ER_UDP_BACKPRESSURE = 37123
	ER_ARDP_INVALID_STATE = 37124
	ER_ARDP_TTL_EXPIRED = 37125
	ER_ARDP_PERSIST_TIMEOUT = 37126
	ER_ARDP_PROBE_TIMEOUT = 37127
	ER_ARDP_REMOTE_CONNECTION_RESET = 37128
	ER_UDP_BUSHELLO = 37129
	ER_UDP_MESSAGE = 37130
	ER_UDP_INVALID = 37131
	ER_UDP_UNSUPPORTED = 37132
	ER_UDP_ENDPOINT_STALLED = 37133
	ER_ARDP_INVALID_RESPONSE = 37134
	ER_ARDP_INVALID_CONNECTION = 37135
	ER_UDP_LOCAL_DISCONNECT = 37136
	ER_UDP_EARLY_EXIT = 37137
	ER_UDP_LOCAL_DISCONNECT_FAIL = 37138
	ER_ARDP_DISCONNECTING = 37139
	ER_ALLJOYN_PING_REPLY_INCOMPATIBLE_REMOTE_ROUTING_NODE = 37140
	ER_ALLJOYN_PING_REPLY_TIMEOUT = 37141
	ER_ALLJOYN_PING_REPLY_UNKNOWN_NAME = 37142
	ER_ALLJOYN_PING_REPLY_FAILED = 37143
	ER_TCP_MAX_UNTRUSTED = 37144
	ER_ALLJOYN_PING_REPLY_IN_PROGRESS = 37145
	ER_LANGUAGE_NOT_SUPPORTED = 37146
	ER_ABOUT_FIELD_ALREADY_SPECIFIED = 37147
	ER_UDP_NOT_DISCONNECTED = 37148
	ER_UDP_ENDPOINT_NOT_STARTED = 37149
	ER_UDP_ENDPOINT_REMOVED = 37150
	ER_ARDP_VERSION_NOT_SUPPORTED = 37151
	ER_CONNECTION_LIMIT_EXCEEDED = 37152
	ER_ARDP_WRITE_BLOCKED = 37153
	ER_PERMISSION_DENIED = 37154
	ER_ABOUT_DEFAULT_LANGUAGE_NOT_SPECIFIED = 37155
	ER_ABOUT_SESSIONPORT_NOT_BOUND = 37156
	ER_ABOUT_ABOUTDATA_MISSING_REQUIRED_FIELD = 37157
	ER_ABOUT_INVALID_ABOUTDATA_LISTENER = 37158
	ER_BUS_PING_GROUP_NOT_FOUND = 37159
	ER_BUS_REMOVED_BY_BINDER_SELF = 37160
	ER_INVALID_CONFIG = 37161
	ER_ABOUT_INVALID_ABOUTDATA_FIELD_VALUE = 37162
	ER_ABOUT_INVALID_ABOUTDATA_FIELD_APPID_SIZE = 37163
	ER_BUS_TRANSPORT_ACCESS_DENIED = 37164
	ER_INVALID_CERTIFICATE = 37165
	ER_CERTIFICATE_NOT_FOUND = 37166
	ER_DUPLICATE_CERTIFICATE = 37167
	ER_UNKNOWN_CERTIFICATE = 37168
	ER_MISSING_DIGEST_IN_CERTIFICATE = 37169
	ER_DIGEST_MISMATCH = 37170
	ER_DUPLICATE_KEY = 37171
	ER_NO_COMMON_TRUST = 37172
	ER_MANIFEST_NOT_FOUND = 37173
	ER_INVALID_CERT_CHAIN = 37174
	ER_NO_TRUST_ANCHOR = 37175
	ER_INVALID_APPLICATION_STATE = 37176
	ER_FEATURE_NOT_AVAILABLE = 37177
	ER_KEY_STORE_ALREADY_INITIALIZED = 37178
	ER_KEY_STORE_ID_NOT_YET_SET = 37179
	ER_POLICY_NOT_NEWER = 37180
	ER_MANIFEST_REJECTED = 37181
	ER_INVALID_CERTIFICATE_USAGE = 37182
	ER_INVALID_SIGNAL_EMISSION_TYPE = 37183
	ER_APPLICATION_STATE_LISTENER_ALREADY_EXISTS = 37184
	ER_APPLICATION_STATE_LISTENER_NO_SUCH_LISTENER = 37185
	ER_MANAGEMENT_ALREADY_STARTED = 37186
	ER_MANAGEMENT_NOT_STARTED = 37187
	ER_BUS_DESCRIPTION_ALREADY_EXISTS = 37188
)

type alljoyn_typeid int32

const (
	ALLJOYN_INVALID = 0
	ALLJOYN_ARRAY = 97
	ALLJOYN_BOOLEAN = 98
	ALLJOYN_DOUBLE = 100
	ALLJOYN_DICT_ENTRY = 101
	ALLJOYN_SIGNATURE = 103
	ALLJOYN_HANDLE = 104
	ALLJOYN_INT32 = 105
	ALLJOYN_INT16 = 110
	ALLJOYN_OBJECT_PATH = 111
	ALLJOYN_UINT16 = 113
	ALLJOYN_STRUCT = 114
	ALLJOYN_STRING = 115
	ALLJOYN_UINT64 = 116
	ALLJOYN_UINT32 = 117
	ALLJOYN_VARIANT = 118
	ALLJOYN_INT64 = 120
	ALLJOYN_BYTE = 121
	ALLJOYN_STRUCT_OPEN = 40
	ALLJOYN_STRUCT_CLOSE = 41
	ALLJOYN_DICT_ENTRY_OPEN = 123
	ALLJOYN_DICT_ENTRY_CLOSE = 125
	ALLJOYN_BOOLEAN_ARRAY = 25185
	ALLJOYN_DOUBLE_ARRAY = 25697
	ALLJOYN_INT32_ARRAY = 26977
	ALLJOYN_INT16_ARRAY = 28257
	ALLJOYN_UINT16_ARRAY = 29025
	ALLJOYN_UINT64_ARRAY = 29793
	ALLJOYN_UINT32_ARRAY = 30049
	ALLJOYN_INT64_ARRAY = 30817
	ALLJOYN_BYTE_ARRAY = 31073
	ALLJOYN_WILDCARD = 42
)

type alljoyn_applicationstate int32

const (
	NOT_CLAIMABLE = 0
	CLAIMABLE = 1
	CLAIMED = 2
	NEED_UPDATE = 3
)

type alljoyn_claimcapability_masks int32

const (
	CAPABLE_ECDHE_NULL = 1
	CAPABLE_ECDHE_ECDSA = 4
	CAPABLE_ECDHE_SPEKE = 8
)

type alljoyn_claimcapabilityadditionalinfo_masks int32

const (
	PASSWORD_GENERATED_BY_SECURITY_MANAGER = 1
	PASSWORD_GENERATED_BY_APPLICATION = 2
)

type alljoyn_messagetype int32

const (
	ALLJOYN_MESSAGE_INVALID = 0
	ALLJOYN_MESSAGE_METHOD_CALL = 1
	ALLJOYN_MESSAGE_METHOD_RET = 2
	ALLJOYN_MESSAGE_ERROR = 3
	ALLJOYN_MESSAGE_SIGNAL = 4
)

type alljoyn_interfacedescription_securitypolicy int32

const (
	AJ_IFC_SECURITY_INHERIT = 0
	AJ_IFC_SECURITY_REQUIRED = 1
	AJ_IFC_SECURITY_OFF = 2
)

type alljoyn_sessionlostreason int32

const (
	ALLJOYN_SESSIONLOST_INVALID = 0
	ALLJOYN_SESSIONLOST_REMOTE_END_LEFT_SESSION = 1
	ALLJOYN_SESSIONLOST_REMOTE_END_CLOSED_ABRUPTLY = 2
	ALLJOYN_SESSIONLOST_REMOVED_BY_BINDER = 3
	ALLJOYN_SESSIONLOST_LINK_TIMEOUT = 4
	ALLJOYN_SESSIONLOST_REASON_OTHER = 5
)


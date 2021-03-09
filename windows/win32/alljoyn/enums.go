// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package alljoyn implements the Windows.Win32.AllJoyn namespace.
package alljoyn

type ALLJOYN_ABOUT_ANNOUNCEFLAG int32

const (
	UNANNOUNCED ALLJOYN_ABOUT_ANNOUNCEFLAG = 0
	ANNOUNCED   ALLJOYN_ABOUT_ANNOUNCEFLAG = 1
)

type QStatus int32

const (
	ER_OK                                                               QStatus = 0
	ER_FAIL                                                             QStatus = 1
	ER_UTF_CONVERSION_FAILED                                            QStatus = 2
	ER_BUFFER_TOO_SMALL                                                 QStatus = 3
	ER_OS_ERROR                                                         QStatus = 4
	ER_OUT_OF_MEMORY                                                    QStatus = 5
	ER_SOCKET_BIND_ERROR                                                QStatus = 6
	ER_INIT_FAILED                                                      QStatus = 7
	ER_WOULDBLOCK                                                       QStatus = 8
	ER_NOT_IMPLEMENTED                                                  QStatus = 9
	ER_TIMEOUT                                                          QStatus = 10
	ER_SOCK_OTHER_END_CLOSED                                            QStatus = 11
	ER_BAD_ARG_1                                                        QStatus = 12
	ER_BAD_ARG_2                                                        QStatus = 13
	ER_BAD_ARG_3                                                        QStatus = 14
	ER_BAD_ARG_4                                                        QStatus = 15
	ER_BAD_ARG_5                                                        QStatus = 16
	ER_BAD_ARG_6                                                        QStatus = 17
	ER_BAD_ARG_7                                                        QStatus = 18
	ER_BAD_ARG_8                                                        QStatus = 19
	ER_INVALID_ADDRESS                                                  QStatus = 20
	ER_INVALID_DATA                                                     QStatus = 21
	ER_READ_ERROR                                                       QStatus = 22
	ER_WRITE_ERROR                                                      QStatus = 23
	ER_OPEN_FAILED                                                      QStatus = 24
	ER_PARSE_ERROR                                                      QStatus = 25
	ER_END_OF_DATA                                                      QStatus = 26
	ER_CONN_REFUSED                                                     QStatus = 27
	ER_BAD_ARG_COUNT                                                    QStatus = 28
	ER_WARNING                                                          QStatus = 29
	ER_EOF                                                              QStatus = 30
	ER_DEADLOCK                                                         QStatus = 31
	ER_COMMON_ERRORS                                                    QStatus = 4096
	ER_STOPPING_THREAD                                                  QStatus = 4097
	ER_ALERTED_THREAD                                                   QStatus = 4098
	ER_XML_MALFORMED                                                    QStatus = 4099
	ER_AUTH_FAIL                                                        QStatus = 4100
	ER_AUTH_USER_REJECT                                                 QStatus = 4101
	ER_NO_SUCH_ALARM                                                    QStatus = 4102
	ER_TIMER_FALLBEHIND                                                 QStatus = 4103
	ER_SSL_ERRORS                                                       QStatus = 4104
	ER_SSL_INIT                                                         QStatus = 4105
	ER_SSL_CONNECT                                                      QStatus = 4106
	ER_SSL_VERIFY                                                       QStatus = 4107
	ER_EXTERNAL_THREAD                                                  QStatus = 4108
	ER_CRYPTO_ERROR                                                     QStatus = 4109
	ER_CRYPTO_TRUNCATED                                                 QStatus = 4110
	ER_CRYPTO_KEY_UNAVAILABLE                                           QStatus = 4111
	ER_BAD_HOSTNAME                                                     QStatus = 4112
	ER_CRYPTO_KEY_UNUSABLE                                              QStatus = 4113
	ER_EMPTY_KEY_BLOB                                                   QStatus = 4114
	ER_CORRUPT_KEYBLOB                                                  QStatus = 4115
	ER_INVALID_KEY_ENCODING                                             QStatus = 4116
	ER_DEAD_THREAD                                                      QStatus = 4117
	ER_THREAD_RUNNING                                                   QStatus = 4118
	ER_THREAD_STOPPING                                                  QStatus = 4119
	ER_BAD_STRING_ENCODING                                              QStatus = 4120
	ER_CRYPTO_INSUFFICIENT_SECURITY                                     QStatus = 4121
	ER_CRYPTO_ILLEGAL_PARAMETERS                                        QStatus = 4122
	ER_CRYPTO_HASH_UNINITIALIZED                                        QStatus = 4123
	ER_THREAD_NO_WAIT                                                   QStatus = 4124
	ER_TIMER_EXITING                                                    QStatus = 4125
	ER_INVALID_GUID                                                     QStatus = 4126
	ER_THREADPOOL_EXHAUSTED                                             QStatus = 4127
	ER_THREADPOOL_STOPPING                                              QStatus = 4128
	ER_INVALID_STREAM                                                   QStatus = 4129
	ER_TIMER_FULL                                                       QStatus = 4130
	ER_IODISPATCH_STOPPING                                              QStatus = 4131
	ER_SLAP_INVALID_PACKET_LEN                                          QStatus = 4132
	ER_SLAP_HDR_CHECKSUM_ERROR                                          QStatus = 4133
	ER_SLAP_INVALID_PACKET_TYPE                                         QStatus = 4134
	ER_SLAP_LEN_MISMATCH                                                QStatus = 4135
	ER_SLAP_PACKET_TYPE_MISMATCH                                        QStatus = 4136
	ER_SLAP_CRC_ERROR                                                   QStatus = 4137
	ER_SLAP_ERROR                                                       QStatus = 4138
	ER_SLAP_OTHER_END_CLOSED                                            QStatus = 4139
	ER_TIMER_NOT_ALLOWED                                                QStatus = 4140
	ER_NOT_CONN                                                         QStatus = 4141
	ER_XML_CONVERTER_ERROR                                              QStatus = 8192
	ER_XML_INVALID_RULES_COUNT                                          QStatus = 8193
	ER_XML_INTERFACE_MEMBERS_MISSING                                    QStatus = 8194
	ER_XML_INVALID_MEMBER_TYPE                                          QStatus = 8195
	ER_XML_INVALID_MEMBER_ACTION                                        QStatus = 8196
	ER_XML_MEMBER_DENY_ACTION_WITH_OTHER                                QStatus = 8197
	ER_XML_INVALID_ANNOTATIONS_COUNT                                    QStatus = 8198
	ER_XML_INVALID_ELEMENT_NAME                                         QStatus = 8199
	ER_XML_INVALID_ATTRIBUTE_VALUE                                      QStatus = 8200
	ER_XML_INVALID_SECURITY_LEVEL_ANNOTATION_VALUE                      QStatus = 8201
	ER_XML_INVALID_ELEMENT_CHILDREN_COUNT                               QStatus = 8202
	ER_XML_INVALID_POLICY_VERSION                                       QStatus = 8203
	ER_XML_INVALID_POLICY_SERIAL_NUMBER                                 QStatus = 8204
	ER_XML_INVALID_ACL_PEER_TYPE                                        QStatus = 8205
	ER_XML_INVALID_ACL_PEER_CHILDREN_COUNT                              QStatus = 8206
	ER_XML_ACL_ALL_TYPE_PEER_WITH_OTHERS                                QStatus = 8207
	ER_XML_INVALID_ACL_PEER_PUBLIC_KEY                                  QStatus = 8208
	ER_XML_ACL_PEER_NOT_UNIQUE                                          QStatus = 8209
	ER_XML_ACL_PEER_PUBLIC_KEY_SET                                      QStatus = 8210
	ER_XML_ACLS_MISSING                                                 QStatus = 8211
	ER_XML_ACL_PEERS_MISSING                                            QStatus = 8212
	ER_XML_INVALID_OBJECT_PATH                                          QStatus = 8213
	ER_XML_INVALID_INTERFACE_NAME                                       QStatus = 8214
	ER_XML_INVALID_MEMBER_NAME                                          QStatus = 8215
	ER_XML_INVALID_MANIFEST_VERSION                                     QStatus = 8216
	ER_XML_INVALID_OID                                                  QStatus = 8217
	ER_XML_INVALID_BASE64                                               QStatus = 8218
	ER_XML_INTERFACE_NAME_NOT_UNIQUE                                    QStatus = 8219
	ER_XML_MEMBER_NAME_NOT_UNIQUE                                       QStatus = 8220
	ER_XML_OBJECT_PATH_NOT_UNIQUE                                       QStatus = 8221
	ER_XML_ANNOTATION_NOT_UNIQUE                                        QStatus = 8222
	ER_NONE                                                             QStatus = 65535
	ER_BUS_ERRORS                                                       QStatus = 36864
	ER_BUS_READ_ERROR                                                   QStatus = 36865
	ER_BUS_WRITE_ERROR                                                  QStatus = 36866
	ER_BUS_BAD_VALUE_TYPE                                               QStatus = 36867
	ER_BUS_BAD_HEADER_FIELD                                             QStatus = 36868
	ER_BUS_BAD_SIGNATURE                                                QStatus = 36869
	ER_BUS_BAD_OBJ_PATH                                                 QStatus = 36870
	ER_BUS_BAD_MEMBER_NAME                                              QStatus = 36871
	ER_BUS_BAD_INTERFACE_NAME                                           QStatus = 36872
	ER_BUS_BAD_ERROR_NAME                                               QStatus = 36873
	ER_BUS_BAD_BUS_NAME                                                 QStatus = 36874
	ER_BUS_NAME_TOO_LONG                                                QStatus = 36875
	ER_BUS_BAD_LENGTH                                                   QStatus = 36876
	ER_BUS_BAD_VALUE                                                    QStatus = 36877
	ER_BUS_BAD_HDR_FLAGS                                                QStatus = 36878
	ER_BUS_BAD_BODY_LEN                                                 QStatus = 36879
	ER_BUS_BAD_HEADER_LEN                                               QStatus = 36880
	ER_BUS_UNKNOWN_SERIAL                                               QStatus = 36881
	ER_BUS_UNKNOWN_PATH                                                 QStatus = 36882
	ER_BUS_UNKNOWN_INTERFACE                                            QStatus = 36883
	ER_BUS_ESTABLISH_FAILED                                             QStatus = 36884
	ER_BUS_UNEXPECTED_SIGNATURE                                         QStatus = 36885
	ER_BUS_INTERFACE_MISSING                                            QStatus = 36886
	ER_BUS_PATH_MISSING                                                 QStatus = 36887
	ER_BUS_MEMBER_MISSING                                               QStatus = 36888
	ER_BUS_REPLY_SERIAL_MISSING                                         QStatus = 36889
	ER_BUS_ERROR_NAME_MISSING                                           QStatus = 36890
	ER_BUS_INTERFACE_NO_SUCH_MEMBER                                     QStatus = 36891
	ER_BUS_NO_SUCH_OBJECT                                               QStatus = 36892
	ER_BUS_OBJECT_NO_SUCH_MEMBER                                        QStatus = 36893
	ER_BUS_OBJECT_NO_SUCH_INTERFACE                                     QStatus = 36894
	ER_BUS_NO_SUCH_INTERFACE                                            QStatus = 36895
	ER_BUS_MEMBER_NO_SUCH_SIGNATURE                                     QStatus = 36896
	ER_BUS_NOT_NUL_TERMINATED                                           QStatus = 36897
	ER_BUS_NO_SUCH_PROPERTY                                             QStatus = 36898
	ER_BUS_SET_WRONG_SIGNATURE                                          QStatus = 36899
	ER_BUS_PROPERTY_VALUE_NOT_SET                                       QStatus = 36900
	ER_BUS_PROPERTY_ACCESS_DENIED                                       QStatus = 36901
	ER_BUS_NO_TRANSPORTS                                                QStatus = 36902
	ER_BUS_BAD_TRANSPORT_ARGS                                           QStatus = 36903
	ER_BUS_NO_ROUTE                                                     QStatus = 36904
	ER_BUS_NO_ENDPOINT                                                  QStatus = 36905
	ER_BUS_BAD_SEND_PARAMETER                                           QStatus = 36906
	ER_BUS_UNMATCHED_REPLY_SERIAL                                       QStatus = 36907
	ER_BUS_BAD_SENDER_ID                                                QStatus = 36908
	ER_BUS_TRANSPORT_NOT_STARTED                                        QStatus = 36909
	ER_BUS_EMPTY_MESSAGE                                                QStatus = 36910
	ER_BUS_NOT_OWNER                                                    QStatus = 36911
	ER_BUS_SET_PROPERTY_REJECTED                                        QStatus = 36912
	ER_BUS_CONNECT_FAILED                                               QStatus = 36913
	ER_BUS_REPLY_IS_ERROR_MESSAGE                                       QStatus = 36914
	ER_BUS_NOT_AUTHENTICATING                                           QStatus = 36915
	ER_BUS_NO_LISTENER                                                  QStatus = 36916
	ER_BUS_NOT_ALLOWED                                                  QStatus = 36918
	ER_BUS_WRITE_QUEUE_FULL                                             QStatus = 36919
	ER_BUS_ENDPOINT_CLOSING                                             QStatus = 36920
	ER_BUS_INTERFACE_MISMATCH                                           QStatus = 36921
	ER_BUS_MEMBER_ALREADY_EXISTS                                        QStatus = 36922
	ER_BUS_PROPERTY_ALREADY_EXISTS                                      QStatus = 36923
	ER_BUS_IFACE_ALREADY_EXISTS                                         QStatus = 36924
	ER_BUS_ERROR_RESPONSE                                               QStatus = 36925
	ER_BUS_BAD_XML                                                      QStatus = 36926
	ER_BUS_BAD_CHILD_PATH                                               QStatus = 36927
	ER_BUS_OBJ_ALREADY_EXISTS                                           QStatus = 36928
	ER_BUS_OBJ_NOT_FOUND                                                QStatus = 36929
	ER_BUS_CANNOT_EXPAND_MESSAGE                                        QStatus = 36930
	ER_BUS_NOT_COMPRESSED                                               QStatus = 36931
	ER_BUS_ALREADY_CONNECTED                                            QStatus = 36932
	ER_BUS_NOT_CONNECTED                                                QStatus = 36933
	ER_BUS_ALREADY_LISTENING                                            QStatus = 36934
	ER_BUS_KEY_UNAVAILABLE                                              QStatus = 36935
	ER_BUS_TRUNCATED                                                    QStatus = 36936
	ER_BUS_KEY_STORE_NOT_LOADED                                         QStatus = 36937
	ER_BUS_NO_AUTHENTICATION_MECHANISM                                  QStatus = 36938
	ER_BUS_BUS_ALREADY_STARTED                                          QStatus = 36939
	ER_BUS_BUS_NOT_STARTED                                              QStatus = 36940
	ER_BUS_KEYBLOB_OP_INVALID                                           QStatus = 36941
	ER_BUS_INVALID_HEADER_CHECKSUM                                      QStatus = 36942
	ER_BUS_MESSAGE_NOT_ENCRYPTED                                        QStatus = 36943
	ER_BUS_INVALID_HEADER_SERIAL                                        QStatus = 36944
	ER_BUS_TIME_TO_LIVE_EXPIRED                                         QStatus = 36945
	ER_BUS_HDR_EXPANSION_INVALID                                        QStatus = 36946
	ER_BUS_MISSING_COMPRESSION_TOKEN                                    QStatus = 36947
	ER_BUS_NO_PEER_GUID                                                 QStatus = 36948
	ER_BUS_MESSAGE_DECRYPTION_FAILED                                    QStatus = 36949
	ER_BUS_SECURITY_FATAL                                               QStatus = 36950
	ER_BUS_KEY_EXPIRED                                                  QStatus = 36951
	ER_BUS_CORRUPT_KEYSTORE                                             QStatus = 36952
	ER_BUS_NO_CALL_FOR_REPLY                                            QStatus = 36953
	ER_BUS_NOT_A_COMPLETE_TYPE                                          QStatus = 36954
	ER_BUS_POLICY_VIOLATION                                             QStatus = 36955
	ER_BUS_NO_SUCH_SERVICE                                              QStatus = 36956
	ER_BUS_TRANSPORT_NOT_AVAILABLE                                      QStatus = 36957
	ER_BUS_INVALID_AUTH_MECHANISM                                       QStatus = 36958
	ER_BUS_KEYSTORE_VERSION_MISMATCH                                    QStatus = 36959
	ER_BUS_BLOCKING_CALL_NOT_ALLOWED                                    QStatus = 36960
	ER_BUS_SIGNATURE_MISMATCH                                           QStatus = 36961
	ER_BUS_STOPPING                                                     QStatus = 36962
	ER_BUS_METHOD_CALL_ABORTED                                          QStatus = 36963
	ER_BUS_CANNOT_ADD_INTERFACE                                         QStatus = 36964
	ER_BUS_CANNOT_ADD_HANDLER                                           QStatus = 36965
	ER_BUS_KEYSTORE_NOT_LOADED                                          QStatus = 36966
	ER_BUS_NO_SUCH_HANDLE                                               QStatus = 36971
	ER_BUS_HANDLES_NOT_ENABLED                                          QStatus = 36972
	ER_BUS_HANDLES_MISMATCH                                             QStatus = 36973
	ER_BUS_NO_SESSION                                                   QStatus = 36975
	ER_BUS_ELEMENT_NOT_FOUND                                            QStatus = 36976
	ER_BUS_NOT_A_DICTIONARY                                             QStatus = 36977
	ER_BUS_WAIT_FAILED                                                  QStatus = 36978
	ER_BUS_BAD_SESSION_OPTS                                             QStatus = 36980
	ER_BUS_CONNECTION_REJECTED                                          QStatus = 36981
	ER_DBUS_REQUEST_NAME_REPLY_PRIMARY_OWNER                            QStatus = 36982
	ER_DBUS_REQUEST_NAME_REPLY_IN_QUEUE                                 QStatus = 36983
	ER_DBUS_REQUEST_NAME_REPLY_EXISTS                                   QStatus = 36984
	ER_DBUS_REQUEST_NAME_REPLY_ALREADY_OWNER                            QStatus = 36985
	ER_DBUS_RELEASE_NAME_REPLY_RELEASED                                 QStatus = 36986
	ER_DBUS_RELEASE_NAME_REPLY_NON_EXISTENT                             QStatus = 36987
	ER_DBUS_RELEASE_NAME_REPLY_NOT_OWNER                                QStatus = 36988
	ER_DBUS_START_REPLY_ALREADY_RUNNING                                 QStatus = 36990
	ER_ALLJOYN_BINDSESSIONPORT_REPLY_ALREADY_EXISTS                     QStatus = 36992
	ER_ALLJOYN_BINDSESSIONPORT_REPLY_FAILED                             QStatus = 36993
	ER_ALLJOYN_JOINSESSION_REPLY_NO_SESSION                             QStatus = 36995
	ER_ALLJOYN_JOINSESSION_REPLY_UNREACHABLE                            QStatus = 36996
	ER_ALLJOYN_JOINSESSION_REPLY_CONNECT_FAILED                         QStatus = 36997
	ER_ALLJOYN_JOINSESSION_REPLY_REJECTED                               QStatus = 36998
	ER_ALLJOYN_JOINSESSION_REPLY_BAD_SESSION_OPTS                       QStatus = 36999
	ER_ALLJOYN_JOINSESSION_REPLY_FAILED                                 QStatus = 37000
	ER_ALLJOYN_LEAVESESSION_REPLY_NO_SESSION                            QStatus = 37002
	ER_ALLJOYN_LEAVESESSION_REPLY_FAILED                                QStatus = 37003
	ER_ALLJOYN_ADVERTISENAME_REPLY_TRANSPORT_NOT_AVAILABLE              QStatus = 37004
	ER_ALLJOYN_ADVERTISENAME_REPLY_ALREADY_ADVERTISING                  QStatus = 37005
	ER_ALLJOYN_ADVERTISENAME_REPLY_FAILED                               QStatus = 37006
	ER_ALLJOYN_CANCELADVERTISENAME_REPLY_FAILED                         QStatus = 37008
	ER_ALLJOYN_FINDADVERTISEDNAME_REPLY_TRANSPORT_NOT_AVAILABLE         QStatus = 37009
	ER_ALLJOYN_FINDADVERTISEDNAME_REPLY_ALREADY_DISCOVERING             QStatus = 37010
	ER_ALLJOYN_FINDADVERTISEDNAME_REPLY_FAILED                          QStatus = 37011
	ER_ALLJOYN_CANCELFINDADVERTISEDNAME_REPLY_FAILED                    QStatus = 37013
	ER_BUS_UNEXPECTED_DISPOSITION                                       QStatus = 37014
	ER_BUS_INTERFACE_ACTIVATED                                          QStatus = 37015
	ER_ALLJOYN_UNBINDSESSIONPORT_REPLY_BAD_PORT                         QStatus = 37016
	ER_ALLJOYN_UNBINDSESSIONPORT_REPLY_FAILED                           QStatus = 37017
	ER_ALLJOYN_BINDSESSIONPORT_REPLY_INVALID_OPTS                       QStatus = 37018
	ER_ALLJOYN_JOINSESSION_REPLY_ALREADY_JOINED                         QStatus = 37019
	ER_BUS_SELF_CONNECT                                                 QStatus = 37020
	ER_BUS_SECURITY_NOT_ENABLED                                         QStatus = 37021
	ER_BUS_LISTENER_ALREADY_SET                                         QStatus = 37022
	ER_BUS_PEER_AUTH_VERSION_MISMATCH                                   QStatus = 37023
	ER_ALLJOYN_SETLINKTIMEOUT_REPLY_NOT_SUPPORTED                       QStatus = 37024
	ER_ALLJOYN_SETLINKTIMEOUT_REPLY_NO_DEST_SUPPORT                     QStatus = 37025
	ER_ALLJOYN_SETLINKTIMEOUT_REPLY_FAILED                              QStatus = 37026
	ER_ALLJOYN_ACCESS_PERMISSION_WARNING                                QStatus = 37027
	ER_ALLJOYN_ACCESS_PERMISSION_ERROR                                  QStatus = 37028
	ER_BUS_DESTINATION_NOT_AUTHENTICATED                                QStatus = 37029
	ER_BUS_ENDPOINT_REDIRECTED                                          QStatus = 37030
	ER_BUS_AUTHENTICATION_PENDING                                       QStatus = 37031
	ER_BUS_NOT_AUTHORIZED                                               QStatus = 37032
	ER_PACKET_BUS_NO_SUCH_CHANNEL                                       QStatus = 37033
	ER_PACKET_BAD_FORMAT                                                QStatus = 37034
	ER_PACKET_CONNECT_TIMEOUT                                           QStatus = 37035
	ER_PACKET_CHANNEL_FAIL                                              QStatus = 37036
	ER_PACKET_TOO_LARGE                                                 QStatus = 37037
	ER_PACKET_BAD_PARAMETER                                             QStatus = 37038
	ER_PACKET_BAD_CRC                                                   QStatus = 37039
	ER_RENDEZVOUS_SERVER_DEACTIVATED_USER                               QStatus = 37067
	ER_RENDEZVOUS_SERVER_UNKNOWN_USER                                   QStatus = 37068
	ER_UNABLE_TO_CONNECT_TO_RENDEZVOUS_SERVER                           QStatus = 37069
	ER_NOT_CONNECTED_TO_RENDEZVOUS_SERVER                               QStatus = 37070
	ER_UNABLE_TO_SEND_MESSAGE_TO_RENDEZVOUS_SERVER                      QStatus = 37071
	ER_INVALID_RENDEZVOUS_SERVER_INTERFACE_MESSAGE                      QStatus = 37072
	ER_INVALID_PERSISTENT_CONNECTION_MESSAGE_RESPONSE                   QStatus = 37073
	ER_INVALID_ON_DEMAND_CONNECTION_MESSAGE_RESPONSE                    QStatus = 37074
	ER_INVALID_HTTP_METHOD_USED_FOR_RENDEZVOUS_SERVER_INTERFACE_MESSAGE QStatus = 37075
	ER_RENDEZVOUS_SERVER_ERR500_INTERNAL_ERROR                          QStatus = 37076
	ER_RENDEZVOUS_SERVER_ERR503_STATUS_UNAVAILABLE                      QStatus = 37077
	ER_RENDEZVOUS_SERVER_ERR401_UNAUTHORIZED_REQUEST                    QStatus = 37078
	ER_RENDEZVOUS_SERVER_UNRECOVERABLE_ERROR                            QStatus = 37079
	ER_RENDEZVOUS_SERVER_ROOT_CERTIFICATE_UNINITIALIZED                 QStatus = 37080
	ER_BUS_NO_SUCH_ANNOTATION                                           QStatus = 37081
	ER_BUS_ANNOTATION_ALREADY_EXISTS                                    QStatus = 37082
	ER_SOCK_CLOSING                                                     QStatus = 37083
	ER_NO_SUCH_DEVICE                                                   QStatus = 37084
	ER_P2P                                                              QStatus = 37085
	ER_P2P_TIMEOUT                                                      QStatus = 37086
	ER_P2P_NOT_CONNECTED                                                QStatus = 37087
	ER_BAD_TRANSPORT_MASK                                               QStatus = 37088
	ER_PROXIMITY_CONNECTION_ESTABLISH_FAIL                              QStatus = 37089
	ER_PROXIMITY_NO_PEERS_FOUND                                         QStatus = 37090
	ER_BUS_OBJECT_NOT_REGISTERED                                        QStatus = 37091
	ER_P2P_DISABLED                                                     QStatus = 37092
	ER_P2P_BUSY                                                         QStatus = 37093
	ER_BUS_INCOMPATIBLE_DAEMON                                          QStatus = 37094
	ER_P2P_NO_GO                                                        QStatus = 37095
	ER_P2P_NO_STA                                                       QStatus = 37096
	ER_P2P_FORBIDDEN                                                    QStatus = 37097
	ER_ALLJOYN_ONAPPSUSPEND_REPLY_FAILED                                QStatus = 37098
	ER_ALLJOYN_ONAPPSUSPEND_REPLY_UNSUPPORTED                           QStatus = 37099
	ER_ALLJOYN_ONAPPRESUME_REPLY_FAILED                                 QStatus = 37100
	ER_ALLJOYN_ONAPPRESUME_REPLY_UNSUPPORTED                            QStatus = 37101
	ER_BUS_NO_SUCH_MESSAGE                                              QStatus = 37102
	ER_ALLJOYN_REMOVESESSIONMEMBER_REPLY_NO_SESSION                     QStatus = 37103
	ER_ALLJOYN_REMOVESESSIONMEMBER_NOT_BINDER                           QStatus = 37104
	ER_ALLJOYN_REMOVESESSIONMEMBER_NOT_MULTIPOINT                       QStatus = 37105
	ER_ALLJOYN_REMOVESESSIONMEMBER_NOT_FOUND                            QStatus = 37106
	ER_ALLJOYN_REMOVESESSIONMEMBER_INCOMPATIBLE_REMOTE_DAEMON           QStatus = 37107
	ER_ALLJOYN_REMOVESESSIONMEMBER_REPLY_FAILED                         QStatus = 37108
	ER_BUS_REMOVED_BY_BINDER                                            QStatus = 37109
	ER_BUS_MATCH_RULE_NOT_FOUND                                         QStatus = 37110
	ER_ALLJOYN_PING_FAILED                                              QStatus = 37111
	ER_ALLJOYN_PING_REPLY_UNREACHABLE                                   QStatus = 37112
	ER_UDP_MSG_TOO_LONG                                                 QStatus = 37113
	ER_UDP_DEMUX_NO_ENDPOINT                                            QStatus = 37114
	ER_UDP_NO_NETWORK                                                   QStatus = 37115
	ER_UDP_UNEXPECTED_LENGTH                                            QStatus = 37116
	ER_UDP_UNEXPECTED_FLOW                                              QStatus = 37117
	ER_UDP_DISCONNECT                                                   QStatus = 37118
	ER_UDP_NOT_IMPLEMENTED                                              QStatus = 37119
	ER_UDP_NO_LISTENER                                                  QStatus = 37120
	ER_UDP_STOPPING                                                     QStatus = 37121
	ER_ARDP_BACKPRESSURE                                                QStatus = 37122
	ER_UDP_BACKPRESSURE                                                 QStatus = 37123
	ER_ARDP_INVALID_STATE                                               QStatus = 37124
	ER_ARDP_TTL_EXPIRED                                                 QStatus = 37125
	ER_ARDP_PERSIST_TIMEOUT                                             QStatus = 37126
	ER_ARDP_PROBE_TIMEOUT                                               QStatus = 37127
	ER_ARDP_REMOTE_CONNECTION_RESET                                     QStatus = 37128
	ER_UDP_BUSHELLO                                                     QStatus = 37129
	ER_UDP_MESSAGE                                                      QStatus = 37130
	ER_UDP_INVALID                                                      QStatus = 37131
	ER_UDP_UNSUPPORTED                                                  QStatus = 37132
	ER_UDP_ENDPOINT_STALLED                                             QStatus = 37133
	ER_ARDP_INVALID_RESPONSE                                            QStatus = 37134
	ER_ARDP_INVALID_CONNECTION                                          QStatus = 37135
	ER_UDP_LOCAL_DISCONNECT                                             QStatus = 37136
	ER_UDP_EARLY_EXIT                                                   QStatus = 37137
	ER_UDP_LOCAL_DISCONNECT_FAIL                                        QStatus = 37138
	ER_ARDP_DISCONNECTING                                               QStatus = 37139
	ER_ALLJOYN_PING_REPLY_INCOMPATIBLE_REMOTE_ROUTING_NODE              QStatus = 37140
	ER_ALLJOYN_PING_REPLY_TIMEOUT                                       QStatus = 37141
	ER_ALLJOYN_PING_REPLY_UNKNOWN_NAME                                  QStatus = 37142
	ER_ALLJOYN_PING_REPLY_FAILED                                        QStatus = 37143
	ER_TCP_MAX_UNTRUSTED                                                QStatus = 37144
	ER_ALLJOYN_PING_REPLY_IN_PROGRESS                                   QStatus = 37145
	ER_LANGUAGE_NOT_SUPPORTED                                           QStatus = 37146
	ER_ABOUT_FIELD_ALREADY_SPECIFIED                                    QStatus = 37147
	ER_UDP_NOT_DISCONNECTED                                             QStatus = 37148
	ER_UDP_ENDPOINT_NOT_STARTED                                         QStatus = 37149
	ER_UDP_ENDPOINT_REMOVED                                             QStatus = 37150
	ER_ARDP_VERSION_NOT_SUPPORTED                                       QStatus = 37151
	ER_CONNECTION_LIMIT_EXCEEDED                                        QStatus = 37152
	ER_ARDP_WRITE_BLOCKED                                               QStatus = 37153
	ER_PERMISSION_DENIED                                                QStatus = 37154
	ER_ABOUT_DEFAULT_LANGUAGE_NOT_SPECIFIED                             QStatus = 37155
	ER_ABOUT_SESSIONPORT_NOT_BOUND                                      QStatus = 37156
	ER_ABOUT_ABOUTDATA_MISSING_REQUIRED_FIELD                           QStatus = 37157
	ER_ABOUT_INVALID_ABOUTDATA_LISTENER                                 QStatus = 37158
	ER_BUS_PING_GROUP_NOT_FOUND                                         QStatus = 37159
	ER_BUS_REMOVED_BY_BINDER_SELF                                       QStatus = 37160
	ER_INVALID_CONFIG                                                   QStatus = 37161
	ER_ABOUT_INVALID_ABOUTDATA_FIELD_VALUE                              QStatus = 37162
	ER_ABOUT_INVALID_ABOUTDATA_FIELD_APPID_SIZE                         QStatus = 37163
	ER_BUS_TRANSPORT_ACCESS_DENIED                                      QStatus = 37164
	ER_INVALID_CERTIFICATE                                              QStatus = 37165
	ER_CERTIFICATE_NOT_FOUND                                            QStatus = 37166
	ER_DUPLICATE_CERTIFICATE                                            QStatus = 37167
	ER_UNKNOWN_CERTIFICATE                                              QStatus = 37168
	ER_MISSING_DIGEST_IN_CERTIFICATE                                    QStatus = 37169
	ER_DIGEST_MISMATCH                                                  QStatus = 37170
	ER_DUPLICATE_KEY                                                    QStatus = 37171
	ER_NO_COMMON_TRUST                                                  QStatus = 37172
	ER_MANIFEST_NOT_FOUND                                               QStatus = 37173
	ER_INVALID_CERT_CHAIN                                               QStatus = 37174
	ER_NO_TRUST_ANCHOR                                                  QStatus = 37175
	ER_INVALID_APPLICATION_STATE                                        QStatus = 37176
	ER_FEATURE_NOT_AVAILABLE                                            QStatus = 37177
	ER_KEY_STORE_ALREADY_INITIALIZED                                    QStatus = 37178
	ER_KEY_STORE_ID_NOT_YET_SET                                         QStatus = 37179
	ER_POLICY_NOT_NEWER                                                 QStatus = 37180
	ER_MANIFEST_REJECTED                                                QStatus = 37181
	ER_INVALID_CERTIFICATE_USAGE                                        QStatus = 37182
	ER_INVALID_SIGNAL_EMISSION_TYPE                                     QStatus = 37183
	ER_APPLICATION_STATE_LISTENER_ALREADY_EXISTS                        QStatus = 37184
	ER_APPLICATION_STATE_LISTENER_NO_SUCH_LISTENER                      QStatus = 37185
	ER_MANAGEMENT_ALREADY_STARTED                                       QStatus = 37186
	ER_MANAGEMENT_NOT_STARTED                                           QStatus = 37187
	ER_BUS_DESCRIPTION_ALREADY_EXISTS                                   QStatus = 37188
)

type ALLJOYN_TYPEID int32

const (
	ALLJOYN_INVALID          ALLJOYN_TYPEID = 0
	ALLJOYN_ARRAY            ALLJOYN_TYPEID = 97
	ALLJOYN_BOOLEAN          ALLJOYN_TYPEID = 98
	ALLJOYN_DOUBLE           ALLJOYN_TYPEID = 100
	ALLJOYN_DICT_ENTRY       ALLJOYN_TYPEID = 101
	ALLJOYN_SIGNATURE        ALLJOYN_TYPEID = 103
	ALLJOYN_HANDLE           ALLJOYN_TYPEID = 104
	ALLJOYN_INT32            ALLJOYN_TYPEID = 105
	ALLJOYN_INT16            ALLJOYN_TYPEID = 110
	ALLJOYN_OBJECT_PATH      ALLJOYN_TYPEID = 111
	ALLJOYN_UINT16           ALLJOYN_TYPEID = 113
	ALLJOYN_STRUCT           ALLJOYN_TYPEID = 114
	ALLJOYN_STRING           ALLJOYN_TYPEID = 115
	ALLJOYN_UINT64           ALLJOYN_TYPEID = 116
	ALLJOYN_UINT32           ALLJOYN_TYPEID = 117
	ALLJOYN_VARIANT          ALLJOYN_TYPEID = 118
	ALLJOYN_INT64            ALLJOYN_TYPEID = 120
	ALLJOYN_BYTE             ALLJOYN_TYPEID = 121
	ALLJOYN_STRUCT_OPEN      ALLJOYN_TYPEID = 40
	ALLJOYN_STRUCT_CLOSE     ALLJOYN_TYPEID = 41
	ALLJOYN_DICT_ENTRY_OPEN  ALLJOYN_TYPEID = 123
	ALLJOYN_DICT_ENTRY_CLOSE ALLJOYN_TYPEID = 125
	ALLJOYN_BOOLEAN_ARRAY    ALLJOYN_TYPEID = 25185
	ALLJOYN_DOUBLE_ARRAY     ALLJOYN_TYPEID = 25697
	ALLJOYN_INT32_ARRAY      ALLJOYN_TYPEID = 26977
	ALLJOYN_INT16_ARRAY      ALLJOYN_TYPEID = 28257
	ALLJOYN_UINT16_ARRAY     ALLJOYN_TYPEID = 29025
	ALLJOYN_UINT64_ARRAY     ALLJOYN_TYPEID = 29793
	ALLJOYN_UINT32_ARRAY     ALLJOYN_TYPEID = 30049
	ALLJOYN_INT64_ARRAY      ALLJOYN_TYPEID = 30817
	ALLJOYN_BYTE_ARRAY       ALLJOYN_TYPEID = 31073
	ALLJOYN_WILDCARD         ALLJOYN_TYPEID = 42
)

type ALLJOYN_APPLICATIONSTATE int32

const (
	NOT_CLAIMABLE ALLJOYN_APPLICATIONSTATE = 0
	CLAIMABLE     ALLJOYN_APPLICATIONSTATE = 1
	CLAIMED       ALLJOYN_APPLICATIONSTATE = 2
	NEED_UPDATE   ALLJOYN_APPLICATIONSTATE = 3
)

type ALLJOYN_CLAIMCAPABILITY_MASKS int32

const (
	CAPABLE_ECDHE_NULL  ALLJOYN_CLAIMCAPABILITY_MASKS = 1
	CAPABLE_ECDHE_ECDSA ALLJOYN_CLAIMCAPABILITY_MASKS = 4
	CAPABLE_ECDHE_SPEKE ALLJOYN_CLAIMCAPABILITY_MASKS = 8
)

type ALLJOYN_CLAIMCAPABILITYADDITIONALINFO_MASKS int32

const (
	PASSWORD_GENERATED_BY_SECURITY_MANAGER ALLJOYN_CLAIMCAPABILITYADDITIONALINFO_MASKS = 1
	PASSWORD_GENERATED_BY_APPLICATION      ALLJOYN_CLAIMCAPABILITYADDITIONALINFO_MASKS = 2
)

type ALLJOYN_MESSAGETYPE int32

const (
	ALLJOYN_MESSAGE_INVALID     ALLJOYN_MESSAGETYPE = 0
	ALLJOYN_MESSAGE_METHOD_CALL ALLJOYN_MESSAGETYPE = 1
	ALLJOYN_MESSAGE_METHOD_RET  ALLJOYN_MESSAGETYPE = 2
	ALLJOYN_MESSAGE_ERROR       ALLJOYN_MESSAGETYPE = 3
	ALLJOYN_MESSAGE_SIGNAL      ALLJOYN_MESSAGETYPE = 4
)

type ALLJOYN_INTERFACEDESCRIPTION_SECURITYPOLICY int32

const (
	AJ_IFC_SECURITY_INHERIT  ALLJOYN_INTERFACEDESCRIPTION_SECURITYPOLICY = 0
	AJ_IFC_SECURITY_REQUIRED ALLJOYN_INTERFACEDESCRIPTION_SECURITYPOLICY = 1
	AJ_IFC_SECURITY_OFF      ALLJOYN_INTERFACEDESCRIPTION_SECURITYPOLICY = 2
)

type ALLJOYN_SESSIONLOSTREASON int32

const (
	ALLJOYN_SESSIONLOST_INVALID                    ALLJOYN_SESSIONLOSTREASON = 0
	ALLJOYN_SESSIONLOST_REMOTE_END_LEFT_SESSION    ALLJOYN_SESSIONLOSTREASON = 1
	ALLJOYN_SESSIONLOST_REMOTE_END_CLOSED_ABRUPTLY ALLJOYN_SESSIONLOSTREASON = 2
	ALLJOYN_SESSIONLOST_REMOVED_BY_BINDER          ALLJOYN_SESSIONLOSTREASON = 3
	ALLJOYN_SESSIONLOST_LINK_TIMEOUT               ALLJOYN_SESSIONLOSTREASON = 4
	ALLJOYN_SESSIONLOST_REASON_OTHER               ALLJOYN_SESSIONLOSTREASON = 5
)

// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package xmlhttpextendedrequest implements the Windows.Win32.XmlHttpExtendedRequest namespace.
package xmlhttpextendedrequest

type SERVERXMLHTTP_OPTION int32

const (
	SXH_OPTION_URL                                SERVERXMLHTTP_OPTION = -1
	SXH_OPTION_URL_CODEPAGE                       SERVERXMLHTTP_OPTION = 0
	SXH_OPTION_ESCAPE_PERCENT_IN_URL              SERVERXMLHTTP_OPTION = 1
	SXH_OPTION_IGNORE_SERVER_SSL_CERT_ERROR_FLAGS SERVERXMLHTTP_OPTION = 2
	SXH_OPTION_SELECT_CLIENT_SSL_CERT             SERVERXMLHTTP_OPTION = 3
)

type SXH_SERVER_CERT_OPTION int32

const (
	SXH_SERVER_CERT_IGNORE_UNKNOWN_CA        SXH_SERVER_CERT_OPTION = 256
	SXH_SERVER_CERT_IGNORE_WRONG_USAGE       SXH_SERVER_CERT_OPTION = 512
	SXH_SERVER_CERT_IGNORE_CERT_CN_INVALID   SXH_SERVER_CERT_OPTION = 4096
	SXH_SERVER_CERT_IGNORE_CERT_DATE_INVALID SXH_SERVER_CERT_OPTION = 8192
	SXH_SERVER_CERT_IGNORE_ALL_SERVER_ERRORS SXH_SERVER_CERT_OPTION = 13056
)

type SXH_PROXY_SETTING int32

const (
	SXH_PROXY_SET_DEFAULT   SXH_PROXY_SETTING = 0
	SXH_PROXY_SET_PRECONFIG SXH_PROXY_SETTING = 0
	SXH_PROXY_SET_DIRECT    SXH_PROXY_SETTING = 1
	SXH_PROXY_SET_PROXY     SXH_PROXY_SETTING = 2
)

type SOMITEMTYPE int32

const (
	SOMITEM_SCHEMA                      SOMITEMTYPE = 4096
	SOMITEM_ATTRIBUTE                   SOMITEMTYPE = 4097
	SOMITEM_ATTRIBUTEGROUP              SOMITEMTYPE = 4098
	SOMITEM_NOTATION                    SOMITEMTYPE = 4099
	SOMITEM_ANNOTATION                  SOMITEMTYPE = 4100
	SOMITEM_IDENTITYCONSTRAINT          SOMITEMTYPE = 4352
	SOMITEM_KEY                         SOMITEMTYPE = 4353
	SOMITEM_KEYREF                      SOMITEMTYPE = 4354
	SOMITEM_UNIQUE                      SOMITEMTYPE = 4355
	SOMITEM_ANYTYPE                     SOMITEMTYPE = 8192
	SOMITEM_DATATYPE                    SOMITEMTYPE = 8448
	SOMITEM_DATATYPE_ANYTYPE            SOMITEMTYPE = 8449
	SOMITEM_DATATYPE_ANYURI             SOMITEMTYPE = 8450
	SOMITEM_DATATYPE_BASE64BINARY       SOMITEMTYPE = 8451
	SOMITEM_DATATYPE_BOOLEAN            SOMITEMTYPE = 8452
	SOMITEM_DATATYPE_BYTE               SOMITEMTYPE = 8453
	SOMITEM_DATATYPE_DATE               SOMITEMTYPE = 8454
	SOMITEM_DATATYPE_DATETIME           SOMITEMTYPE = 8455
	SOMITEM_DATATYPE_DAY                SOMITEMTYPE = 8456
	SOMITEM_DATATYPE_DECIMAL            SOMITEMTYPE = 8457
	SOMITEM_DATATYPE_DOUBLE             SOMITEMTYPE = 8458
	SOMITEM_DATATYPE_DURATION           SOMITEMTYPE = 8459
	SOMITEM_DATATYPE_ENTITIES           SOMITEMTYPE = 8460
	SOMITEM_DATATYPE_ENTITY             SOMITEMTYPE = 8461
	SOMITEM_DATATYPE_FLOAT              SOMITEMTYPE = 8462
	SOMITEM_DATATYPE_HEXBINARY          SOMITEMTYPE = 8463
	SOMITEM_DATATYPE_ID                 SOMITEMTYPE = 8464
	SOMITEM_DATATYPE_IDREF              SOMITEMTYPE = 8465
	SOMITEM_DATATYPE_IDREFS             SOMITEMTYPE = 8466
	SOMITEM_DATATYPE_INT                SOMITEMTYPE = 8467
	SOMITEM_DATATYPE_INTEGER            SOMITEMTYPE = 8468
	SOMITEM_DATATYPE_LANGUAGE           SOMITEMTYPE = 8469
	SOMITEM_DATATYPE_LONG               SOMITEMTYPE = 8470
	SOMITEM_DATATYPE_MONTH              SOMITEMTYPE = 8471
	SOMITEM_DATATYPE_MONTHDAY           SOMITEMTYPE = 8472
	SOMITEM_DATATYPE_NAME               SOMITEMTYPE = 8473
	SOMITEM_DATATYPE_NCNAME             SOMITEMTYPE = 8474
	SOMITEM_DATATYPE_NEGATIVEINTEGER    SOMITEMTYPE = 8475
	SOMITEM_DATATYPE_NMTOKEN            SOMITEMTYPE = 8476
	SOMITEM_DATATYPE_NMTOKENS           SOMITEMTYPE = 8477
	SOMITEM_DATATYPE_NONNEGATIVEINTEGER SOMITEMTYPE = 8478
	SOMITEM_DATATYPE_NONPOSITIVEINTEGER SOMITEMTYPE = 8479
	SOMITEM_DATATYPE_NORMALIZEDSTRING   SOMITEMTYPE = 8480
	SOMITEM_DATATYPE_NOTATION           SOMITEMTYPE = 8481
	SOMITEM_DATATYPE_POSITIVEINTEGER    SOMITEMTYPE = 8482
	SOMITEM_DATATYPE_QNAME              SOMITEMTYPE = 8483
	SOMITEM_DATATYPE_SHORT              SOMITEMTYPE = 8484
	SOMITEM_DATATYPE_STRING             SOMITEMTYPE = 8485
	SOMITEM_DATATYPE_TIME               SOMITEMTYPE = 8486
	SOMITEM_DATATYPE_TOKEN              SOMITEMTYPE = 8487
	SOMITEM_DATATYPE_UNSIGNEDBYTE       SOMITEMTYPE = 8488
	SOMITEM_DATATYPE_UNSIGNEDINT        SOMITEMTYPE = 8489
	SOMITEM_DATATYPE_UNSIGNEDLONG       SOMITEMTYPE = 8490
	SOMITEM_DATATYPE_UNSIGNEDSHORT      SOMITEMTYPE = 8491
	SOMITEM_DATATYPE_YEAR               SOMITEMTYPE = 8492
	SOMITEM_DATATYPE_YEARMONTH          SOMITEMTYPE = 8493
	SOMITEM_DATATYPE_ANYSIMPLETYPE      SOMITEMTYPE = 8703
	SOMITEM_SIMPLETYPE                  SOMITEMTYPE = 8704
	SOMITEM_COMPLEXTYPE                 SOMITEMTYPE = 9216
	SOMITEM_PARTICLE                    SOMITEMTYPE = 16384
	SOMITEM_ANY                         SOMITEMTYPE = 16385
	SOMITEM_ANYATTRIBUTE                SOMITEMTYPE = 16386
	SOMITEM_ELEMENT                     SOMITEMTYPE = 16387
	SOMITEM_GROUP                       SOMITEMTYPE = 16640
	SOMITEM_ALL                         SOMITEMTYPE = 16641
	SOMITEM_CHOICE                      SOMITEMTYPE = 16642
	SOMITEM_SEQUENCE                    SOMITEMTYPE = 16643
	SOMITEM_EMPTYPARTICLE               SOMITEMTYPE = 16644
	SOMITEM_NULL                        SOMITEMTYPE = 2048
	SOMITEM_NULL_TYPE                   SOMITEMTYPE = 10240
	SOMITEM_NULL_ANY                    SOMITEMTYPE = 18433
	SOMITEM_NULL_ANYATTRIBUTE           SOMITEMTYPE = 18434
	SOMITEM_NULL_ELEMENT                SOMITEMTYPE = 18435
)

type SCHEMAUSE int32

const (
	SCHEMAUSE_OPTIONAL   SCHEMAUSE = 0
	SCHEMAUSE_PROHIBITED SCHEMAUSE = 1
	SCHEMAUSE_REQUIRED   SCHEMAUSE = 2
)

type SCHEMADERIVATIONMETHOD int32

const (
	SCHEMADERIVATIONMETHOD_EMPTY        SCHEMADERIVATIONMETHOD = 0
	SCHEMADERIVATIONMETHOD_SUBSTITUTION SCHEMADERIVATIONMETHOD = 1
	SCHEMADERIVATIONMETHOD_EXTENSION    SCHEMADERIVATIONMETHOD = 2
	SCHEMADERIVATIONMETHOD_RESTRICTION  SCHEMADERIVATIONMETHOD = 4
	SCHEMADERIVATIONMETHOD_LIST         SCHEMADERIVATIONMETHOD = 8
	SCHEMADERIVATIONMETHOD_UNION        SCHEMADERIVATIONMETHOD = 16
	SCHEMADERIVATIONMETHOD_ALL          SCHEMADERIVATIONMETHOD = 255
	SCHEMADERIVATIONMETHOD_NONE         SCHEMADERIVATIONMETHOD = 256
)

type SCHEMACONTENTTYPE int32

const (
	SCHEMACONTENTTYPE_EMPTY       SCHEMACONTENTTYPE = 0
	SCHEMACONTENTTYPE_TEXTONLY    SCHEMACONTENTTYPE = 1
	SCHEMACONTENTTYPE_ELEMENTONLY SCHEMACONTENTTYPE = 2
	SCHEMACONTENTTYPE_MIXED       SCHEMACONTENTTYPE = 3
)

type SCHEMAPROCESSCONTENTS int32

const (
	SCHEMAPROCESSCONTENTS_NONE   SCHEMAPROCESSCONTENTS = 0
	SCHEMAPROCESSCONTENTS_SKIP   SCHEMAPROCESSCONTENTS = 1
	SCHEMAPROCESSCONTENTS_LAX    SCHEMAPROCESSCONTENTS = 2
	SCHEMAPROCESSCONTENTS_STRICT SCHEMAPROCESSCONTENTS = 3
)

type SCHEMAWHITESPACE int32

const (
	SCHEMAWHITESPACE_NONE     SCHEMAWHITESPACE = -1
	SCHEMAWHITESPACE_PRESERVE SCHEMAWHITESPACE = 0
	SCHEMAWHITESPACE_REPLACE  SCHEMAWHITESPACE = 1
	SCHEMAWHITESPACE_COLLAPSE SCHEMAWHITESPACE = 2
)

type SCHEMATYPEVARIETY int32

const (
	SCHEMATYPEVARIETY_NONE   SCHEMATYPEVARIETY = -1
	SCHEMATYPEVARIETY_ATOMIC SCHEMATYPEVARIETY = 0
	SCHEMATYPEVARIETY_LIST   SCHEMATYPEVARIETY = 1
	SCHEMATYPEVARIETY_UNION  SCHEMATYPEVARIETY = 2
)

type XHR_COOKIE_STATE int32

const (
	XHR_COOKIE_STATE_UNKNOWN   XHR_COOKIE_STATE = 0
	XHR_COOKIE_STATE_ACCEPT    XHR_COOKIE_STATE = 1
	XHR_COOKIE_STATE_PROMPT    XHR_COOKIE_STATE = 2
	XHR_COOKIE_STATE_LEASH     XHR_COOKIE_STATE = 3
	XHR_COOKIE_STATE_DOWNGRADE XHR_COOKIE_STATE = 4
	XHR_COOKIE_STATE_REJECT    XHR_COOKIE_STATE = 5
)

type XHR_COOKIE_FLAG int32

const (
	XHR_COOKIE_IS_SECURE       XHR_COOKIE_FLAG = 1
	XHR_COOKIE_IS_SESSION      XHR_COOKIE_FLAG = 2
	XHR_COOKIE_THIRD_PARTY     XHR_COOKIE_FLAG = 16
	XHR_COOKIE_PROMPT_REQUIRED XHR_COOKIE_FLAG = 32
	XHR_COOKIE_EVALUATE_P3P    XHR_COOKIE_FLAG = 64
	XHR_COOKIE_APPLY_P3P       XHR_COOKIE_FLAG = 128
	XHR_COOKIE_P3P_ENABLED     XHR_COOKIE_FLAG = 256
	XHR_COOKIE_IS_RESTRICTED   XHR_COOKIE_FLAG = 512
	XHR_COOKIE_IE6             XHR_COOKIE_FLAG = 1024
	XHR_COOKIE_IS_LEGACY       XHR_COOKIE_FLAG = 2048
	XHR_COOKIE_NON_SCRIPT      XHR_COOKIE_FLAG = 4096
	XHR_COOKIE_HTTPONLY        XHR_COOKIE_FLAG = 8192
)

type XHR_CRED_PROMPT int32

const (
	XHR_CRED_PROMPT_ALL   XHR_CRED_PROMPT = 0
	XHR_CRED_PROMPT_NONE  XHR_CRED_PROMPT = 1
	XHR_CRED_PROMPT_PROXY XHR_CRED_PROMPT = 2
)

type XHR_AUTH int32

const (
	XHR_AUTH_ALL   XHR_AUTH = 0
	XHR_AUTH_NONE  XHR_AUTH = 1
	XHR_AUTH_PROXY XHR_AUTH = 2
)

type XHR_PROPERTY int32

const (
	XHR_PROP_NO_CRED_PROMPT         XHR_PROPERTY = 0
	XHR_PROP_NO_AUTH                XHR_PROPERTY = 1
	XHR_PROP_TIMEOUT                XHR_PROPERTY = 2
	XHR_PROP_NO_DEFAULT_HEADERS     XHR_PROPERTY = 3
	XHR_PROP_REPORT_REDIRECT_STATUS XHR_PROPERTY = 4
	XHR_PROP_NO_CACHE               XHR_PROPERTY = 5
	XHR_PROP_EXTENDED_ERROR         XHR_PROPERTY = 6
	XHR_PROP_QUERY_STRING_UTF8      XHR_PROPERTY = 7
	XHR_PROP_IGNORE_CERT_ERRORS     XHR_PROPERTY = 8
	XHR_PROP_ONDATA_THRESHOLD       XHR_PROPERTY = 9
	XHR_PROP_SET_ENTERPRISEID       XHR_PROPERTY = 10
	XHR_PROP_MAX_CONNECTIONS        XHR_PROPERTY = 11
)

type XHR_CERT_IGNORE_FLAG uint32

const (
	XHR_CERT_IGNORE_REVOCATION_FAILED XHR_CERT_IGNORE_FLAG = 128
	XHR_CERT_IGNORE_UNKNOWN_CA        XHR_CERT_IGNORE_FLAG = 256
	XHR_CERT_IGNORE_CERT_CN_INVALID   XHR_CERT_IGNORE_FLAG = 4096
	XHR_CERT_IGNORE_CERT_DATE_INVALID XHR_CERT_IGNORE_FLAG = 8192
	XHR_CERT_IGNORE_ALL_SERVER_ERRORS XHR_CERT_IGNORE_FLAG = 12672
)

type XHR_CERT_ERROR_FLAG uint32

const (
	XHR_CERT_ERROR_REVOCATION_FAILED XHR_CERT_ERROR_FLAG = 8388608
	XHR_CERT_ERROR_UNKNOWN_CA        XHR_CERT_ERROR_FLAG = 16777216
	XHR_CERT_ERROR_CERT_CN_INVALID   XHR_CERT_ERROR_FLAG = 33554432
	XHR_CERT_ERROR_CERT_DATE_INVALID XHR_CERT_ERROR_FLAG = 67108864
	XHR_CERT_ERROR_ALL_SERVER_ERRORS XHR_CERT_ERROR_FLAG = 125829120
)

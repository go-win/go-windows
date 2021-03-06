// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package automation implements the Windows.Win32.Automation namespace.
package automation

type SF_TYPE int32

const (
	SF_ERROR    SF_TYPE = 10
	SF_I1       SF_TYPE = 16
	SF_I2       SF_TYPE = 2
	SF_I4       SF_TYPE = 3
	SF_I8       SF_TYPE = 20
	SF_BSTR     SF_TYPE = 8
	SF_UNKNOWN  SF_TYPE = 13
	SF_DISPATCH SF_TYPE = 9
	SF_VARIANT  SF_TYPE = 12
	SF_RECORD   SF_TYPE = 36
	SF_HAVEIID  SF_TYPE = 32781
)

type TYPEKIND int32

const (
	TKIND_ENUM      TYPEKIND = 0
	TKIND_RECORD    TYPEKIND = 1
	TKIND_MODULE    TYPEKIND = 2
	TKIND_INTERFACE TYPEKIND = 3
	TKIND_DISPATCH  TYPEKIND = 4
	TKIND_COCLASS   TYPEKIND = 5
	TKIND_ALIAS     TYPEKIND = 6
	TKIND_UNION     TYPEKIND = 7
	TKIND_MAX       TYPEKIND = 8
)

type CALLCONV int32

const (
	CC_FASTCALL   CALLCONV = 0
	CC_CDECL      CALLCONV = 1
	CC_MSCPASCAL  CALLCONV = 2
	CC_PASCAL     CALLCONV = 2
	CC_MACPASCAL  CALLCONV = 3
	CC_STDCALL    CALLCONV = 4
	CC_FPFASTCALL CALLCONV = 5
	CC_SYSCALL    CALLCONV = 6
	CC_MPWCDECL   CALLCONV = 7
	CC_MPWPASCAL  CALLCONV = 8
	CC_MAX        CALLCONV = 9
)

type FUNCKIND int32

const (
	FUNC_VIRTUAL     FUNCKIND = 0
	FUNC_PUREVIRTUAL FUNCKIND = 1
	FUNC_NONVIRTUAL  FUNCKIND = 2
	FUNC_STATIC      FUNCKIND = 3
	FUNC_DISPATCH    FUNCKIND = 4
)

type INVOKEKIND int32

const (
	INVOKE_FUNC           INVOKEKIND = 1
	INVOKE_PROPERTYGET    INVOKEKIND = 2
	INVOKE_PROPERTYPUT    INVOKEKIND = 4
	INVOKE_PROPERTYPUTREF INVOKEKIND = 8
)

type VARKIND int32

const (
	VAR_PERINSTANCE VARKIND = 0
	VAR_STATIC      VARKIND = 1
	VAR_CONST       VARKIND = 2
	VAR_DISPATCH    VARKIND = 3
)

type TYPEFLAGS int32

const (
	TYPEFLAG_FAPPOBJECT     TYPEFLAGS = 1
	TYPEFLAG_FCANCREATE     TYPEFLAGS = 2
	TYPEFLAG_FLICENSED      TYPEFLAGS = 4
	TYPEFLAG_FPREDECLID     TYPEFLAGS = 8
	TYPEFLAG_FHIDDEN        TYPEFLAGS = 16
	TYPEFLAG_FCONTROL       TYPEFLAGS = 32
	TYPEFLAG_FDUAL          TYPEFLAGS = 64
	TYPEFLAG_FNONEXTENSIBLE TYPEFLAGS = 128
	TYPEFLAG_FOLEAUTOMATION TYPEFLAGS = 256
	TYPEFLAG_FRESTRICTED    TYPEFLAGS = 512
	TYPEFLAG_FAGGREGATABLE  TYPEFLAGS = 1024
	TYPEFLAG_FREPLACEABLE   TYPEFLAGS = 2048
	TYPEFLAG_FDISPATCHABLE  TYPEFLAGS = 4096
	TYPEFLAG_FREVERSEBIND   TYPEFLAGS = 8192
	TYPEFLAG_FPROXY         TYPEFLAGS = 16384
)

type FUNCFLAGS int32

const (
	FUNCFLAG_FRESTRICTED       FUNCFLAGS = 1
	FUNCFLAG_FSOURCE           FUNCFLAGS = 2
	FUNCFLAG_FBINDABLE         FUNCFLAGS = 4
	FUNCFLAG_FREQUESTEDIT      FUNCFLAGS = 8
	FUNCFLAG_FDISPLAYBIND      FUNCFLAGS = 16
	FUNCFLAG_FDEFAULTBIND      FUNCFLAGS = 32
	FUNCFLAG_FHIDDEN           FUNCFLAGS = 64
	FUNCFLAG_FUSESGETLASTERROR FUNCFLAGS = 128
	FUNCFLAG_FDEFAULTCOLLELEM  FUNCFLAGS = 256
	FUNCFLAG_FUIDEFAULT        FUNCFLAGS = 512
	FUNCFLAG_FNONBROWSABLE     FUNCFLAGS = 1024
	FUNCFLAG_FREPLACEABLE      FUNCFLAGS = 2048
	FUNCFLAG_FIMMEDIATEBIND    FUNCFLAGS = 4096
)

type VARFLAGS int32

const (
	VARFLAG_FREADONLY        VARFLAGS = 1
	VARFLAG_FSOURCE          VARFLAGS = 2
	VARFLAG_FBINDABLE        VARFLAGS = 4
	VARFLAG_FREQUESTEDIT     VARFLAGS = 8
	VARFLAG_FDISPLAYBIND     VARFLAGS = 16
	VARFLAG_FDEFAULTBIND     VARFLAGS = 32
	VARFLAG_FHIDDEN          VARFLAGS = 64
	VARFLAG_FRESTRICTED      VARFLAGS = 128
	VARFLAG_FDEFAULTCOLLELEM VARFLAGS = 256
	VARFLAG_FUIDEFAULT       VARFLAGS = 512
	VARFLAG_FNONBROWSABLE    VARFLAGS = 1024
	VARFLAG_FREPLACEABLE     VARFLAGS = 2048
	VARFLAG_FIMMEDIATEBIND   VARFLAGS = 4096
)

type DESCKIND int32

const (
	DESCKIND_NONE           DESCKIND = 0
	DESCKIND_FUNCDESC       DESCKIND = 1
	DESCKIND_VARDESC        DESCKIND = 2
	DESCKIND_TYPECOMP       DESCKIND = 3
	DESCKIND_IMPLICITAPPOBJ DESCKIND = 4
	DESCKIND_MAX            DESCKIND = 5
)

type SYSKIND int32

const (
	SYS_WIN16 SYSKIND = 0
	SYS_WIN32 SYSKIND = 1
	SYS_MAC   SYSKIND = 2
	SYS_WIN64 SYSKIND = 3
)

type LIBFLAGS int32

const (
	LIBFLAG_FRESTRICTED   LIBFLAGS = 1
	LIBFLAG_FCONTROL      LIBFLAGS = 2
	LIBFLAG_FHIDDEN       LIBFLAGS = 4
	LIBFLAG_FHASDISKIMAGE LIBFLAGS = 8
)

type CHANGEKIND int32

const (
	CHANGEKIND_ADDMEMBER        CHANGEKIND = 0
	CHANGEKIND_DELETEMEMBER     CHANGEKIND = 1
	CHANGEKIND_SETNAMES         CHANGEKIND = 2
	CHANGEKIND_SETDOCUMENTATION CHANGEKIND = 3
	CHANGEKIND_GENERAL          CHANGEKIND = 4
	CHANGEKIND_INVALIDATE       CHANGEKIND = 5
	CHANGEKIND_CHANGEFAILED     CHANGEKIND = 6
	CHANGEKIND_MAX              CHANGEKIND = 7
)

type REGKIND int32

const (
	REGKIND_DEFAULT  REGKIND = 0
	REGKIND_REGISTER REGKIND = 1
	REGKIND_NONE     REGKIND = 2
)

type VARENUM int32

const (
	VT_EMPTY            VARENUM = 0
	VT_NULL             VARENUM = 1
	VT_I2               VARENUM = 2
	VT_I4               VARENUM = 3
	VT_R4               VARENUM = 4
	VT_R8               VARENUM = 5
	VT_CY               VARENUM = 6
	VT_DATE             VARENUM = 7
	VT_BSTR             VARENUM = 8
	VT_DISPATCH         VARENUM = 9
	VT_ERROR            VARENUM = 10
	VT_BOOL             VARENUM = 11
	VT_VARIANT          VARENUM = 12
	VT_UNKNOWN          VARENUM = 13
	VT_DECIMAL          VARENUM = 14
	VT_I1               VARENUM = 16
	VT_UI1              VARENUM = 17
	VT_UI2              VARENUM = 18
	VT_UI4              VARENUM = 19
	VT_I8               VARENUM = 20
	VT_UI8              VARENUM = 21
	VT_INT              VARENUM = 22
	VT_UINT             VARENUM = 23
	VT_VOID             VARENUM = 24
	VT_HRESULT          VARENUM = 25
	VT_PTR              VARENUM = 26
	VT_SAFEARRAY        VARENUM = 27
	VT_CARRAY           VARENUM = 28
	VT_USERDEFINED      VARENUM = 29
	VT_LPSTR            VARENUM = 30
	VT_LPWSTR           VARENUM = 31
	VT_RECORD           VARENUM = 36
	VT_INT_PTR          VARENUM = 37
	VT_UINT_PTR         VARENUM = 38
	VT_FILETIME         VARENUM = 64
	VT_BLOB             VARENUM = 65
	VT_STREAM           VARENUM = 66
	VT_STORAGE          VARENUM = 67
	VT_STREAMED_OBJECT  VARENUM = 68
	VT_STORED_OBJECT    VARENUM = 69
	VT_BLOB_OBJECT      VARENUM = 70
	VT_CF               VARENUM = 71
	VT_CLSID            VARENUM = 72
	VT_VERSIONED_STREAM VARENUM = 73
	VT_BSTR_BLOB        VARENUM = 4095
	VT_VECTOR           VARENUM = 4096
	VT_ARRAY            VARENUM = 8192
	VT_BYREF            VARENUM = 16384
	VT_RESERVED         VARENUM = 32768
	VT_ILLEGAL          VARENUM = 65535
	VT_ILLEGALMASKED    VARENUM = 4095
	VT_TYPEMASK         VARENUM = 4095
)

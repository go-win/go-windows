// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package applicationinstallationandservicing implements the Windows.Win32.ApplicationInstallationAndServicing namespace.
package applicationinstallationandservicing

type ACTCTX_REQUESTED_RUN_LEVEL int32

const (
	ACTCTX_RUN_LEVEL_UNSPECIFIED       ACTCTX_REQUESTED_RUN_LEVEL = 0
	ACTCTX_RUN_LEVEL_AS_INVOKER        ACTCTX_REQUESTED_RUN_LEVEL = 1
	ACTCTX_RUN_LEVEL_HIGHEST_AVAILABLE ACTCTX_REQUESTED_RUN_LEVEL = 2
	ACTCTX_RUN_LEVEL_REQUIRE_ADMIN     ACTCTX_REQUESTED_RUN_LEVEL = 3
	ACTCTX_RUN_LEVEL_NUMBERS           ACTCTX_REQUESTED_RUN_LEVEL = 4
)

type ACTCTX_COMPATIBILITY_ELEMENT_TYPE int32

const (
	ACTCTX_COMPATIBILITY_ELEMENT_TYPE_UNKNOWN          ACTCTX_COMPATIBILITY_ELEMENT_TYPE = 0
	ACTCTX_COMPATIBILITY_ELEMENT_TYPE_OS               ACTCTX_COMPATIBILITY_ELEMENT_TYPE = 1
	ACTCTX_COMPATIBILITY_ELEMENT_TYPE_MITIGATION       ACTCTX_COMPATIBILITY_ELEMENT_TYPE = 2
	ACTCTX_COMPATIBILITY_ELEMENT_TYPE_MAXVERSIONTESTED ACTCTX_COMPATIBILITY_ELEMENT_TYPE = 3
)

type RESULTTYPES int32

const (
	IEUNKNOWN RESULTTYPES = 0
	IEERROR   RESULTTYPES = 1
	IEWARNING RESULTTYPES = 2
	IEINFO    RESULTTYPES = 3
)

type STATUSTYPES int32

const (
	IESTATUSGETCUB       STATUSTYPES = 0
	IESTATUSICECOUNT     STATUSTYPES = 1
	IESTATUSMERGE        STATUSTYPES = 2
	IESTATUSSUMMARYINFO  STATUSTYPES = 3
	IESTATUSCREATEENGINE STATUSTYPES = 4
	IESTATUSSTARTING     STATUSTYPES = 5
	IESTATUSRUNICE       STATUSTYPES = 6
	IESTATUSSHUTDOWN     STATUSTYPES = 7
	IESTATUSSUCCESS      STATUSTYPES = 8
	IESTATUSFAIL         STATUSTYPES = 9
	IESTATUSCANCEL       STATUSTYPES = 10
)

type MSMERRORTYPE int32

const (
	MSMERRORLANGUAGEUNSUPPORTED MSMERRORTYPE = 1
	MSMERRORLANGUAGEFAILED      MSMERRORTYPE = 2
	MSMERROREXCLUSION           MSMERRORTYPE = 3
	MSMERRORTABLEMERGE          MSMERRORTYPE = 4
	MSMERRORRESEQUENCEMERGE     MSMERRORTYPE = 5
	MSMERRORFILECREATE          MSMERRORTYPE = 6
	MSMERRORDIRCREATE           MSMERRORTYPE = 7
	MSMERRORFEATUREREQUIRED     MSMERRORTYPE = 8
)

type INSTALLMESSAGE int32

const (
	INSTALLMESSAGE_FATALEXIT      INSTALLMESSAGE = 0
	INSTALLMESSAGE_ERROR          INSTALLMESSAGE = 16777216
	INSTALLMESSAGE_WARNING        INSTALLMESSAGE = 33554432
	INSTALLMESSAGE_USER           INSTALLMESSAGE = 50331648
	INSTALLMESSAGE_INFO           INSTALLMESSAGE = 67108864
	INSTALLMESSAGE_FILESINUSE     INSTALLMESSAGE = 83886080
	INSTALLMESSAGE_RESOLVESOURCE  INSTALLMESSAGE = 100663296
	INSTALLMESSAGE_OUTOFDISKSPACE INSTALLMESSAGE = 117440512
	INSTALLMESSAGE_ACTIONSTART    INSTALLMESSAGE = 134217728
	INSTALLMESSAGE_ACTIONDATA     INSTALLMESSAGE = 150994944
	INSTALLMESSAGE_PROGRESS       INSTALLMESSAGE = 167772160
	INSTALLMESSAGE_COMMONDATA     INSTALLMESSAGE = 184549376
	INSTALLMESSAGE_INITIALIZE     INSTALLMESSAGE = 201326592
	INSTALLMESSAGE_TERMINATE      INSTALLMESSAGE = 218103808
	INSTALLMESSAGE_SHOWDIALOG     INSTALLMESSAGE = 234881024
	INSTALLMESSAGE_PERFORMANCE    INSTALLMESSAGE = 251658240
	INSTALLMESSAGE_RMFILESINUSE   INSTALLMESSAGE = 419430400
	INSTALLMESSAGE_INSTALLSTART   INSTALLMESSAGE = 436207616
	INSTALLMESSAGE_INSTALLEND     INSTALLMESSAGE = 452984832
)

type INSTALLUILEVEL int32

const (
	INSTALLUILEVEL_NOCHANGE      INSTALLUILEVEL = 0
	INSTALLUILEVEL_DEFAULT       INSTALLUILEVEL = 1
	INSTALLUILEVEL_NONE          INSTALLUILEVEL = 2
	INSTALLUILEVEL_BASIC         INSTALLUILEVEL = 3
	INSTALLUILEVEL_REDUCED       INSTALLUILEVEL = 4
	INSTALLUILEVEL_FULL          INSTALLUILEVEL = 5
	INSTALLUILEVEL_ENDDIALOG     INSTALLUILEVEL = 128
	INSTALLUILEVEL_PROGRESSONLY  INSTALLUILEVEL = 64
	INSTALLUILEVEL_HIDECANCEL    INSTALLUILEVEL = 32
	INSTALLUILEVEL_SOURCERESONLY INSTALLUILEVEL = 256
	INSTALLUILEVEL_UACONLY       INSTALLUILEVEL = 512
)

type INSTALLSTATE int32

const (
	INSTALLSTATE_NOTUSED      INSTALLSTATE = -7
	INSTALLSTATE_BADCONFIG    INSTALLSTATE = -6
	INSTALLSTATE_INCOMPLETE   INSTALLSTATE = -5
	INSTALLSTATE_SOURCEABSENT INSTALLSTATE = -4
	INSTALLSTATE_MOREDATA     INSTALLSTATE = -3
	INSTALLSTATE_INVALIDARG   INSTALLSTATE = -2
	INSTALLSTATE_UNKNOWN      INSTALLSTATE = -1
	INSTALLSTATE_BROKEN       INSTALLSTATE = 0
	INSTALLSTATE_ADVERTISED   INSTALLSTATE = 1
	INSTALLSTATE_REMOVED      INSTALLSTATE = 1
	INSTALLSTATE_ABSENT       INSTALLSTATE = 2
	INSTALLSTATE_LOCAL        INSTALLSTATE = 3
	INSTALLSTATE_SOURCE       INSTALLSTATE = 4
	INSTALLSTATE_DEFAULT      INSTALLSTATE = 5
)

type USERINFOSTATE int32

const (
	USERINFOSTATE_MOREDATA   USERINFOSTATE = -3
	USERINFOSTATE_INVALIDARG USERINFOSTATE = -2
	USERINFOSTATE_UNKNOWN    USERINFOSTATE = -1
	USERINFOSTATE_ABSENT     USERINFOSTATE = 0
	USERINFOSTATE_PRESENT    USERINFOSTATE = 1
)

type INSTALLLEVEL int32

const (
	INSTALLLEVEL_DEFAULT INSTALLLEVEL = 0
	INSTALLLEVEL_MINIMUM INSTALLLEVEL = 1
	INSTALLLEVEL_MAXIMUM INSTALLLEVEL = 65535
)

type REINSTALLMODE int32

const (
	REINSTALLMODE_REPAIR           REINSTALLMODE = 1
	REINSTALLMODE_FILEMISSING      REINSTALLMODE = 2
	REINSTALLMODE_FILEOLDERVERSION REINSTALLMODE = 4
	REINSTALLMODE_FILEEQUALVERSION REINSTALLMODE = 8
	REINSTALLMODE_FILEEXACT        REINSTALLMODE = 16
	REINSTALLMODE_FILEVERIFY       REINSTALLMODE = 32
	REINSTALLMODE_FILEREPLACE      REINSTALLMODE = 64
	REINSTALLMODE_MACHINEDATA      REINSTALLMODE = 128
	REINSTALLMODE_USERDATA         REINSTALLMODE = 256
	REINSTALLMODE_SHORTCUT         REINSTALLMODE = 512
	REINSTALLMODE_PACKAGE          REINSTALLMODE = 1024
)

type TAGINSTALLOGMODE int32

const (
	INSTALLLOGMODE_FATALEXIT      TAGINSTALLOGMODE = 1
	INSTALLLOGMODE_ERROR          TAGINSTALLOGMODE = 2
	INSTALLLOGMODE_WARNING        TAGINSTALLOGMODE = 4
	INSTALLLOGMODE_USER           TAGINSTALLOGMODE = 8
	INSTALLLOGMODE_INFO           TAGINSTALLOGMODE = 16
	INSTALLLOGMODE_RESOLVESOURCE  TAGINSTALLOGMODE = 64
	INSTALLLOGMODE_OUTOFDISKSPACE TAGINSTALLOGMODE = 128
	INSTALLLOGMODE_ACTIONSTART    TAGINSTALLOGMODE = 256
	INSTALLLOGMODE_ACTIONDATA     TAGINSTALLOGMODE = 512
	INSTALLLOGMODE_COMMONDATA     TAGINSTALLOGMODE = 2048
	INSTALLLOGMODE_PROPERTYDUMP   TAGINSTALLOGMODE = 1024
	INSTALLLOGMODE_VERBOSE        TAGINSTALLOGMODE = 4096
	INSTALLLOGMODE_EXTRADEBUG     TAGINSTALLOGMODE = 8192
	INSTALLLOGMODE_LOGONLYONERROR TAGINSTALLOGMODE = 16384
	INSTALLLOGMODE_LOGPERFORMANCE TAGINSTALLOGMODE = 32768
	INSTALLLOGMODE_PROGRESS       TAGINSTALLOGMODE = 1024
	INSTALLLOGMODE_INITIALIZE     TAGINSTALLOGMODE = 4096
	INSTALLLOGMODE_TERMINATE      TAGINSTALLOGMODE = 8192
	INSTALLLOGMODE_SHOWDIALOG     TAGINSTALLOGMODE = 16384
	INSTALLLOGMODE_FILESINUSE     TAGINSTALLOGMODE = 32
	INSTALLLOGMODE_RMFILESINUSE   TAGINSTALLOGMODE = 33554432
	INSTALLLOGMODE_INSTALLSTART   TAGINSTALLOGMODE = 67108864
	INSTALLLOGMODE_INSTALLEND     TAGINSTALLOGMODE = 134217728
)

type INSTALLLOGATTRIBUTES int32

const (
	INSTALLLOGATTRIBUTES_APPEND        INSTALLLOGATTRIBUTES = 1
	INSTALLLOGATTRIBUTES_FLUSHEACHLINE INSTALLLOGATTRIBUTES = 2
)

type INSTALLFEATUREATTRIBUTE int32

const (
	INSTALLFEATUREATTRIBUTE_FAVORLOCAL             INSTALLFEATUREATTRIBUTE = 1
	INSTALLFEATUREATTRIBUTE_FAVORSOURCE            INSTALLFEATUREATTRIBUTE = 2
	INSTALLFEATUREATTRIBUTE_FOLLOWPARENT           INSTALLFEATUREATTRIBUTE = 4
	INSTALLFEATUREATTRIBUTE_FAVORADVERTISE         INSTALLFEATUREATTRIBUTE = 8
	INSTALLFEATUREATTRIBUTE_DISALLOWADVERTISE      INSTALLFEATUREATTRIBUTE = 16
	INSTALLFEATUREATTRIBUTE_NOUNSUPPORTEDADVERTISE INSTALLFEATUREATTRIBUTE = 32
)

type INSTALLMODE int32

const (
	INSTALLMODE_NODETECTION_ANY    INSTALLMODE = -4
	INSTALLMODE_NOSOURCERESOLUTION INSTALLMODE = -3
	INSTALLMODE_NODETECTION        INSTALLMODE = -2
	INSTALLMODE_EXISTING           INSTALLMODE = -1
	INSTALLMODE_DEFAULT            INSTALLMODE = 0
)

type MSIPATCHSTATE int32

const (
	MSIPATCHSTATE_INVALID    MSIPATCHSTATE = 0
	MSIPATCHSTATE_APPLIED    MSIPATCHSTATE = 1
	MSIPATCHSTATE_SUPERSEDED MSIPATCHSTATE = 2
	MSIPATCHSTATE_OBSOLETED  MSIPATCHSTATE = 4
	MSIPATCHSTATE_REGISTERED MSIPATCHSTATE = 8
	MSIPATCHSTATE_ALL        MSIPATCHSTATE = 15
)

type MSIINSTALLCONTEXT int32

const (
	MSIINSTALLCONTEXT_FIRSTVISIBLE   MSIINSTALLCONTEXT = 0
	MSIINSTALLCONTEXT_NONE           MSIINSTALLCONTEXT = 0
	MSIINSTALLCONTEXT_USERMANAGED    MSIINSTALLCONTEXT = 1
	MSIINSTALLCONTEXT_USERUNMANAGED  MSIINSTALLCONTEXT = 2
	MSIINSTALLCONTEXT_MACHINE        MSIINSTALLCONTEXT = 4
	MSIINSTALLCONTEXT_ALL            MSIINSTALLCONTEXT = 7
	MSIINSTALLCONTEXT_ALLUSERMANAGED MSIINSTALLCONTEXT = 8
)

type MSIPATCHDATATYPE int32

const (
	MSIPATCH_DATATYPE_PATCHFILE MSIPATCHDATATYPE = 0
	MSIPATCH_DATATYPE_XMLPATH   MSIPATCHDATATYPE = 1
	MSIPATCH_DATATYPE_XMLBLOB   MSIPATCHDATATYPE = 2
)

type SCRIPTFLAGS int32

const (
	SCRIPTFLAGS_CACHEINFO                SCRIPTFLAGS = 1
	SCRIPTFLAGS_SHORTCUTS                SCRIPTFLAGS = 4
	SCRIPTFLAGS_MACHINEASSIGN            SCRIPTFLAGS = 8
	SCRIPTFLAGS_REGDATA_CNFGINFO         SCRIPTFLAGS = 32
	SCRIPTFLAGS_VALIDATE_TRANSFORMS_LIST SCRIPTFLAGS = 64
	SCRIPTFLAGS_REGDATA_CLASSINFO        SCRIPTFLAGS = 128
	SCRIPTFLAGS_REGDATA_EXTENSIONINFO    SCRIPTFLAGS = 256
	SCRIPTFLAGS_REGDATA_APPINFO          SCRIPTFLAGS = 384
	SCRIPTFLAGS_REGDATA                  SCRIPTFLAGS = 416
)

type ADVERTISEFLAGS int32

const (
	ADVERTISEFLAGS_MACHINEASSIGN ADVERTISEFLAGS = 0
	ADVERTISEFLAGS_USERASSIGN    ADVERTISEFLAGS = 1
)

type INSTALLTYPE int32

const (
	INSTALLTYPE_DEFAULT         INSTALLTYPE = 0
	INSTALLTYPE_NETWORK_IMAGE   INSTALLTYPE = 1
	INSTALLTYPE_SINGLE_INSTANCE INSTALLTYPE = 2
)

type MSIARCHITECTUREFLAGS int32

const (
	MSIARCHITECTUREFLAGS_X86   MSIARCHITECTUREFLAGS = 1
	MSIARCHITECTUREFLAGS_IA64  MSIARCHITECTUREFLAGS = 2
	MSIARCHITECTUREFLAGS_AMD64 MSIARCHITECTUREFLAGS = 4
	MSIARCHITECTUREFLAGS_ARM   MSIARCHITECTUREFLAGS = 8
)

type MSIOPENPACKAGEFLAGS int32

const (
	MSIOPENPACKAGEFLAGS_IGNOREMACHINESTATE MSIOPENPACKAGEFLAGS = 1
)

type MSIADVERTISEOPTIONFLAGS int32

const (
	MSIADVERTISEOPTIONFLAGS_INSTANCE MSIADVERTISEOPTIONFLAGS = 1
)

type MSISOURCETYPE int32

const (
	MSISOURCETYPE_UNKNOWN MSISOURCETYPE = 0
	MSISOURCETYPE_NETWORK MSISOURCETYPE = 1
	MSISOURCETYPE_URL     MSISOURCETYPE = 2
	MSISOURCETYPE_MEDIA   MSISOURCETYPE = 4
)

type MSICODE int32

const (
	MSICODE_PRODUCT MSICODE = 0
	MSICODE_PATCH   MSICODE = 1073741824
)

type MSITRANSACTION int32

const (
	MSITRANSACTION_CHAIN_EMBEDDEDUI         MSITRANSACTION = 1
	MSITRANSACTION_JOIN_EXISTING_EMBEDDEDUI MSITRANSACTION = 2
)

type MSITRANSACTIONSTATE int32

const (
	MSITRANSACTIONSTATE_ROLLBACK MSITRANSACTIONSTATE = 0
	MSITRANSACTIONSTATE_COMMIT   MSITRANSACTIONSTATE = 1
)

type MSIDBSTATE int32

const (
	MSIDBSTATE_ERROR MSIDBSTATE = -1
	MSIDBSTATE_READ  MSIDBSTATE = 0
	MSIDBSTATE_WRITE MSIDBSTATE = 1
)

type MSIMODIFY int32

const (
	MSIMODIFY_SEEK             MSIMODIFY = -1
	MSIMODIFY_REFRESH          MSIMODIFY = 0
	MSIMODIFY_INSERT           MSIMODIFY = 1
	MSIMODIFY_UPDATE           MSIMODIFY = 2
	MSIMODIFY_ASSIGN           MSIMODIFY = 3
	MSIMODIFY_REPLACE          MSIMODIFY = 4
	MSIMODIFY_MERGE            MSIMODIFY = 5
	MSIMODIFY_DELETE           MSIMODIFY = 6
	MSIMODIFY_INSERT_TEMPORARY MSIMODIFY = 7
	MSIMODIFY_VALIDATE         MSIMODIFY = 8
	MSIMODIFY_VALIDATE_NEW     MSIMODIFY = 9
	MSIMODIFY_VALIDATE_FIELD   MSIMODIFY = 10
	MSIMODIFY_VALIDATE_DELETE  MSIMODIFY = 11
)

type MSICOLINFO int32

const (
	MSICOLINFO_NAMES MSICOLINFO = 0
	MSICOLINFO_TYPES MSICOLINFO = 1
)

type MSICONDITION int32

const (
	MSICONDITION_FALSE MSICONDITION = 0
	MSICONDITION_TRUE  MSICONDITION = 1
	MSICONDITION_NONE  MSICONDITION = 2
	MSICONDITION_ERROR MSICONDITION = 3
)

type MSICOSTTREE int32

const (
	MSICOSTTREE_SELFONLY MSICOSTTREE = 0
	MSICOSTTREE_CHILDREN MSICOSTTREE = 1
	MSICOSTTREE_PARENTS  MSICOSTTREE = 2
	MSICOSTTREE_RESERVED MSICOSTTREE = 3
)

type MSIDBERROR int32

const (
	MSIDBERROR_INVALIDARG        MSIDBERROR = -3
	MSIDBERROR_MOREDATA          MSIDBERROR = -2
	MSIDBERROR_FUNCTIONERROR     MSIDBERROR = -1
	MSIDBERROR_NOERROR           MSIDBERROR = 0
	MSIDBERROR_DUPLICATEKEY      MSIDBERROR = 1
	MSIDBERROR_REQUIRED          MSIDBERROR = 2
	MSIDBERROR_BADLINK           MSIDBERROR = 3
	MSIDBERROR_OVERFLOW          MSIDBERROR = 4
	MSIDBERROR_UNDERFLOW         MSIDBERROR = 5
	MSIDBERROR_NOTINSET          MSIDBERROR = 6
	MSIDBERROR_BADVERSION        MSIDBERROR = 7
	MSIDBERROR_BADCASE           MSIDBERROR = 8
	MSIDBERROR_BADGUID           MSIDBERROR = 9
	MSIDBERROR_BADWILDCARD       MSIDBERROR = 10
	MSIDBERROR_BADIDENTIFIER     MSIDBERROR = 11
	MSIDBERROR_BADLANGUAGE       MSIDBERROR = 12
	MSIDBERROR_BADFILENAME       MSIDBERROR = 13
	MSIDBERROR_BADPATH           MSIDBERROR = 14
	MSIDBERROR_BADCONDITION      MSIDBERROR = 15
	MSIDBERROR_BADFORMATTED      MSIDBERROR = 16
	MSIDBERROR_BADTEMPLATE       MSIDBERROR = 17
	MSIDBERROR_BADDEFAULTDIR     MSIDBERROR = 18
	MSIDBERROR_BADREGPATH        MSIDBERROR = 19
	MSIDBERROR_BADCUSTOMSOURCE   MSIDBERROR = 20
	MSIDBERROR_BADPROPERTY       MSIDBERROR = 21
	MSIDBERROR_MISSINGDATA       MSIDBERROR = 22
	MSIDBERROR_BADCATEGORY       MSIDBERROR = 23
	MSIDBERROR_BADKEYTABLE       MSIDBERROR = 24
	MSIDBERROR_BADMAXMINVALUES   MSIDBERROR = 25
	MSIDBERROR_BADCABINET        MSIDBERROR = 26
	MSIDBERROR_BADSHORTCUT       MSIDBERROR = 27
	MSIDBERROR_STRINGOVERFLOW    MSIDBERROR = 28
	MSIDBERROR_BADLOCALIZEATTRIB MSIDBERROR = 29
)

type MSIRUNMODE int32

const (
	MSIRUNMODE_ADMIN            MSIRUNMODE = 0
	MSIRUNMODE_ADVERTISE        MSIRUNMODE = 1
	MSIRUNMODE_MAINTENANCE      MSIRUNMODE = 2
	MSIRUNMODE_ROLLBACKENABLED  MSIRUNMODE = 3
	MSIRUNMODE_LOGENABLED       MSIRUNMODE = 4
	MSIRUNMODE_OPERATIONS       MSIRUNMODE = 5
	MSIRUNMODE_REBOOTATEND      MSIRUNMODE = 6
	MSIRUNMODE_REBOOTNOW        MSIRUNMODE = 7
	MSIRUNMODE_CABINET          MSIRUNMODE = 8
	MSIRUNMODE_SOURCESHORTNAMES MSIRUNMODE = 9
	MSIRUNMODE_TARGETSHORTNAMES MSIRUNMODE = 10
	MSIRUNMODE_RESERVED11       MSIRUNMODE = 11
	MSIRUNMODE_WINDOWS9X        MSIRUNMODE = 12
	MSIRUNMODE_ZAWENABLED       MSIRUNMODE = 13
	MSIRUNMODE_RESERVED14       MSIRUNMODE = 14
	MSIRUNMODE_RESERVED15       MSIRUNMODE = 15
	MSIRUNMODE_SCHEDULED        MSIRUNMODE = 16
	MSIRUNMODE_ROLLBACK         MSIRUNMODE = 17
	MSIRUNMODE_COMMIT           MSIRUNMODE = 18
)

type MSITRANSFORM_ERROR int32

const (
	MSITRANSFORM_ERROR_ADDEXISTINGROW   MSITRANSFORM_ERROR = 1
	MSITRANSFORM_ERROR_DELMISSINGROW    MSITRANSFORM_ERROR = 2
	MSITRANSFORM_ERROR_ADDEXISTINGTABLE MSITRANSFORM_ERROR = 4
	MSITRANSFORM_ERROR_DELMISSINGTABLE  MSITRANSFORM_ERROR = 8
	MSITRANSFORM_ERROR_UPDATEMISSINGROW MSITRANSFORM_ERROR = 16
	MSITRANSFORM_ERROR_CHANGECODEPAGE   MSITRANSFORM_ERROR = 32
	MSITRANSFORM_ERROR_VIEWTRANSFORM    MSITRANSFORM_ERROR = 256
)

type MSITRANSFORM_VALIDATE int32

const (
	MSITRANSFORM_VALIDATE_LANGUAGE                   MSITRANSFORM_VALIDATE = 1
	MSITRANSFORM_VALIDATE_PRODUCT                    MSITRANSFORM_VALIDATE = 2
	MSITRANSFORM_VALIDATE_PLATFORM                   MSITRANSFORM_VALIDATE = 4
	MSITRANSFORM_VALIDATE_MAJORVERSION               MSITRANSFORM_VALIDATE = 8
	MSITRANSFORM_VALIDATE_MINORVERSION               MSITRANSFORM_VALIDATE = 16
	MSITRANSFORM_VALIDATE_UPDATEVERSION              MSITRANSFORM_VALIDATE = 32
	MSITRANSFORM_VALIDATE_NEWLESSBASEVERSION         MSITRANSFORM_VALIDATE = 64
	MSITRANSFORM_VALIDATE_NEWLESSEQUALBASEVERSION    MSITRANSFORM_VALIDATE = 128
	MSITRANSFORM_VALIDATE_NEWEQUALBASEVERSION        MSITRANSFORM_VALIDATE = 256
	MSITRANSFORM_VALIDATE_NEWGREATEREQUALBASEVERSION MSITRANSFORM_VALIDATE = 512
	MSITRANSFORM_VALIDATE_NEWGREATERBASEVERSION      MSITRANSFORM_VALIDATE = 1024
	MSITRANSFORM_VALIDATE_UPGRADECODE                MSITRANSFORM_VALIDATE = 2048
)

type ASM_NAME int32

const (
	ASM_NAME_PUBLIC_KEY            ASM_NAME = 0
	ASM_NAME_PUBLIC_KEY_TOKEN      ASM_NAME = 1
	ASM_NAME_HASH_VALUE            ASM_NAME = 2
	ASM_NAME_NAME                  ASM_NAME = 3
	ASM_NAME_MAJOR_VERSION         ASM_NAME = 4
	ASM_NAME_MINOR_VERSION         ASM_NAME = 5
	ASM_NAME_BUILD_NUMBER          ASM_NAME = 6
	ASM_NAME_REVISION_NUMBER       ASM_NAME = 7
	ASM_NAME_CULTURE               ASM_NAME = 8
	ASM_NAME_PROCESSOR_ID_ARRAY    ASM_NAME = 9
	ASM_NAME_OSINFO_ARRAY          ASM_NAME = 10
	ASM_NAME_HASH_ALGID            ASM_NAME = 11
	ASM_NAME_ALIAS                 ASM_NAME = 12
	ASM_NAME_CODEBASE_URL          ASM_NAME = 13
	ASM_NAME_CODEBASE_LASTMOD      ASM_NAME = 14
	ASM_NAME_NULL_PUBLIC_KEY       ASM_NAME = 15
	ASM_NAME_NULL_PUBLIC_KEY_TOKEN ASM_NAME = 16
	ASM_NAME_CUSTOM                ASM_NAME = 17
	ASM_NAME_NULL_CUSTOM           ASM_NAME = 18
	ASM_NAME_MVID                  ASM_NAME = 19
	ASM_NAME_MAX_PARAMS            ASM_NAME = 20
)

type MIDL_IAssemblyName_0002 int32

const (
	ASM_BINDF_FORCE_CACHE_INSTALL MIDL_IAssemblyName_0002 = 1
	ASM_BINDF_RFS_INTEGRITY_CHECK MIDL_IAssemblyName_0002 = 2
	ASM_BINDF_RFS_MODULE_CHECK    MIDL_IAssemblyName_0002 = 4
	ASM_BINDF_BINPATH_PROBE_ONLY  MIDL_IAssemblyName_0002 = 8
	ASM_BINDF_SHARED_BINPATH_HINT MIDL_IAssemblyName_0002 = 16
	ASM_BINDF_PARENT_ASM_HINT     MIDL_IAssemblyName_0002 = 32
)

type ASM_DISPLAY_FLAGS int32

const (
	ASM_DISPLAYF_VERSION               ASM_DISPLAY_FLAGS = 1
	ASM_DISPLAYF_CULTURE               ASM_DISPLAY_FLAGS = 2
	ASM_DISPLAYF_PUBLIC_KEY_TOKEN      ASM_DISPLAY_FLAGS = 4
	ASM_DISPLAYF_PUBLIC_KEY            ASM_DISPLAY_FLAGS = 8
	ASM_DISPLAYF_CUSTOM                ASM_DISPLAY_FLAGS = 16
	ASM_DISPLAYF_PROCESSORARCHITECTURE ASM_DISPLAY_FLAGS = 32
	ASM_DISPLAYF_LANGUAGEID            ASM_DISPLAY_FLAGS = 64
)

type ASM_CMP_FLAGS int32

const (
	ASM_CMPF_NAME             ASM_CMP_FLAGS = 1
	ASM_CMPF_MAJOR_VERSION    ASM_CMP_FLAGS = 2
	ASM_CMPF_MINOR_VERSION    ASM_CMP_FLAGS = 4
	ASM_CMPF_BUILD_NUMBER     ASM_CMP_FLAGS = 8
	ASM_CMPF_REVISION_NUMBER  ASM_CMP_FLAGS = 16
	ASM_CMPF_PUBLIC_KEY_TOKEN ASM_CMP_FLAGS = 32
	ASM_CMPF_CULTURE          ASM_CMP_FLAGS = 64
	ASM_CMPF_CUSTOM           ASM_CMP_FLAGS = 128
	ASM_CMPF_ALL              ASM_CMP_FLAGS = 255
	ASM_CMPF_DEFAULT          ASM_CMP_FLAGS = 256
)

type CREATE_ASM_NAME_OBJ_FLAGS int32

const (
	CANOF_PARSE_DISPLAY_NAME CREATE_ASM_NAME_OBJ_FLAGS = 1
	CANOF_SET_DEFAULT_VALUES CREATE_ASM_NAME_OBJ_FLAGS = 2
)

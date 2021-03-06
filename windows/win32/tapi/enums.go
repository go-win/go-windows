// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package tapi implements the Windows.Win32.Tapi namespace.
package tapi

type TAPI_TONEMODE int32

const (
	TTM_RINGBACK TAPI_TONEMODE = 2
	TTM_BUSY     TAPI_TONEMODE = 4
	TTM_BEEP     TAPI_TONEMODE = 8
	TTM_BILLING  TAPI_TONEMODE = 16
)

type TAPI_GATHERTERM int32

const (
	TGT_BUFFERFULL   TAPI_GATHERTERM = 1
	TGT_TERMDIGIT    TAPI_GATHERTERM = 2
	TGT_FIRSTTIMEOUT TAPI_GATHERTERM = 4
	TGT_INTERTIMEOUT TAPI_GATHERTERM = 8
	TGT_CANCEL       TAPI_GATHERTERM = 16
)

type ADDRESS_EVENT int32

const (
	AE_STATE          ADDRESS_EVENT = 0
	AE_CAPSCHANGE     ADDRESS_EVENT = 1
	AE_RINGING        ADDRESS_EVENT = 2
	AE_CONFIGCHANGE   ADDRESS_EVENT = 3
	AE_FORWARD        ADDRESS_EVENT = 4
	AE_NEWTERMINAL    ADDRESS_EVENT = 5
	AE_REMOVETERMINAL ADDRESS_EVENT = 6
	AE_MSGWAITON      ADDRESS_EVENT = 7
	AE_MSGWAITOFF     ADDRESS_EVENT = 8
	AE_LASTITEM       ADDRESS_EVENT = 8
)

type ADDRESS_STATE int32

const (
	AS_INSERVICE    ADDRESS_STATE = 0
	AS_OUTOFSERVICE ADDRESS_STATE = 1
)

type CALL_STATE int32

const (
	CS_IDLE         CALL_STATE = 0
	CS_INPROGRESS   CALL_STATE = 1
	CS_CONNECTED    CALL_STATE = 2
	CS_DISCONNECTED CALL_STATE = 3
	CS_OFFERING     CALL_STATE = 4
	CS_HOLD         CALL_STATE = 5
	CS_QUEUED       CALL_STATE = 6
	CS_LASTITEM     CALL_STATE = 6
)

type CALL_STATE_EVENT_CAUSE int32

const (
	CEC_NONE                  CALL_STATE_EVENT_CAUSE = 0
	CEC_DISCONNECT_NORMAL     CALL_STATE_EVENT_CAUSE = 1
	CEC_DISCONNECT_BUSY       CALL_STATE_EVENT_CAUSE = 2
	CEC_DISCONNECT_BADADDRESS CALL_STATE_EVENT_CAUSE = 3
	CEC_DISCONNECT_NOANSWER   CALL_STATE_EVENT_CAUSE = 4
	CEC_DISCONNECT_CANCELLED  CALL_STATE_EVENT_CAUSE = 5
	CEC_DISCONNECT_REJECTED   CALL_STATE_EVENT_CAUSE = 6
	CEC_DISCONNECT_FAILED     CALL_STATE_EVENT_CAUSE = 7
	CEC_DISCONNECT_BLOCKED    CALL_STATE_EVENT_CAUSE = 8
)

type CALL_MEDIA_EVENT int32

const (
	CME_NEW_STREAM      CALL_MEDIA_EVENT = 0
	CME_STREAM_FAIL     CALL_MEDIA_EVENT = 1
	CME_TERMINAL_FAIL   CALL_MEDIA_EVENT = 2
	CME_STREAM_NOT_USED CALL_MEDIA_EVENT = 3
	CME_STREAM_ACTIVE   CALL_MEDIA_EVENT = 4
	CME_STREAM_INACTIVE CALL_MEDIA_EVENT = 5
	CME_LASTITEM        CALL_MEDIA_EVENT = 5
)

type CALL_MEDIA_EVENT_CAUSE int32

const (
	CMC_UNKNOWN            CALL_MEDIA_EVENT_CAUSE = 0
	CMC_BAD_DEVICE         CALL_MEDIA_EVENT_CAUSE = 1
	CMC_CONNECT_FAIL       CALL_MEDIA_EVENT_CAUSE = 2
	CMC_LOCAL_REQUEST      CALL_MEDIA_EVENT_CAUSE = 3
	CMC_REMOTE_REQUEST     CALL_MEDIA_EVENT_CAUSE = 4
	CMC_MEDIA_TIMEOUT      CALL_MEDIA_EVENT_CAUSE = 5
	CMC_MEDIA_RECOVERED    CALL_MEDIA_EVENT_CAUSE = 6
	CMC_QUALITY_OF_SERVICE CALL_MEDIA_EVENT_CAUSE = 7
)

type DISCONNECT_CODE int32

const (
	DC_NORMAL   DISCONNECT_CODE = 0
	DC_NOANSWER DISCONNECT_CODE = 1
	DC_REJECTED DISCONNECT_CODE = 2
)

type TERMINAL_STATE int32

const (
	TS_INUSE    TERMINAL_STATE = 0
	TS_NOTINUSE TERMINAL_STATE = 1
)

type TERMINAL_DIRECTION int32

const (
	TD_CAPTURE          TERMINAL_DIRECTION = 0
	TD_RENDER           TERMINAL_DIRECTION = 1
	TD_BIDIRECTIONAL    TERMINAL_DIRECTION = 2
	TD_MULTITRACK_MIXED TERMINAL_DIRECTION = 3
	TD_NONE             TERMINAL_DIRECTION = 4
)

type TERMINAL_TYPE int32

const (
	TT_STATIC  TERMINAL_TYPE = 0
	TT_DYNAMIC TERMINAL_TYPE = 1
)

type CALL_PRIVILEGE int32

const (
	CP_OWNER   CALL_PRIVILEGE = 0
	CP_MONITOR CALL_PRIVILEGE = 1
)

type TAPI_EVENT int32

const (
	TE_TAPIOBJECT         TAPI_EVENT = 1
	TE_ADDRESS            TAPI_EVENT = 2
	TE_CALLNOTIFICATION   TAPI_EVENT = 4
	TE_CALLSTATE          TAPI_EVENT = 8
	TE_CALLMEDIA          TAPI_EVENT = 16
	TE_CALLHUB            TAPI_EVENT = 32
	TE_CALLINFOCHANGE     TAPI_EVENT = 64
	TE_PRIVATE            TAPI_EVENT = 128
	TE_REQUEST            TAPI_EVENT = 256
	TE_AGENT              TAPI_EVENT = 512
	TE_AGENTSESSION       TAPI_EVENT = 1024
	TE_QOSEVENT           TAPI_EVENT = 2048
	TE_AGENTHANDLER       TAPI_EVENT = 4096
	TE_ACDGROUP           TAPI_EVENT = 8192
	TE_QUEUE              TAPI_EVENT = 16384
	TE_DIGITEVENT         TAPI_EVENT = 32768
	TE_GENERATEEVENT      TAPI_EVENT = 65536
	TE_ASRTERMINAL        TAPI_EVENT = 131072
	TE_TTSTERMINAL        TAPI_EVENT = 262144
	TE_FILETERMINAL       TAPI_EVENT = 524288
	TE_TONETERMINAL       TAPI_EVENT = 1048576
	TE_PHONEEVENT         TAPI_EVENT = 2097152
	TE_TONEEVENT          TAPI_EVENT = 4194304
	TE_GATHERDIGITS       TAPI_EVENT = 8388608
	TE_ADDRESSDEVSPECIFIC TAPI_EVENT = 16777216
	TE_PHONEDEVSPECIFIC   TAPI_EVENT = 33554432
)

type CALL_NOTIFICATION_EVENT int32

const (
	CNE_OWNER    CALL_NOTIFICATION_EVENT = 0
	CNE_MONITOR  CALL_NOTIFICATION_EVENT = 1
	CNE_LASTITEM CALL_NOTIFICATION_EVENT = 1
)

type CALLHUB_EVENT int32

const (
	CHE_CALLJOIN    CALLHUB_EVENT = 0
	CHE_CALLLEAVE   CALLHUB_EVENT = 1
	CHE_CALLHUBNEW  CALLHUB_EVENT = 2
	CHE_CALLHUBIDLE CALLHUB_EVENT = 3
	CHE_LASTITEM    CALLHUB_EVENT = 3
)

type CALLHUB_STATE int32

const (
	CHS_ACTIVE CALLHUB_STATE = 0
	CHS_IDLE   CALLHUB_STATE = 1
)

type TAPIOBJECT_EVENT int32

const (
	TE_ADDRESSCREATE   TAPIOBJECT_EVENT = 0
	TE_ADDRESSREMOVE   TAPIOBJECT_EVENT = 1
	TE_REINIT          TAPIOBJECT_EVENT = 2
	TE_TRANSLATECHANGE TAPIOBJECT_EVENT = 3
	TE_ADDRESSCLOSE    TAPIOBJECT_EVENT = 4
	TE_PHONECREATE     TAPIOBJECT_EVENT = 5
	TE_PHONEREMOVE     TAPIOBJECT_EVENT = 6
)

type TAPI_OBJECT_TYPE int32

const (
	TOT_NONE     TAPI_OBJECT_TYPE = 0
	TOT_TAPI     TAPI_OBJECT_TYPE = 1
	TOT_ADDRESS  TAPI_OBJECT_TYPE = 2
	TOT_TERMINAL TAPI_OBJECT_TYPE = 3
	TOT_CALL     TAPI_OBJECT_TYPE = 4
	TOT_CALLHUB  TAPI_OBJECT_TYPE = 5
	TOT_PHONE    TAPI_OBJECT_TYPE = 6
)

type QOS_SERVICE_LEVEL int32

const (
	QSL_NEEDED       QOS_SERVICE_LEVEL = 1
	QSL_IF_AVAILABLE QOS_SERVICE_LEVEL = 2
	QSL_BEST_EFFORT  QOS_SERVICE_LEVEL = 3
)

type QOS_EVENT int32

const (
	QE_NOQOS            QOS_EVENT = 1
	QE_ADMISSIONFAILURE QOS_EVENT = 2
	QE_POLICYFAILURE    QOS_EVENT = 3
	QE_GENERICERROR     QOS_EVENT = 4
	QE_LASTITEM         QOS_EVENT = 4
)

type CALLINFOCHANGE_CAUSE int32

const (
	CIC_OTHER         CALLINFOCHANGE_CAUSE = 0
	CIC_DEVSPECIFIC   CALLINFOCHANGE_CAUSE = 1
	CIC_BEARERMODE    CALLINFOCHANGE_CAUSE = 2
	CIC_RATE          CALLINFOCHANGE_CAUSE = 3
	CIC_APPSPECIFIC   CALLINFOCHANGE_CAUSE = 4
	CIC_CALLID        CALLINFOCHANGE_CAUSE = 5
	CIC_RELATEDCALLID CALLINFOCHANGE_CAUSE = 6
	CIC_ORIGIN        CALLINFOCHANGE_CAUSE = 7
	CIC_REASON        CALLINFOCHANGE_CAUSE = 8
	CIC_COMPLETIONID  CALLINFOCHANGE_CAUSE = 9
	CIC_NUMOWNERINCR  CALLINFOCHANGE_CAUSE = 10
	CIC_NUMOWNERDECR  CALLINFOCHANGE_CAUSE = 11
	CIC_NUMMONITORS   CALLINFOCHANGE_CAUSE = 12
	CIC_TRUNK         CALLINFOCHANGE_CAUSE = 13
	CIC_CALLERID      CALLINFOCHANGE_CAUSE = 14
	CIC_CALLEDID      CALLINFOCHANGE_CAUSE = 15
	CIC_CONNECTEDID   CALLINFOCHANGE_CAUSE = 16
	CIC_REDIRECTIONID CALLINFOCHANGE_CAUSE = 17
	CIC_REDIRECTINGID CALLINFOCHANGE_CAUSE = 18
	CIC_USERUSERINFO  CALLINFOCHANGE_CAUSE = 19
	CIC_HIGHLEVELCOMP CALLINFOCHANGE_CAUSE = 20
	CIC_LOWLEVELCOMP  CALLINFOCHANGE_CAUSE = 21
	CIC_CHARGINGINFO  CALLINFOCHANGE_CAUSE = 22
	CIC_TREATMENT     CALLINFOCHANGE_CAUSE = 23
	CIC_CALLDATA      CALLINFOCHANGE_CAUSE = 24
	CIC_PRIVILEGE     CALLINFOCHANGE_CAUSE = 25
	CIC_MEDIATYPE     CALLINFOCHANGE_CAUSE = 26
	CIC_LASTITEM      CALLINFOCHANGE_CAUSE = 26
)

type CALLINFO_LONG int32

const (
	CIL_MEDIATYPESAVAILABLE      CALLINFO_LONG = 0
	CIL_BEARERMODE               CALLINFO_LONG = 1
	CIL_CALLERIDADDRESSTYPE      CALLINFO_LONG = 2
	CIL_CALLEDIDADDRESSTYPE      CALLINFO_LONG = 3
	CIL_CONNECTEDIDADDRESSTYPE   CALLINFO_LONG = 4
	CIL_REDIRECTIONIDADDRESSTYPE CALLINFO_LONG = 5
	CIL_REDIRECTINGIDADDRESSTYPE CALLINFO_LONG = 6
	CIL_ORIGIN                   CALLINFO_LONG = 7
	CIL_REASON                   CALLINFO_LONG = 8
	CIL_APPSPECIFIC              CALLINFO_LONG = 9
	CIL_CALLPARAMSFLAGS          CALLINFO_LONG = 10
	CIL_CALLTREATMENT            CALLINFO_LONG = 11
	CIL_MINRATE                  CALLINFO_LONG = 12
	CIL_MAXRATE                  CALLINFO_LONG = 13
	CIL_COUNTRYCODE              CALLINFO_LONG = 14
	CIL_CALLID                   CALLINFO_LONG = 15
	CIL_RELATEDCALLID            CALLINFO_LONG = 16
	CIL_COMPLETIONID             CALLINFO_LONG = 17
	CIL_NUMBEROFOWNERS           CALLINFO_LONG = 18
	CIL_NUMBEROFMONITORS         CALLINFO_LONG = 19
	CIL_TRUNK                    CALLINFO_LONG = 20
	CIL_RATE                     CALLINFO_LONG = 21
	CIL_GENERATEDIGITDURATION    CALLINFO_LONG = 22
	CIL_MONITORDIGITMODES        CALLINFO_LONG = 23
	CIL_MONITORMEDIAMODES        CALLINFO_LONG = 24
)

type CALLINFO_STRING int32

const (
	CIS_CALLERIDNAME            CALLINFO_STRING = 0
	CIS_CALLERIDNUMBER          CALLINFO_STRING = 1
	CIS_CALLEDIDNAME            CALLINFO_STRING = 2
	CIS_CALLEDIDNUMBER          CALLINFO_STRING = 3
	CIS_CONNECTEDIDNAME         CALLINFO_STRING = 4
	CIS_CONNECTEDIDNUMBER       CALLINFO_STRING = 5
	CIS_REDIRECTIONIDNAME       CALLINFO_STRING = 6
	CIS_REDIRECTIONIDNUMBER     CALLINFO_STRING = 7
	CIS_REDIRECTINGIDNAME       CALLINFO_STRING = 8
	CIS_REDIRECTINGIDNUMBER     CALLINFO_STRING = 9
	CIS_CALLEDPARTYFRIENDLYNAME CALLINFO_STRING = 10
	CIS_COMMENT                 CALLINFO_STRING = 11
	CIS_DISPLAYABLEADDRESS      CALLINFO_STRING = 12
	CIS_CALLINGPARTYID          CALLINFO_STRING = 13
)

type CALLINFO_BUFFER int32

const (
	CIB_USERUSERINFO                 CALLINFO_BUFFER = 0
	CIB_DEVSPECIFICBUFFER            CALLINFO_BUFFER = 1
	CIB_CALLDATABUFFER               CALLINFO_BUFFER = 2
	CIB_CHARGINGINFOBUFFER           CALLINFO_BUFFER = 3
	CIB_HIGHLEVELCOMPATIBILITYBUFFER CALLINFO_BUFFER = 4
	CIB_LOWLEVELCOMPATIBILITYBUFFER  CALLINFO_BUFFER = 5
)

type ADDRESS_CAPABILITY int32

const (
	AC_ADDRESSTYPES                 ADDRESS_CAPABILITY = 0
	AC_BEARERMODES                  ADDRESS_CAPABILITY = 1
	AC_MAXACTIVECALLS               ADDRESS_CAPABILITY = 2
	AC_MAXONHOLDCALLS               ADDRESS_CAPABILITY = 3
	AC_MAXONHOLDPENDINGCALLS        ADDRESS_CAPABILITY = 4
	AC_MAXNUMCONFERENCE             ADDRESS_CAPABILITY = 5
	AC_MAXNUMTRANSCONF              ADDRESS_CAPABILITY = 6
	AC_MONITORDIGITSUPPORT          ADDRESS_CAPABILITY = 7
	AC_GENERATEDIGITSUPPORT         ADDRESS_CAPABILITY = 8
	AC_GENERATETONEMODES            ADDRESS_CAPABILITY = 9
	AC_GENERATETONEMAXNUMFREQ       ADDRESS_CAPABILITY = 10
	AC_MONITORTONEMAXNUMFREQ        ADDRESS_CAPABILITY = 11
	AC_MONITORTONEMAXNUMENTRIES     ADDRESS_CAPABILITY = 12
	AC_DEVCAPFLAGS                  ADDRESS_CAPABILITY = 13
	AC_ANSWERMODES                  ADDRESS_CAPABILITY = 14
	AC_LINEFEATURES                 ADDRESS_CAPABILITY = 15
	AC_SETTABLEDEVSTATUS            ADDRESS_CAPABILITY = 16
	AC_PARKSUPPORT                  ADDRESS_CAPABILITY = 17
	AC_CALLERIDSUPPORT              ADDRESS_CAPABILITY = 18
	AC_CALLEDIDSUPPORT              ADDRESS_CAPABILITY = 19
	AC_CONNECTEDIDSUPPORT           ADDRESS_CAPABILITY = 20
	AC_REDIRECTIONIDSUPPORT         ADDRESS_CAPABILITY = 21
	AC_REDIRECTINGIDSUPPORT         ADDRESS_CAPABILITY = 22
	AC_ADDRESSCAPFLAGS              ADDRESS_CAPABILITY = 23
	AC_CALLFEATURES1                ADDRESS_CAPABILITY = 24
	AC_CALLFEATURES2                ADDRESS_CAPABILITY = 25
	AC_REMOVEFROMCONFCAPS           ADDRESS_CAPABILITY = 26
	AC_REMOVEFROMCONFSTATE          ADDRESS_CAPABILITY = 27
	AC_TRANSFERMODES                ADDRESS_CAPABILITY = 28
	AC_ADDRESSFEATURES              ADDRESS_CAPABILITY = 29
	AC_PREDICTIVEAUTOTRANSFERSTATES ADDRESS_CAPABILITY = 30
	AC_MAXCALLDATASIZE              ADDRESS_CAPABILITY = 31
	AC_LINEID                       ADDRESS_CAPABILITY = 32
	AC_ADDRESSID                    ADDRESS_CAPABILITY = 33
	AC_FORWARDMODES                 ADDRESS_CAPABILITY = 34
	AC_MAXFORWARDENTRIES            ADDRESS_CAPABILITY = 35
	AC_MAXSPECIFICENTRIES           ADDRESS_CAPABILITY = 36
	AC_MINFWDNUMRINGS               ADDRESS_CAPABILITY = 37
	AC_MAXFWDNUMRINGS               ADDRESS_CAPABILITY = 38
	AC_MAXCALLCOMPLETIONS           ADDRESS_CAPABILITY = 39
	AC_CALLCOMPLETIONCONDITIONS     ADDRESS_CAPABILITY = 40
	AC_CALLCOMPLETIONMODES          ADDRESS_CAPABILITY = 41
	AC_PERMANENTDEVICEID            ADDRESS_CAPABILITY = 42
	AC_GATHERDIGITSMINTIMEOUT       ADDRESS_CAPABILITY = 43
	AC_GATHERDIGITSMAXTIMEOUT       ADDRESS_CAPABILITY = 44
	AC_GENERATEDIGITMINDURATION     ADDRESS_CAPABILITY = 45
	AC_GENERATEDIGITMAXDURATION     ADDRESS_CAPABILITY = 46
	AC_GENERATEDIGITDEFAULTDURATION ADDRESS_CAPABILITY = 47
)

type ADDRESS_CAPABILITY_STRING int32

const (
	ACS_PROTOCOL              ADDRESS_CAPABILITY_STRING = 0
	ACS_ADDRESSDEVICESPECIFIC ADDRESS_CAPABILITY_STRING = 1
	ACS_LINEDEVICESPECIFIC    ADDRESS_CAPABILITY_STRING = 2
	ACS_PROVIDERSPECIFIC      ADDRESS_CAPABILITY_STRING = 3
	ACS_SWITCHSPECIFIC        ADDRESS_CAPABILITY_STRING = 4
	ACS_PERMANENTDEVICEGUID   ADDRESS_CAPABILITY_STRING = 5
)

type FULLDUPLEX_SUPPORT int32

const (
	FDS_SUPPORTED    FULLDUPLEX_SUPPORT = 0
	FDS_NOTSUPPORTED FULLDUPLEX_SUPPORT = 1
	FDS_UNKNOWN      FULLDUPLEX_SUPPORT = 2
)

type FINISH_MODE int32

const (
	FM_ASTRANSFER   FINISH_MODE = 0
	FM_ASCONFERENCE FINISH_MODE = 1
)

type PHONE_PRIVILEGE int32

const (
	PP_OWNER   PHONE_PRIVILEGE = 0
	PP_MONITOR PHONE_PRIVILEGE = 1
)

type PHONE_HOOK_SWITCH_DEVICE int32

const (
	PHSD_HANDSET      PHONE_HOOK_SWITCH_DEVICE = 1
	PHSD_SPEAKERPHONE PHONE_HOOK_SWITCH_DEVICE = 2
	PHSD_HEADSET      PHONE_HOOK_SWITCH_DEVICE = 4
)

type PHONE_HOOK_SWITCH_STATE int32

const (
	PHSS_ONHOOK               PHONE_HOOK_SWITCH_STATE = 1
	PHSS_OFFHOOK_MIC_ONLY     PHONE_HOOK_SWITCH_STATE = 2
	PHSS_OFFHOOK_SPEAKER_ONLY PHONE_HOOK_SWITCH_STATE = 4
	PHSS_OFFHOOK              PHONE_HOOK_SWITCH_STATE = 8
)

type PHONE_LAMP_MODE int32

const (
	LM_DUMMY         PHONE_LAMP_MODE = 1
	LM_OFF           PHONE_LAMP_MODE = 2
	LM_STEADY        PHONE_LAMP_MODE = 4
	LM_WINK          PHONE_LAMP_MODE = 8
	LM_FLASH         PHONE_LAMP_MODE = 16
	LM_FLUTTER       PHONE_LAMP_MODE = 32
	LM_BROKENFLUTTER PHONE_LAMP_MODE = 64
	LM_UNKNOWN       PHONE_LAMP_MODE = 128
)

type PHONECAPS_LONG int32

const (
	PCL_HOOKSWITCHES                PHONECAPS_LONG = 0
	PCL_HANDSETHOOKSWITCHMODES      PHONECAPS_LONG = 1
	PCL_HEADSETHOOKSWITCHMODES      PHONECAPS_LONG = 2
	PCL_SPEAKERPHONEHOOKSWITCHMODES PHONECAPS_LONG = 3
	PCL_DISPLAYNUMROWS              PHONECAPS_LONG = 4
	PCL_DISPLAYNUMCOLUMNS           PHONECAPS_LONG = 5
	PCL_NUMRINGMODES                PHONECAPS_LONG = 6
	PCL_NUMBUTTONLAMPS              PHONECAPS_LONG = 7
	PCL_GENERICPHONE                PHONECAPS_LONG = 8
)

type PHONECAPS_STRING int32

const (
	PCS_PHONENAME    PHONECAPS_STRING = 0
	PCS_PHONEINFO    PHONECAPS_STRING = 1
	PCS_PROVIDERINFO PHONECAPS_STRING = 2
)

type PHONECAPS_BUFFER int32

const (
	PCB_DEVSPECIFICBUFFER PHONECAPS_BUFFER = 0
)

type PHONE_BUTTON_STATE int32

const (
	PBS_UP      PHONE_BUTTON_STATE = 1
	PBS_DOWN    PHONE_BUTTON_STATE = 2
	PBS_UNKNOWN PHONE_BUTTON_STATE = 4
	PBS_UNAVAIL PHONE_BUTTON_STATE = 8
)

type PHONE_BUTTON_MODE int32

const (
	PBM_DUMMY   PHONE_BUTTON_MODE = 0
	PBM_CALL    PHONE_BUTTON_MODE = 1
	PBM_FEATURE PHONE_BUTTON_MODE = 2
	PBM_KEYPAD  PHONE_BUTTON_MODE = 3
	PBM_LOCAL   PHONE_BUTTON_MODE = 4
	PBM_DISPLAY PHONE_BUTTON_MODE = 5
)

type PHONE_BUTTON_FUNCTION int32

const (
	PBF_UNKNOWN      PHONE_BUTTON_FUNCTION = 0
	PBF_CONFERENCE   PHONE_BUTTON_FUNCTION = 1
	PBF_TRANSFER     PHONE_BUTTON_FUNCTION = 2
	PBF_DROP         PHONE_BUTTON_FUNCTION = 3
	PBF_HOLD         PHONE_BUTTON_FUNCTION = 4
	PBF_RECALL       PHONE_BUTTON_FUNCTION = 5
	PBF_DISCONNECT   PHONE_BUTTON_FUNCTION = 6
	PBF_CONNECT      PHONE_BUTTON_FUNCTION = 7
	PBF_MSGWAITON    PHONE_BUTTON_FUNCTION = 8
	PBF_MSGWAITOFF   PHONE_BUTTON_FUNCTION = 9
	PBF_SELECTRING   PHONE_BUTTON_FUNCTION = 10
	PBF_ABBREVDIAL   PHONE_BUTTON_FUNCTION = 11
	PBF_FORWARD      PHONE_BUTTON_FUNCTION = 12
	PBF_PICKUP       PHONE_BUTTON_FUNCTION = 13
	PBF_RINGAGAIN    PHONE_BUTTON_FUNCTION = 14
	PBF_PARK         PHONE_BUTTON_FUNCTION = 15
	PBF_REJECT       PHONE_BUTTON_FUNCTION = 16
	PBF_REDIRECT     PHONE_BUTTON_FUNCTION = 17
	PBF_MUTE         PHONE_BUTTON_FUNCTION = 18
	PBF_VOLUMEUP     PHONE_BUTTON_FUNCTION = 19
	PBF_VOLUMEDOWN   PHONE_BUTTON_FUNCTION = 20
	PBF_SPEAKERON    PHONE_BUTTON_FUNCTION = 21
	PBF_SPEAKEROFF   PHONE_BUTTON_FUNCTION = 22
	PBF_FLASH        PHONE_BUTTON_FUNCTION = 23
	PBF_DATAON       PHONE_BUTTON_FUNCTION = 24
	PBF_DATAOFF      PHONE_BUTTON_FUNCTION = 25
	PBF_DONOTDISTURB PHONE_BUTTON_FUNCTION = 26
	PBF_INTERCOM     PHONE_BUTTON_FUNCTION = 27
	PBF_BRIDGEDAPP   PHONE_BUTTON_FUNCTION = 28
	PBF_BUSY         PHONE_BUTTON_FUNCTION = 29
	PBF_CALLAPP      PHONE_BUTTON_FUNCTION = 30
	PBF_DATETIME     PHONE_BUTTON_FUNCTION = 31
	PBF_DIRECTORY    PHONE_BUTTON_FUNCTION = 32
	PBF_COVER        PHONE_BUTTON_FUNCTION = 33
	PBF_CALLID       PHONE_BUTTON_FUNCTION = 34
	PBF_LASTNUM      PHONE_BUTTON_FUNCTION = 35
	PBF_NIGHTSRV     PHONE_BUTTON_FUNCTION = 36
	PBF_SENDCALLS    PHONE_BUTTON_FUNCTION = 37
	PBF_MSGINDICATOR PHONE_BUTTON_FUNCTION = 38
	PBF_REPDIAL      PHONE_BUTTON_FUNCTION = 39
	PBF_SETREPDIAL   PHONE_BUTTON_FUNCTION = 40
	PBF_SYSTEMSPEED  PHONE_BUTTON_FUNCTION = 41
	PBF_STATIONSPEED PHONE_BUTTON_FUNCTION = 42
	PBF_CAMPON       PHONE_BUTTON_FUNCTION = 43
	PBF_SAVEREPEAT   PHONE_BUTTON_FUNCTION = 44
	PBF_QUEUECALL    PHONE_BUTTON_FUNCTION = 45
	PBF_NONE         PHONE_BUTTON_FUNCTION = 46
	PBF_SEND         PHONE_BUTTON_FUNCTION = 47
)

type PHONE_TONE int32

const (
	PT_KEYPADZERO       PHONE_TONE = 0
	PT_KEYPADONE        PHONE_TONE = 1
	PT_KEYPADTWO        PHONE_TONE = 2
	PT_KEYPADTHREE      PHONE_TONE = 3
	PT_KEYPADFOUR       PHONE_TONE = 4
	PT_KEYPADFIVE       PHONE_TONE = 5
	PT_KEYPADSIX        PHONE_TONE = 6
	PT_KEYPADSEVEN      PHONE_TONE = 7
	PT_KEYPADEIGHT      PHONE_TONE = 8
	PT_KEYPADNINE       PHONE_TONE = 9
	PT_KEYPADSTAR       PHONE_TONE = 10
	PT_KEYPADPOUND      PHONE_TONE = 11
	PT_KEYPADA          PHONE_TONE = 12
	PT_KEYPADB          PHONE_TONE = 13
	PT_KEYPADC          PHONE_TONE = 14
	PT_KEYPADD          PHONE_TONE = 15
	PT_NORMALDIALTONE   PHONE_TONE = 16
	PT_EXTERNALDIALTONE PHONE_TONE = 17
	PT_BUSY             PHONE_TONE = 18
	PT_RINGBACK         PHONE_TONE = 19
	PT_ERRORTONE        PHONE_TONE = 20
	PT_SILENCE          PHONE_TONE = 21
)

type PHONE_EVENT int32

const (
	PE_DISPLAY        PHONE_EVENT = 0
	PE_LAMPMODE       PHONE_EVENT = 1
	PE_RINGMODE       PHONE_EVENT = 2
	PE_RINGVOLUME     PHONE_EVENT = 3
	PE_HOOKSWITCH     PHONE_EVENT = 4
	PE_CAPSCHANGE     PHONE_EVENT = 5
	PE_BUTTON         PHONE_EVENT = 6
	PE_CLOSE          PHONE_EVENT = 7
	PE_NUMBERGATHERED PHONE_EVENT = 8
	PE_DIALING        PHONE_EVENT = 9
	PE_ANSWER         PHONE_EVENT = 10
	PE_DISCONNECT     PHONE_EVENT = 11
	PE_LASTITEM       PHONE_EVENT = 11
)

type TERMINAL_MEDIA_STATE int32

const (
	TMS_IDLE     TERMINAL_MEDIA_STATE = 0
	TMS_ACTIVE   TERMINAL_MEDIA_STATE = 1
	TMS_PAUSED   TERMINAL_MEDIA_STATE = 2
	TMS_LASTITEM TERMINAL_MEDIA_STATE = 2
)

type FT_STATE_EVENT_CAUSE int32

const (
	FTEC_NORMAL      FT_STATE_EVENT_CAUSE = 0
	FTEC_END_OF_FILE FT_STATE_EVENT_CAUSE = 1
	FTEC_READ_ERROR  FT_STATE_EVENT_CAUSE = 2
	FTEC_WRITE_ERROR FT_STATE_EVENT_CAUSE = 3
)

type AGENT_EVENT int32

const (
	AE_NOT_READY     AGENT_EVENT = 0
	AE_READY         AGENT_EVENT = 1
	AE_BUSY_ACD      AGENT_EVENT = 2
	AE_BUSY_INCOMING AGENT_EVENT = 3
	AE_BUSY_OUTGOING AGENT_EVENT = 4
	AE_UNKNOWN       AGENT_EVENT = 5
)

type AGENT_STATE int32

const (
	AS_NOT_READY     AGENT_STATE = 0
	AS_READY         AGENT_STATE = 1
	AS_BUSY_ACD      AGENT_STATE = 2
	AS_BUSY_INCOMING AGENT_STATE = 3
	AS_BUSY_OUTGOING AGENT_STATE = 4
	AS_UNKNOWN       AGENT_STATE = 5
)

type AGENT_SESSION_EVENT int32

const (
	ASE_NEW_SESSION AGENT_SESSION_EVENT = 0
	ASE_NOT_READY   AGENT_SESSION_EVENT = 1
	ASE_READY       AGENT_SESSION_EVENT = 2
	ASE_BUSY        AGENT_SESSION_EVENT = 3
	ASE_WRAPUP      AGENT_SESSION_EVENT = 4
	ASE_END         AGENT_SESSION_EVENT = 5
)

type AGENT_SESSION_STATE int32

const (
	ASST_NOT_READY     AGENT_SESSION_STATE = 0
	ASST_READY         AGENT_SESSION_STATE = 1
	ASST_BUSY_ON_CALL  AGENT_SESSION_STATE = 2
	ASST_BUSY_WRAPUP   AGENT_SESSION_STATE = 3
	ASST_SESSION_ENDED AGENT_SESSION_STATE = 4
)

type AGENTHANDLER_EVENT int32

const (
	AHE_NEW_AGENTHANDLER     AGENTHANDLER_EVENT = 0
	AHE_AGENTHANDLER_REMOVED AGENTHANDLER_EVENT = 1
)

type ACDGROUP_EVENT int32

const (
	ACDGE_NEW_GROUP     ACDGROUP_EVENT = 0
	ACDGE_GROUP_REMOVED ACDGROUP_EVENT = 1
)

type ACDQUEUE_EVENT int32

const (
	ACDQE_NEW_QUEUE     ACDQUEUE_EVENT = 0
	ACDQE_QUEUE_REMOVED ACDQUEUE_EVENT = 1
)

type MSP_ADDRESS_EVENT int32

const (
	ADDRESS_TERMINAL_AVAILABLE   MSP_ADDRESS_EVENT = 0
	ADDRESS_TERMINAL_UNAVAILABLE MSP_ADDRESS_EVENT = 1
)

type MSP_CALL_EVENT int32

const (
	CALL_NEW_STREAM      MSP_CALL_EVENT = 0
	CALL_STREAM_FAIL     MSP_CALL_EVENT = 1
	CALL_TERMINAL_FAIL   MSP_CALL_EVENT = 2
	CALL_STREAM_NOT_USED MSP_CALL_EVENT = 3
	CALL_STREAM_ACTIVE   MSP_CALL_EVENT = 4
	CALL_STREAM_INACTIVE MSP_CALL_EVENT = 5
)

type MSP_CALL_EVENT_CAUSE int32

const (
	CALL_CAUSE_UNKNOWN            MSP_CALL_EVENT_CAUSE = 0
	CALL_CAUSE_BAD_DEVICE         MSP_CALL_EVENT_CAUSE = 1
	CALL_CAUSE_CONNECT_FAIL       MSP_CALL_EVENT_CAUSE = 2
	CALL_CAUSE_LOCAL_REQUEST      MSP_CALL_EVENT_CAUSE = 3
	CALL_CAUSE_REMOTE_REQUEST     MSP_CALL_EVENT_CAUSE = 4
	CALL_CAUSE_MEDIA_TIMEOUT      MSP_CALL_EVENT_CAUSE = 5
	CALL_CAUSE_MEDIA_RECOVERED    MSP_CALL_EVENT_CAUSE = 6
	CALL_CAUSE_QUALITY_OF_SERVICE MSP_CALL_EVENT_CAUSE = 7
)

type MSP_EVENT int32

const (
	ME_ADDRESS_EVENT       MSP_EVENT = 0
	ME_CALL_EVENT          MSP_EVENT = 1
	ME_TSP_DATA            MSP_EVENT = 2
	ME_PRIVATE_EVENT       MSP_EVENT = 3
	ME_ASR_TERMINAL_EVENT  MSP_EVENT = 4
	ME_TTS_TERMINAL_EVENT  MSP_EVENT = 5
	ME_FILE_TERMINAL_EVENT MSP_EVENT = 6
	ME_TONE_TERMINAL_EVENT MSP_EVENT = 7
)

type DIRECTORY_TYPE int32

const (
	DT_NTDS DIRECTORY_TYPE = 1
	DT_ILS  DIRECTORY_TYPE = 2
)

type DIRECTORY_OBJECT_TYPE int32

const (
	OT_CONFERENCE DIRECTORY_OBJECT_TYPE = 1
	OT_USER       DIRECTORY_OBJECT_TYPE = 2
)

type RND_ADVERTISING_SCOPE int32

const (
	RAS_LOCAL  RND_ADVERTISING_SCOPE = 1
	RAS_SITE   RND_ADVERTISING_SCOPE = 2
	RAS_REGION RND_ADVERTISING_SCOPE = 3
	RAS_WORLD  RND_ADVERTISING_SCOPE = 4
)

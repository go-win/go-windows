// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsmediaformat implements the Windows.Win32.WindowsMediaFormat namespace.
package windowsmediaformat

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0001 int32

const (
	WEBSTREAM_SAMPLE_TYPE_FILE   __MIDL___MIDL_itf_wmsdkidl_0000_0000_0001 = 1
	WEBSTREAM_SAMPLE_TYPE_RENDER __MIDL___MIDL_itf_wmsdkidl_0000_0000_0001 = 2
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0002 int32

const (
	WM_SF_CLEANPOINT    __MIDL___MIDL_itf_wmsdkidl_0000_0000_0002 = 1
	WM_SF_DISCONTINUITY __MIDL___MIDL_itf_wmsdkidl_0000_0000_0002 = 2
	WM_SF_DATALOSS      __MIDL___MIDL_itf_wmsdkidl_0000_0000_0002 = 4
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0003 int32

const (
	WM_SFEX_NOTASYNCPOINT __MIDL___MIDL_itf_wmsdkidl_0000_0000_0003 = 2
	WM_SFEX_DATALOSS      __MIDL___MIDL_itf_wmsdkidl_0000_0000_0003 = 4
)

type WMT_STATUS int32

const (
	WMT_ERROR                       WMT_STATUS = 0
	WMT_OPENED                      WMT_STATUS = 1
	WMT_BUFFERING_START             WMT_STATUS = 2
	WMT_BUFFERING_STOP              WMT_STATUS = 3
	WMT_EOF                         WMT_STATUS = 4
	WMT_END_OF_FILE                 WMT_STATUS = 4
	WMT_END_OF_SEGMENT              WMT_STATUS = 5
	WMT_END_OF_STREAMING            WMT_STATUS = 6
	WMT_LOCATING                    WMT_STATUS = 7
	WMT_CONNECTING                  WMT_STATUS = 8
	WMT_NO_RIGHTS                   WMT_STATUS = 9
	WMT_MISSING_CODEC               WMT_STATUS = 10
	WMT_STARTED                     WMT_STATUS = 11
	WMT_STOPPED                     WMT_STATUS = 12
	WMT_CLOSED                      WMT_STATUS = 13
	WMT_STRIDING                    WMT_STATUS = 14
	WMT_TIMER                       WMT_STATUS = 15
	WMT_INDEX_PROGRESS              WMT_STATUS = 16
	WMT_SAVEAS_START                WMT_STATUS = 17
	WMT_SAVEAS_STOP                 WMT_STATUS = 18
	WMT_NEW_SOURCEFLAGS             WMT_STATUS = 19
	WMT_NEW_METADATA                WMT_STATUS = 20
	WMT_BACKUPRESTORE_BEGIN         WMT_STATUS = 21
	WMT_SOURCE_SWITCH               WMT_STATUS = 22
	WMT_ACQUIRE_LICENSE             WMT_STATUS = 23
	WMT_INDIVIDUALIZE               WMT_STATUS = 24
	WMT_NEEDS_INDIVIDUALIZATION     WMT_STATUS = 25
	WMT_NO_RIGHTS_EX                WMT_STATUS = 26
	WMT_BACKUPRESTORE_END           WMT_STATUS = 27
	WMT_BACKUPRESTORE_CONNECTING    WMT_STATUS = 28
	WMT_BACKUPRESTORE_DISCONNECTING WMT_STATUS = 29
	WMT_ERROR_WITHURL               WMT_STATUS = 30
	WMT_RESTRICTED_LICENSE          WMT_STATUS = 31
	WMT_CLIENT_CONNECT              WMT_STATUS = 32
	WMT_CLIENT_DISCONNECT           WMT_STATUS = 33
	WMT_NATIVE_OUTPUT_PROPS_CHANGED WMT_STATUS = 34
	WMT_RECONNECT_START             WMT_STATUS = 35
	WMT_RECONNECT_END               WMT_STATUS = 36
	WMT_CLIENT_CONNECT_EX           WMT_STATUS = 37
	WMT_CLIENT_DISCONNECT_EX        WMT_STATUS = 38
	WMT_SET_FEC_SPAN                WMT_STATUS = 39
	WMT_PREROLL_READY               WMT_STATUS = 40
	WMT_PREROLL_COMPLETE            WMT_STATUS = 41
	WMT_CLIENT_PROPERTIES           WMT_STATUS = 42
	WMT_LICENSEURL_SIGNATURE_STATE  WMT_STATUS = 43
	WMT_INIT_PLAYLIST_BURN          WMT_STATUS = 44
	WMT_TRANSCRYPTOR_INIT           WMT_STATUS = 45
	WMT_TRANSCRYPTOR_SEEKED         WMT_STATUS = 46
	WMT_TRANSCRYPTOR_READ           WMT_STATUS = 47
	WMT_TRANSCRYPTOR_CLOSED         WMT_STATUS = 48
	WMT_PROXIMITY_RESULT            WMT_STATUS = 49
	WMT_PROXIMITY_COMPLETED         WMT_STATUS = 50
	WMT_CONTENT_ENABLER             WMT_STATUS = 51
)

type WMT_STREAM_SELECTION int32

const (
	WMT_OFF             WMT_STREAM_SELECTION = 0
	WMT_CLEANPOINT_ONLY WMT_STREAM_SELECTION = 1
	WMT_ON              WMT_STREAM_SELECTION = 2
)

type WMT_IMAGE_TYPE int32

const (
	WMT_IT_NONE   WMT_IMAGE_TYPE = 0
	WMT_IT_BITMAP WMT_IMAGE_TYPE = 1
	WMT_IT_JPEG   WMT_IMAGE_TYPE = 2
	WMT_IT_GIF    WMT_IMAGE_TYPE = 3
)

type WMT_ATTR_DATATYPE int32

const (
	WMT_TYPE_DWORD  WMT_ATTR_DATATYPE = 0
	WMT_TYPE_STRING WMT_ATTR_DATATYPE = 1
	WMT_TYPE_BINARY WMT_ATTR_DATATYPE = 2
	WMT_TYPE_BOOL   WMT_ATTR_DATATYPE = 3
	WMT_TYPE_QWORD  WMT_ATTR_DATATYPE = 4
	WMT_TYPE_WORD   WMT_ATTR_DATATYPE = 5
	WMT_TYPE_GUID   WMT_ATTR_DATATYPE = 6
)

type WMT_ATTR_IMAGETYPE int32

const (
	WMT_IMAGETYPE_BITMAP WMT_ATTR_IMAGETYPE = 1
	WMT_IMAGETYPE_JPEG   WMT_ATTR_IMAGETYPE = 2
	WMT_IMAGETYPE_GIF    WMT_ATTR_IMAGETYPE = 3
)

type WMT_VERSION int32

const (
	WMT_VER_4_0 WMT_VERSION = 262144
	WMT_VER_7_0 WMT_VERSION = 458752
	WMT_VER_8_0 WMT_VERSION = 524288
	WMT_VER_9_0 WMT_VERSION = 589824
)

type WMT_STORAGE_FORMAT int32

const (
	WMT_Storage_Format_MP3 WMT_STORAGE_FORMAT = 0
	WMT_Storage_Format_V1  WMT_STORAGE_FORMAT = 1
)

type WMT_DRMLA_TRUST int32

const (
	WMT_DRMLA_UNTRUSTED WMT_DRMLA_TRUST = 0
	WMT_DRMLA_TRUSTED   WMT_DRMLA_TRUST = 1
	WMT_DRMLA_TAMPERED  WMT_DRMLA_TRUST = 2
)

type WMT_TRANSPORT_TYPE int32

const (
	WMT_Transport_Type_Unreliable WMT_TRANSPORT_TYPE = 0
	WMT_Transport_Type_Reliable   WMT_TRANSPORT_TYPE = 1
)

type WMT_NET_PROTOCOL int32

const (
	WMT_PROTOCOL_HTTP WMT_NET_PROTOCOL = 0
)

type WMT_PLAY_MODE int32

const (
	WMT_PLAY_MODE_AUTOSELECT WMT_PLAY_MODE = 0
	WMT_PLAY_MODE_LOCAL      WMT_PLAY_MODE = 1
	WMT_PLAY_MODE_DOWNLOAD   WMT_PLAY_MODE = 2
	WMT_PLAY_MODE_STREAMING  WMT_PLAY_MODE = 3
)

type WMT_PROXY_SETTINGS int32

const (
	WMT_PROXY_SETTING_NONE    WMT_PROXY_SETTINGS = 0
	WMT_PROXY_SETTING_MANUAL  WMT_PROXY_SETTINGS = 1
	WMT_PROXY_SETTING_AUTO    WMT_PROXY_SETTINGS = 2
	WMT_PROXY_SETTING_BROWSER WMT_PROXY_SETTINGS = 3
	WMT_PROXY_SETTING_MAX     WMT_PROXY_SETTINGS = 4
)

type WMT_CODEC_INFO_TYPE int32

const (
	WMT_CODECINFO_AUDIO   WMT_CODEC_INFO_TYPE = 0
	WMT_CODECINFO_VIDEO   WMT_CODEC_INFO_TYPE = 1
	WMT_CODECINFO_UNKNOWN WMT_CODEC_INFO_TYPE = -1
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0004 int32

const (
	WM_DM_NOTINTERLACED                          __MIDL___MIDL_itf_wmsdkidl_0000_0000_0004 = 0
	WM_DM_DEINTERLACE_NORMAL                     __MIDL___MIDL_itf_wmsdkidl_0000_0000_0004 = 1
	WM_DM_DEINTERLACE_HALFSIZE                   __MIDL___MIDL_itf_wmsdkidl_0000_0000_0004 = 2
	WM_DM_DEINTERLACE_HALFSIZEDOUBLERATE         __MIDL___MIDL_itf_wmsdkidl_0000_0000_0004 = 3
	WM_DM_DEINTERLACE_INVERSETELECINE            __MIDL___MIDL_itf_wmsdkidl_0000_0000_0004 = 4
	WM_DM_DEINTERLACE_VERTICALHALFSIZEDOUBLERATE __MIDL___MIDL_itf_wmsdkidl_0000_0000_0004 = 5
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 int32

const (
	WM_DM_IT_DISABLE_COHERENT_MODE            __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 = 0
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_AA_TOP    __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 = 1
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BB_TOP    __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 = 2
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BC_TOP    __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 = 3
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_CD_TOP    __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 = 4
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_DD_TOP    __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 = 5
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_AA_BOTTOM __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 = 6
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BB_BOTTOM __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 = 7
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BC_BOTTOM __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 = 8
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_CD_BOTTOM __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 = 9
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_DD_BOTTOM __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 = 10
)

type WMT_OFFSET_FORMAT int32

const (
	WMT_OFFSET_FORMAT_100NS             WMT_OFFSET_FORMAT = 0
	WMT_OFFSET_FORMAT_FRAME_NUMBERS     WMT_OFFSET_FORMAT = 1
	WMT_OFFSET_FORMAT_PLAYLIST_OFFSET   WMT_OFFSET_FORMAT = 2
	WMT_OFFSET_FORMAT_TIMECODE          WMT_OFFSET_FORMAT = 3
	WMT_OFFSET_FORMAT_100NS_APPROXIMATE WMT_OFFSET_FORMAT = 4
)

type WMT_INDEXER_TYPE int32

const (
	WMT_IT_PRESENTATION_TIME WMT_INDEXER_TYPE = 0
	WMT_IT_FRAME_NUMBERS     WMT_INDEXER_TYPE = 1
	WMT_IT_TIMECODE          WMT_INDEXER_TYPE = 2
)

type WMT_INDEX_TYPE int32

const (
	WMT_IT_NEAREST_DATA_UNIT   WMT_INDEX_TYPE = 1
	WMT_IT_NEAREST_OBJECT      WMT_INDEX_TYPE = 2
	WMT_IT_NEAREST_CLEAN_POINT WMT_INDEX_TYPE = 3
)

type WMT_FILESINK_MODE int32

const (
	WMT_FM_SINGLE_BUFFERS      WMT_FILESINK_MODE = 1
	WMT_FM_FILESINK_DATA_UNITS WMT_FILESINK_MODE = 2
	WMT_FM_FILESINK_UNBUFFERED WMT_FILESINK_MODE = 4
)

type WMT_MUSICSPEECH_CLASS_MODE int32

const (
	WMT_MS_CLASS_MUSIC  WMT_MUSICSPEECH_CLASS_MODE = 0
	WMT_MS_CLASS_SPEECH WMT_MUSICSPEECH_CLASS_MODE = 1
	WMT_MS_CLASS_MIXED  WMT_MUSICSPEECH_CLASS_MODE = 2
)

type WMT_WATERMARK_ENTRY_TYPE int32

const (
	WMT_WMETYPE_AUDIO WMT_WATERMARK_ENTRY_TYPE = 1
	WMT_WMETYPE_VIDEO WMT_WATERMARK_ENTRY_TYPE = 2
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0006 int32

const (
	WM_PLAYBACK_DRC_HIGH   __MIDL___MIDL_itf_wmsdkidl_0000_0000_0006 = 0
	WM_PLAYBACK_DRC_MEDIUM __MIDL___MIDL_itf_wmsdkidl_0000_0000_0006 = 1
	WM_PLAYBACK_DRC_LOW    __MIDL___MIDL_itf_wmsdkidl_0000_0000_0006 = 2
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0007 int32

const (
	WMT_TIMECODE_FRAMERATE_30     __MIDL___MIDL_itf_wmsdkidl_0000_0000_0007 = 0
	WMT_TIMECODE_FRAMERATE_30DROP __MIDL___MIDL_itf_wmsdkidl_0000_0000_0007 = 1
	WMT_TIMECODE_FRAMERATE_25     __MIDL___MIDL_itf_wmsdkidl_0000_0000_0007 = 2
	WMT_TIMECODE_FRAMERATE_24     __MIDL___MIDL_itf_wmsdkidl_0000_0000_0007 = 3
)

type WMT_CREDENTIAL_FLAGS int32

const (
	WMT_CREDENTIAL_SAVE       WMT_CREDENTIAL_FLAGS = 1
	WMT_CREDENTIAL_DONT_CACHE WMT_CREDENTIAL_FLAGS = 2
	WMT_CREDENTIAL_CLEAR_TEXT WMT_CREDENTIAL_FLAGS = 4
	WMT_CREDENTIAL_PROXY      WMT_CREDENTIAL_FLAGS = 8
	WMT_CREDENTIAL_ENCRYPT    WMT_CREDENTIAL_FLAGS = 16
)

type WM_AETYPE int32

const (
	WM_AETYPE_INCLUDE WM_AETYPE = 105
	WM_AETYPE_EXCLUDE WM_AETYPE = 101
)

type WMT_RIGHTS int32

const (
	WMT_RIGHT_PLAYBACK                WMT_RIGHTS = 1
	WMT_RIGHT_COPY_TO_NON_SDMI_DEVICE WMT_RIGHTS = 2
	WMT_RIGHT_COPY_TO_CD              WMT_RIGHTS = 8
	WMT_RIGHT_COPY_TO_SDMI_DEVICE     WMT_RIGHTS = 16
	WMT_RIGHT_ONE_TIME                WMT_RIGHTS = 32
	WMT_RIGHT_SAVE_STREAM_PROTECTED   WMT_RIGHTS = 64
	WMT_RIGHT_COPY                    WMT_RIGHTS = 128
	WMT_RIGHT_COLLABORATIVE_PLAY      WMT_RIGHTS = 256
	WMT_RIGHT_SDMI_TRIGGER            WMT_RIGHTS = 65536
	WMT_RIGHT_SDMI_NOMORECOPIES       WMT_RIGHTS = 131072
)

type NETSOURCE_URLCREDPOLICY_SETTINGS int32

const (
	NETSOURCE_URLCREDPOLICY_SETTING_SILENTLOGONOK  NETSOURCE_URLCREDPOLICY_SETTINGS = 0
	NETSOURCE_URLCREDPOLICY_SETTING_MUSTPROMPTUSER NETSOURCE_URLCREDPOLICY_SETTINGS = 1
	NETSOURCE_URLCREDPOLICY_SETTING_ANONYMOUSONLY  NETSOURCE_URLCREDPOLICY_SETTINGS = 2
)

type _AM_ASFWRITERCONFIG_PARAM int32

const (
	AM_CONFIGASFWRITER_PARAM_AUTOINDEX    _AM_ASFWRITERCONFIG_PARAM = 1
	AM_CONFIGASFWRITER_PARAM_MULTIPASS    _AM_ASFWRITERCONFIG_PARAM = 2
	AM_CONFIGASFWRITER_PARAM_DONTCOMPRESS _AM_ASFWRITERCONFIG_PARAM = 3
)

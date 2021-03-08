// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsmediaformat implements the Windows.Win32.WindowsMediaFormat namespace.
package windowsmediaformat

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0001 int32

const (
	WEBSTREAM_SAMPLE_TYPE_FILE = 1
	WEBSTREAM_SAMPLE_TYPE_RENDER = 2
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0002 int32

const (
	WM_SF_CLEANPOINT = 1
	WM_SF_DISCONTINUITY = 2
	WM_SF_DATALOSS = 4
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0003 int32

const (
	WM_SFEX_NOTASYNCPOINT = 2
	WM_SFEX_DATALOSS = 4
)

type WMT_STATUS int32

const (
	WMT_ERROR = 0
	WMT_OPENED = 1
	WMT_BUFFERING_START = 2
	WMT_BUFFERING_STOP = 3
	WMT_EOF = 4
	WMT_END_OF_FILE = 4
	WMT_END_OF_SEGMENT = 5
	WMT_END_OF_STREAMING = 6
	WMT_LOCATING = 7
	WMT_CONNECTING = 8
	WMT_NO_RIGHTS = 9
	WMT_MISSING_CODEC = 10
	WMT_STARTED = 11
	WMT_STOPPED = 12
	WMT_CLOSED = 13
	WMT_STRIDING = 14
	WMT_TIMER = 15
	WMT_INDEX_PROGRESS = 16
	WMT_SAVEAS_START = 17
	WMT_SAVEAS_STOP = 18
	WMT_NEW_SOURCEFLAGS = 19
	WMT_NEW_METADATA = 20
	WMT_BACKUPRESTORE_BEGIN = 21
	WMT_SOURCE_SWITCH = 22
	WMT_ACQUIRE_LICENSE = 23
	WMT_INDIVIDUALIZE = 24
	WMT_NEEDS_INDIVIDUALIZATION = 25
	WMT_NO_RIGHTS_EX = 26
	WMT_BACKUPRESTORE_END = 27
	WMT_BACKUPRESTORE_CONNECTING = 28
	WMT_BACKUPRESTORE_DISCONNECTING = 29
	WMT_ERROR_WITHURL = 30
	WMT_RESTRICTED_LICENSE = 31
	WMT_CLIENT_CONNECT = 32
	WMT_CLIENT_DISCONNECT = 33
	WMT_NATIVE_OUTPUT_PROPS_CHANGED = 34
	WMT_RECONNECT_START = 35
	WMT_RECONNECT_END = 36
	WMT_CLIENT_CONNECT_EX = 37
	WMT_CLIENT_DISCONNECT_EX = 38
	WMT_SET_FEC_SPAN = 39
	WMT_PREROLL_READY = 40
	WMT_PREROLL_COMPLETE = 41
	WMT_CLIENT_PROPERTIES = 42
	WMT_LICENSEURL_SIGNATURE_STATE = 43
	WMT_INIT_PLAYLIST_BURN = 44
	WMT_TRANSCRYPTOR_INIT = 45
	WMT_TRANSCRYPTOR_SEEKED = 46
	WMT_TRANSCRYPTOR_READ = 47
	WMT_TRANSCRYPTOR_CLOSED = 48
	WMT_PROXIMITY_RESULT = 49
	WMT_PROXIMITY_COMPLETED = 50
	WMT_CONTENT_ENABLER = 51
)

type WMT_STREAM_SELECTION int32

const (
	WMT_OFF = 0
	WMT_CLEANPOINT_ONLY = 1
	WMT_ON = 2
)

type WMT_IMAGE_TYPE int32

const (
	WMT_IT_NONE = 0
	WMT_IT_BITMAP = 1
	WMT_IT_JPEG = 2
	WMT_IT_GIF = 3
)

type WMT_ATTR_DATATYPE int32

const (
	WMT_TYPE_DWORD = 0
	WMT_TYPE_STRING = 1
	WMT_TYPE_BINARY = 2
	WMT_TYPE_BOOL = 3
	WMT_TYPE_QWORD = 4
	WMT_TYPE_WORD = 5
	WMT_TYPE_GUID = 6
)

type WMT_ATTR_IMAGETYPE int32

const (
	WMT_IMAGETYPE_BITMAP = 1
	WMT_IMAGETYPE_JPEG = 2
	WMT_IMAGETYPE_GIF = 3
)

type WMT_VERSION int32

const (
	WMT_VER_4_0 = 262144
	WMT_VER_7_0 = 458752
	WMT_VER_8_0 = 524288
	WMT_VER_9_0 = 589824
)

type WMT_STORAGE_FORMAT int32

const (
	WMT_Storage_Format_MP3 = 0
	WMT_Storage_Format_V1 = 1
)

type WMT_DRMLA_TRUST int32

const (
	WMT_DRMLA_UNTRUSTED = 0
	WMT_DRMLA_TRUSTED = 1
	WMT_DRMLA_TAMPERED = 2
)

type WMT_TRANSPORT_TYPE int32

const (
	WMT_Transport_Type_Unreliable = 0
	WMT_Transport_Type_Reliable = 1
)

type WMT_NET_PROTOCOL int32

const (
	WMT_PROTOCOL_HTTP = 0
)

type WMT_PLAY_MODE int32

const (
	WMT_PLAY_MODE_AUTOSELECT = 0
	WMT_PLAY_MODE_LOCAL = 1
	WMT_PLAY_MODE_DOWNLOAD = 2
	WMT_PLAY_MODE_STREAMING = 3
)

type WMT_PROXY_SETTINGS int32

const (
	WMT_PROXY_SETTING_NONE = 0
	WMT_PROXY_SETTING_MANUAL = 1
	WMT_PROXY_SETTING_AUTO = 2
	WMT_PROXY_SETTING_BROWSER = 3
	WMT_PROXY_SETTING_MAX = 4
)

type WMT_CODEC_INFO_TYPE int32

const (
	WMT_CODECINFO_AUDIO = 0
	WMT_CODECINFO_VIDEO = 1
	WMT_CODECINFO_UNKNOWN = -1
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0004 int32

const (
	WM_DM_NOTINTERLACED = 0
	WM_DM_DEINTERLACE_NORMAL = 1
	WM_DM_DEINTERLACE_HALFSIZE = 2
	WM_DM_DEINTERLACE_HALFSIZEDOUBLERATE = 3
	WM_DM_DEINTERLACE_INVERSETELECINE = 4
	WM_DM_DEINTERLACE_VERTICALHALFSIZEDOUBLERATE = 5
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0005 int32

const (
	WM_DM_IT_DISABLE_COHERENT_MODE = 0
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_AA_TOP = 1
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BB_TOP = 2
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BC_TOP = 3
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_CD_TOP = 4
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_DD_TOP = 5
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_AA_BOTTOM = 6
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BB_BOTTOM = 7
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BC_BOTTOM = 8
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_CD_BOTTOM = 9
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_DD_BOTTOM = 10
)

type WMT_OFFSET_FORMAT int32

const (
	WMT_OFFSET_FORMAT_100NS = 0
	WMT_OFFSET_FORMAT_FRAME_NUMBERS = 1
	WMT_OFFSET_FORMAT_PLAYLIST_OFFSET = 2
	WMT_OFFSET_FORMAT_TIMECODE = 3
	WMT_OFFSET_FORMAT_100NS_APPROXIMATE = 4
)

type WMT_INDEXER_TYPE int32

const (
	WMT_IT_PRESENTATION_TIME = 0
	WMT_IT_FRAME_NUMBERS = 1
	WMT_IT_TIMECODE = 2
)

type WMT_INDEX_TYPE int32

const (
	WMT_IT_NEAREST_DATA_UNIT = 1
	WMT_IT_NEAREST_OBJECT = 2
	WMT_IT_NEAREST_CLEAN_POINT = 3
)

type WMT_FILESINK_MODE int32

const (
	WMT_FM_SINGLE_BUFFERS = 1
	WMT_FM_FILESINK_DATA_UNITS = 2
	WMT_FM_FILESINK_UNBUFFERED = 4
)

type WMT_MUSICSPEECH_CLASS_MODE int32

const (
	WMT_MS_CLASS_MUSIC = 0
	WMT_MS_CLASS_SPEECH = 1
	WMT_MS_CLASS_MIXED = 2
)

type WMT_WATERMARK_ENTRY_TYPE int32

const (
	WMT_WMETYPE_AUDIO = 1
	WMT_WMETYPE_VIDEO = 2
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0006 int32

const (
	WM_PLAYBACK_DRC_HIGH = 0
	WM_PLAYBACK_DRC_MEDIUM = 1
	WM_PLAYBACK_DRC_LOW = 2
)

type __MIDL___MIDL_itf_wmsdkidl_0000_0000_0007 int32

const (
	WMT_TIMECODE_FRAMERATE_30 = 0
	WMT_TIMECODE_FRAMERATE_30DROP = 1
	WMT_TIMECODE_FRAMERATE_25 = 2
	WMT_TIMECODE_FRAMERATE_24 = 3
)

type WMT_CREDENTIAL_FLAGS int32

const (
	WMT_CREDENTIAL_SAVE = 1
	WMT_CREDENTIAL_DONT_CACHE = 2
	WMT_CREDENTIAL_CLEAR_TEXT = 4
	WMT_CREDENTIAL_PROXY = 8
	WMT_CREDENTIAL_ENCRYPT = 16
)

type WM_AETYPE int32

const (
	WM_AETYPE_INCLUDE = 105
	WM_AETYPE_EXCLUDE = 101
)

type WMT_RIGHTS int32

const (
	WMT_RIGHT_PLAYBACK = 1
	WMT_RIGHT_COPY_TO_NON_SDMI_DEVICE = 2
	WMT_RIGHT_COPY_TO_CD = 8
	WMT_RIGHT_COPY_TO_SDMI_DEVICE = 16
	WMT_RIGHT_ONE_TIME = 32
	WMT_RIGHT_SAVE_STREAM_PROTECTED = 64
	WMT_RIGHT_COPY = 128
	WMT_RIGHT_COLLABORATIVE_PLAY = 256
	WMT_RIGHT_SDMI_TRIGGER = 65536
	WMT_RIGHT_SDMI_NOMORECOPIES = 131072
)

type NETSOURCE_URLCREDPOLICY_SETTINGS int32

const (
	NETSOURCE_URLCREDPOLICY_SETTING_SILENTLOGONOK = 0
	NETSOURCE_URLCREDPOLICY_SETTING_MUSTPROMPTUSER = 1
	NETSOURCE_URLCREDPOLICY_SETTING_ANONYMOUSONLY = 2
)

type _AM_ASFWRITERCONFIG_PARAM int32

const (
	AM_CONFIGASFWRITER_PARAM_AUTOINDEX = 1
	AM_CONFIGASFWRITER_PARAM_MULTIPASS = 2
	AM_CONFIGASFWRITER_PARAM_DONTCOMPRESS = 3
)

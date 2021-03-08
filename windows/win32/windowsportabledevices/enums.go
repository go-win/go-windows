// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsportabledevices implements the Windows.Win32.WindowsPortableDevices namespace.
package windowsportabledevices

type DELETE_OBJECT_OPTIONS int32

const (
	PORTABLE_DEVICE_DELETE_NO_RECURSION   DELETE_OBJECT_OPTIONS = 0
	PORTABLE_DEVICE_DELETE_WITH_RECURSION DELETE_OBJECT_OPTIONS = 1
)

type WPD_DEVICE_TYPES int32

const (
	WPD_DEVICE_TYPE_GENERIC                      WPD_DEVICE_TYPES = 0
	WPD_DEVICE_TYPE_CAMERA                       WPD_DEVICE_TYPES = 1
	WPD_DEVICE_TYPE_MEDIA_PLAYER                 WPD_DEVICE_TYPES = 2
	WPD_DEVICE_TYPE_PHONE                        WPD_DEVICE_TYPES = 3
	WPD_DEVICE_TYPE_VIDEO                        WPD_DEVICE_TYPES = 4
	WPD_DEVICE_TYPE_PERSONAL_INFORMATION_MANAGER WPD_DEVICE_TYPES = 5
	WPD_DEVICE_TYPE_AUDIO_RECORDER               WPD_DEVICE_TYPES = 6
)

type WpdAttributeForm int32

const (
	WPD_PROPERTY_ATTRIBUTE_FORM_UNSPECIFIED        WpdAttributeForm = 0
	WPD_PROPERTY_ATTRIBUTE_FORM_RANGE              WpdAttributeForm = 1
	WPD_PROPERTY_ATTRIBUTE_FORM_ENUMERATION        WpdAttributeForm = 2
	WPD_PROPERTY_ATTRIBUTE_FORM_REGULAR_EXPRESSION WpdAttributeForm = 3
	WPD_PROPERTY_ATTRIBUTE_FORM_OBJECT_IDENTIFIER  WpdAttributeForm = 4
)

type WpdParameterAttributeForm int32

const (
	WPD_PARAMETER_ATTRIBUTE_FORM_UNSPECIFIED        WpdParameterAttributeForm = 0
	WPD_PARAMETER_ATTRIBUTE_FORM_RANGE              WpdParameterAttributeForm = 1
	WPD_PARAMETER_ATTRIBUTE_FORM_ENUMERATION        WpdParameterAttributeForm = 2
	WPD_PARAMETER_ATTRIBUTE_FORM_REGULAR_EXPRESSION WpdParameterAttributeForm = 3
	WPD_PARAMETER_ATTRIBUTE_FORM_OBJECT_IDENTIFIER  WpdParameterAttributeForm = 4
)

type WPD_DEVICE_TRANSPORTS int32

const (
	WPD_DEVICE_TRANSPORT_UNSPECIFIED WPD_DEVICE_TRANSPORTS = 0
	WPD_DEVICE_TRANSPORT_USB         WPD_DEVICE_TRANSPORTS = 1
	WPD_DEVICE_TRANSPORT_IP          WPD_DEVICE_TRANSPORTS = 2
	WPD_DEVICE_TRANSPORT_BLUETOOTH   WPD_DEVICE_TRANSPORTS = 3
)

type WPD_STORAGE_TYPE_VALUES int32

const (
	WPD_STORAGE_TYPE_UNDEFINED     WPD_STORAGE_TYPE_VALUES = 0
	WPD_STORAGE_TYPE_FIXED_ROM     WPD_STORAGE_TYPE_VALUES = 1
	WPD_STORAGE_TYPE_REMOVABLE_ROM WPD_STORAGE_TYPE_VALUES = 2
	WPD_STORAGE_TYPE_FIXED_RAM     WPD_STORAGE_TYPE_VALUES = 3
	WPD_STORAGE_TYPE_REMOVABLE_RAM WPD_STORAGE_TYPE_VALUES = 4
)

type WPD_STORAGE_ACCESS_CAPABILITY_VALUES int32

const (
	WPD_STORAGE_ACCESS_CAPABILITY_READWRITE                         WPD_STORAGE_ACCESS_CAPABILITY_VALUES = 0
	WPD_STORAGE_ACCESS_CAPABILITY_READ_ONLY_WITHOUT_OBJECT_DELETION WPD_STORAGE_ACCESS_CAPABILITY_VALUES = 1
	WPD_STORAGE_ACCESS_CAPABILITY_READ_ONLY_WITH_OBJECT_DELETION    WPD_STORAGE_ACCESS_CAPABILITY_VALUES = 2
)

type WPD_SMS_ENCODING_TYPES int32

const (
	SMS_ENCODING_7_BIT  WPD_SMS_ENCODING_TYPES = 0
	SMS_ENCODING_8_BIT  WPD_SMS_ENCODING_TYPES = 1
	SMS_ENCODING_UTF_16 WPD_SMS_ENCODING_TYPES = 2
)

type SMS_MESSAGE_TYPES int32

const (
	SMS_TEXT_MESSAGE   SMS_MESSAGE_TYPES = 0
	SMS_BINARY_MESSAGE SMS_MESSAGE_TYPES = 1
)

type WPD_POWER_SOURCES int32

const (
	WPD_POWER_SOURCE_BATTERY  WPD_POWER_SOURCES = 0
	WPD_POWER_SOURCE_EXTERNAL WPD_POWER_SOURCES = 1
)

type WPD_WHITE_BALANCE_SETTINGS int32

const (
	WPD_WHITE_BALANCE_UNDEFINED          WPD_WHITE_BALANCE_SETTINGS = 0
	WPD_WHITE_BALANCE_MANUAL             WPD_WHITE_BALANCE_SETTINGS = 1
	WPD_WHITE_BALANCE_AUTOMATIC          WPD_WHITE_BALANCE_SETTINGS = 2
	WPD_WHITE_BALANCE_ONE_PUSH_AUTOMATIC WPD_WHITE_BALANCE_SETTINGS = 3
	WPD_WHITE_BALANCE_DAYLIGHT           WPD_WHITE_BALANCE_SETTINGS = 4
	WPD_WHITE_BALANCE_FLORESCENT         WPD_WHITE_BALANCE_SETTINGS = 5
	WPD_WHITE_BALANCE_TUNGSTEN           WPD_WHITE_BALANCE_SETTINGS = 6
	WPD_WHITE_BALANCE_FLASH              WPD_WHITE_BALANCE_SETTINGS = 7
)

type WPD_FOCUS_MODES int32

const (
	WPD_FOCUS_UNDEFINED       WPD_FOCUS_MODES = 0
	WPD_FOCUS_MANUAL          WPD_FOCUS_MODES = 1
	WPD_FOCUS_AUTOMATIC       WPD_FOCUS_MODES = 2
	WPD_FOCUS_AUTOMATIC_MACRO WPD_FOCUS_MODES = 3
)

type WPD_EXPOSURE_METERING_MODES int32

const (
	WPD_EXPOSURE_METERING_MODE_UNDEFINED               WPD_EXPOSURE_METERING_MODES = 0
	WPD_EXPOSURE_METERING_MODE_AVERAGE                 WPD_EXPOSURE_METERING_MODES = 1
	WPD_EXPOSURE_METERING_MODE_CENTER_WEIGHTED_AVERAGE WPD_EXPOSURE_METERING_MODES = 2
	WPD_EXPOSURE_METERING_MODE_MULTI_SPOT              WPD_EXPOSURE_METERING_MODES = 3
	WPD_EXPOSURE_METERING_MODE_CENTER_SPOT             WPD_EXPOSURE_METERING_MODES = 4
)

type WPD_FLASH_MODES int32

const (
	WPD_FLASH_MODE_UNDEFINED     WPD_FLASH_MODES = 0
	WPD_FLASH_MODE_AUTO          WPD_FLASH_MODES = 1
	WPD_FLASH_MODE_OFF           WPD_FLASH_MODES = 2
	WPD_FLASH_MODE_FILL          WPD_FLASH_MODES = 3
	WPD_FLASH_MODE_RED_EYE_AUTO  WPD_FLASH_MODES = 4
	WPD_FLASH_MODE_RED_EYE_FILL  WPD_FLASH_MODES = 5
	WPD_FLASH_MODE_EXTERNAL_SYNC WPD_FLASH_MODES = 6
)

type WPD_EXPOSURE_PROGRAM_MODES int32

const (
	WPD_EXPOSURE_PROGRAM_MODE_UNDEFINED         WPD_EXPOSURE_PROGRAM_MODES = 0
	WPD_EXPOSURE_PROGRAM_MODE_MANUAL            WPD_EXPOSURE_PROGRAM_MODES = 1
	WPD_EXPOSURE_PROGRAM_MODE_AUTO              WPD_EXPOSURE_PROGRAM_MODES = 2
	WPD_EXPOSURE_PROGRAM_MODE_APERTURE_PRIORITY WPD_EXPOSURE_PROGRAM_MODES = 3
	WPD_EXPOSURE_PROGRAM_MODE_SHUTTER_PRIORITY  WPD_EXPOSURE_PROGRAM_MODES = 4
	WPD_EXPOSURE_PROGRAM_MODE_CREATIVE          WPD_EXPOSURE_PROGRAM_MODES = 5
	WPD_EXPOSURE_PROGRAM_MODE_ACTION            WPD_EXPOSURE_PROGRAM_MODES = 6
	WPD_EXPOSURE_PROGRAM_MODE_PORTRAIT          WPD_EXPOSURE_PROGRAM_MODES = 7
)

type WPD_CAPTURE_MODES int32

const (
	WPD_CAPTURE_MODE_UNDEFINED WPD_CAPTURE_MODES = 0
	WPD_CAPTURE_MODE_NORMAL    WPD_CAPTURE_MODES = 1
	WPD_CAPTURE_MODE_BURST     WPD_CAPTURE_MODES = 2
	WPD_CAPTURE_MODE_TIMELAPSE WPD_CAPTURE_MODES = 3
)

type WPD_EFFECT_MODES int32

const (
	WPD_EFFECT_MODE_UNDEFINED       WPD_EFFECT_MODES = 0
	WPD_EFFECT_MODE_COLOR           WPD_EFFECT_MODES = 1
	WPD_EFFECT_MODE_BLACK_AND_WHITE WPD_EFFECT_MODES = 2
	WPD_EFFECT_MODE_SEPIA           WPD_EFFECT_MODES = 3
)

type WPD_FOCUS_METERING_MODES int32

const (
	WPD_FOCUS_METERING_MODE_UNDEFINED   WPD_FOCUS_METERING_MODES = 0
	WPD_FOCUS_METERING_MODE_CENTER_SPOT WPD_FOCUS_METERING_MODES = 1
	WPD_FOCUS_METERING_MODE_MULTI_SPOT  WPD_FOCUS_METERING_MODES = 2
)

type WPD_BITRATE_TYPES int32

const (
	WPD_BITRATE_TYPE_UNUSED   WPD_BITRATE_TYPES = 0
	WPD_BITRATE_TYPE_DISCRETE WPD_BITRATE_TYPES = 1
	WPD_BITRATE_TYPE_VARIABLE WPD_BITRATE_TYPES = 2
	WPD_BITRATE_TYPE_FREE     WPD_BITRATE_TYPES = 3
)

type WPD_META_GENRES int32

const (
	WPD_META_GENRE_UNUSED                           WPD_META_GENRES = 0
	WPD_META_GENRE_GENERIC_MUSIC_AUDIO_FILE         WPD_META_GENRES = 1
	WPD_META_GENRE_GENERIC_NON_MUSIC_AUDIO_FILE     WPD_META_GENRES = 17
	WPD_META_GENRE_SPOKEN_WORD_AUDIO_BOOK_FILES     WPD_META_GENRES = 18
	WPD_META_GENRE_SPOKEN_WORD_FILES_NON_AUDIO_BOOK WPD_META_GENRES = 19
	WPD_META_GENRE_SPOKEN_WORD_NEWS                 WPD_META_GENRES = 20
	WPD_META_GENRE_SPOKEN_WORD_TALK_SHOWS           WPD_META_GENRES = 21
	WPD_META_GENRE_GENERIC_VIDEO_FILE               WPD_META_GENRES = 33
	WPD_META_GENRE_NEWS_VIDEO_FILE                  WPD_META_GENRES = 34
	WPD_META_GENRE_MUSIC_VIDEO_FILE                 WPD_META_GENRES = 35
	WPD_META_GENRE_HOME_VIDEO_FILE                  WPD_META_GENRES = 36
	WPD_META_GENRE_FEATURE_FILM_VIDEO_FILE          WPD_META_GENRES = 37
	WPD_META_GENRE_TELEVISION_VIDEO_FILE            WPD_META_GENRES = 38
	WPD_META_GENRE_TRAINING_EDUCATIONAL_VIDEO_FILE  WPD_META_GENRES = 39
	WPD_META_GENRE_PHOTO_MONTAGE_VIDEO_FILE         WPD_META_GENRES = 40
	WPD_META_GENRE_GENERIC_NON_AUDIO_NON_VIDEO      WPD_META_GENRES = 48
	WPD_META_GENRE_AUDIO_PODCAST                    WPD_META_GENRES = 64
	WPD_META_GENRE_VIDEO_PODCAST                    WPD_META_GENRES = 65
	WPD_META_GENRE_MIXED_PODCAST                    WPD_META_GENRES = 66
)

type WPD_CROPPED_STATUS_VALUES int32

const (
	WPD_CROPPED_STATUS_NOT_CROPPED           WPD_CROPPED_STATUS_VALUES = 0
	WPD_CROPPED_STATUS_CROPPED               WPD_CROPPED_STATUS_VALUES = 1
	WPD_CROPPED_STATUS_SHOULD_NOT_BE_CROPPED WPD_CROPPED_STATUS_VALUES = 2
)

type WPD_COLOR_CORRECTED_STATUS_VALUES int32

const (
	WPD_COLOR_CORRECTED_STATUS_NOT_CORRECTED           WPD_COLOR_CORRECTED_STATUS_VALUES = 0
	WPD_COLOR_CORRECTED_STATUS_CORRECTED               WPD_COLOR_CORRECTED_STATUS_VALUES = 1
	WPD_COLOR_CORRECTED_STATUS_SHOULD_NOT_BE_CORRECTED WPD_COLOR_CORRECTED_STATUS_VALUES = 2
)

type WPD_VIDEO_SCAN_TYPES int32

const (
	WPD_VIDEO_SCAN_TYPE_UNUSED                          WPD_VIDEO_SCAN_TYPES = 0
	WPD_VIDEO_SCAN_TYPE_PROGRESSIVE                     WPD_VIDEO_SCAN_TYPES = 1
	WPD_VIDEO_SCAN_TYPE_FIELD_INTERLEAVED_UPPER_FIRST   WPD_VIDEO_SCAN_TYPES = 2
	WPD_VIDEO_SCAN_TYPE_FIELD_INTERLEAVED_LOWER_FIRST   WPD_VIDEO_SCAN_TYPES = 3
	WPD_VIDEO_SCAN_TYPE_FIELD_SINGLE_UPPER_FIRST        WPD_VIDEO_SCAN_TYPES = 4
	WPD_VIDEO_SCAN_TYPE_FIELD_SINGLE_LOWER_FIRST        WPD_VIDEO_SCAN_TYPES = 5
	WPD_VIDEO_SCAN_TYPE_MIXED_INTERLACE                 WPD_VIDEO_SCAN_TYPES = 6
	WPD_VIDEO_SCAN_TYPE_MIXED_INTERLACE_AND_PROGRESSIVE WPD_VIDEO_SCAN_TYPES = 7
)

type WPD_OPERATION_STATES int32

const (
	WPD_OPERATION_STATE_UNSPECIFIED WPD_OPERATION_STATES = 0
	WPD_OPERATION_STATE_STARTED     WPD_OPERATION_STATES = 1
	WPD_OPERATION_STATE_RUNNING     WPD_OPERATION_STATES = 2
	WPD_OPERATION_STATE_PAUSED      WPD_OPERATION_STATES = 3
	WPD_OPERATION_STATE_CANCELLED   WPD_OPERATION_STATES = 4
	WPD_OPERATION_STATE_FINISHED    WPD_OPERATION_STATES = 5
	WPD_OPERATION_STATE_ABORTED     WPD_OPERATION_STATES = 6
)

type WPD_SECTION_DATA_UNITS_VALUES int32

const (
	WPD_SECTION_DATA_UNITS_BYTES        WPD_SECTION_DATA_UNITS_VALUES = 0
	WPD_SECTION_DATA_UNITS_MILLISECONDS WPD_SECTION_DATA_UNITS_VALUES = 1
)

type WPD_RENDERING_INFORMATION_PROFILE_ENTRY_TYPES int32

const (
	WPD_RENDERING_INFORMATION_PROFILE_ENTRY_TYPE_OBJECT   WPD_RENDERING_INFORMATION_PROFILE_ENTRY_TYPES = 0
	WPD_RENDERING_INFORMATION_PROFILE_ENTRY_TYPE_RESOURCE WPD_RENDERING_INFORMATION_PROFILE_ENTRY_TYPES = 1
)

type WPD_COMMAND_ACCESS_TYPES int32

const (
	WPD_COMMAND_ACCESS_READ                              WPD_COMMAND_ACCESS_TYPES = 1
	WPD_COMMAND_ACCESS_READWRITE                         WPD_COMMAND_ACCESS_TYPES = 3
	WPD_COMMAND_ACCESS_FROM_PROPERTY_WITH_STGM_ACCESS    WPD_COMMAND_ACCESS_TYPES = 4
	WPD_COMMAND_ACCESS_FROM_PROPERTY_WITH_FILE_ACCESS    WPD_COMMAND_ACCESS_TYPES = 8
	WPD_COMMAND_ACCESS_FROM_ATTRIBUTE_WITH_METHOD_ACCESS WPD_COMMAND_ACCESS_TYPES = 16
)

type WPD_SERVICE_INHERITANCE_TYPES int32

const (
	WPD_SERVICE_INHERITANCE_IMPLEMENTATION WPD_SERVICE_INHERITANCE_TYPES = 0
)

type WPD_PARAMETER_USAGE_TYPES int32

const (
	WPD_PARAMETER_USAGE_RETURN WPD_PARAMETER_USAGE_TYPES = 0
	WPD_PARAMETER_USAGE_IN     WPD_PARAMETER_USAGE_TYPES = 1
	WPD_PARAMETER_USAGE_OUT    WPD_PARAMETER_USAGE_TYPES = 2
	WPD_PARAMETER_USAGE_INOUT  WPD_PARAMETER_USAGE_TYPES = 3
)

type WPD_STREAM_UNITS int32

const (
	WPD_STREAM_UNITS_BYTES        WPD_STREAM_UNITS = 0
	WPD_STREAM_UNITS_FRAMES       WPD_STREAM_UNITS = 1
	WPD_STREAM_UNITS_ROWS         WPD_STREAM_UNITS = 2
	WPD_STREAM_UNITS_MILLISECONDS WPD_STREAM_UNITS = 4
	WPD_STREAM_UNITS_MICROSECONDS WPD_STREAM_UNITS = 8
)

// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package imapi implements the Windows.Win32.IMapi namespace.
package imapi

type IMAPI_MEDIA_PHYSICAL_TYPE int32

const (
	IMAPI_MEDIA_TYPE_UNKNOWN             IMAPI_MEDIA_PHYSICAL_TYPE = 0
	IMAPI_MEDIA_TYPE_CDROM               IMAPI_MEDIA_PHYSICAL_TYPE = 1
	IMAPI_MEDIA_TYPE_CDR                 IMAPI_MEDIA_PHYSICAL_TYPE = 2
	IMAPI_MEDIA_TYPE_CDRW                IMAPI_MEDIA_PHYSICAL_TYPE = 3
	IMAPI_MEDIA_TYPE_DVDROM              IMAPI_MEDIA_PHYSICAL_TYPE = 4
	IMAPI_MEDIA_TYPE_DVDRAM              IMAPI_MEDIA_PHYSICAL_TYPE = 5
	IMAPI_MEDIA_TYPE_DVDPLUSR            IMAPI_MEDIA_PHYSICAL_TYPE = 6
	IMAPI_MEDIA_TYPE_DVDPLUSRW           IMAPI_MEDIA_PHYSICAL_TYPE = 7
	IMAPI_MEDIA_TYPE_DVDPLUSR_DUALLAYER  IMAPI_MEDIA_PHYSICAL_TYPE = 8
	IMAPI_MEDIA_TYPE_DVDDASHR            IMAPI_MEDIA_PHYSICAL_TYPE = 9
	IMAPI_MEDIA_TYPE_DVDDASHRW           IMAPI_MEDIA_PHYSICAL_TYPE = 10
	IMAPI_MEDIA_TYPE_DVDDASHR_DUALLAYER  IMAPI_MEDIA_PHYSICAL_TYPE = 11
	IMAPI_MEDIA_TYPE_DISK                IMAPI_MEDIA_PHYSICAL_TYPE = 12
	IMAPI_MEDIA_TYPE_DVDPLUSRW_DUALLAYER IMAPI_MEDIA_PHYSICAL_TYPE = 13
	IMAPI_MEDIA_TYPE_HDDVDROM            IMAPI_MEDIA_PHYSICAL_TYPE = 14
	IMAPI_MEDIA_TYPE_HDDVDR              IMAPI_MEDIA_PHYSICAL_TYPE = 15
	IMAPI_MEDIA_TYPE_HDDVDRAM            IMAPI_MEDIA_PHYSICAL_TYPE = 16
	IMAPI_MEDIA_TYPE_BDROM               IMAPI_MEDIA_PHYSICAL_TYPE = 17
	IMAPI_MEDIA_TYPE_BDR                 IMAPI_MEDIA_PHYSICAL_TYPE = 18
	IMAPI_MEDIA_TYPE_BDRE                IMAPI_MEDIA_PHYSICAL_TYPE = 19
	IMAPI_MEDIA_TYPE_MAX                 IMAPI_MEDIA_PHYSICAL_TYPE = 19
)

type IMAPI_MEDIA_WRITE_PROTECT_STATE int32

const (
	IMAPI_WRITEPROTECTED_UNTIL_POWERDOWN           IMAPI_MEDIA_WRITE_PROTECT_STATE = 1
	IMAPI_WRITEPROTECTED_BY_CARTRIDGE              IMAPI_MEDIA_WRITE_PROTECT_STATE = 2
	IMAPI_WRITEPROTECTED_BY_MEDIA_SPECIFIC_REASON  IMAPI_MEDIA_WRITE_PROTECT_STATE = 4
	IMAPI_WRITEPROTECTED_BY_SOFTWARE_WRITE_PROTECT IMAPI_MEDIA_WRITE_PROTECT_STATE = 8
	IMAPI_WRITEPROTECTED_BY_DISC_CONTROL_BLOCK     IMAPI_MEDIA_WRITE_PROTECT_STATE = 16
	IMAPI_WRITEPROTECTED_READ_ONLY_MEDIA           IMAPI_MEDIA_WRITE_PROTECT_STATE = 16384
)

type IMAPI_READ_TRACK_ADDRESS_TYPE int32

const (
	IMAPI_READ_TRACK_ADDRESS_TYPE_LBA     IMAPI_READ_TRACK_ADDRESS_TYPE = 0
	IMAPI_READ_TRACK_ADDRESS_TYPE_TRACK   IMAPI_READ_TRACK_ADDRESS_TYPE = 1
	IMAPI_READ_TRACK_ADDRESS_TYPE_SESSION IMAPI_READ_TRACK_ADDRESS_TYPE = 2
)

type IMAPI_MODE_PAGE_REQUEST_TYPE int32

const (
	IMAPI_MODE_PAGE_REQUEST_TYPE_CURRENT_VALUES    IMAPI_MODE_PAGE_REQUEST_TYPE = 0
	IMAPI_MODE_PAGE_REQUEST_TYPE_CHANGEABLE_VALUES IMAPI_MODE_PAGE_REQUEST_TYPE = 1
	IMAPI_MODE_PAGE_REQUEST_TYPE_DEFAULT_VALUES    IMAPI_MODE_PAGE_REQUEST_TYPE = 2
	IMAPI_MODE_PAGE_REQUEST_TYPE_SAVED_VALUES      IMAPI_MODE_PAGE_REQUEST_TYPE = 3
)

type IMAPI_MODE_PAGE_TYPE int32

const (
	IMAPI_MODE_PAGE_TYPE_READ_WRITE_ERROR_RECOVERY IMAPI_MODE_PAGE_TYPE = 1
	IMAPI_MODE_PAGE_TYPE_MRW                       IMAPI_MODE_PAGE_TYPE = 3
	IMAPI_MODE_PAGE_TYPE_WRITE_PARAMETERS          IMAPI_MODE_PAGE_TYPE = 5
	IMAPI_MODE_PAGE_TYPE_CACHING                   IMAPI_MODE_PAGE_TYPE = 8
	IMAPI_MODE_PAGE_TYPE_INFORMATIONAL_EXCEPTIONS  IMAPI_MODE_PAGE_TYPE = 28
	IMAPI_MODE_PAGE_TYPE_TIMEOUT_AND_PROTECT       IMAPI_MODE_PAGE_TYPE = 29
	IMAPI_MODE_PAGE_TYPE_POWER_CONDITION           IMAPI_MODE_PAGE_TYPE = 26
	IMAPI_MODE_PAGE_TYPE_LEGACY_CAPABILITIES       IMAPI_MODE_PAGE_TYPE = 42
)

type IMAPI_FEATURE_PAGE_TYPE int32

const (
	IMAPI_FEATURE_PAGE_TYPE_PROFILE_LIST                   IMAPI_FEATURE_PAGE_TYPE = 0
	IMAPI_FEATURE_PAGE_TYPE_CORE                           IMAPI_FEATURE_PAGE_TYPE = 1
	IMAPI_FEATURE_PAGE_TYPE_MORPHING                       IMAPI_FEATURE_PAGE_TYPE = 2
	IMAPI_FEATURE_PAGE_TYPE_REMOVABLE_MEDIUM               IMAPI_FEATURE_PAGE_TYPE = 3
	IMAPI_FEATURE_PAGE_TYPE_WRITE_PROTECT                  IMAPI_FEATURE_PAGE_TYPE = 4
	IMAPI_FEATURE_PAGE_TYPE_RANDOMLY_READABLE              IMAPI_FEATURE_PAGE_TYPE = 16
	IMAPI_FEATURE_PAGE_TYPE_CD_MULTIREAD                   IMAPI_FEATURE_PAGE_TYPE = 29
	IMAPI_FEATURE_PAGE_TYPE_CD_READ                        IMAPI_FEATURE_PAGE_TYPE = 30
	IMAPI_FEATURE_PAGE_TYPE_DVD_READ                       IMAPI_FEATURE_PAGE_TYPE = 31
	IMAPI_FEATURE_PAGE_TYPE_RANDOMLY_WRITABLE              IMAPI_FEATURE_PAGE_TYPE = 32
	IMAPI_FEATURE_PAGE_TYPE_INCREMENTAL_STREAMING_WRITABLE IMAPI_FEATURE_PAGE_TYPE = 33
	IMAPI_FEATURE_PAGE_TYPE_SECTOR_ERASABLE                IMAPI_FEATURE_PAGE_TYPE = 34
	IMAPI_FEATURE_PAGE_TYPE_FORMATTABLE                    IMAPI_FEATURE_PAGE_TYPE = 35
	IMAPI_FEATURE_PAGE_TYPE_HARDWARE_DEFECT_MANAGEMENT     IMAPI_FEATURE_PAGE_TYPE = 36
	IMAPI_FEATURE_PAGE_TYPE_WRITE_ONCE                     IMAPI_FEATURE_PAGE_TYPE = 37
	IMAPI_FEATURE_PAGE_TYPE_RESTRICTED_OVERWRITE           IMAPI_FEATURE_PAGE_TYPE = 38
	IMAPI_FEATURE_PAGE_TYPE_CDRW_CAV_WRITE                 IMAPI_FEATURE_PAGE_TYPE = 39
	IMAPI_FEATURE_PAGE_TYPE_MRW                            IMAPI_FEATURE_PAGE_TYPE = 40
	IMAPI_FEATURE_PAGE_TYPE_ENHANCED_DEFECT_REPORTING      IMAPI_FEATURE_PAGE_TYPE = 41
	IMAPI_FEATURE_PAGE_TYPE_DVD_PLUS_RW                    IMAPI_FEATURE_PAGE_TYPE = 42
	IMAPI_FEATURE_PAGE_TYPE_DVD_PLUS_R                     IMAPI_FEATURE_PAGE_TYPE = 43
	IMAPI_FEATURE_PAGE_TYPE_RIGID_RESTRICTED_OVERWRITE     IMAPI_FEATURE_PAGE_TYPE = 44
	IMAPI_FEATURE_PAGE_TYPE_CD_TRACK_AT_ONCE               IMAPI_FEATURE_PAGE_TYPE = 45
	IMAPI_FEATURE_PAGE_TYPE_CD_MASTERING                   IMAPI_FEATURE_PAGE_TYPE = 46
	IMAPI_FEATURE_PAGE_TYPE_DVD_DASH_WRITE                 IMAPI_FEATURE_PAGE_TYPE = 47
	IMAPI_FEATURE_PAGE_TYPE_DOUBLE_DENSITY_CD_READ         IMAPI_FEATURE_PAGE_TYPE = 48
	IMAPI_FEATURE_PAGE_TYPE_DOUBLE_DENSITY_CD_R_WRITE      IMAPI_FEATURE_PAGE_TYPE = 49
	IMAPI_FEATURE_PAGE_TYPE_DOUBLE_DENSITY_CD_RW_WRITE     IMAPI_FEATURE_PAGE_TYPE = 50
	IMAPI_FEATURE_PAGE_TYPE_LAYER_JUMP_RECORDING           IMAPI_FEATURE_PAGE_TYPE = 51
	IMAPI_FEATURE_PAGE_TYPE_CD_RW_MEDIA_WRITE_SUPPORT      IMAPI_FEATURE_PAGE_TYPE = 55
	IMAPI_FEATURE_PAGE_TYPE_BD_PSEUDO_OVERWRITE            IMAPI_FEATURE_PAGE_TYPE = 56
	IMAPI_FEATURE_PAGE_TYPE_DVD_PLUS_R_DUAL_LAYER          IMAPI_FEATURE_PAGE_TYPE = 59
	IMAPI_FEATURE_PAGE_TYPE_BD_READ                        IMAPI_FEATURE_PAGE_TYPE = 64
	IMAPI_FEATURE_PAGE_TYPE_BD_WRITE                       IMAPI_FEATURE_PAGE_TYPE = 65
	IMAPI_FEATURE_PAGE_TYPE_HD_DVD_READ                    IMAPI_FEATURE_PAGE_TYPE = 80
	IMAPI_FEATURE_PAGE_TYPE_HD_DVD_WRITE                   IMAPI_FEATURE_PAGE_TYPE = 81
	IMAPI_FEATURE_PAGE_TYPE_POWER_MANAGEMENT               IMAPI_FEATURE_PAGE_TYPE = 256
	IMAPI_FEATURE_PAGE_TYPE_SMART                          IMAPI_FEATURE_PAGE_TYPE = 257
	IMAPI_FEATURE_PAGE_TYPE_EMBEDDED_CHANGER               IMAPI_FEATURE_PAGE_TYPE = 258
	IMAPI_FEATURE_PAGE_TYPE_CD_ANALOG_PLAY                 IMAPI_FEATURE_PAGE_TYPE = 259
	IMAPI_FEATURE_PAGE_TYPE_MICROCODE_UPDATE               IMAPI_FEATURE_PAGE_TYPE = 260
	IMAPI_FEATURE_PAGE_TYPE_TIMEOUT                        IMAPI_FEATURE_PAGE_TYPE = 261
	IMAPI_FEATURE_PAGE_TYPE_DVD_CSS                        IMAPI_FEATURE_PAGE_TYPE = 262
	IMAPI_FEATURE_PAGE_TYPE_REAL_TIME_STREAMING            IMAPI_FEATURE_PAGE_TYPE = 263
	IMAPI_FEATURE_PAGE_TYPE_LOGICAL_UNIT_SERIAL_NUMBER     IMAPI_FEATURE_PAGE_TYPE = 264
	IMAPI_FEATURE_PAGE_TYPE_MEDIA_SERIAL_NUMBER            IMAPI_FEATURE_PAGE_TYPE = 265
	IMAPI_FEATURE_PAGE_TYPE_DISC_CONTROL_BLOCKS            IMAPI_FEATURE_PAGE_TYPE = 266
	IMAPI_FEATURE_PAGE_TYPE_DVD_CPRM                       IMAPI_FEATURE_PAGE_TYPE = 267
	IMAPI_FEATURE_PAGE_TYPE_FIRMWARE_INFORMATION           IMAPI_FEATURE_PAGE_TYPE = 268
	IMAPI_FEATURE_PAGE_TYPE_AACS                           IMAPI_FEATURE_PAGE_TYPE = 269
	IMAPI_FEATURE_PAGE_TYPE_VCPS                           IMAPI_FEATURE_PAGE_TYPE = 272
)

type IMAPI_PROFILE_TYPE int32

const (
	IMAPI_PROFILE_TYPE_INVALID                    IMAPI_PROFILE_TYPE = 0
	IMAPI_PROFILE_TYPE_NON_REMOVABLE_DISK         IMAPI_PROFILE_TYPE = 1
	IMAPI_PROFILE_TYPE_REMOVABLE_DISK             IMAPI_PROFILE_TYPE = 2
	IMAPI_PROFILE_TYPE_MO_ERASABLE                IMAPI_PROFILE_TYPE = 3
	IMAPI_PROFILE_TYPE_MO_WRITE_ONCE              IMAPI_PROFILE_TYPE = 4
	IMAPI_PROFILE_TYPE_AS_MO                      IMAPI_PROFILE_TYPE = 5
	IMAPI_PROFILE_TYPE_CDROM                      IMAPI_PROFILE_TYPE = 8
	IMAPI_PROFILE_TYPE_CD_RECORDABLE              IMAPI_PROFILE_TYPE = 9
	IMAPI_PROFILE_TYPE_CD_REWRITABLE              IMAPI_PROFILE_TYPE = 10
	IMAPI_PROFILE_TYPE_DVDROM                     IMAPI_PROFILE_TYPE = 16
	IMAPI_PROFILE_TYPE_DVD_DASH_RECORDABLE        IMAPI_PROFILE_TYPE = 17
	IMAPI_PROFILE_TYPE_DVD_RAM                    IMAPI_PROFILE_TYPE = 18
	IMAPI_PROFILE_TYPE_DVD_DASH_REWRITABLE        IMAPI_PROFILE_TYPE = 19
	IMAPI_PROFILE_TYPE_DVD_DASH_RW_SEQUENTIAL     IMAPI_PROFILE_TYPE = 20
	IMAPI_PROFILE_TYPE_DVD_DASH_R_DUAL_SEQUENTIAL IMAPI_PROFILE_TYPE = 21
	IMAPI_PROFILE_TYPE_DVD_DASH_R_DUAL_LAYER_JUMP IMAPI_PROFILE_TYPE = 22
	IMAPI_PROFILE_TYPE_DVD_PLUS_RW                IMAPI_PROFILE_TYPE = 26
	IMAPI_PROFILE_TYPE_DVD_PLUS_R                 IMAPI_PROFILE_TYPE = 27
	IMAPI_PROFILE_TYPE_DDCDROM                    IMAPI_PROFILE_TYPE = 32
	IMAPI_PROFILE_TYPE_DDCD_RECORDABLE            IMAPI_PROFILE_TYPE = 33
	IMAPI_PROFILE_TYPE_DDCD_REWRITABLE            IMAPI_PROFILE_TYPE = 34
	IMAPI_PROFILE_TYPE_DVD_PLUS_RW_DUAL           IMAPI_PROFILE_TYPE = 42
	IMAPI_PROFILE_TYPE_DVD_PLUS_R_DUAL            IMAPI_PROFILE_TYPE = 43
	IMAPI_PROFILE_TYPE_BD_ROM                     IMAPI_PROFILE_TYPE = 64
	IMAPI_PROFILE_TYPE_BD_R_SEQUENTIAL            IMAPI_PROFILE_TYPE = 65
	IMAPI_PROFILE_TYPE_BD_R_RANDOM_RECORDING      IMAPI_PROFILE_TYPE = 66
	IMAPI_PROFILE_TYPE_BD_REWRITABLE              IMAPI_PROFILE_TYPE = 67
	IMAPI_PROFILE_TYPE_HD_DVD_ROM                 IMAPI_PROFILE_TYPE = 80
	IMAPI_PROFILE_TYPE_HD_DVD_RECORDABLE          IMAPI_PROFILE_TYPE = 81
	IMAPI_PROFILE_TYPE_HD_DVD_RAM                 IMAPI_PROFILE_TYPE = 82
	IMAPI_PROFILE_TYPE_NON_STANDARD               IMAPI_PROFILE_TYPE = 65535
)

type IMAPI_FORMAT2_DATA_WRITE_ACTION int32

const (
	IMAPI_FORMAT2_DATA_WRITE_ACTION_VALIDATING_MEDIA      IMAPI_FORMAT2_DATA_WRITE_ACTION = 0
	IMAPI_FORMAT2_DATA_WRITE_ACTION_FORMATTING_MEDIA      IMAPI_FORMAT2_DATA_WRITE_ACTION = 1
	IMAPI_FORMAT2_DATA_WRITE_ACTION_INITIALIZING_HARDWARE IMAPI_FORMAT2_DATA_WRITE_ACTION = 2
	IMAPI_FORMAT2_DATA_WRITE_ACTION_CALIBRATING_POWER     IMAPI_FORMAT2_DATA_WRITE_ACTION = 3
	IMAPI_FORMAT2_DATA_WRITE_ACTION_WRITING_DATA          IMAPI_FORMAT2_DATA_WRITE_ACTION = 4
	IMAPI_FORMAT2_DATA_WRITE_ACTION_FINALIZATION          IMAPI_FORMAT2_DATA_WRITE_ACTION = 5
	IMAPI_FORMAT2_DATA_WRITE_ACTION_COMPLETED             IMAPI_FORMAT2_DATA_WRITE_ACTION = 6
	IMAPI_FORMAT2_DATA_WRITE_ACTION_VERIFYING             IMAPI_FORMAT2_DATA_WRITE_ACTION = 7
)

type IMAPI_FORMAT2_DATA_MEDIA_STATE int32

const (
	IMAPI_FORMAT2_DATA_MEDIA_STATE_UNKNOWN            IMAPI_FORMAT2_DATA_MEDIA_STATE = 0
	IMAPI_FORMAT2_DATA_MEDIA_STATE_INFORMATIONAL_MASK IMAPI_FORMAT2_DATA_MEDIA_STATE = 15
	IMAPI_FORMAT2_DATA_MEDIA_STATE_UNSUPPORTED_MASK   IMAPI_FORMAT2_DATA_MEDIA_STATE = 64512
	IMAPI_FORMAT2_DATA_MEDIA_STATE_OVERWRITE_ONLY     IMAPI_FORMAT2_DATA_MEDIA_STATE = 1
	IMAPI_FORMAT2_DATA_MEDIA_STATE_RANDOMLY_WRITABLE  IMAPI_FORMAT2_DATA_MEDIA_STATE = 1
	IMAPI_FORMAT2_DATA_MEDIA_STATE_BLANK              IMAPI_FORMAT2_DATA_MEDIA_STATE = 2
	IMAPI_FORMAT2_DATA_MEDIA_STATE_APPENDABLE         IMAPI_FORMAT2_DATA_MEDIA_STATE = 4
	IMAPI_FORMAT2_DATA_MEDIA_STATE_FINAL_SESSION      IMAPI_FORMAT2_DATA_MEDIA_STATE = 8
	IMAPI_FORMAT2_DATA_MEDIA_STATE_DAMAGED            IMAPI_FORMAT2_DATA_MEDIA_STATE = 1024
	IMAPI_FORMAT2_DATA_MEDIA_STATE_ERASE_REQUIRED     IMAPI_FORMAT2_DATA_MEDIA_STATE = 2048
	IMAPI_FORMAT2_DATA_MEDIA_STATE_NON_EMPTY_SESSION  IMAPI_FORMAT2_DATA_MEDIA_STATE = 4096
	IMAPI_FORMAT2_DATA_MEDIA_STATE_WRITE_PROTECTED    IMAPI_FORMAT2_DATA_MEDIA_STATE = 8192
	IMAPI_FORMAT2_DATA_MEDIA_STATE_FINALIZED          IMAPI_FORMAT2_DATA_MEDIA_STATE = 16384
	IMAPI_FORMAT2_DATA_MEDIA_STATE_UNSUPPORTED_MEDIA  IMAPI_FORMAT2_DATA_MEDIA_STATE = 32768
)

type IMAPI_FORMAT2_TAO_WRITE_ACTION int32

const (
	IMAPI_FORMAT2_TAO_WRITE_ACTION_UNKNOWN   IMAPI_FORMAT2_TAO_WRITE_ACTION = 0
	IMAPI_FORMAT2_TAO_WRITE_ACTION_PREPARING IMAPI_FORMAT2_TAO_WRITE_ACTION = 1
	IMAPI_FORMAT2_TAO_WRITE_ACTION_WRITING   IMAPI_FORMAT2_TAO_WRITE_ACTION = 2
	IMAPI_FORMAT2_TAO_WRITE_ACTION_FINISHING IMAPI_FORMAT2_TAO_WRITE_ACTION = 3
	IMAPI_FORMAT2_TAO_WRITE_ACTION_VERIFYING IMAPI_FORMAT2_TAO_WRITE_ACTION = 4
)

type IMAPI_FORMAT2_RAW_CD_DATA_SECTOR_TYPE int32

const (
	IMAPI_FORMAT2_RAW_CD_SUBCODE_PQ_ONLY   IMAPI_FORMAT2_RAW_CD_DATA_SECTOR_TYPE = 1
	IMAPI_FORMAT2_RAW_CD_SUBCODE_IS_COOKED IMAPI_FORMAT2_RAW_CD_DATA_SECTOR_TYPE = 2
	IMAPI_FORMAT2_RAW_CD_SUBCODE_IS_RAW    IMAPI_FORMAT2_RAW_CD_DATA_SECTOR_TYPE = 3
)

type IMAPI_FORMAT2_RAW_CD_WRITE_ACTION int32

const (
	IMAPI_FORMAT2_RAW_CD_WRITE_ACTION_UNKNOWN   IMAPI_FORMAT2_RAW_CD_WRITE_ACTION = 0
	IMAPI_FORMAT2_RAW_CD_WRITE_ACTION_PREPARING IMAPI_FORMAT2_RAW_CD_WRITE_ACTION = 1
	IMAPI_FORMAT2_RAW_CD_WRITE_ACTION_WRITING   IMAPI_FORMAT2_RAW_CD_WRITE_ACTION = 2
	IMAPI_FORMAT2_RAW_CD_WRITE_ACTION_FINISHING IMAPI_FORMAT2_RAW_CD_WRITE_ACTION = 3
)

type IMAPI_CD_SECTOR_TYPE int32

const (
	IMAPI_CD_SECTOR_AUDIO         IMAPI_CD_SECTOR_TYPE = 0
	IMAPI_CD_SECTOR_MODE_ZERO     IMAPI_CD_SECTOR_TYPE = 1
	IMAPI_CD_SECTOR_MODE1         IMAPI_CD_SECTOR_TYPE = 2
	IMAPI_CD_SECTOR_MODE2FORM0    IMAPI_CD_SECTOR_TYPE = 3
	IMAPI_CD_SECTOR_MODE2FORM1    IMAPI_CD_SECTOR_TYPE = 4
	IMAPI_CD_SECTOR_MODE2FORM2    IMAPI_CD_SECTOR_TYPE = 5
	IMAPI_CD_SECTOR_MODE1RAW      IMAPI_CD_SECTOR_TYPE = 6
	IMAPI_CD_SECTOR_MODE2FORM0RAW IMAPI_CD_SECTOR_TYPE = 7
	IMAPI_CD_SECTOR_MODE2FORM1RAW IMAPI_CD_SECTOR_TYPE = 8
	IMAPI_CD_SECTOR_MODE2FORM2RAW IMAPI_CD_SECTOR_TYPE = 9
)

type IMAPI_CD_TRACK_DIGITAL_COPY_SETTING int32

const (
	IMAPI_CD_TRACK_DIGITAL_COPY_PERMITTED  IMAPI_CD_TRACK_DIGITAL_COPY_SETTING = 0
	IMAPI_CD_TRACK_DIGITAL_COPY_PROHIBITED IMAPI_CD_TRACK_DIGITAL_COPY_SETTING = 1
	IMAPI_CD_TRACK_DIGITAL_COPY_SCMS       IMAPI_CD_TRACK_DIGITAL_COPY_SETTING = 2
)

type IMAPI_BURN_VERIFICATION_LEVEL int32

const (
	IMAPI_BURN_VERIFICATION_NONE  IMAPI_BURN_VERIFICATION_LEVEL = 0
	IMAPI_BURN_VERIFICATION_QUICK IMAPI_BURN_VERIFICATION_LEVEL = 1
	IMAPI_BURN_VERIFICATION_FULL  IMAPI_BURN_VERIFICATION_LEVEL = 2
)

type FsiItemType int32

const (
	FsiItemNotFound  FsiItemType = 0
	FsiItemDirectory FsiItemType = 1
	FsiItemFile      FsiItemType = 2
)

type FsiFileSystems int32

const (
	FsiFileSystemNone    FsiFileSystems = 0
	FsiFileSystemISO9660 FsiFileSystems = 1
	FsiFileSystemJoliet  FsiFileSystems = 2
	FsiFileSystemUDF     FsiFileSystems = 4
	FsiFileSystemUnknown FsiFileSystems = 1073741824
)

type EmulationType int32

const (
	EmulationNone       EmulationType = 0
	Emulation12MFloppy  EmulationType = 1
	Emulation144MFloppy EmulationType = 2
	Emulation288MFloppy EmulationType = 3
	EmulationHardDisk   EmulationType = 4
)

type PlatformId int32

const (
	PlatformX86     PlatformId = 0
	PlatformPowerPC PlatformId = 1
	PlatformMac     PlatformId = 2
	PlatformEFI     PlatformId = 239
)

type MEDIA_TYPES int32

const (
	MEDIA_CDDA_CDROM MEDIA_TYPES = 1
	MEDIA_CD_ROM_XA  MEDIA_TYPES = 2
	MEDIA_CD_I       MEDIA_TYPES = 3
	MEDIA_CD_EXTRA   MEDIA_TYPES = 4
	MEDIA_CD_OTHER   MEDIA_TYPES = 5
	MEDIA_SPECIAL    MEDIA_TYPES = 6
)

type MEDIA_FLAGS int32

const (
	MEDIA_BLANK                    MEDIA_FLAGS = 1
	MEDIA_RW                       MEDIA_FLAGS = 2
	MEDIA_WRITABLE                 MEDIA_FLAGS = 4
	MEDIA_FORMAT_UNUSABLE_BY_IMAPI MEDIA_FLAGS = 8
)

type RECORDER_TYPES int32

const (
	RECORDER_CDR  RECORDER_TYPES = 1
	RECORDER_CDRW RECORDER_TYPES = 2
)

// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package directmanipulation implements the Windows.Win32.DirectManipulation namespace.
package directmanipulation

type DIRECTMANIPULATION_STATUS int32

const (
	DIRECTMANIPULATION_BUILDING  DIRECTMANIPULATION_STATUS = 0
	DIRECTMANIPULATION_ENABLED   DIRECTMANIPULATION_STATUS = 1
	DIRECTMANIPULATION_DISABLED  DIRECTMANIPULATION_STATUS = 2
	DIRECTMANIPULATION_RUNNING   DIRECTMANIPULATION_STATUS = 3
	DIRECTMANIPULATION_INERTIA   DIRECTMANIPULATION_STATUS = 4
	DIRECTMANIPULATION_READY     DIRECTMANIPULATION_STATUS = 5
	DIRECTMANIPULATION_SUSPENDED DIRECTMANIPULATION_STATUS = 6
)

type DIRECTMANIPULATION_HITTEST_TYPE int32

const (
	DIRECTMANIPULATION_HITTEST_TYPE_ASYNCHRONOUS     DIRECTMANIPULATION_HITTEST_TYPE = 0
	DIRECTMANIPULATION_HITTEST_TYPE_SYNCHRONOUS      DIRECTMANIPULATION_HITTEST_TYPE = 1
	DIRECTMANIPULATION_HITTEST_TYPE_AUTO_SYNCHRONOUS DIRECTMANIPULATION_HITTEST_TYPE = 2
)

type DIRECTMANIPULATION_CONFIGURATION int32

const (
	DIRECTMANIPULATION_CONFIGURATION_NONE                DIRECTMANIPULATION_CONFIGURATION = 0
	DIRECTMANIPULATION_CONFIGURATION_INTERACTION         DIRECTMANIPULATION_CONFIGURATION = 1
	DIRECTMANIPULATION_CONFIGURATION_TRANSLATION_X       DIRECTMANIPULATION_CONFIGURATION = 2
	DIRECTMANIPULATION_CONFIGURATION_TRANSLATION_Y       DIRECTMANIPULATION_CONFIGURATION = 4
	DIRECTMANIPULATION_CONFIGURATION_SCALING             DIRECTMANIPULATION_CONFIGURATION = 16
	DIRECTMANIPULATION_CONFIGURATION_TRANSLATION_INERTIA DIRECTMANIPULATION_CONFIGURATION = 32
	DIRECTMANIPULATION_CONFIGURATION_SCALING_INERTIA     DIRECTMANIPULATION_CONFIGURATION = 128
	DIRECTMANIPULATION_CONFIGURATION_RAILS_X             DIRECTMANIPULATION_CONFIGURATION = 256
	DIRECTMANIPULATION_CONFIGURATION_RAILS_Y             DIRECTMANIPULATION_CONFIGURATION = 512
)

type DIRECTMANIPULATION_GESTURE_CONFIGURATION int32

const (
	DIRECTMANIPULATION_GESTURE_NONE                   DIRECTMANIPULATION_GESTURE_CONFIGURATION = 0
	DIRECTMANIPULATION_GESTURE_DEFAULT                DIRECTMANIPULATION_GESTURE_CONFIGURATION = 0
	DIRECTMANIPULATION_GESTURE_CROSS_SLIDE_VERTICAL   DIRECTMANIPULATION_GESTURE_CONFIGURATION = 8
	DIRECTMANIPULATION_GESTURE_CROSS_SLIDE_HORIZONTAL DIRECTMANIPULATION_GESTURE_CONFIGURATION = 16
	DIRECTMANIPULATION_GESTURE_PINCH_ZOOM             DIRECTMANIPULATION_GESTURE_CONFIGURATION = 32
)

type DIRECTMANIPULATION_MOTION_TYPES int32

const (
	DIRECTMANIPULATION_MOTION_NONE       DIRECTMANIPULATION_MOTION_TYPES = 0
	DIRECTMANIPULATION_MOTION_TRANSLATEX DIRECTMANIPULATION_MOTION_TYPES = 1
	DIRECTMANIPULATION_MOTION_TRANSLATEY DIRECTMANIPULATION_MOTION_TYPES = 2
	DIRECTMANIPULATION_MOTION_ZOOM       DIRECTMANIPULATION_MOTION_TYPES = 4
	DIRECTMANIPULATION_MOTION_CENTERX    DIRECTMANIPULATION_MOTION_TYPES = 16
	DIRECTMANIPULATION_MOTION_CENTERY    DIRECTMANIPULATION_MOTION_TYPES = 32
	DIRECTMANIPULATION_MOTION_ALL        DIRECTMANIPULATION_MOTION_TYPES = 55
)

type DIRECTMANIPULATION_VIEWPORT_OPTIONS int32

const (
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_DEFAULT              DIRECTMANIPULATION_VIEWPORT_OPTIONS = 0
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_AUTODISABLE          DIRECTMANIPULATION_VIEWPORT_OPTIONS = 1
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_MANUALUPDATE         DIRECTMANIPULATION_VIEWPORT_OPTIONS = 2
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_INPUT                DIRECTMANIPULATION_VIEWPORT_OPTIONS = 4
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_EXPLICITHITTEST      DIRECTMANIPULATION_VIEWPORT_OPTIONS = 8
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_DISABLEPIXELSNAPPING DIRECTMANIPULATION_VIEWPORT_OPTIONS = 16
)

type DIRECTMANIPULATION_SNAPPOINT_TYPE int32

const (
	DIRECTMANIPULATION_SNAPPOINT_MANDATORY        DIRECTMANIPULATION_SNAPPOINT_TYPE = 0
	DIRECTMANIPULATION_SNAPPOINT_OPTIONAL         DIRECTMANIPULATION_SNAPPOINT_TYPE = 1
	DIRECTMANIPULATION_SNAPPOINT_MANDATORY_SINGLE DIRECTMANIPULATION_SNAPPOINT_TYPE = 2
	DIRECTMANIPULATION_SNAPPOINT_OPTIONAL_SINGLE  DIRECTMANIPULATION_SNAPPOINT_TYPE = 3
)

type DIRECTMANIPULATION_SNAPPOINT_COORDINATE int32

const (
	DIRECTMANIPULATION_COORDINATE_BOUNDARY DIRECTMANIPULATION_SNAPPOINT_COORDINATE = 0
	DIRECTMANIPULATION_COORDINATE_ORIGIN   DIRECTMANIPULATION_SNAPPOINT_COORDINATE = 1
	DIRECTMANIPULATION_COORDINATE_MIRRORED DIRECTMANIPULATION_SNAPPOINT_COORDINATE = 16
)

type DIRECTMANIPULATION_HORIZONTALALIGNMENT int32

const (
	DIRECTMANIPULATION_HORIZONTALALIGNMENT_NONE         DIRECTMANIPULATION_HORIZONTALALIGNMENT = 0
	DIRECTMANIPULATION_HORIZONTALALIGNMENT_LEFT         DIRECTMANIPULATION_HORIZONTALALIGNMENT = 1
	DIRECTMANIPULATION_HORIZONTALALIGNMENT_CENTER       DIRECTMANIPULATION_HORIZONTALALIGNMENT = 2
	DIRECTMANIPULATION_HORIZONTALALIGNMENT_RIGHT        DIRECTMANIPULATION_HORIZONTALALIGNMENT = 4
	DIRECTMANIPULATION_HORIZONTALALIGNMENT_UNLOCKCENTER DIRECTMANIPULATION_HORIZONTALALIGNMENT = 8
)

type DIRECTMANIPULATION_VERTICALALIGNMENT int32

const (
	DIRECTMANIPULATION_VERTICALALIGNMENT_NONE         DIRECTMANIPULATION_VERTICALALIGNMENT = 0
	DIRECTMANIPULATION_VERTICALALIGNMENT_TOP          DIRECTMANIPULATION_VERTICALALIGNMENT = 1
	DIRECTMANIPULATION_VERTICALALIGNMENT_CENTER       DIRECTMANIPULATION_VERTICALALIGNMENT = 2
	DIRECTMANIPULATION_VERTICALALIGNMENT_BOTTOM       DIRECTMANIPULATION_VERTICALALIGNMENT = 4
	DIRECTMANIPULATION_VERTICALALIGNMENT_UNLOCKCENTER DIRECTMANIPULATION_VERTICALALIGNMENT = 8
)

type DIRECTMANIPULATION_INPUT_MODE int32

const (
	DIRECTMANIPULATION_INPUT_MODE_AUTOMATIC DIRECTMANIPULATION_INPUT_MODE = 0
	DIRECTMANIPULATION_INPUT_MODE_MANUAL    DIRECTMANIPULATION_INPUT_MODE = 1
)

type DIRECTMANIPULATION_DRAG_DROP_STATUS int32

const (
	DIRECTMANIPULATION_DRAG_DROP_READY     DIRECTMANIPULATION_DRAG_DROP_STATUS = 0
	DIRECTMANIPULATION_DRAG_DROP_PRESELECT DIRECTMANIPULATION_DRAG_DROP_STATUS = 1
	DIRECTMANIPULATION_DRAG_DROP_SELECTING DIRECTMANIPULATION_DRAG_DROP_STATUS = 2
	DIRECTMANIPULATION_DRAG_DROP_DRAGGING  DIRECTMANIPULATION_DRAG_DROP_STATUS = 3
	DIRECTMANIPULATION_DRAG_DROP_CANCELLED DIRECTMANIPULATION_DRAG_DROP_STATUS = 4
	DIRECTMANIPULATION_DRAG_DROP_COMMITTED DIRECTMANIPULATION_DRAG_DROP_STATUS = 5
)

type DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION int32

const (
	DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION_VERTICAL    DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION = 1
	DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION_HORIZONTAL  DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION = 2
	DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION_SELECT_ONLY DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION = 16
	DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION_SELECT_DRAG DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION = 32
	DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION_HOLD_DRAG   DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION = 64
)

type DIRECTMANIPULATION_INTERACTION_TYPE int32

const (
	DIRECTMANIPULATION_INTERACTION_BEGIN                    DIRECTMANIPULATION_INTERACTION_TYPE = 0
	DIRECTMANIPULATION_INTERACTION_TYPE_MANIPULATION        DIRECTMANIPULATION_INTERACTION_TYPE = 1
	DIRECTMANIPULATION_INTERACTION_TYPE_GESTURE_TAP         DIRECTMANIPULATION_INTERACTION_TYPE = 2
	DIRECTMANIPULATION_INTERACTION_TYPE_GESTURE_HOLD        DIRECTMANIPULATION_INTERACTION_TYPE = 3
	DIRECTMANIPULATION_INTERACTION_TYPE_GESTURE_CROSS_SLIDE DIRECTMANIPULATION_INTERACTION_TYPE = 4
	DIRECTMANIPULATION_INTERACTION_TYPE_GESTURE_PINCH_ZOOM  DIRECTMANIPULATION_INTERACTION_TYPE = 5
	DIRECTMANIPULATION_INTERACTION_END                      DIRECTMANIPULATION_INTERACTION_TYPE = 100
)

type DIRECTMANIPULATION_AUTOSCROLL_CONFIGURATION int32

const (
	DIRECTMANIPULATION_AUTOSCROLL_CONFIGURATION_STOP    DIRECTMANIPULATION_AUTOSCROLL_CONFIGURATION = 0
	DIRECTMANIPULATION_AUTOSCROLL_CONFIGURATION_FORWARD DIRECTMANIPULATION_AUTOSCROLL_CONFIGURATION = 1
	DIRECTMANIPULATION_AUTOSCROLL_CONFIGURATION_REVERSE DIRECTMANIPULATION_AUTOSCROLL_CONFIGURATION = 2
)

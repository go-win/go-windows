// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package directmanipulation implements the Windows.Win32.DirectManipulation namespace.
package directmanipulation

type DIRECTMANIPULATION_STATUS int32

const (
	DIRECTMANIPULATION_BUILDING = 0
	DIRECTMANIPULATION_ENABLED = 1
	DIRECTMANIPULATION_DISABLED = 2
	DIRECTMANIPULATION_RUNNING = 3
	DIRECTMANIPULATION_INERTIA = 4
	DIRECTMANIPULATION_READY = 5
	DIRECTMANIPULATION_SUSPENDED = 6
)

type DIRECTMANIPULATION_HITTEST_TYPE int32

const (
	DIRECTMANIPULATION_HITTEST_TYPE_ASYNCHRONOUS = 0
	DIRECTMANIPULATION_HITTEST_TYPE_SYNCHRONOUS = 1
	DIRECTMANIPULATION_HITTEST_TYPE_AUTO_SYNCHRONOUS = 2
)

type DIRECTMANIPULATION_CONFIGURATION int32

const (
	DIRECTMANIPULATION_CONFIGURATION_NONE = 0
	DIRECTMANIPULATION_CONFIGURATION_INTERACTION = 1
	DIRECTMANIPULATION_CONFIGURATION_TRANSLATION_X = 2
	DIRECTMANIPULATION_CONFIGURATION_TRANSLATION_Y = 4
	DIRECTMANIPULATION_CONFIGURATION_SCALING = 16
	DIRECTMANIPULATION_CONFIGURATION_TRANSLATION_INERTIA = 32
	DIRECTMANIPULATION_CONFIGURATION_SCALING_INERTIA = 128
	DIRECTMANIPULATION_CONFIGURATION_RAILS_X = 256
	DIRECTMANIPULATION_CONFIGURATION_RAILS_Y = 512
)

type DIRECTMANIPULATION_GESTURE_CONFIGURATION int32

const (
	DIRECTMANIPULATION_GESTURE_NONE = 0
	DIRECTMANIPULATION_GESTURE_DEFAULT = 0
	DIRECTMANIPULATION_GESTURE_CROSS_SLIDE_VERTICAL = 8
	DIRECTMANIPULATION_GESTURE_CROSS_SLIDE_HORIZONTAL = 16
	DIRECTMANIPULATION_GESTURE_PINCH_ZOOM = 32
)

type DIRECTMANIPULATION_MOTION_TYPES int32

const (
	DIRECTMANIPULATION_MOTION_NONE = 0
	DIRECTMANIPULATION_MOTION_TRANSLATEX = 1
	DIRECTMANIPULATION_MOTION_TRANSLATEY = 2
	DIRECTMANIPULATION_MOTION_ZOOM = 4
	DIRECTMANIPULATION_MOTION_CENTERX = 16
	DIRECTMANIPULATION_MOTION_CENTERY = 32
	DIRECTMANIPULATION_MOTION_ALL = 55
)

type DIRECTMANIPULATION_VIEWPORT_OPTIONS int32

const (
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_DEFAULT = 0
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_AUTODISABLE = 1
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_MANUALUPDATE = 2
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_INPUT = 4
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_EXPLICITHITTEST = 8
	DIRECTMANIPULATION_VIEWPORT_OPTIONS_DISABLEPIXELSNAPPING = 16
)

type DIRECTMANIPULATION_SNAPPOINT_TYPE int32

const (
	DIRECTMANIPULATION_SNAPPOINT_MANDATORY = 0
	DIRECTMANIPULATION_SNAPPOINT_OPTIONAL = 1
	DIRECTMANIPULATION_SNAPPOINT_MANDATORY_SINGLE = 2
	DIRECTMANIPULATION_SNAPPOINT_OPTIONAL_SINGLE = 3
)

type DIRECTMANIPULATION_SNAPPOINT_COORDINATE int32

const (
	DIRECTMANIPULATION_COORDINATE_BOUNDARY = 0
	DIRECTMANIPULATION_COORDINATE_ORIGIN = 1
	DIRECTMANIPULATION_COORDINATE_MIRRORED = 16
)

type DIRECTMANIPULATION_HORIZONTALALIGNMENT int32

const (
	DIRECTMANIPULATION_HORIZONTALALIGNMENT_NONE = 0
	DIRECTMANIPULATION_HORIZONTALALIGNMENT_LEFT = 1
	DIRECTMANIPULATION_HORIZONTALALIGNMENT_CENTER = 2
	DIRECTMANIPULATION_HORIZONTALALIGNMENT_RIGHT = 4
	DIRECTMANIPULATION_HORIZONTALALIGNMENT_UNLOCKCENTER = 8
)

type DIRECTMANIPULATION_VERTICALALIGNMENT int32

const (
	DIRECTMANIPULATION_VERTICALALIGNMENT_NONE = 0
	DIRECTMANIPULATION_VERTICALALIGNMENT_TOP = 1
	DIRECTMANIPULATION_VERTICALALIGNMENT_CENTER = 2
	DIRECTMANIPULATION_VERTICALALIGNMENT_BOTTOM = 4
	DIRECTMANIPULATION_VERTICALALIGNMENT_UNLOCKCENTER = 8
)

type DIRECTMANIPULATION_INPUT_MODE int32

const (
	DIRECTMANIPULATION_INPUT_MODE_AUTOMATIC = 0
	DIRECTMANIPULATION_INPUT_MODE_MANUAL = 1
)

type DIRECTMANIPULATION_DRAG_DROP_STATUS int32

const (
	DIRECTMANIPULATION_DRAG_DROP_READY = 0
	DIRECTMANIPULATION_DRAG_DROP_PRESELECT = 1
	DIRECTMANIPULATION_DRAG_DROP_SELECTING = 2
	DIRECTMANIPULATION_DRAG_DROP_DRAGGING = 3
	DIRECTMANIPULATION_DRAG_DROP_CANCELLED = 4
	DIRECTMANIPULATION_DRAG_DROP_COMMITTED = 5
)

type DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION int32

const (
	DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION_VERTICAL = 1
	DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION_HORIZONTAL = 2
	DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION_SELECT_ONLY = 16
	DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION_SELECT_DRAG = 32
	DIRECTMANIPULATION_DRAG_DROP_CONFIGURATION_HOLD_DRAG = 64
)

type DIRECTMANIPULATION_INTERACTION_TYPE int32

const (
	DIRECTMANIPULATION_INTERACTION_BEGIN = 0
	DIRECTMANIPULATION_INTERACTION_TYPE_MANIPULATION = 1
	DIRECTMANIPULATION_INTERACTION_TYPE_GESTURE_TAP = 2
	DIRECTMANIPULATION_INTERACTION_TYPE_GESTURE_HOLD = 3
	DIRECTMANIPULATION_INTERACTION_TYPE_GESTURE_CROSS_SLIDE = 4
	DIRECTMANIPULATION_INTERACTION_TYPE_GESTURE_PINCH_ZOOM = 5
	DIRECTMANIPULATION_INTERACTION_END = 100
)

type DIRECTMANIPULATION_AUTOSCROLL_CONFIGURATION int32

const (
	DIRECTMANIPULATION_AUTOSCROLL_CONFIGURATION_STOP = 0
	DIRECTMANIPULATION_AUTOSCROLL_CONFIGURATION_FORWARD = 1
	DIRECTMANIPULATION_AUTOSCROLL_CONFIGURATION_REVERSE = 2
)


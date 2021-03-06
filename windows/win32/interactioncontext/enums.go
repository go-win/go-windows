// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package interactioncontext implements the Windows.Win32.InteractionContext namespace.
package interactioncontext

type INTERACTION_ID int32

const (
	INTERACTION_ID_NONE          INTERACTION_ID = 0
	INTERACTION_ID_MANIPULATION  INTERACTION_ID = 1
	INTERACTION_ID_TAP           INTERACTION_ID = 2
	INTERACTION_ID_SECONDARY_TAP INTERACTION_ID = 3
	INTERACTION_ID_HOLD          INTERACTION_ID = 4
	INTERACTION_ID_DRAG          INTERACTION_ID = 5
	INTERACTION_ID_CROSS_SLIDE   INTERACTION_ID = 6
	INTERACTION_ID_MAX           INTERACTION_ID = -1
)

type INTERACTION_FLAGS int32

const (
	INTERACTION_FLAG_NONE    INTERACTION_FLAGS = 0
	INTERACTION_FLAG_BEGIN   INTERACTION_FLAGS = 1
	INTERACTION_FLAG_END     INTERACTION_FLAGS = 2
	INTERACTION_FLAG_CANCEL  INTERACTION_FLAGS = 4
	INTERACTION_FLAG_INERTIA INTERACTION_FLAGS = 8
	INTERACTION_FLAG_MAX     INTERACTION_FLAGS = -1
)

type INTERACTION_CONFIGURATION_FLAGS int32

const (
	INTERACTION_CONFIGURATION_FLAG_NONE                                 INTERACTION_CONFIGURATION_FLAGS = 0
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION                         INTERACTION_CONFIGURATION_FLAGS = 1
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION_TRANSLATION_X           INTERACTION_CONFIGURATION_FLAGS = 2
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION_TRANSLATION_Y           INTERACTION_CONFIGURATION_FLAGS = 4
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION_ROTATION                INTERACTION_CONFIGURATION_FLAGS = 8
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION_SCALING                 INTERACTION_CONFIGURATION_FLAGS = 16
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION_TRANSLATION_INERTIA     INTERACTION_CONFIGURATION_FLAGS = 32
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION_ROTATION_INERTIA        INTERACTION_CONFIGURATION_FLAGS = 64
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION_SCALING_INERTIA         INTERACTION_CONFIGURATION_FLAGS = 128
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION_RAILS_X                 INTERACTION_CONFIGURATION_FLAGS = 256
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION_RAILS_Y                 INTERACTION_CONFIGURATION_FLAGS = 512
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION_EXACT                   INTERACTION_CONFIGURATION_FLAGS = 1024
	INTERACTION_CONFIGURATION_FLAG_MANIPULATION_MULTIPLE_FINGER_PANNING INTERACTION_CONFIGURATION_FLAGS = 2048
	INTERACTION_CONFIGURATION_FLAG_CROSS_SLIDE                          INTERACTION_CONFIGURATION_FLAGS = 1
	INTERACTION_CONFIGURATION_FLAG_CROSS_SLIDE_HORIZONTAL               INTERACTION_CONFIGURATION_FLAGS = 2
	INTERACTION_CONFIGURATION_FLAG_CROSS_SLIDE_SELECT                   INTERACTION_CONFIGURATION_FLAGS = 4
	INTERACTION_CONFIGURATION_FLAG_CROSS_SLIDE_SPEED_BUMP               INTERACTION_CONFIGURATION_FLAGS = 8
	INTERACTION_CONFIGURATION_FLAG_CROSS_SLIDE_REARRANGE                INTERACTION_CONFIGURATION_FLAGS = 16
	INTERACTION_CONFIGURATION_FLAG_CROSS_SLIDE_EXACT                    INTERACTION_CONFIGURATION_FLAGS = 32
	INTERACTION_CONFIGURATION_FLAG_TAP                                  INTERACTION_CONFIGURATION_FLAGS = 1
	INTERACTION_CONFIGURATION_FLAG_TAP_DOUBLE                           INTERACTION_CONFIGURATION_FLAGS = 2
	INTERACTION_CONFIGURATION_FLAG_TAP_MULTIPLE_FINGER                  INTERACTION_CONFIGURATION_FLAGS = 4
	INTERACTION_CONFIGURATION_FLAG_SECONDARY_TAP                        INTERACTION_CONFIGURATION_FLAGS = 1
	INTERACTION_CONFIGURATION_FLAG_HOLD                                 INTERACTION_CONFIGURATION_FLAGS = 1
	INTERACTION_CONFIGURATION_FLAG_HOLD_MOUSE                           INTERACTION_CONFIGURATION_FLAGS = 2
	INTERACTION_CONFIGURATION_FLAG_HOLD_MULTIPLE_FINGER                 INTERACTION_CONFIGURATION_FLAGS = 4
	INTERACTION_CONFIGURATION_FLAG_DRAG                                 INTERACTION_CONFIGURATION_FLAGS = 1
	INTERACTION_CONFIGURATION_FLAG_MAX                                  INTERACTION_CONFIGURATION_FLAGS = -1
)

type INERTIA_PARAMETER int32

const (
	INERTIA_PARAMETER_TRANSLATION_DECELERATION INERTIA_PARAMETER = 1
	INERTIA_PARAMETER_TRANSLATION_DISPLACEMENT INERTIA_PARAMETER = 2
	INERTIA_PARAMETER_ROTATION_DECELERATION    INERTIA_PARAMETER = 3
	INERTIA_PARAMETER_ROTATION_ANGLE           INERTIA_PARAMETER = 4
	INERTIA_PARAMETER_EXPANSION_DECELERATION   INERTIA_PARAMETER = 5
	INERTIA_PARAMETER_EXPANSION_EXPANSION      INERTIA_PARAMETER = 6
	INERTIA_PARAMETER_MAX                      INERTIA_PARAMETER = -1
)

type INTERACTION_STATE int32

const (
	INTERACTION_STATE_IDLE                INTERACTION_STATE = 0
	INTERACTION_STATE_IN_INTERACTION      INTERACTION_STATE = 1
	INTERACTION_STATE_POSSIBLE_DOUBLE_TAP INTERACTION_STATE = 2
	INTERACTION_STATE_MAX                 INTERACTION_STATE = -1
)

type INTERACTION_CONTEXT_PROPERTY int32

const (
	INTERACTION_CONTEXT_PROPERTY_MEASUREMENT_UNITS       INTERACTION_CONTEXT_PROPERTY = 1
	INTERACTION_CONTEXT_PROPERTY_INTERACTION_UI_FEEDBACK INTERACTION_CONTEXT_PROPERTY = 2
	INTERACTION_CONTEXT_PROPERTY_FILTER_POINTERS         INTERACTION_CONTEXT_PROPERTY = 3
	INTERACTION_CONTEXT_PROPERTY_MAX                     INTERACTION_CONTEXT_PROPERTY = -1
)

type CROSS_SLIDE_THRESHOLD int32

const (
	CROSS_SLIDE_THRESHOLD_SELECT_START     CROSS_SLIDE_THRESHOLD = 0
	CROSS_SLIDE_THRESHOLD_SPEED_BUMP_START CROSS_SLIDE_THRESHOLD = 1
	CROSS_SLIDE_THRESHOLD_SPEED_BUMP_END   CROSS_SLIDE_THRESHOLD = 2
	CROSS_SLIDE_THRESHOLD_REARRANGE_START  CROSS_SLIDE_THRESHOLD = 3
	CROSS_SLIDE_THRESHOLD_COUNT            CROSS_SLIDE_THRESHOLD = 4
	CROSS_SLIDE_THRESHOLD_MAX              CROSS_SLIDE_THRESHOLD = -1
)

type CROSS_SLIDE_FLAGS int32

const (
	CROSS_SLIDE_FLAGS_NONE       CROSS_SLIDE_FLAGS = 0
	CROSS_SLIDE_FLAGS_SELECT     CROSS_SLIDE_FLAGS = 1
	CROSS_SLIDE_FLAGS_SPEED_BUMP CROSS_SLIDE_FLAGS = 2
	CROSS_SLIDE_FLAGS_REARRANGE  CROSS_SLIDE_FLAGS = 4
	CROSS_SLIDE_FLAGS_MAX        CROSS_SLIDE_FLAGS = -1
)

type MOUSE_WHEEL_PARAMETER int32

const (
	MOUSE_WHEEL_PARAMETER_CHAR_TRANSLATION_X MOUSE_WHEEL_PARAMETER = 1
	MOUSE_WHEEL_PARAMETER_CHAR_TRANSLATION_Y MOUSE_WHEEL_PARAMETER = 2
	MOUSE_WHEEL_PARAMETER_DELTA_SCALE        MOUSE_WHEEL_PARAMETER = 3
	MOUSE_WHEEL_PARAMETER_DELTA_ROTATION     MOUSE_WHEEL_PARAMETER = 4
	MOUSE_WHEEL_PARAMETER_PAGE_TRANSLATION_X MOUSE_WHEEL_PARAMETER = 5
	MOUSE_WHEEL_PARAMETER_PAGE_TRANSLATION_Y MOUSE_WHEEL_PARAMETER = 6
	MOUSE_WHEEL_PARAMETER_MAX                MOUSE_WHEEL_PARAMETER = -1
)

type TAP_PARAMETER int32

const (
	TAP_PARAMETER_MIN_CONTACT_COUNT TAP_PARAMETER = 0
	TAP_PARAMETER_MAX_CONTACT_COUNT TAP_PARAMETER = 1
	TAP_PARAMETER_MAX               TAP_PARAMETER = -1
)

type HOLD_PARAMETER int32

const (
	HOLD_PARAMETER_MIN_CONTACT_COUNT     HOLD_PARAMETER = 0
	HOLD_PARAMETER_MAX_CONTACT_COUNT     HOLD_PARAMETER = 1
	HOLD_PARAMETER_THRESHOLD_RADIUS      HOLD_PARAMETER = 2
	HOLD_PARAMETER_THRESHOLD_START_DELAY HOLD_PARAMETER = 3
	HOLD_PARAMETER_MAX                   HOLD_PARAMETER = -1
)

type TRANSLATION_PARAMETER int32

const (
	TRANSLATION_PARAMETER_MIN_CONTACT_COUNT TRANSLATION_PARAMETER = 0
	TRANSLATION_PARAMETER_MAX_CONTACT_COUNT TRANSLATION_PARAMETER = 1
	TRANSLATION_PARAMETER_MAX               TRANSLATION_PARAMETER = -1
)

type MANIPULATION_RAILS_STATE int32

const (
	MANIPULATION_RAILS_STATE_UNDECIDED MANIPULATION_RAILS_STATE = 0
	MANIPULATION_RAILS_STATE_FREE      MANIPULATION_RAILS_STATE = 1
	MANIPULATION_RAILS_STATE_RAILED    MANIPULATION_RAILS_STATE = 2
	MANIPULATION_RAILS_STATE_MAX       MANIPULATION_RAILS_STATE = -1
)

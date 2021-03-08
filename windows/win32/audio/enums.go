// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package audio implements the Windows.Win32.Audio namespace.
package audio

type APO_CONNECTION_BUFFER_TYPE int32

const (
	APO_CONNECTION_BUFFER_TYPE_ALLOCATED = 0
	APO_CONNECTION_BUFFER_TYPE_EXTERNAL = 1
	APO_CONNECTION_BUFFER_TYPE_DEPENDANT = 2
)

type APO_FLAG int32

const (
	APO_FLAG_NONE = 0
	APO_FLAG_INPLACE = 1
	APO_FLAG_SAMPLESPERFRAME_MUST_MATCH = 2
	APO_FLAG_FRAMESPERSECOND_MUST_MATCH = 4
	APO_FLAG_BITSPERSAMPLE_MUST_MATCH = 8
	APO_FLAG_MIXER = 16
	APO_FLAG_DEFAULT = 14
)

type AUDIO_FLOW_TYPE int32

const (
	AUDIO_FLOW_PULL = 0
	AUDIO_FLOW_PUSH = 1
)

type EAudioConstriction int32

const (
	eAudioConstrictionOff = 0
	eAudioConstriction48_16 = 1
	eAudioConstriction44_16 = 2
	eAudioConstriction14_14 = 3
	eAudioConstrictionMute = 4
)

type DMUS_CLOCKTYPE int32

const (
	DMUS_CLOCK_SYSTEM = 0
	DMUS_CLOCK_WAVE = 1
)

type KSPROPERTY_AUDIOEFFECTSDISCOVERY int32

const (
	KSPROPERTY_AUDIOEFFECTSDISCOVERY_EFFECTSLIST = 1
)

package winmd

// CorElementType is an enumeration of possible primitives.
type CorElementType uint8

// These values each represent a type of primitive.
const (
	CorElementTypeEnd     CorElementType = 0x00
	CorElementTypeVoid    CorElementType = 0x01
	CorElementTypeBool    CorElementType = 0x02
	CorElementTypeChar    CorElementType = 0x03
	CorElementTypeInt8    CorElementType = 0x04
	CorElementTypeUint8   CorElementType = 0x05
	CorElementTypeInt16   CorElementType = 0x06
	CorElementTypeUint16  CorElementType = 0x07
	CorElementTypeInt32   CorElementType = 0x08
	CorElementTypeUint32  CorElementType = 0x09
	CorElementTypeInt64   CorElementType = 0x0a
	CorElementTypeUint64  CorElementType = 0x0b
	CorElementTypeFloat32 CorElementType = 0x0c
	CorElementTypeFloat64 CorElementType = 0x0d
	CorElementTypeString  CorElementType = 0x0e

	CorElementTypePtr   CorElementType = 0x0f
	CorElementTypeByRef CorElementType = 0x10

	CorElementTypeValueType   CorElementType = 0x11
	CorElementTypeClass       CorElementType = 0x12
	CorElementTypeVar         CorElementType = 0x13
	CorElementTypeArray       CorElementType = 0x14
	CorElementTypeGenericInst CorElementType = 0x15
	CorElementTypeTypedByRef  CorElementType = 0x16

	CorElementTypeInt     CorElementType = 0x18
	CorElementTypeUint    CorElementType = 0x19
	CorElementTypeFnPtr   CorElementType = 0x1b
	CorElementTypeObject  CorElementType = 0x1c
	CorElementTypeSzArray CorElementType = 0x1d
	CorElementTypeMVar    CorElementType = 0x1e

	CorElementTypeCModReqD CorElementType = 0x1f
	CorElementTypeCModOpt  CorElementType = 0x20

	CorElementTypeInternal CorElementType = 0x21
	CorElementTypeMax      CorElementType = 0x22

	CorElementTypeModifier CorElementType = 0x40
	CorElementTypeSentinal CorElementType = CorElementTypeModifier | 0x01
	CorElementTypePinned   CorElementType = CorElementTypeModifier | 0x05
)

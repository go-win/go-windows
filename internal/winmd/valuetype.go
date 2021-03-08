package winmd

// ValueType is an enumeration of possible primitives.
type ValueType uint8

// These values each represent a type of primitive.
const (
	Void   ValueType = 0x01
	Bool   ValueType = 0x02
	Char   ValueType = 0x03
	I8     ValueType = 0x04
	U8     ValueType = 0x05
	I16    ValueType = 0x06
	U16    ValueType = 0x07
	I32    ValueType = 0x08
	U32    ValueType = 0x09
	I64    ValueType = 0x0a
	U64    ValueType = 0x0b
	F32    ValueType = 0x0c
	F64    ValueType = 0x0d
	String ValueType = 0x0e
	ISize  ValueType = 0x18
	USize  ValueType = 0x19
)

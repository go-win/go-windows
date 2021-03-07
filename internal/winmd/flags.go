package winmd

// MethodFlags are the flags on methods.
type MethodFlags uint16

// TypeFlags are the flags on types.
type TypeFlags uint32

// ParamFlags are the flags on params.
type ParamFlags uint16

// FieldFlags are the flags on fields.
type FieldFlags uint16

// Special returns if the method is 'special' (such as a constructor.)
func (f MethodFlags) Special() bool {
	return f&0b1000_0000_0000 != 0
}

// WinRT returns whether or not the type is a WinRT type.
func (f TypeFlags) WinRT() bool {
	return f&0b100_0000_0000_0000 != 0
}

// Interface returns whether or not the type is an interface.
func (f TypeFlags) Interface() bool {
	return f&0b10_0000 != 0
}

// Explicit returns whether or not the type has an explicit layout.
func (f TypeFlags) Explicit() bool {
	return f&0b1_0000 != 0
}

// Input returns true if the parameter is an input ([In]) parameter.
func (f ParamFlags) Input() bool {
	return f&0x0001 != 0
}

// Output returns true if the parameter is an output ([Out]) parameter.
func (f ParamFlags) Output() bool {
	return f&0x0002 != 0
}

// Optional returns true if the parameter is "optional".
func (f ParamFlags) Optional() bool {
	return f&0x0010 != 0
}

// Literal returns true if the field value is a compile time constant.
func (f FieldFlags) Literal() bool {
	return f&0b100_0000 != 0
}

// Static returns true if the field is per-type instead of per-instance.
func (f FieldFlags) Static() bool {
	return f&0b1_0000 != 0
}

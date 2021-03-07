package winmd

import (
	"bytes"
	"context"
	"io"
)

func readidx(ctx context.Context, size uint32, r io.Reader) (uint32, error) {
	offset := uint32(0)
	switch size {
	case 2:
		offset16 := uint16(0)
		if err := read(ctx, r, &offset16); err != nil {
			return 0, err
		}
		offset = uint32(offset16)
	case 4:
		offset32 := uint32(0)
		if err := read(ctx, r, &offset32); err != nil {
			return 0, err
		}
		offset = uint32(offset32)
	default:
		panic("unhandled index size")
	}
	return offset, nil
}

//
// Table Indices
//

// FieldIndex is an index into the field table.
type FieldIndex uint32

func (index *FieldIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.field, r)
	*index = FieldIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index FieldIndex) Lookup(f *File) *Field {
	if int(index) >= len(f.Field) {
		return nil
	}
	return &f.Field[index]
}

// MethodDefIndex is an index into the methoddef table.
type MethodDefIndex uint32

func (index *MethodDefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.methoddef, r)
	*index = MethodDefIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index MethodDefIndex) Lookup(f *File) *MethodDef {
	if int(index) >= len(f.MethodDef) {
		return nil
	}
	return &f.MethodDef[index]
}

// TypeDefIndex is an index into the typedef table.
type TypeDefIndex uint32

func (index *TypeDefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.typedef, r)
	*index = TypeDefIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index TypeDefIndex) Lookup(f *File) *TypeDef {
	if int(index) >= len(f.TypeDef) {
		return nil
	}
	return &f.TypeDef[index]
}

// ParamIndex is an index into the param table.
type ParamIndex uint32

func (index *ParamIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.param, r)
	*index = ParamIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index ParamIndex) Lookup(f *File) *Param {
	if int(index) >= len(f.Param) {
		return nil
	}
	return &f.Param[index]
}

// EventIndex is an index into the param table.
type EventIndex uint32

func (index *EventIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.event, r)
	*index = EventIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index EventIndex) Lookup(f *File) *Event {
	if int(index) >= len(f.Event) {
		return nil
	}
	return &f.Event[index]
}

// PropertyIndex is an index into the param table.
type PropertyIndex uint32

func (index *PropertyIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.property, r)
	*index = PropertyIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index PropertyIndex) Lookup(f *File) *Property {
	if int(index) >= len(f.Property) {
		return nil
	}
	return &f.Property[index]
}

// ModuleRefIndex is an index into the param table.
type ModuleRefIndex uint32

func (index *ModuleRefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.moduleref, r)
	*index = ModuleRefIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index ModuleRefIndex) Lookup(f *File) *ModuleRef {
	if int(index) >= len(f.ModuleRef) {
		return nil
	}
	return &f.ModuleRef[index]
}

// AssemblyRefIndex is an index into the param table.
type AssemblyRefIndex uint32

func (index *AssemblyRefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.assemblyref, r)
	*index = AssemblyRefIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index AssemblyRefIndex) Lookup(f *File) *AssemblyRef {
	if int(index) >= len(f.AssemblyRef) {
		return nil
	}
	return &f.AssemblyRef[index]
}

// GenericParamIndex is an index into the generic param table.
type GenericParamIndex uint32

func (index *GenericParamIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.genericparam, r)
	*index = GenericParamIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index GenericParamIndex) Lookup(f *File) *GenericParam {
	if int(index) >= len(f.GenericParam) {
		return nil
	}
	return &f.GenericParam[index]
}

//
// Composite Table Indices
//

// TypeDefOrRefIndex is a typedef or typeref index.
type TypeDefOrRefIndex uint32

func (index *TypeDefOrRefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.typedeforref, r)
	*index = TypeDefOrRefIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index TypeDefOrRefIndex) Lookup(f *File) interface{} {
	table := index & 0b11
	row := index >> 2
	switch table {
	case 0:
		if int(row) >= len(f.TypeDef) {
			return nil
		}
		return &f.TypeDef[index]
	case 1:
		if int(row) >= len(f.TypeRef) {
			return nil
		}
		return &f.TypeRef[index]
	case 2:
		if int(row) >= len(f.TypeSpec) {
			return nil
		}
		return &f.TypeSpec[index]
	default:
		panic("invalid table in composite index")
	}
}

// HasConstantIndex is a constant index.
type HasConstantIndex uint32

func (index *HasConstantIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.hasconstant, r)
	*index = HasConstantIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index HasConstantIndex) Lookup(f *File) interface{} {
	table := index & 0b11
	row := index >> 2
	switch table {
	case 0:
		if int(row) >= len(f.Field) {
			return nil
		}
		return &f.Field[index]
	case 1:
		if int(row) >= len(f.Param) {
			return nil
		}
		return &f.Param[index]
	case 2:
		if int(row) >= len(f.Property) {
			return nil
		}
		return &f.Property[index]
	default:
		panic("invalid table in composite index")
	}
}

// HasCustomAttributeIndex is a custom attribute index.
type HasCustomAttributeIndex uint32

func (index *HasCustomAttributeIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.hascustomattribute, r)
	*index = HasCustomAttributeIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index HasCustomAttributeIndex) Lookup(f *File) interface{} {
	table := index & 0b11111
	row := index >> 5
	switch table {
	case 0:
		if int(row) >= len(f.MethodDef) {
			return nil
		}
		return &f.MethodDef[index]
	case 1:
		if int(row) >= len(f.Field) {
			return nil
		}
		return &f.Field[index]
	case 2:
		if int(row) >= len(f.TypeRef) {
			return nil
		}
		return &f.TypeRef[index]
	case 3:
		if int(row) >= len(f.TypeDef) {
			return nil
		}
		return &f.TypeDef[index]
	case 4:
		if int(row) >= len(f.Param) {
			return nil
		}
		return &f.Param[index]
	case 5:
		if int(row) >= len(f.InterfaceImpl) {
			return nil
		}
		return &f.InterfaceImpl[index]
	case 6:
		if int(row) >= len(f.MemberRef) {
			return nil
		}
		return &f.MemberRef[index]
	case 7:
		if int(row) >= len(f.Module) {
			return nil
		}
		return &f.Module[index]
	case 8:
		if int(row) >= len(f.Property) {
			return nil
		}
		return &f.Property[index]
	case 9:
		if int(row) >= len(f.Event) {
			return nil
		}
		return &f.Event[index]
	case 10:
		if int(row) >= len(f.StandaloneSig) {
			return nil
		}
		return &f.StandaloneSig[index]
	case 11:
		if int(row) >= len(f.ModuleRef) {
			return nil
		}
		return &f.ModuleRef[index]
	case 12:
		if int(row) >= len(f.TypeSpec) {
			return nil
		}
		return &f.TypeSpec[index]
	case 13:
		if int(row) >= len(f.Assembly) {
			return nil
		}
		return &f.Assembly[index]
	case 14:
		if int(row) >= len(f.AssemblyRef) {
			return nil
		}
		return &f.AssemblyRef[index]
	case 15:
		if int(row) >= len(f.File) {
			return nil
		}
		return &f.File[index]
	case 16:
		if int(row) >= len(f.ExportedType) {
			return nil
		}
		return &f.ExportedType[index]
	case 17:
		if int(row) >= len(f.ManifestResource) {
			return nil
		}
		return &f.ManifestResource[index]
	case 18:
		if int(row) >= len(f.GenericParam) {
			return nil
		}
		return &f.GenericParam[index]
	case 19:
		if int(row) >= len(f.GenericParamConstraint) {
			return nil
		}
		return &f.GenericParamConstraint[index]
	case 20:
		if int(row) >= len(f.MethodSpec) {
			return nil
		}
		return &f.MethodSpec[index]
	default:
		panic("invalid table in composite index")
	}
}

// HasFieldMarshalIndex is a field marshal index.
type HasFieldMarshalIndex uint32

func (index *HasFieldMarshalIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.hasfieldmarshal, r)
	*index = HasFieldMarshalIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index HasFieldMarshalIndex) Lookup(f *File) interface{} {
	table := index & 0b1
	row := index >> 1
	switch table {
	case 0:
		if int(row) >= len(f.Field) {
			return nil
		}
		return &f.Field[index]
	case 1:
		if int(row) >= len(f.Param) {
			return nil
		}
		return &f.Param[index]
	default:
		panic("invalid table in composite index")
	}
}

// HasDeclSecurityIndex is a decl security index.
type HasDeclSecurityIndex uint32

func (index *HasDeclSecurityIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.hasdeclsecurity, r)
	*index = HasDeclSecurityIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index HasDeclSecurityIndex) Lookup(f *File) interface{} {
	table := index & 0b11
	row := index >> 2
	switch table {
	case 0:
		if int(row) >= len(f.TypeDef) {
			return nil
		}
		return &f.TypeDef[index]
	case 1:
		if int(row) >= len(f.MethodDef) {
			return nil
		}
		return &f.MethodDef[index]
	case 2:
		if int(row) >= len(f.Assembly) {
			return nil
		}
		return &f.Assembly[index]
	default:
		panic("invalid table in composite index")
	}
}

// MemberRefParentIndex is a member ref parent index.
type MemberRefParentIndex uint32

func (index *MemberRefParentIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.memberrefparent, r)
	*index = MemberRefParentIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index MemberRefParentIndex) Lookup(f *File) interface{} {
	table := index & 0b111
	row := index >> 3
	switch table {
	case 0:
		if int(row) >= len(f.TypeDef) {
			return nil
		}
		return &f.TypeDef[index]
	case 1:
		if int(row) >= len(f.TypeRef) {
			return nil
		}
		return &f.TypeRef[index]
	case 2:
		if int(row) >= len(f.ModuleRef) {
			return nil
		}
		return &f.ModuleRef[index]
	case 3:
		if int(row) >= len(f.MethodDef) {
			return nil
		}
		return &f.MethodDef[index]
	case 4:
		if int(row) >= len(f.TypeSpec) {
			return nil
		}
		return &f.TypeSpec[index]
	default:
		panic("invalid table in composite index")
	}
}

// HasSemanticsIndex is a semantic index.
type HasSemanticsIndex uint32

func (index *HasSemanticsIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.hassemantics, r)
	*index = HasSemanticsIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index HasSemanticsIndex) Lookup(f *File) interface{} {
	table := index & 0b1
	row := index >> 1
	switch table {
	case 0:
		if int(row) >= len(f.Event) {
			return nil
		}
		return &f.Event[index]
	case 1:
		if int(row) >= len(f.Property) {
			return nil
		}
		return &f.Property[index]
	default:
		panic("invalid table in composite index")
	}
}

// MethodDefOrRefIndex is a method def or method ref index.
type MethodDefOrRefIndex uint32

func (index *MethodDefOrRefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.methoddeforref, r)
	*index = MethodDefOrRefIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index MethodDefOrRefIndex) Lookup(f *File) interface{} {
	table := index & 0b1
	row := index >> 1
	switch table {
	case 0:
		if int(row) >= len(f.MethodDef) {
			return nil
		}
		return &f.MethodDef[index]
	case 1:
		if int(row) >= len(f.MemberRef) {
			return nil
		}
		return &f.MemberRef[index]
	default:
		panic("invalid table in composite index")
	}
}

// MemberForwardedIndex is a member forwarded index.
type MemberForwardedIndex uint32

func (index *MemberForwardedIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.memberforwarded, r)
	*index = MemberForwardedIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index MemberForwardedIndex) Lookup(f *File) interface{} {
	table := index & 0b1
	row := index >> 1
	switch table {
	case 0:
		if int(row) >= len(f.Field) {
			return nil
		}
		return &f.Field[index]
	case 1:
		if int(row) >= len(f.MethodDef) {
			return nil
		}
		return &f.MethodDef[index]
	default:
		panic("invalid table in composite index")
	}
}

// ImplementationIndex is an implementation index.
type ImplementationIndex uint32

func (index *ImplementationIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.implementation, r)
	*index = ImplementationIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index ImplementationIndex) Lookup(f *File) interface{} {
	table := index & 0b11
	row := index >> 2
	switch table {
	case 0:
		if int(row) >= len(f.File) {
			return nil
		}
		return &f.File[index]
	case 1:
		if int(row) >= len(f.AssemblyRef) {
			return nil
		}
		return &f.AssemblyRef[index]
	case 2:
		if int(row) >= len(f.ExportedType) {
			return nil
		}
		return &f.ExportedType[index]
	default:
		panic("invalid table in composite index")
	}
}

// CustomAttributeTypeIndex is a custom attribute type index.
type CustomAttributeTypeIndex uint32

func (index *CustomAttributeTypeIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.customattributetype, r)
	*index = CustomAttributeTypeIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index CustomAttributeTypeIndex) Lookup(f *File) interface{} {
	table := index & 0b111
	row := index >> 3
	switch table {
	case 0:
		if int(row) >= len(f.MethodDef) {
			return nil
		}
		return &f.MethodDef[index]
	case 1:
		if int(row) >= len(f.MemberRef) {
			return nil
		}
		return &f.MemberRef[index]
	default:
		panic("invalid table in composite index")
	}
}

// ResolutionScopeIndex is a resolution scope index.
type ResolutionScopeIndex uint32

func (index *ResolutionScopeIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.resolutionscope, r)
	*index = ResolutionScopeIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index ResolutionScopeIndex) Lookup(f *File) interface{} {
	table := index & 0b11
	row := index >> 2
	switch table {
	case 0:
		if int(row) >= len(f.Module) {
			return nil
		}
		return &f.Module[index]
	case 1:
		if int(row) >= len(f.ModuleRef) {
			return nil
		}
		return &f.ModuleRef[index]
	case 2:
		if int(row) >= len(f.AssemblyRef) {
			return nil
		}
		return &f.AssemblyRef[index]
	case 3:
		if int(row) >= len(f.TypeRef) {
			return nil
		}
		return &f.TypeRef[index]
	default:
		panic("invalid table in composite index")
	}
}

// TypeOrMethodDefIndex is a type or method def index.
type TypeOrMethodDefIndex uint32

func (index *TypeOrMethodDefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.typeormethoddef, r)
	*index = TypeOrMethodDefIndex(i)
	return err
}

// Lookup performs a lookup for this index against assembly f.
func (index TypeOrMethodDefIndex) Lookup(f *File) interface{} {
	table := index & 0b1
	row := index >> 1
	switch table {
	case 0:
		if int(row) >= len(f.TypeDef) {
			return nil
		}
		return &f.TypeDef[index]
	case 1:
		if int(row) >= len(f.MethodDef) {
			return nil
		}
		return &f.MethodDef[index]
	default:
		panic("invalid table in composite index")
	}
}

//
// Data Stream Indices
//

// IndexedString is a field in a WinMD that is a string indexed from the #Strings stream.
type IndexedString string

func (s *IndexedString) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	offset, err := readidx(ctx, ctxdata.stringindexsize, r)
	if err != nil {
		return err
	}
	len := uint32(bytes.IndexByte(ctxdata.stringdata[offset:], 0))
	*s = IndexedString(ctxdata.stringdata[offset : offset+len])
	return nil
}

// IndexedGUID is a field in a WinMD that is indexed from the #GUID stream.
type IndexedGUID [16]byte

func (guid *IndexedGUID) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	offset, err := readidx(ctx, ctxdata.guidindexsize, r)
	if err != nil {
		return err
	}
	copy((*guid)[:], ctxdata.guiddata[offset:])
	return nil
}

// IndexedBlob is a field in WinMD that is indexed from the #Blob stream.
type IndexedBlob []byte

func (blob *IndexedBlob) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	offset, err := readidx(ctx, ctxdata.blobindexsize, r)
	if err != nil {
		return err
	}
	size := uint32(ctxdata.blobdata[offset])
	*blob = IndexedBlob(ctxdata.blobdata[offset+1 : offset+1+size])
	return nil
}

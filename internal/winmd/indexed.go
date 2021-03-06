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

// MethodDefIndex is an index into the methoddef table.
type MethodDefIndex uint32

func (index *MethodDefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.methoddef, r)
	*index = MethodDefIndex(i)
	return err
}

// TypeDefIndex is an index into the typedef table.
type TypeDefIndex uint32

func (index *TypeDefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.typedef, r)
	*index = TypeDefIndex(i)
	return err
}

// ParamIndex is an index into the param table.
type ParamIndex uint32

func (index *ParamIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.param, r)
	*index = ParamIndex(i)
	return err
}

// EventIndex is an index into the param table.
type EventIndex uint32

func (index *EventIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.event, r)
	*index = EventIndex(i)
	return err
}

// PropertyIndex is an index into the param table.
type PropertyIndex uint32

func (index *PropertyIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.property, r)
	*index = PropertyIndex(i)
	return err
}

// ModuleRefIndex is an index into the param table.
type ModuleRefIndex uint32

func (index *ModuleRefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.moduleref, r)
	*index = ModuleRefIndex(i)
	return err
}

// AssemblyRefIndex is an index into the param table.
type AssemblyRefIndex uint32

func (index *AssemblyRefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.assemblyref, r)
	*index = AssemblyRefIndex(i)
	return err
}

// GenericParamIndex is an index into the generic param table.
type GenericParamIndex uint32

func (index *GenericParamIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.genericparam, r)
	*index = GenericParamIndex(i)
	return err
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

// HasConstantIndex is a constant index.
type HasConstantIndex uint32

func (index *HasConstantIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.hasconstant, r)
	*index = HasConstantIndex(i)
	return err
}

// HasCustomAttributeIndex is a custom attribute index.
type HasCustomAttributeIndex uint32

func (index *HasCustomAttributeIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.hascustomattribute, r)
	*index = HasCustomAttributeIndex(i)
	return err
}

// HasFieldMarshalIndex is a field marshal index.
type HasFieldMarshalIndex uint32

func (index *HasFieldMarshalIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.hasfieldmarshal, r)
	*index = HasFieldMarshalIndex(i)
	return err
}

// HasDeclSecurityIndex is a decl security index.
type HasDeclSecurityIndex uint32

func (index *HasDeclSecurityIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.hasdeclsecurity, r)
	*index = HasDeclSecurityIndex(i)
	return err
}

// MemberRefParentIndex is a member ref parent index.
type MemberRefParentIndex uint32

func (index *MemberRefParentIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.memberrefparent, r)
	*index = MemberRefParentIndex(i)
	return err
}

// HasSemanticsIndex is a semantic index.
type HasSemanticsIndex uint32

func (index *HasSemanticsIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.hassemantics, r)
	*index = HasSemanticsIndex(i)
	return err
}

// MethodDefOrRefIndex is a method def or method ref index.
type MethodDefOrRefIndex uint32

func (index *MethodDefOrRefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.methoddeforref, r)
	*index = MethodDefOrRefIndex(i)
	return err
}

// MemberForwardedIndex is a member forwarded index.
type MemberForwardedIndex uint32

func (index *MemberForwardedIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.memberforwarded, r)
	*index = MemberForwardedIndex(i)
	return err
}

// ImplementationIndex is an implementation index.
type ImplementationIndex uint32

func (index *ImplementationIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.implementation, r)
	*index = ImplementationIndex(i)
	return err
}

// CustomAttributeTypeIndex is a custom attribute type index.
type CustomAttributeTypeIndex uint32

func (index *CustomAttributeTypeIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.customattributetype, r)
	*index = CustomAttributeTypeIndex(i)
	return err
}

// ResolutionScopeIndex is a resolution scope index.
type ResolutionScopeIndex uint32

func (index *ResolutionScopeIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.resolutionscope, r)
	*index = ResolutionScopeIndex(i)
	return err
}

// TypeOrMethodDefIndex is a type or method def index.
type TypeOrMethodDefIndex uint32

func (index *TypeOrMethodDefIndex) unpack(ctx context.Context, r io.Reader) error {
	ctxdata := ctx.Value(ctxdatakey).(ctxdata)
	i, err := readidx(ctx, ctxdata.typeormethoddef, r)
	*index = TypeOrMethodDefIndex(i)
	return err
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

// Package winmd implements parsing for the WinMD file format. WinMD files are
// essentially .NET assemblies, and as such this package can parse the metadata
// out of theoretically any .NET assembly.
package winmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/bits"
)

// File represents a WinMD file in memory.
type File struct {
	Module                 []Module
	TypeRef                []TypeRef
	TypeDef                []TypeDef
	Field                  []Field
	MethodDef              []MethodDef
	Param                  []Param
	InterfaceImpl          []InterfaceImpl
	MemberRef              []MemberRef
	Constant               []Constant
	CustomAttribute        []CustomAttribute
	FieldMarshal           []FieldMarshal
	DeclSecurity           []DeclSecurity
	ClassLayout            []ClassLayout
	FieldLayout            []FieldLayout
	StandaloneSig          []StandaloneSig
	EventMap               []EventMap
	Event                  []Event
	PropertyMap            []PropertyMap
	Property               []Property
	MethodSemantics        []MethodSemantics
	MethodImpl             []MethodImpl
	ModuleRef              []ModuleRef
	TypeSpec               []TypeSpec
	ImplMap                []ImplMap
	FieldRVA               []FieldRVA
	Assembly               []Assembly
	AssemblyProcessor      []AssemblyProcessor
	AssemblyOS             []AssemblyOS
	AssemblyRef            []AssemblyRef
	AssemblyRefProcessor   []AssemblyRefProcessor
	AssemblyRefOS          []AssemblyRefOS
	FileRow                []FileRow
	ExportedType           []ExportedType
	ManifestResource       []ManifestResource
	NestedClass            []NestedClass
	GenericParam           []GenericParam
	MethodSpec             []MethodSpec
	GenericParamConstraint []GenericParamConstraint
}

// Load loads a WinMD file from a readseeker.
func Load(r io.ReadSeeker) (*File, error) {
	ctx := context.Background()
	f := new(File)

	// Read MZ header.
	mz := mz{}
	if err := read(ctx, r, &mz); err != nil {
		return nil, fmt.Errorf("reading MZ header: %w", err)
	}
	if mz.Signature != mzsig {
		return nil, fmt.Errorf("expected MZ signature %04x, got %04x", mzsig, mz.Signature)
	}

	// Read PE header.
	pe := pe{}
	if _, err := r.Seek(int64(mz.PeOffset), io.SeekStart); err != nil {
		return nil, fmt.Errorf("seeking to PE header at 0x%08x: %w", mz.PeOffset, err)
	}
	if err := read(ctx, r, &pe); err != nil {
		return nil, fmt.Errorf("reading PE header: %w", err)
	}
	if pe.Signature != pesig {
		return nil, fmt.Errorf("expected PE signature %08x, got %08x", pesig, pe.Signature)
	}
	if _, err := r.Seek(int64(pe.OptSize), io.SeekCurrent); err != nil {
		return nil, fmt.Errorf("seeking past optional headers of size 0x%08x: %w", pe.OptSize, err)
	}
	sections := make([]section, pe.NumSection)
	for i := 0; i < int(pe.NumSection); i++ {
		if err := read(ctx, r, &sections[i]); err != nil {
			return nil, err
		}
	}

	// Read COM data directory.
	netdir := datadir{}
	netdiroff := int64(mz.PeOffset) + 0xE8
	if _, err := r.Seek(netdiroff, io.SeekStart); err != nil {
		return nil, fmt.Errorf("seeking to .NET data directory at 0x%08x: %w", netdiroff, err)
	}
	if err := read(ctx, r, &netdir); err != nil {
		return nil, fmt.Errorf("reading .NET data directory: %w", err)
	}

	// Read COR ('Common Object Runtime') header.
	corhead := corhead{}
	corheadoff, err := rva2off(sections, netdir.RVA)
	if err != nil {
		return nil, fmt.Errorf("finding COR header offset: %w", err)
	}
	if _, err := r.Seek(int64(corheadoff), io.SeekStart); err != nil {
		return nil, fmt.Errorf("seeking to COR header at 0x%08x: %w", corheadoff, err)
	}
	if err := read(ctx, r, &corhead); err != nil {
		return nil, fmt.Errorf("reading COR header: %w", err)
	}

	// Read metadata header.
	metahead := metahead{}
	metaheadoff, err := rva2off(sections, corhead.Metadata.RVA)
	if err != nil {
		return nil, fmt.Errorf("finding metadata header offset: %w", err)
	}
	if _, err := r.Seek(int64(metaheadoff), io.SeekStart); err != nil {
		return nil, fmt.Errorf("seeking to metadata header at 0x%08x: %w", metaheadoff, err)
	}
	if err := read(ctx, r, &metahead); err != nil {
		return nil, fmt.Errorf("reading metadata header: %w", err)
	}

	// Stream entries
	tablesoff := int64(0)
	stringsoff := int64(0)
	stringslen := int64(0)
	guidlen := int64(0)
	guidoff := int64(0)
	bloblen := int64(0)
	bloboff := int64(0)
	for i := 0; i < int(metahead.NumberOfStreams); i++ {
		streamentry := streamentry{}
		if err := read(ctx, r, &streamentry); err != nil {
			return nil, err
		}
		switch string(streamentry.Name) {
		case "#~":
			tablesoff = int64(metaheadoff + streamentry.Offset)
		case "#Strings":
			stringsoff = int64(metaheadoff + streamentry.Offset)
			stringslen = int64(streamentry.Size)
		case "#Blob":
			bloboff = int64(metaheadoff + streamentry.Offset)
			bloblen = int64(streamentry.Size)
		case "#GUID":
			guidoff = int64(metaheadoff + streamentry.Offset)
			guidlen = int64(streamentry.Size)
		}
	}

	ctxdata := ctxdata{}

	// Read string data
	ctxdata.stringdata = make([]byte, stringslen)
	if _, err := r.Seek(int64(stringsoff), io.SeekStart); err != nil {
		return nil, fmt.Errorf("seeking to string stream at 0x%08x: %w", stringsoff, err)
	}
	if err := read(ctx, r, ctxdata.stringdata); err != nil {
		return nil, fmt.Errorf("reading string stream: %w", err)
	}

	// Read blob data
	ctxdata.guiddata = make([]byte, guidlen)
	if _, err := r.Seek(int64(guidoff), io.SeekStart); err != nil {
		return nil, fmt.Errorf("seeking to blob stream at 0x%08x: %w", guidoff, err)
	}
	if err := read(ctx, r, ctxdata.guiddata); err != nil {
		return nil, fmt.Errorf("reading blob stream: %w", err)
	}

	// Read blob data
	ctxdata.blobdata = make([]byte, bloblen)
	if _, err := r.Seek(int64(bloboff), io.SeekStart); err != nil {
		return nil, fmt.Errorf("seeking to blob stream at 0x%08x: %w", bloboff, err)
	}
	if err := read(ctx, r, ctxdata.blobdata); err != nil {
		return nil, fmt.Errorf("reading blob stream: %w", err)
	}

	// Read tables header
	tableheader := tableheader{}
	if _, err := r.Seek(int64(tablesoff), io.SeekStart); err != nil {
		return nil, fmt.Errorf("seeking to table header at 0x%08x: %w", tablesoff, err)
	}
	if err := read(ctx, r, &tableheader); err != nil {
		return nil, fmt.Errorf("reading table header: %w", err)
	}

	// Read table row counts.
	rows := rowcounts{}
	for i := 0; i < 64; i++ {
		if (int(tableheader.MaskValid) >> i & 1) == 0 {
			continue
		}

		rowcount := uint32(0)
		if err := read(ctx, r, &rowcount); err != nil {
			return nil, fmt.Errorf("reading table rowcount: %w", err)
		}

		switch i {
		case 0x00:
			rows.module = rowcount
		case 0x01:
			rows.typeref = rowcount
		case 0x02:
			rows.typedef = rowcount
		case 0x04:
			rows.field = rowcount
		case 0x06:
			rows.methoddef = rowcount
		case 0x08:
			rows.param = rowcount
		case 0x09:
			rows.interfaceimpl = rowcount
		case 0x0a:
			rows.memberref = rowcount
		case 0x0b:
			rows.constant = rowcount
		case 0x0c:
			rows.customattribute = rowcount
		case 0x0d:
			rows.fieldmarshal = rowcount
		case 0x0e:
			rows.declsecurity = rowcount
		case 0x0f:
			rows.classlayout = rowcount
		case 0x10:
			rows.fieldlayout = rowcount
		case 0x11:
			rows.standalonesig = rowcount
		case 0x12:
			rows.eventmap = rowcount
		case 0x14:
			rows.event = rowcount
		case 0x15:
			rows.propertymap = rowcount
		case 0x17:
			rows.property = rowcount
		case 0x18:
			rows.methodsemantics = rowcount
		case 0x19:
			rows.methodimpl = rowcount
		case 0x1a:
			rows.moduleref = rowcount
		case 0x1b:
			rows.typespec = rowcount
		case 0x1c:
			rows.implmap = rowcount
		case 0x1d:
			rows.fieldrva = rowcount
		case 0x20:
			rows.assembly = rowcount
		case 0x21:
			rows.assemblyprocessor = rowcount
		case 0x22:
			rows.assemblyos = rowcount
		case 0x23:
			rows.assemblyref = rowcount
		case 0x24:
			rows.assemblyrefprocessor = rowcount
		case 0x25:
			rows.assemblyrefos = rowcount
		case 0x26:
			rows.file = rowcount
		case 0x27:
			rows.exportedtype = rowcount
		case 0x28:
			rows.manifestresource = rowcount
		case 0x29:
			rows.nestedclass = rowcount
		case 0x2a:
			rows.genericparam = rowcount
		case 0x2b:
			rows.methodspec = rowcount
		case 0x2c:
			rows.genericparamconstraint = rowcount
		default:
			panic("unknown table")
		}
	}

	// Calculate table index sizes
	ctxdata.field = indexsize(rows.field)
	ctxdata.methoddef = indexsize(rows.methoddef)
	ctxdata.typedef = indexsize(rows.typedef)
	ctxdata.param = indexsize(rows.param)
	ctxdata.event = indexsize(rows.event)
	ctxdata.property = indexsize(rows.property)
	ctxdata.moduleref = indexsize(rows.moduleref)
	ctxdata.assemblyref = indexsize(rows.assemblyref)
	ctxdata.genericparam = indexsize(rows.genericparam)

	// Calculate composite index sizes
	ctxdata.typedeforref = indexsize(rows.typedef, rows.typeref, rows.typespec)
	ctxdata.hasconstant = indexsize(rows.field, rows.param, rows.property)
	ctxdata.hascustomattribute = indexsize(rows.methoddef, rows.field, rows.typeref, rows.typedef, rows.param, rows.interfaceimpl, rows.memberref, rows.module, rows.property, rows.event, rows.standalonesig, rows.moduleref, rows.typespec, rows.assembly, rows.assemblyref, rows.file, rows.exportedtype, rows.manifestresource, rows.genericparam, rows.genericparamconstraint, rows.methodspec)
	ctxdata.hasfieldmarshal = indexsize(rows.field, rows.param)
	ctxdata.hasdeclsecurity = indexsize(rows.typedef, rows.methoddef, rows.assembly)
	ctxdata.memberrefparent = indexsize(rows.typedef, rows.typeref, rows.moduleref, rows.methoddef, rows.typespec)
	ctxdata.hassemantics = indexsize(rows.event, rows.property)
	ctxdata.methoddeforref = indexsize(rows.methoddef, rows.memberref)
	ctxdata.memberforwarded = indexsize(rows.field, rows.methoddef)
	ctxdata.implementation = indexsize(rows.file, rows.assemblyref, rows.exportedtype)
	ctxdata.customattributetype = indexsize(rows.methoddef, rows.memberref, 0, 0, 0)
	ctxdata.resolutionscope = indexsize(rows.module, rows.moduleref, rows.assemblyref, rows.typeref)
	ctxdata.typeormethoddef = indexsize(rows.typedef, rows.methoddef)

	// Determine data stream index sizes
	ctxdata.stringindexsize = uint32(2)
	ctxdata.guidindexsize = uint32(2)
	ctxdata.blobindexsize = uint32(2)
	if (tableheader.HeapOffsetSizes>>0)&1 == 1 {
		ctxdata.stringindexsize = 4
	}
	if (tableheader.HeapOffsetSizes>>1)&1 == 1 {
		ctxdata.guidindexsize = 4
	}
	if (tableheader.HeapOffsetSizes>>2)&1 == 1 {
		ctxdata.blobindexsize = 4
	}

	// Setup context information (for propagating index sizes)
	ctx = context.WithValue(ctx, ctxdatakey, ctxdata)

	// Allocate table rows.
	f.Module = make([]Module, rows.module)
	f.TypeRef = make([]TypeRef, rows.typeref)
	f.TypeDef = make([]TypeDef, rows.typedef)
	f.Field = make([]Field, rows.field)
	f.MethodDef = make([]MethodDef, rows.methoddef)
	f.Param = make([]Param, rows.param)
	f.InterfaceImpl = make([]InterfaceImpl, rows.interfaceimpl)
	f.MemberRef = make([]MemberRef, rows.memberref)
	f.Constant = make([]Constant, rows.constant)
	f.CustomAttribute = make([]CustomAttribute, rows.customattribute)
	f.FieldMarshal = make([]FieldMarshal, rows.fieldmarshal)
	f.DeclSecurity = make([]DeclSecurity, rows.declsecurity)
	f.ClassLayout = make([]ClassLayout, rows.classlayout)
	f.FieldLayout = make([]FieldLayout, rows.fieldlayout)
	f.StandaloneSig = make([]StandaloneSig, rows.standalonesig)
	f.EventMap = make([]EventMap, rows.eventmap)
	f.Event = make([]Event, rows.event)
	f.PropertyMap = make([]PropertyMap, rows.propertymap)
	f.Property = make([]Property, rows.property)
	f.MethodSemantics = make([]MethodSemantics, rows.methodsemantics)
	f.MethodImpl = make([]MethodImpl, rows.methodimpl)
	f.ModuleRef = make([]ModuleRef, rows.moduleref)
	f.TypeSpec = make([]TypeSpec, rows.typespec)
	f.ImplMap = make([]ImplMap, rows.implmap)
	f.FieldRVA = make([]FieldRVA, rows.fieldrva)
	f.Assembly = make([]Assembly, rows.assembly)
	f.AssemblyProcessor = make([]AssemblyProcessor, rows.assemblyprocessor)
	f.AssemblyOS = make([]AssemblyOS, rows.assemblyos)
	f.AssemblyRef = make([]AssemblyRef, rows.assemblyref)
	f.AssemblyRefProcessor = make([]AssemblyRefProcessor, rows.assemblyrefprocessor)
	f.AssemblyRefOS = make([]AssemblyRefOS, rows.assemblyrefos)
	f.FileRow = make([]FileRow, rows.file)
	f.ExportedType = make([]ExportedType, rows.exportedtype)
	f.ManifestResource = make([]ManifestResource, rows.manifestresource)
	f.NestedClass = make([]NestedClass, rows.nestedclass)
	f.GenericParam = make([]GenericParam, rows.genericparam)
	f.MethodSpec = make([]MethodSpec, rows.methodspec)
	f.GenericParamConstraint = make([]GenericParamConstraint, rows.genericparamconstraint)

	// Read tables. Using a buffered reader improves performance significantly
	// (~4x in my testing.)
	if err := read(ctx, bufio.NewReader(r), f); err != nil {
		return nil, err
	}

	return f, nil
}

// Converts an RVA into a file offset.
func rva2off(sections []section, rva uint32) (uint32, error) {
	for _, s := range sections {
		if rva >= s.VirtualAddr && rva < s.VirtualAddr+s.VirtualSize {
			return rva - s.VirtualAddr + s.DataOffset, nil
		}
	}
	return 0, fmt.Errorf("could not find section for rva: 0x%08x", rva)
}

// Calculates an index size. When multiple paramters are present, calculates a
// composite index size.
func indexsize(rowcounts ...uint32) uint32 {
	bits := bits.Len(uint(len(rowcounts) - 1))
	for _, rowcount := range rowcounts {
		if uint64(rowcount) >= uint64(1<<(16-bits)) {
			return 4
		}
	}
	return 2
}

// mz represents the DOS MZ stub header.
type mz struct {
	Signature uint16
	Unused    [58]byte
	PeOffset  uint32
}

// mzsig is the valid value for the DOS MZ header signature.
const mzsig = uint16(0x5a4d)

// pe represents the PE EXE stub header.
type pe struct {
	Signature       uint32
	Machine         uint16
	NumSection      uint16
	TimeDate        uint32
	SymTab          uint32
	NumSym          uint32
	OptSize         uint16
	Characteristics uint16
}

// pesig is the valid value for the PE EXE header signature.
const pesig = uint32(0x4550)

// section is a PE section header.
type section struct {
	Name            [8]byte
	VirtualSize     uint32
	VirtualAddr     uint32
	SizeRaw         uint32
	DataOffset      uint32
	RelocOffset     uint32
	LineNoOffset    uint32
	NumReloc        uint16
	NumLineNo       uint16
	Characteristics uint32
}

// datadir is a PE data directory.
type datadir struct {
	RVA  uint32
	Size uint32
}

// corhead is the COR (common object runtime) header.
type corhead struct {
	HeaderSize             uint32
	MajorRuntimeVer        uint16
	MinorRuntimeVer        uint16
	Metadata               datadir
	Flags                  uint32
	Entrypoint             uint32
	Resources              datadir
	StrongNameSignature    datadir
	CodeManagerTable       datadir
	VTableFixups           datadir
	ExportAddressTableJmps datadir
	ManagedNativeHeader    datadir
}

// metahead is the .NET BSJB metadata header.
type metahead struct {
	Signature       uint32
	MajorVersion    uint16
	MinorVersion    uint16
	Reserved        uint32
	Version         pstr32
	Flags           uint16
	NumberOfStreams uint16
}

// streamentry is a CLR stream header
type streamentry struct {
	Offset uint32
	Size   uint32
	Name   padstr
}

// tableheader is the header for the CLR tables stream
type tableheader struct {
	Reserved1       uint32
	MajorVersion    byte
	MinorVersion    byte
	HeapOffsetSizes byte
	Reserved2       byte
	MaskValid       uint64
	MaskSorted      uint64
}

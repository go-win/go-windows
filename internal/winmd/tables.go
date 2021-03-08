package winmd

// Module is a row in the module table.
type Module struct {
	Generation uint16
	Name       IndexedString
	MVID       IndexedGUID
	EncID      IndexedGUID
	EncBaseID  IndexedGUID
}

// TypeRef is a row in the type ref table.
type TypeRef struct {
	ResolutionScope ResolutionScopeIndex
	Name            IndexedString
	Namespace       IndexedString
}

// TypeDef is a row in the type def table.
type TypeDef struct {
	Flags      TypeFlags
	Name       IndexedString
	Namespace  IndexedString
	Extends    TypeDefOrRefIndex
	FieldList  FieldIndex
	MethodList MethodDefIndex
}

// Field is a row in the field table.
type Field struct {
	Flags     FieldFlags
	Name      IndexedString
	Signature IndexedBlob
}

// MethodDef is a row in the method table.
type MethodDef struct {
	RVA       uint32
	ImplFlags uint16
	Flags     MethodFlags
	Name      IndexedString
	Signature IndexedBlob
	ParamList ParamIndex
}

// Param is a row in the param table.
type Param struct {
	Flags    ParamFlags
	Sequence uint16
	Name     IndexedString
}

// InterfaceImpl is a row in the interface impl table.
type InterfaceImpl struct {
	Class     TypeDefIndex
	Interface TypeDefOrRefIndex
}

// MemberRef is a row in the member ref table.
type MemberRef struct {
	Class     MemberRefParentIndex
	Name      IndexedString
	Signature IndexedBlob
}

// Constant is a row in the constant table.
type Constant struct {
	Type        ValueType
	PaddingZero byte
	Parent      HasConstantIndex
	Value       IndexedBlob
}

// CustomAttribute is a row in the custom attribute table.
type CustomAttribute struct {
	Parent HasCustomAttributeIndex
	Type   CustomAttributeTypeIndex
	Value  IndexedBlob
}

// FieldMarshal is a row in the field marshal table.
type FieldMarshal struct {
	Parent     HasFieldMarshalIndex
	NativeType IndexedBlob
}

// DeclSecurity is a row in the declsecurity table.
type DeclSecurity struct {
	Action        uint16
	Parent        HasDeclSecurityIndex
	PermissionSet IndexedBlob
}

// ClassLayout is a row in the class layout table.
type ClassLayout struct {
	PackingSize uint16
	ClassSize   uint32
	Parent      TypeDefIndex
}

// FieldLayout is a row in the field layout table.
type FieldLayout struct {
	Offset uint32
	Field  FieldIndex
}

// StandaloneSig is a row in the standalone signature table.
type StandaloneSig struct {
	Signature IndexedBlob
}

// EventMap is a row in the event map table.
type EventMap struct {
	Parent    TypeDefIndex
	EventList EventIndex
}

// Event is a row in the event table.
type Event struct {
	EventFlags uint16
	Name       IndexedString
	EventType  TypeDefOrRefIndex
}

// PropertyMap is a row in the property map table.
type PropertyMap struct {
	Parent       TypeDefIndex
	PropertyList PropertyIndex
}

// Property is a row in the property table.
type Property struct {
	Flags uint16
	Name  IndexedString
	Type  IndexedBlob
}

// MethodSemantics is a row in the method semantics table.
type MethodSemantics struct {
	Semantics   uint16
	Method      MethodDefIndex
	Association HasSemanticsIndex
}

// MethodImpl is a row in the method impl table.
type MethodImpl struct {
	Class             TypeDefIndex
	MethodBody        MethodDefOrRefIndex
	MethodDeclaration MethodDefOrRefIndex
}

// ModuleRef is a row in the module ref table.
type ModuleRef struct {
	Name IndexedString
}

// TypeSpec is a row in the typespec table.
type TypeSpec struct {
	Signature IndexedBlob
}

// ImplMap is a row in the implmap table.
type ImplMap struct {
	MappingFlags    uint16
	MemberForwarded MemberForwardedIndex
	ImportName      IndexedString
	ImportScope     ModuleRefIndex
}

// FieldRVA is a row in the field rva table.
type FieldRVA struct {
	RVA   uint32
	Field FieldIndex
}

// Assembly is a row in the assembly table.
type Assembly struct {
	HashAlgID      uint32
	MajorVersion   uint16
	MinorVersion   uint16
	BuildNumber    uint16
	RevisionNumber uint16
	Flags          uint32
	PublicKey      IndexedBlob
	Name           IndexedString
	Culture        IndexedString
}

// AssemblyProcessor is a row in the assembly processor table.
type AssemblyProcessor struct {
	Processor uint32
}

// AssemblyOS is a row in the assembly OS table.
type AssemblyOS struct {
	OSPlatformID   uint32
	OSMajorVersion uint32
	OSMinorVersion uint32
}

// AssemblyRef is a row in the assembly ref table.
type AssemblyRef struct {
	MajorVersion     uint16
	MinorVersion     uint16
	BuildNumber      uint16
	RevisionNumber   uint16
	Flags            uint32
	PublicKeyOrToken IndexedBlob
	Name             IndexedString
	Culture          IndexedString
	HashValue        IndexedBlob
}

// AssemblyRefProcessor is a row in the assembly ref processor table.
type AssemblyRefProcessor struct {
	Processor   uint32
	AssemblyRef AssemblyRefIndex
}

// AssemblyRefOS is a row in the assembly ref OS table.
type AssemblyRefOS struct {
	OSPlatformID   uint32
	OSMajorVersion uint32
	OSMinorVersion uint32
	AssemblyRef    AssemblyRefIndex
}

// FileRow is a row in the file table.
type FileRow struct {
	Flags     uint32
	Name      IndexedString
	HashValue IndexedString
}

// ExportedType is a row in the exported type table.
type ExportedType struct {
	Flags          uint32
	TypeDefID      uint32
	TypeName       IndexedString
	TypeNamespace  IndexedString
	Implementation ImplementationIndex
}

// ManifestResource is a row in the manifest resource table.
type ManifestResource struct {
	Offset         uint32
	Flags          uint32
	Name           IndexedString
	Implementation ImplementationIndex
}

// NestedClass is a row in the nested class table.
type NestedClass struct {
	NestedClass    TypeDefIndex
	EnclosingClass TypeDefIndex
}

// GenericParam is a row in the generic param table.
type GenericParam struct {
	Number uint16
	Flags  uint16
	Owner  TypeOrMethodDefIndex
	Name   IndexedString
}

// MethodSpec is a row in the methdo spec table.
type MethodSpec struct {
	Method    MethodDefOrRefIndex
	Signature IndexedBlob
}

// GenericParamConstraint is a row in the generic param constraint table.
type GenericParamConstraint struct {
	Owner      GenericParamIndex
	Constraint TypeDefOrRefIndex
}

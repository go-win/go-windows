package winmd

type contextkey string

var ctxdatakey contextkey = "ctxdata"

type rowcounts struct {
	module                 uint32
	typeref                uint32
	typedef                uint32
	field                  uint32
	methoddef              uint32
	param                  uint32
	interfaceimpl          uint32
	memberref              uint32
	constant               uint32
	customattribute        uint32
	fieldmarshal           uint32
	declsecurity           uint32
	classlayout            uint32
	fieldlayout            uint32
	standalonesig          uint32
	eventmap               uint32
	event                  uint32
	propertymap            uint32
	property               uint32
	methodsemantics        uint32
	methodimpl             uint32
	moduleref              uint32
	typespec               uint32
	implmap                uint32
	fieldrva               uint32
	assembly               uint32
	assemblyprocessor      uint32
	assemblyos             uint32
	assemblyref            uint32
	assemblyrefprocessor   uint32
	assemblyrefos          uint32
	file                   uint32
	exportedtype           uint32
	manifestresource       uint32
	nestedclass            uint32
	genericparam           uint32
	methodspec             uint32
	genericparamconstraint uint32
}

type ctxdata struct {
	// table index sizes
	field        uint32
	methoddef    uint32
	typedef      uint32
	param        uint32
	event        uint32
	property     uint32
	moduleref    uint32
	assemblyref  uint32
	genericparam uint32

	// composite table index sizes
	typedeforref        uint32
	hasconstant         uint32
	hascustomattribute  uint32
	hasfieldmarshal     uint32
	hasdeclsecurity     uint32
	memberrefparent     uint32
	hassemantics        uint32
	methoddeforref      uint32
	memberforwarded     uint32
	implementation      uint32
	customattributetype uint32
	resolutionscope     uint32
	typeormethoddef     uint32

	// data index sizes
	stringindexsize uint32
	guidindexsize   uint32
	blobindexsize   uint32

	// data streams
	stringdata []byte
	guiddata   []byte
	blobdata   []byte
}

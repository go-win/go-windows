// Package winmd implements parsing for the WinMD file format.
package winmd

import (
	"fmt"
	"io"
)

// File represents a WinMD file in memory.
type File struct {
}

// Load loads a WinMD file from a readseeker.
func Load(r io.ReadSeeker) (*File, error) {
	f := new(File)

	// Read MZ header.
	mz := mz{}
	if err := read(r, &mz); err != nil {
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
	if err := read(r, &pe); err != nil {
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
		if err := read(r, &sections[i]); err != nil {
			return nil, err
		}
	}

	// Read COM data directory.
	netdir := datadir{}
	netdiroff := int64(mz.PeOffset) + 0xE8
	if _, err := r.Seek(netdiroff, io.SeekStart); err != nil {
		return nil, fmt.Errorf("seeking to .NET data directory at 0x%08x: %w", netdiroff, err)
	}
	if err := read(r, &netdir); err != nil {
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
	if err := read(r, &corhead); err != nil {
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
	if err := read(r, &metahead); err != nil {
		return nil, fmt.Errorf("reading metadata header: %w", err)
	}

	// Stream entries
	tablesoff := int64(0)
	stringsoff := int64(0)
	blobsoff := int64(0)
	for i := 0; i < int(metahead.NumberOfStreams); i++ {
		streamentry := streamentry{}
		if err := read(r, &streamentry); err != nil {
			return nil, err
		}
		switch string(streamentry.Name) {
		case "#~":
			tablesoff = int64(metaheadoff + streamentry.Offset)
		case "#Strings":
			stringsoff = int64(metaheadoff + streamentry.Offset)
		case "#Blob":
			blobsoff = int64(metaheadoff + streamentry.Offset)
		}
	}

	_, _, _ = tablesoff, stringsoff, blobsoff

	return f, nil
}

func rva2off(sections []section, rva uint32) (uint32, error) {
	for _, s := range sections {
		if rva >= s.VirtualAddr && rva < s.VirtualAddr+s.VirtualSize {
			return rva - s.VirtualAddr + s.DataOffset, nil
		}
	}
	return 0, fmt.Errorf("could not find section for rva: 0x%08x", rva)
}

type mz struct {
	Signature uint16
	Unused    [58]byte
	PeOffset  uint32
}

const mzsig = uint16(0x5a4d)

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

const pesig = uint32(0x4550)

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

type datadir struct {
	RVA  uint32
	Size uint32
}

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

type metahead struct {
	Signature       uint32
	MajorVersion    uint16
	MinorVersion    uint16
	Reserved        uint32
	Version         pstr32
	Flags           uint16
	NumberOfStreams uint16
}

type streamentry struct {
	Offset uint32
	Size   uint32
	Name   padstr
}

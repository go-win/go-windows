package winmd

import (
	"context"
	"encoding/binary"
	"errors"
	"io"
)

// CompressedUint represents a compressed unsigned integer, used in some places
// in the metadata.
type CompressedUint uint32

func (field *CompressedUint) unpack(ctx context.Context, r io.Reader) error {
	b := [4]byte{}
	if _, err := io.ReadFull(r, b[0:1]); err != nil {
		return err
	}

	switch {
	case b[0]&0x80 == 0x00:
		*field = CompressedUint(b[0])
		return nil
	case b[0]&0xc0 == 0x80:
		if _, err := io.ReadFull(r, b[1:2]); err != nil {
			return err
		}
		b[0] &^= 0xC0
		*field = CompressedUint(binary.LittleEndian.Uint16(b[0:2]))
		return nil
	case b[0]&0xe0 == 0xc0:
		if _, err := io.ReadFull(r, b[1:4]); err != nil {
			return err
		}
		b[0] &^= 0xe0
		*field = CompressedUint(binary.LittleEndian.Uint32(b[0:4]))
		return nil
	default:
		return errors.New("unsupported compressed int encoding")
	}
}

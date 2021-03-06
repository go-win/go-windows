package winmd

import (
	"context"
	"encoding/binary"
	"io"
	"reflect"
	"sync"
)

var structSize sync.Map

type decoder struct {
	r io.Reader
}
type unpacker interface {
	unpack(context.Context, io.Reader) error
}
type pstr32 string
type padstr string

func (s *pstr32) unpack(ctx context.Context, r io.Reader) error {
	l, err := readuint32(r)
	if err != nil {
		return err
	}
	b := make([]byte, l)
	_, err = io.ReadFull(r, b)
	*s = pstr32(string(b))
	return err
}

func (s *padstr) unpack(ctx context.Context, r io.Reader) error {
	b := []byte{}
	for {
		b = append(b, 0, 0, 0, 0)
		if _, err := io.ReadFull(r, b[len(b)-4:]); err != nil {
			return err
		}
		switch {
		case b[len(b)-4] == 0:
			b = b[:len(b)-1]
			fallthrough
		case b[len(b)-3] == 0:
			b = b[:len(b)-1]
			fallthrough
		case b[len(b)-2] == 0:
			b = b[:len(b)-1]
			fallthrough
		case b[len(b)-1] == 0:
			b = b[:len(b)-1]
			*s = padstr(string(b))
			return nil
		}
	}
}

func readbool(r io.Reader) (bool, error) {
	buf := [1]byte{}
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return false, err
	}
	return buf[0] != 0, nil
}
func readuint8(r io.Reader) (uint8, error) {
	buf := [1]byte{}
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return 0, err
	}
	return buf[0], nil
}
func readuint16(r io.Reader) (uint16, error) {
	buf := [2]byte{}
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(buf[:]), nil
}
func readuint32(r io.Reader) (uint32, error) {
	buf := [4]byte{}
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(buf[:]), nil
}
func readuint64(r io.Reader) (uint64, error) {
	buf := [8]byte{}
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(buf[:]), nil
}
func readint8(r io.Reader) (int8, error) {
	x, err := readuint8(r)
	if err != nil {
		return 0, err
	}
	return int8(x), nil
}
func readint16(r io.Reader) (int16, error) {
	x, err := readuint16(r)
	if err != nil {
		return 0, err
	}
	return int16(x), nil
}
func readint32(r io.Reader) (int32, error) {
	x, err := readuint32(r)
	if err != nil {
		return 0, err
	}
	return int32(x), nil
}
func readint64(r io.Reader) (int64, error) {
	x, err := readuint64(r)
	if err != nil {
		return 0, err
	}
	return int64(x), nil
}
func readvalue(ctx context.Context, r io.Reader, v reflect.Value) error {
	if s, ok := v.Interface().(unpacker); ok {
		return s.unpack(ctx, r)
	}
	if v.CanAddr() {
		if s, ok := v.Addr().Interface().(unpacker); ok {
			return s.unpack(ctx, r)
		}
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// slice fast paths
	if v.CanInterface() && v.Kind() == reflect.Slice {
		switch data := v.Interface().(type) {
		case []int8:
			b := make([]byte, len(data))
			if _, err := r.Read(b); err != nil {
				return nil
			}
			for i, n := range b {
				data[i] = int8(n)
			}
			return nil
		case []uint8:
			if _, err := r.Read(data); err != nil {
				return nil
			}
			return nil
		case []int16:
			b := make([]byte, len(data)*2)
			if _, err := r.Read(b); err != nil {
				return nil
			}
			for i := range data {
				data[i] = int16(binary.LittleEndian.Uint16(b[2*i:]))
			}
			return nil
		case []uint16:
			b := make([]byte, len(data)*2)
			if _, err := r.Read(b); err != nil {
				return nil
			}
			for i := range data {
				data[i] = binary.LittleEndian.Uint16(b[2*i:])
			}
			return nil
		case []int32:
			b := make([]byte, len(data)*4)
			if _, err := r.Read(b); err != nil {
				return nil
			}
			for i := range data {
				data[i] = int32(binary.LittleEndian.Uint32(b[4*i:]))
			}
			return nil
		case []uint32:
			b := make([]byte, len(data)*4)
			if _, err := r.Read(b); err != nil {
				return nil
			}
			for i := range data {
				data[i] = binary.LittleEndian.Uint32(b[4*i:])
			}
			return nil
		case []int64:
			b := make([]byte, len(data)*8)
			if _, err := r.Read(b); err != nil {
				return nil
			}
			for i := range data {
				data[i] = int64(binary.LittleEndian.Uint64(b[8*i:]))
			}
			return nil
		case []uint64:
			b := make([]byte, len(data)*8)
			if _, err := r.Read(b); err != nil {
				return nil
			}
			for i := range data {
				data[i] = binary.LittleEndian.Uint64(b[8*i:])
			}
			return nil
		}
	}

	switch v.Kind() {
	case reflect.Array:
		for i, l := 0, v.Len(); i < l; i++ {
			if err := readvalue(ctx, r, v.Index(i)); err != nil {
				return err
			}
		}
	case reflect.Struct:
		for i, l := 0, v.NumField(); i < l; i++ {
			if err := readvalue(ctx, r, v.Field(i)); err != nil {
				return err
			}
		}
	case reflect.Slice:
		for i, l := 0, v.Len(); i < l; i++ {
			if err := readvalue(ctx, r, v.Index(i)); err != nil {
				return err
			}
		}
	case reflect.Bool:
		x, err := readbool(r)
		if err != nil {
			return err
		}
		v.SetBool(x)
	case reflect.Int8:
		x, err := readint8(r)
		if err != nil {
			return err
		}
		v.SetInt(int64(x))
	case reflect.Int16:
		x, err := readint16(r)
		if err != nil {
			return err
		}
		v.SetInt(int64(x))
	case reflect.Int32:
		x, err := readint32(r)
		if err != nil {
			return err
		}
		v.SetInt(int64(x))
	case reflect.Int64:
		x, err := readint64(r)
		if err != nil {
			return err
		}
		v.SetInt(int64(x))
	case reflect.Uint8:
		x, err := readuint8(r)
		if err != nil {
			return err
		}
		v.SetUint(uint64(x))
	case reflect.Uint16:
		x, err := readuint16(r)
		if err != nil {
			return err
		}
		v.SetUint(uint64(x))
	case reflect.Uint32:
		x, err := readuint32(r)
		if err != nil {
			return err
		}
		v.SetUint(uint64(x))
	case reflect.Uint64:
		x, err := readuint64(r)
		if err != nil {
			return err
		}
		v.SetUint(uint64(x))
	}
	return nil
}

func read(ctx context.Context, r io.Reader, data interface{}) error {
	return readvalue(ctx, r, reflect.ValueOf(data))
}

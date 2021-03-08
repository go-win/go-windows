// Package main implements a code generator for a Win32 language projection.
package main

import (
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-win/go-windows/internal/winmd"
)

func gen(f *winmd.File) error {
	type EnumValue struct {
		Name string
		Type winmd.ValueType
		Data []byte
	}

	type Enum struct {
		Namespace string
		Name      string
		Values    []*EnumValue
	}

	var systemEnum *winmd.TypeRef

	namespaces := map[string][]*Enum{}

	enumValueMap := map[*winmd.Field]*EnumValue{}

	for i := range f.TypeRef {
		ref := &f.TypeRef[i]
		if ref.Namespace == "System" {
			if ref.Name == "Enum" {
				systemEnum = ref
			}
		}
	}

	if systemEnum == nil {
		return errors.New("expected to find System.Enum in typeref table")
	}

	for i := range f.TypeDef {
		def := &f.TypeDef[i]
		fieldEnd := len(f.Field)
		methodEnd := len(f.MethodDef)
		if i+1 < len(f.TypeDef) {
			fieldEnd = int(f.TypeDef[i+1].FieldList) - 1
			methodEnd = int(f.TypeDef[i+1].MethodList) - 1
		}
		if def.Extends.Lookup(f) == systemEnum {
			enum := &Enum{
				Namespace: string(def.Namespace),
				Name:      string(def.Name),
			}
			fields := f.Field[def.FieldList:fieldEnd]
			for i := range fields {
				field := &fields[i]
				value := &EnumValue{
					Name: string(field.Name),
				}
				enum.Values = append(enum.Values, value)
				enumValueMap[field] = value
			}
			namespaces[string(def.Namespace)] = append(namespaces[string(def.Namespace)], enum)
		}
		_ = methodEnd
	}

	for i := range f.Constant {
		constant := &f.Constant[i]
		switch parent := constant.Parent.Lookup(f).(type) {
		case *winmd.Field:
			enumValue, ok := enumValueMap[parent]
			if !ok {
				break
			}
			enumValue.Type = constant.Type
			enumValue.Data = constant.Value
		}
	}

	for namespace, enums := range namespaces {
		ns := strings.Split(strings.ToLower(namespace), ".")
		pkg := ns[len(ns)-1]
		folderpath := strings.ReplaceAll(strings.ToLower(namespace), ".", "/")
		filepath := folderpath + "/enums.go"
		if err := os.MkdirAll(folderpath, 0755); err != nil {
			return err
		}
		f, err := os.Create(filepath)
		if err != nil {
			return err
		}

		names := map[string]*Enum{}
		conflictingEnums := map[*Enum]struct{}{}

		for _, enum := range enums {
			for _, value := range enum.Values {
				if other, ok := names[value.Name]; ok {
					conflictingEnums[enum] = struct{}{}
					conflictingEnums[other] = struct{}{}
				}
				names[value.Name] = enum
			}
		}

		fmt.Println("Generating", filepath)
		fmt.Fprintf(f, "// AUTOGENERATED - DO NOT EDIT\n// Generated by go-windows.\n\n// Package %s implements the %s namespace.\npackage %s\n\n", pkg, namespace, pkg)
		for _, enum := range enums {
			if len(enum.Values) == 0 {
				fmt.Fprintf(f, "type %s int\n\n", enum.Name)
				continue
			}
			switch enum.Values[0].Type {
			case winmd.U16:
				fmt.Fprintf(f, "type %s uint16\n\n", enum.Name)
			case winmd.U32:
				fmt.Fprintf(f, "type %s uint32\n\n", enum.Name)
			case winmd.I32:
				fmt.Fprintf(f, "type %s int32\n\n", enum.Name)
			}
			fmt.Fprintf(f, "const (\n")
			prefix := ""
			if _, ok := conflictingEnums[enum]; ok {
				prefix = enum.Name + "_"
			}
			for _, value := range enum.Values {
				switch value.Type {
				case winmd.U16:
					fmt.Fprintf(f, "\t%s %s = %d\n", prefix+value.Name, enum.Name, binary.LittleEndian.Uint16(value.Data))
				case winmd.U32:
					fmt.Fprintf(f, "\t%s %s = %d\n", prefix+value.Name, enum.Name, binary.LittleEndian.Uint32(value.Data))
				case winmd.I32:
					fmt.Fprintf(f, "\t%s %s = %d\n", prefix+value.Name, enum.Name, int32(binary.LittleEndian.Uint32(value.Data)))
				}
			}
			fmt.Fprintf(f, ")\n\n")
		}
	}

	return nil
}

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		f, err := os.Open(arg)
		if err != nil {
			log.Fatalln(err)
		}
		w, err := winmd.Load(f)
		if err != nil {
			log.Fatalln(err)
		}
		if err := gen(w); err != nil {
			log.Fatalln(err)
		}
	}
}

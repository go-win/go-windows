// Package main implements a code generator for a Win32 language projection.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/go-win/go-windows/internal/winmd"
)

func gen(file *winmd.File) error {
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

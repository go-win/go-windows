// Package main implements a code generator for a Win32 language projection.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/go-win/go-windows/internal/winmd"
)

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		f, err := os.Open(arg)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = winmd.Load(f)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

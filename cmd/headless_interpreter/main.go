package main

import (
	"flag"
	"fmt"
	"os"

	"git.adyxax.org/adyxax/gofunge/pkg/field"
	"git.adyxax.org/adyxax/gofunge/pkg/interpreter"
	"git.adyxax.org/adyxax/gofunge/pkg/pointer"
)

func main() {
	filename := flag.String("f", "", "b98 file to interpret")
	help := flag.Bool("h", false, "display this help message")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *filename == "" {
		fmt.Println("Error : no b98 file to interpret")
		flag.Usage()
		os.Exit(1)
	}
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Failed to open file %s : %+v", *filename, err)
		os.Exit(2)
	}
	defer file.Close()

	f, err := field.Load(file)
	if err != nil {
		fmt.Printf("Failed to load file %s : %+v", *filename, err)
		os.Exit(3)
	}
	p := pointer.NewPointer()
	p.Argv = []string{*filename}
	v := interpreter.NewInterpreter(f, p).Run()
	os.Exit(v)
}

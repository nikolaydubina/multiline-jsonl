package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nikolaydubina/multiline-jsonl/mjsonl"
)

func main() {
	var (
		expand bool
	)

	flag.BoolVar(&expand, "expand", false, "if true will pretty print each JSON")
	flag.Parse()
	flag.Usage = func() {
		doc := "This tool reads multiple JSONs from STDIN and writes formatted output to STDOUT.\n\nUsage of %s:\n"
		fmt.Fprintf(flag.CommandLine.Output(), doc, os.Args[0])
		flag.PrintDefaults()
	}

	if err := mjsonl.FormatJSONL(os.Stdin, os.Stdout, expand); err != nil {
		log.Fatal(err.Error())
	}
}

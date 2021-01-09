package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nickng/bibtex"
)

func print_yaml(bib *bibtex.BibTex) {
	for _, entry := range bib.Entries {
		fmt.Printf("- id: %s\n", entry.CiteName)
		for key, field := range entry.Fields {
			fmt.Printf("  %s: %s\n", key, field.String())
		}
		fmt.Printf("\n")
	}
}

func main() {
	infile := flag.String("i", "", "Input file (default: stdin)")
	reader := os.Stdin

	flag.Parse()

	if *infile != "" {
		rdFile, err := os.Open(*infile)
		if err != nil {
			log.Fatal(err)
		}
		defer rdFile.Close()
		reader = rdFile
	}

	bib, err := bibtex.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}

	print_yaml(bib)

}

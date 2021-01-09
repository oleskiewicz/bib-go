package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nickng/bibtex"
)

func print_refer(bib *bibtex.BibTex) {
	name_to_key := map[string]rune{"author": 'A', "title": 'T', "year": 'Y'}
	for _, entry := range bib.Entries {
		fmt.Printf("%%K %s\n", entry.CiteName)
		for name, field := range entry.Fields {
			if key, ok := name_to_key[name]; ok {
				fmt.Printf("%%%c %s\n", key, field.String())
			}
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

	print_refer(bib)

}

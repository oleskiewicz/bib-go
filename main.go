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

func print_tsv(bib *bibtex.BibTex) {
	title, author := "", ""
	for _, entry := range bib.Entries {
		for key, field := range entry.Fields {
			if key == "author" {
				author = field.String()
			}
			if key == "title" {
				title = field.String()
			}
		}
		fmt.Printf("%s\t%s\n", author, title)
	}
}

func main() {
	infile := flag.String("i", "", "Input file (default: stdin)")
	outFmt := flag.String("o", "b", "Output format (default: b)")
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

	switch *outFmt {
	case "b":
		fmt.Printf(bib.PrettyString())
	case "r":
		print_refer(bib)
	case "t":
		print_tsv(bib)
	case "y":
		print_yaml(bib)
	default:
		fmt.Printf("unercognised output format\n")
	}

}

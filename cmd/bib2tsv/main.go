package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nickng/bibtex"
)

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

	print_tsv(bib)

}

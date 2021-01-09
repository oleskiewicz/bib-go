package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nickng/bibtex"
)

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

	fmt.Printf(bib.PrettyString())

}

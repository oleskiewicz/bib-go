bib
===
BibTex bibliography tools, written in Go using `nickng/bibtex` library.

The following utilities are included:

- `bibfmt`: format BibTex
- `bib2ref`: convert BibTex to GNU Refer format, to use with `groff`
- `bib2tsv`: convert BibTex to TSV (only author, title columns for now)
- `bib2yml`: convert BibTex to Yaml, to use with `pandoc`


Usage
-----

	BIN < FILE.bib

Every tool reads from a filename or from a standard input and writes to standard output.

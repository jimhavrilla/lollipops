package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	uniprot = flag.String("U", "", "Uniprot accession instead of GENE_SYMBOL")
	output  = flag.String("o", "", "output SVG file (default GENE_SYMBOL.svg)")
	width   = flag.Int("w", 0, "SVG output width (default automatic fit labels)")
	maf     = flag.Int("m", 0, "adds maf frequencies to plot and varies height by rarity if you change it to 1") // new flag for maf
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] GENE_SYMBOL [PROTEIN CHANGES ...]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Where options are:")
		flag.PrintDefaults()
	}
	flag.Parse()

	var err error
	varStart := 0
	acc := ""
	geneSymbol := ""
	out := []byte{0}
	args := []string{flag.Arg(0)}
	if *uniprot == "" && flag.NArg() > 0 {
		cmd := "idconv.py"
		out, err = exec.Command(cmd, args...).Output()
		acc = string(out)
		varStart = 1

		fmt.Fprintln(os.Stderr, "HGNC Symbol: ", flag.Arg(0))

		//acc, err = GetProtID(flag.Arg(0))
		//if err != nil {
		//	fmt.Fprintln(os.Stderr, err)
		//	os.Exit(1)
		//}

		fmt.Fprintln(os.Stderr, "Uniprot/SwissProt Accession: ", acc)
	}

	if *uniprot != "" {
		acc = *uniprot
	}

	if flag.NArg() == 0 && *uniprot == "" {
		flag.Usage()
		os.Exit(1)
	}

	data, err := GetPfamGraphicData(acc)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if geneSymbol == "" {
		geneSymbol = data.Metadata.Identifier
		fmt.Fprintln(os.Stderr, "Pfam Symbol: ", geneSymbol)
	}

	if *output == "" {
		*output = geneSymbol + ".svg"
	}

	f, err := os.OpenFile(*output, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Fprintln(os.Stderr, "Drawing diagram to", *output)
	DrawSVG(f, *width, flag.Args()[varStart:], *maf, data) // added new field for maf

}

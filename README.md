This is a fork of the original lollipops diagram generator to be found at [Jeremy Jay's world of magic!](https://github.com/pbnjay/lollipops)  My version uses a Python executable to translate the HGNC Symbols into Uniprot numbers for Pfam domain accession instead of using the UniProt REST API (which works fine, and Jeremy uses) or the Biomart API (which had plenty of issues in the past).  Additionally, I allow for height modulation by rarity of a variant, and contributed the color changing of variants by nature of variant (stop lost, stop gained, missense, silent, etc.) to his original repo.  I plan on updating the program with further insertions to possibly allow for CATH domains to be transcribed, and to add more examples of how my version is used, and how it differs.  If you would rather use the original, simply check out the link above!

lollipops
---------

This is a simple 'lollipop' mutation diagram generator that tries to make things
as automated as possible. It uses the [Pfam API](http://pfam.xfam.org/help#tabview=tab9)
to retrieve domains and colors, and mygene (https://pypi.python.org/pypi/mygene)
to translate HGNC Gene Symbols into Uniprot/SwissProt Accession numbers. If variant changes
are provided, it will also annotate them to the diagram using the "lollipops" markers that
give the tool its name.

Example
-------

    ./lollipops TP53  E343Q R342Q F338 R335C R283H R283C R282W R248Q G245S C242 N235S P223H P222 P222L V216 R213 L206 L194 D186E S185 S185N R158 R156H Y107H

Results in the following SVG image:

![TP53 Lollipop diagram with 3 marked mutations](tp53.png?raw=true)

Usage
-----

Usage: ``lollipops [options] GENE_SYMBOL [PROTEIN CHANGES ...]``

Where ``GENE_SYMBOL`` is the official HGNC symbol and ``PROTEIN CHANGES``
is a list of amino acid changes of the format "(amino-code)(position)..."
Amino-code can be either the 1- or 3-character code for the amino acid.
Only the first position in each change is used for plotting even if the
change contains a range. All characters after the position are ignored.

    -o=out.svg         SVG output filename (default GENE_SYMBOL.svg)
    -labels            draw labels for each mutation
    -hide-axis         do not draw the aa position axis
    -hide-disordered   do not draw disordered regions
    -hide-motifs       do not draw motifs
    -w=700             SVG output width (default=automatic)

If you are working with non-human data, or know the Uniprot Accession
already, You can specify it with `-U UNIPROTID` instead of GENE_SYMBOL,
for example the following mouse query works for gene `Mobp`:

    ./lollipops -U Q9D2P8

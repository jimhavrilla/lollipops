#!/usr/bin/env python
import mygene
import sys

mg = mygene.MyGeneInfo()
xli = sys.argv[1]
out = mg.query(xli, scopes='symbol', fields='uniprot', species='human')
a = str(out['hits'][0]['uniprot']['Swiss-Prot'])
sys.stdout.write(a)

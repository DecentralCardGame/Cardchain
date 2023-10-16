import tendermint_chunked_genesis_download as tcgd
f = open('syncnode.txt', 'r')

tcgd.download_genesis('http://'+f.read()+':26657/')
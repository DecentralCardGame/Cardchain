import json
import requests
import tendermint_chunked_genesis_download as tcgd

with open('peer_nodes.json', 'r') as f:
  data = json.load(f)

for url in data['rpcs']:
	try:
	    page = requests.get(url)
	    print(url, page.status_code)
	    tcgd.download_genesis(url)
	    break
	except (requests.exceptions.HTTPError, requests.exceptions.ConnectionError):
	    print("Error")
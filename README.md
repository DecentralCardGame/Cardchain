# Decentral-card.net Cardchain
## Crowd Control Chain

install go version 1.11 or higher

Then set path variables:
> echo "export GOPATH=$HOME/go" >> ~/.bashrc \
source ~/.bashrc \
echo "export GOBIN=$GOPATH/bin" >> ~/.bashrc \
source ~/.bashrc \
echo "export PATH=$PATH:$GOBIN" >> ~/.bashrc \
source ~/.bashrc

install jq

build the project:
> make install

this runs all init commands of nsd and nscli to setup accounts bob and alice:
> sh scripts/setup.sh

this starts a blockchain node:
> csd start

you can start a restserver via:
> sh scripts/startRESTserver.sh

In the scripts folder the most useful commands can be found in .sh files. In addition the client gives help and shows all possible transaction and queries:
> cscli help

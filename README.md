Decentral-card.net Cardchain // Crowd Control Chain

install go version 1.11 or higher

echo "export GOPATH=$HOME/go" >> ~/.bashrc

source ~/.bashrc

echo "export GOBIN=$GOPATH/bin" >> ~/.bashrc

source ~/.bashrc

echo "export PATH=$PATH:$GOBIN" >> ~/.bashrc

source ~/.bashrc

make install

sh scripts/setup.sh

csd start

cscli help

see /scripts (jq needed)

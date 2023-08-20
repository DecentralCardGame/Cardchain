
FROM ignitehq/cli:v0.27.1


USER root
RUN apt-get -y -qq update && \
	apt-get install -y -qq apt-transport-https curl wget unzip screen bash jq && \
	apt-get clean
#
# install jq to parse json within bash scripts
#RUN curl -o /usr/local/bin/jq https://github.com/jqlang/jq/releases/download/jq-1.6/jq-linux64 && \
#  chmod +x /usr/local/bin/jq
#RUN jq #this will crash in case jq is not properly installed

# install correct go version
RUN wget https://go.dev/dl/go1.20.2.linux-amd64.tar.gz
RUN tar -xvf go1.20.2.linux-amd64.tar.gz
RUN rm /usr/local/go -rf
RUN mv go /usr/local


USER tendermint
WORKDIR /home/tendermints

RUN export GOPATH=$HOME/go
RUN wget https://github.com/DecentralCardGame/go-faucet/archive/master.zip && \
	unzip master.zip -d . && cd go-faucet-master && go build

EXPOSE 1317
EXPOSE 26657
EXPOSE 26658
EXPOSE 9090
EXPOSE 4500

COPY --chown=tendermint:tendermint . .


RUN ignite chain build
RUN ignite chain init

#RUN	mkdir -p $HOME/.Cardchain/config
RUN wget -O $HOME/.Cardchain/config/genesis.json "https://raw.githubusercontent.com/DecentralCardGame/Testnet/Testnet4/genesis.json"
RUN	wget -O $HOME/.Cardchain/config/addrbook.json "https://raw.githubusercontent.com/DecentralCardGame/Testnet/Testnet4/addrbook.json"


RUN chmod +x ./docker-run.sh
ENTRYPOINT bash docker-run.sh
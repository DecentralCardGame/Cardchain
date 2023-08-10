
FROM ignitehq/cli:v0.27.1


USER root
RUN apt-get -y -qq update && \
	apt-get install -y -qq apt-transport-https curl wget unzip screen && \
	apt-get clean
#
# install jq to parse json within bash scripts
RUN curl -o /usr/local/bin/jq http://stedolan.github.io/jq/download/linux64/jq && \
  chmod +x /usr/local/bin/jq



USER tendermint
WORKDIR /home/tendermint

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

RUN echo $HOME

RUN chmod +x ./docker-run.sh
ENTRYPOINT ./docker-run.sh
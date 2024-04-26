
FROM ignitehq/cli:v0.26.1


USER root
RUN apt-get -y -qq update && \
	apt-get install -y -qq apt-transport-https curl wget unzip screen bash jq python3 pip && \
	apt-get clean


# install python script to download genesis
RUN pip install tendermint-chunked-genesis-download


# install correct go version
RUN if [ $(uname -m) = "x86_64" ]; then \
        wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz; \
        tar -xvf go1.21.3.linux-amd64.tar.gz; \
        rm /usr/local/go -rf; \
        mv go /usr/local; \
    elif [ $(uname -m) = "aarch64" ]; then \
        wget https://go.dev/dl/go1.21.3.linux-arm64.tar.gz; \
        tar -xvf go1.21.3.linux-arm64.tar.gz; \
        rm /usr/local/go -rf; \
        mv go /usr/local; \
    else \
        echo "what the hell is your OS? Go will not work that way."; \
    fi


USER tendermint
WORKDIR /home/tendermint

RUN export GOPATH=$HOME/go
RUN wget https://github.com/DecentralCardGame/go-faucet/archive/master.zip && \
	unzip master.zip -d . && cd go-faucet-master && go build

EXPOSE 1317
EXPOSE 26657
EXPOSE 26658
EXPOSE 9090
EXPOSE 9091
EXPOSE 4500

COPY --chown=tendermint:tendermint . .

RUN ignite chain build
RUN ignite chain init

RUN mv $HOME/.Cardchain $HOME/.cardchaind

COPY scripts/download_genesis.py download_genesis.py
RUN python3 download_genesis.py
RUN mv genesis.json $HOME/.cardchaind/config/genesis.json
RUN	wget -O $HOME/.cardchaind/config/addrbook.json "https://raw.githubusercontent.com/DecentralCardGame/Testnet/main/addrbook.json"

RUN chmod +x ./docker-run.sh
ENTRYPOINT bash docker-run.sh

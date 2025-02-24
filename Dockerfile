
#FROM ignitehq/cli:v0.26.1
FROM debian:bookworm

USER root
RUN apt-get -y -qq update && \
    apt-get install -y -qq apt-transport-https curl wget unzip screen bash jq python3 pip python3.11-venv && \
    apt-get clean

# install python script to download genesis
RUN python3 -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"
RUN pip install tendermint-chunked-genesis-download


# install correct go version
RUN if [ $(uname -m) = "x86_64" ]; then \
    wget https://go.dev/dl/go1.23.6.linux-amd64.tar.gz; \
    tar -xvf go1.23.6.linux-amd64.tar.gz; \
    rm /usr/local/go -rf; \
    mv go /usr/local; \
    elif [ $(uname -m) = "aarch64" ]; then \
    wget https://go.dev/dl/go1.23.6.linux-arm64.tar.gz; \
    tar -xvf go1.23.6.linux-arm64.tar.gz; \
    rm /usr/local/go -rf; \
    mv go /usr/local; \
    else \
    echo "what the hell is your OS? Go will not work that way."; \
    fi

RUN wget https://github.com/DecentralCardGame/Cardchain/releases/download/v0.17.0/cardchaind -O /usr/local/go/bin/cardchaind
RUN chmod +x /usr/local/go/bin/cardchaind

RUN useradd -ms /bin/bash tendermint
USER tendermint
WORKDIR /home/tendermint

ENV PATH="$PATH:/usr/local/go/bin"
RUN go version
RUN wget https://github.com/DecentralCardGame/go-faucet/archive/master.zip && \
    unzip master.zip -d . && cd go-faucet-master && go build

EXPOSE 1317
EXPOSE 26657
EXPOSE 26658
EXPOSE 9090
EXPOSE 9091
EXPOSE 4500

COPY --chown=tendermint:tendermint . .

#RUN ignite chain build --skip-proto
#RUN ignite chain init --skip-proto
#RUN mv $HOME/.Cardchain $HOME/.cardchaind


RUN cardchaind init gonninode

COPY scripts/download_genesis.py download_genesis.py
RUN python3 download_genesis.py
RUN mv genesis.json $HOME/.cardchaind/config/genesis.json

RUN chmod +x ./docker-run.sh
ENTRYPOINT bash docker-run.sh

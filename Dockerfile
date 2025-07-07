# Use Go 1.23 bookworm as base image
FROM golang:1.23-bookworm AS base

RUN apt-get -y -qq update && \
    apt-get install -y -qq apt-transport-https curl wget unzip screen bash jq python3 python3-pip python3-venv && \
    rm -rf /var/lib/apt/lists/*

# install python script to download genesis
RUN python3 -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"
RUN pip install tendermint-chunked-genesis-download

RUN pip install tendermint-chunked-genesis-download

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY . .

RUN wget https://github.com/DecentralCardGame/go-faucet/archive/master.zip && \
    unzip master.zip -d . && cd go-faucet-master && go build

# Install dependencies
RUN go install github.com/lxgr-linux/ignite-vm@latest

# Build the application
RUN ignite-vm install v0.26.2
RUN ignite-vm set v0.26.2

RUN ~/.local/bin/ignite chain init --home /build/.cardchaind
RUN ~/.local/bin/ignite chain build

# Document the port that may need to be published
EXPOSE 1317
EXPOSE 26657
EXPOSE 26658
EXPOSE 9090
EXPOSE 9091
EXPOSE 4500

COPY scripts/download_genesis.py download_genesis.py
RUN python3 download_genesis.py
RUN mv genesis.json /build/.cardchaind/config/genesis.json

# Start the application
RUN chmod +x ./docker-run.sh
ENTRYPOINT bash docker-run.sh
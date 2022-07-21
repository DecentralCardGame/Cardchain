FROM ignitehq/cli:0.20.4

EXPOSE 1317
EXPOSE 26657
EXPOSE 26658
EXPOSE 9090
EXPOSE 4500

WORKDIR .
COPY --chown=tendermint:tendermint . .

RUN chmod +x ./docker-run.sh

RUN ignite chain build
#RUN ignite chain init
#RUN python3 ./scripts/migrate_with_data.py ./blockchain-data/exported_genesis.json ~/.Cardchain/config/genesis.json

ENTRYPOINT ./docker-run.sh

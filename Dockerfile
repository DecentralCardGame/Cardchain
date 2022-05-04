FROM ignitehq/cli:0.20.4

EXPOSE 1318
EXPOSE 26659

WORKDIR .
COPY --chown=tendermint:tendermint . .

RUN chmod +x ./docker-run.sh

RUN ignite chain build

ENTRYPOINT ./docker-run.sh

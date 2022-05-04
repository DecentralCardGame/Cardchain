FROM ignitehq/cli:0.20.4

EXPOSE 1318
EXPOSE 26659

WORKDIR .
COPY --chown=tendermint:tendermint . .

RUN ignite chain build

ENTRYPOINT ./docker-run.sh

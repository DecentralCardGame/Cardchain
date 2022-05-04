FROM ignitehq/cli:0.20.4

EXPOSE 1318
EXPOSE 26659

WORKDIR .
COPY . .

RUN ignite chain build


CMD ./docker-run.sh

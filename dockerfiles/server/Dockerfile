FROM alpine

WORKDIR /app

COPY ./quality-trace-server /app/quality-trace-server

EXPOSE 11633/tcp

ENTRYPOINT ["/app/quality-trace-server", "serve"]
FROM alpine

WORKDIR /app

COPY ./quality-trace-server /app/quality-trace-server

# Adding /app folder on $PATH to allow users to call tracetest cli on docker
ENV PATH="$PATH:/app"

EXPOSE 11633/tcp

ENTRYPOINT ["/app/quality-trace-server", "serve"]
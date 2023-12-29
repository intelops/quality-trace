FROM alpine

WORKDIR /app

COPY ./quality-trace-server /app/quality-trace-server
COPY ./quality-trace /app/quality-trace

# Adding /app folder on $PATH to allow users to call quality-trace cli on docker
ENV PATH="$PATH:/app"

EXPOSE 11633/tcp

ENTRYPOINT ["/app/quality-trace-server", "serve"]
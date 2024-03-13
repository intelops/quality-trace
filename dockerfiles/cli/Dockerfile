FROM alpine

WORKDIR /app

COPY ./quality-trace /app/quality-trace

# Adding /app folder on $PATH to allow users to call quality-trace cli on docker
ENV PATH="$PATH:/app"

ENTRYPOINT ["/app/quality-trace"]
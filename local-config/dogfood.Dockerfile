FROM alpine

RUN apk --update add bash jq curl

WORKDIR /app
COPY ./dist/quality-trace /app/quality-trace
COPY ./testing/server-qualitytracing ./qualitytracing

WORKDIR /app/qualitytracing
CMD ["/bin/sh", "/app/qualitytracing/run.bash"]

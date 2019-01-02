FROM alpine:3.8

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
COPY tapper-server .
COPY tapper.tgz .
COPY script/* ./script/

ENV GIN_MODE=release
ENV PORT=80

ENTRYPOINT ["./tapper-server"]
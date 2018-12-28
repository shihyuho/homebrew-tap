FROM alpine:3.8

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
ADD tapper /app/tapper

ENV GIN_MODE=release
ENV PORT=80

ENTRYPOINT ["./tapper"]
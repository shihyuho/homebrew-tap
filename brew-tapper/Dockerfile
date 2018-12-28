FROM alpine:3.8

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
ADD brew-tapper /app/brew-tapper

ENV GIN_MODE=release
ENV PORT=80

ENTRYPOINT ["./brew-tapper"]
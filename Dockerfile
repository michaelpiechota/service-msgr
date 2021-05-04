FROM golang:1.16 as builder

LABEL maintainer="Michael Piechota <mepiechota@gmail.com>"

WORKDIR /usr/src/app

COPY go.mod go.sum makefile ./

RUN make deps

COPY . .

RUN make build-linux

FROM alpine:latest

LABEL maintainer="Michael Piechota <mepiechota@gmail.com>"

RUN addgroup -S app \
    && adduser -S -g app app

RUN apk --no-cache add ca-certificates

WORKDIR /usr/src/app

COPY  --from=builder /usr/src/app/bin/serve bin
COPY .env.template .

RUN chown -R app:app /usr/src/app

EXPOSE 3000

USER app

CMD ["./bin"]
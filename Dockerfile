ARG GO_VERSION=1.19
ARG ALPINE_VERSION=3.17
ARG DOORAY_HOOK_URL="https://hook.dooray.com/services/example"

FROM golang:${GO_VERSION}-alpine AS builder

ENV CGO_ENABLED=0

WORKDIR "/go/src/dooray-bot"

COPY clients ./clients
COPY cmd ./cmd
COPY vendor ./vendor/
COPY go.* ./

RUN CGO_ENABLED=0 go build -mod=vendor -o bin/dooray-bot -v ./cmd/main.go

FROM alpine:${ALPINE_VERSION}
LABEL maintainer="deepdiveinwinter <deepdiveinwinter@gmail.com>"

ENV DOORAY_HOOK_URL=${DOORAY_HOOK_URL}

COPY --from=builder /go/src/dooray-bot/bin/dooray-bot /usr/bin/dooray-bot

ENTRYPOINT ["dooray-bot"]
FROM golang:1.22.1-alpine3.18 as builder

RUN apk add --no-cache make musl-dev linux-headers gcc git jq bash


WORKDIR /app/aggregator

# build node with local monorepo go modules
COPY ./cmd/aggregator ./cmd/aggregator
COPY aggregator ./aggregator
COPY contracts/bindings ./contracts/bindings
COPY utils ./utils
COPY go.mod ./
COPY go.sum ./
COPY vendor ./vendor


RUN go build -o ./bin/aggregator ./cmd/aggregator

FROM alpine:3.18

COPY --from=builder /app/aggregator/bin/aggregator /usr/local/bin

ENTRYPOINT ["aggregator"]

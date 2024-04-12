FROM golang:1.22.1-alpine3.18 as builder

RUN apk add --no-cache make musl-dev linux-headers gcc git jq bash


WORKDIR /app/oprtool

# build node with local monorepo go modules
COPY ./cmd/oprtool ./cmd/oprtool
COPY operator ./operator
COPY aggregator ./aggregator
COPY contracts/bindings ./contracts/bindings
COPY utils ./utils
COPY go.mod ./
COPY go.sum ./


RUN go build -o ./bin/oprtool ./cmd/oprtool

FROM alpine:3.18

COPY --from=builder /app/oprtool/bin/oprtool /usr/local/bin

ENTRYPOINT ["oprtool"]

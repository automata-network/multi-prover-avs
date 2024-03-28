FROM golang:1.22.1-alpine3.18 as builder

RUN apk add --no-cache make musl-dev linux-headers gcc git jq bash


WORKDIR /app/operator

# build node with local monorepo go modules
COPY ./cmd/operator ./cmd/operator
COPY aggregator ./aggregator
COPY operator ./operator
COPY contracts/bindings ./contracts/bindings
COPY utils ./utils
COPY go.mod ./
COPY go.sum ./


RUN go build -o ./bin/operator ./cmd/operator

FROM alpine:3.18

COPY --from=builder /app/operator/bin/operator /usr/local/bin

ENTRYPOINT ["operator"]

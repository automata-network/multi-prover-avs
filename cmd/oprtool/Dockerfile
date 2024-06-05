FROM golang:1.22.1-alpine3.18 as builder

RUN apk add --no-cache make musl-dev linux-headers gcc git jq bash


WORKDIR /app/oprtool

# build node with local monorepo go modules
COPY ./cmd/oprtool ./cmd/oprtool
COPY operator ./operator
COPY aggregator ./aggregator
COPY contracts/bindings ./contracts/bindings
COPY utils ./utils
COPY xmetric ./xmetric
COPY xtask ./xtask
COPY go.mod ./
COPY go.sum ./
COPY vendor ./vendor

ARG BUILD_TAG
ARG GIT_COMMIT
ARG GIT_DATE
ENV BUILD_TAG $BUILD_TAG
ENV GIT_COMMIT $GIT_COMMIT
ENV GIT_DATE $GIT_DATE

RUN go build -ldflags "-X main.SemVer=${BUILD_TAG} -X main.GitCommit=${GIT_COMMIT} -X main.GitDate=${GIT_DATE}" -o ./bin/oprtool ./cmd/oprtool

FROM alpine:3.18

COPY --from=builder /app/oprtool/bin/oprtool /usr/local/bin

ENTRYPOINT ["oprtool"]

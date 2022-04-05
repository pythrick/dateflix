FROM golang:1.18 as gobuild
ARG VERSION=latest
WORKDIR /go/src/github.com/pythrick/dateflix

COPY go.mod go.sum ./
COPY cmd ./cmd
COPY internal ./internal
COPY build ./build

RUN cd cmd/dateflix-api && make build
RUN mv cmd/dateflix-api/dateflix-api build/dateflix-api/

FROM gcr.io/distroless/base
WORKDIR /root/
COPY --from=gobuild /go/src/github.com/pythrick/dateflix/build/dateflix-api/ .
ENTRYPOINT ["./dateflix-api"]

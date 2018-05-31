PROJECT?=github.com/hinrikg/gophercon
APP?=gophercon
PORT?=8000

RELEASE?=0.0.0
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

default: build test

clean:
	rm -f ./bin/${APP}

build: clean
	go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ./bin/${APP} ${PROJECT}/cmd/gophercon

run: build
	SERVICE_PORT=${PORT} ./bin/${APP}

test:
	go test -race ./...

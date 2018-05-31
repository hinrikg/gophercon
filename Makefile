PROJECT?=github.com/hinrikg/gophercon
APP?=gophercon
PORT?=8000
INTERNAL_PORT?=8001

RELEASE?=0.0.0
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

CONTAINER_IMAGE?=${APP}

GOOS?=linux
GOARCH?=amd64

default: build test

clean:
	rm -f ./bin/${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
	go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ./bin/${APP} ${PROJECT}/cmd/gophercon

run: container
	docker stop ${CONTAINER_IMAGE}:${RELEASE} || true && docker rm ${CONTAINER_IMAGE}:${RELEASE} || true
	docker run --name ${CONTAINER_IMAGE} -p ${PORT}:${PORT} -p ${INTERNAL_PORT}:${INTERNAL_PORT} --rm \
		-e "SERVICE_PORT=${PORT}" -e "INTERNAL_PORT=${INTERNAL_PORT}" ${CONTAINER_IMAGE}:${RELEASE}

container: build
	docker build -t ${CONTAINER_IMAGE}:${RELEASE} .

test:
	go test -race ./...

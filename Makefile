PROJECT?=github.com/hinrikg/gophercon
APP?=gophercon
PORT?=8000

default: build test

clean:
	rm -f ./bin/${APP}

build: clean
	go build -o ./bin/${APP} ${PROJECT}/cmd/gophercon

run: build
	SERVICE_PORT=${PORT} ./bin/${APP}

test:
	go test -race ./...

.PHONY: dev test sample clean

LDFLAGS=-ldflags "-s -w"

dev:
	docker run -it --name gotag \
		-v $(PWD):/go/src/github.com/MOXA-ISD/gotag \
		-w /go/src/github.com/MOXA-ISD/gotag \
		golang:1.13.5-stretch \
		bash

test:
	go test -v --cover -count=1 .

sample:
	go build $(LDFLAGS) -o client sample/client.go

clean:
	rm -rf client

.PHONY: test sample clean

LDFLAGS=-ldflags "-s -w"

test:
	go test -v --cover .

sample:
	go build $(LDFLAGS) -o client sample/client.go

clean:
	rm -rf client

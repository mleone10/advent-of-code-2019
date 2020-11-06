DAYS := $(shell ls ./cmd/)

clean:
	rm -rf ./bin

test:
	go test ./...

build: clean test
	for d in $(DAYS); do \
		go build -o ./bin/$$d ./cmd/$$d; \
	done

DAYS := $(shell ls ./cmd/)

clean:
	rm -rf ./bin

build: clean
	for d in $(DAYS); do \
		go build -o ./bin/$$d ./cmd/$$d; \
	done

default: build

clean:
	rm -rf ./.cli-unit

deps:
	go get

build:
	go build

test:
	go test
	./cli-unitw.sh *_test.md

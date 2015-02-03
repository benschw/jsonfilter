VERSION := $(shell cat VERSION)
ITTERATION := $(shell date +%s)

# # drone build
# sudo apt-get update
# sudo apt-get install ruby-dev build-essential rubygems wget curl
# sudo gem install fpm
# make deps test build deb gzip

all: build

deps:
	go get -t -v ./...

test:
	go test
	/bin/bash ./cli-unitw.sh README.md *_test.md

build:
	mkdir -p build/output
	mkdir -p build/root/usr/bin
	go build -o build/root/usr/bin/jsonfilter

install:
	install -t /usr/bin build/root/usr/bin/jsonfilter

clean:
	rm -rf ./.cli-unit
	rm -rf build

packages: clean build deb gzip

gzip:
	cp build/root/usr/bin/jsonfilter build/output/jsonfilter
	gzip build/output/jsonfilter

# sudo apt-get install ruby-dev build-essential
# sudo gem install fpm
# 
# creates a debian package
# `sudo dpkg -i jsonfilter.deb`
deb:
	fpm -s dir -t deb -n jsonfilter -v $(VERSION) -p build/output/jsonfilter.deb \
		--deb-priority optional \
		--category util \
		--force \
		--iteration $(ITTERATION) \
		--deb-compression bzip2 \
		--url https://github.com/benschw/jsonfilter \
		--description "jsonfilter json parsing and filtering" \
		-m "Ben Schwartz <benschw@gmail.com>" \
		--license "Apache License 2.0" \
		--vendor "fliglio.com" -a amd64 \
		build/root/=/

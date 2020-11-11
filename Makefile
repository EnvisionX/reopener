.PHONY: all test clean

export GOPATH=/usr/share/gocode
export HOME = $(shell pwd)

all:
	go build -o reopener

test:
	/bin/sh -x ./test.sh

clean:
	rm -f reopener *~ *_flymake
	rm -rf test_tmp

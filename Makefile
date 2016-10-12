.PHONY: all test clean

all:
	go build -o reopener

test:
	/bin/sh -x ./test.sh

clean:
	rm -f reopener *~ *_flymake
	rm -rf test_tmp

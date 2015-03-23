.PHONY: all test clean

all:

test: clean
	/bin/sh -x ./test.sh

clean:
	rm -f -- *~ *.pyc *.pyo
	rm -rf test_tmp

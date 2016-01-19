#!/bin/sh

set -e

spam(){
    for i in `seq 1 30`; do
        echo $i
        sleep $1
    done
}

# prepare environment
TMP=test_tmp
rm -rf "$TMP"
mkdir "$TMP"

# spawn a "logger" process
export REOPENER_PERIOD=1
spam 0.1 | ./reopener "$TMP"/log &

# rotate the log periodically
for i in 1 2 3; do
    sleep 1
    [ -f "$TMP"/log ] && mv "$TMP"/log "$TMP"/log.$i
done

# wait for the logger to stop and collect all the logs to a one file
wait
for i in 1 2 3; do
    [ -f "$TMP"/log.$i ] && cat "$TMP"/log.$i >> "$TMP"/result
done
[ -f "$TMP"/log ] && cat "$TMP"/log >> "$TMP"/result

# create a file with expected data
spam 0 > "$TMP"/expect

# compare the written logs with our expectations
if diff -q "$TMP"/expect "$TMP"/result; then
    echo OK
else
    echo FAIL
    diff -u "$TMP"/expect "$TMP"/result
    exit 1
fi

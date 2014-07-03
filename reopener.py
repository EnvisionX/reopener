#!/usr/bin/env python

"""
Pass all stdout line by line to an file, reopening it periodically.

It provides an easy way to rotate log files for applications
which are not clever enough to reopen logfiles for themselves.
"""

import os
import sys
import time

# ----------------------------------------------------------------------
# Internal definitions

PERIOD = 60  # one minute
if 'REOPENER_PERIOD' in os.environ:
    # for test purposes
    PERIOD = int(os.environ['REOPENER_PERIOD'])


def usage():
    """
    Show the program usage synopsis and exit.
    """
    sys.stdout.write('Usage: {0} <filename>\n'.format(sys.argv[0]))
    sys.exit(1)


def loop(filename):
    """
    Read the stdin line by line and append the file, reopening
    it every PERIOD seconds.

    :param filename: path to target file
    :type filename: string
    """
    while True:
        open_time = time.time()
        with open(filename, 'a') as fdescr:
            line = sys.stdin.readline()
            while line:
                fdescr.write(line)
                if time.time() - open_time >= PERIOD:
                    break
                line = sys.stdin.readline()
            if not line:
                break

if __name__ == '__main__':
    try:
        if len(sys.argv) != 2:
            usage()
        loop(sys.argv[1])
    except KeyboardInterrupt:
        pass

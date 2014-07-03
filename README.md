# Reopener

Pass all stdout line by line to an file, reopening it periodically.

Provides an easy way to rotate log files for applications
which are not clever enough to reopen logfiles for themselves.

## Usage example

    /usr/bin/stupid-daemon | reopener /var/log/stupid-daemon.log

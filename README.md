# Reopener

Pass all stdout line by line to an file, reopening it periodically.

Provides an easy way to rotate log files for applications
which are not clever enough to reopen logfiles for themselves.

## Second life of the program.

Reopener was obsolete since Release 5.0 OldPew (16 Feb 2016), when we
completely switched to systemd with journald as logging facility.

But first decade of March 2016 shows the Journald isn't the best solution
for reliable logging. For example, it can die when memory of the node
is consumed by another daemon and experience shows it doesn't recover
after such issue.

So we decided to switch our daemons back to the simple pipe writer with
old good logrotate tool.

## Usage example

    /usr/bin/stupid-daemon | reopener /var/log/stupid-daemon.log

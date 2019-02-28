Text files for the `man` command.

## Writing a manual file

* The file name should be the name of the command with no extension
* Max line length of 68 characters
* The file should be in the following format:

```
                  ARP-SH1 GENERAL COMMANDS MANUAL

NAME
    $CMD -- $SHORT_DESCRIPTION

SYNOPSIS
    $CMD [-f] [file ...]
    $CMD

DESCRIPTION
    $LONG_DESCRIPTION

SEE ALSO
    $OTHER_CMD
```
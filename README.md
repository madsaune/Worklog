#worklog

A simple cmdline tool for tracking time spent on a particular task.

## Features

- [X] Use CTRL+C to stop tracking
- [X] Be able to add metadata
- [X] Write to sqlite database
- [X] Create database in `~/.worklog`
- [ ] Export function
- [ ] `start`, `list`, `show`, `delete` arguments

## Example

```bash
worklog "Sample Project" "customer=ACME" "type=meeting"
Started tracking...

Press CTRL+D to stop.
^C

Title     : Sample Project
Start     : 12:00:00
Stop      : 13:37:00
Duration  : 01h, 37m
```
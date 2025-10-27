# timer

 Simple CLI command to setup timer that triggers notification when done.

## Usage

Sets timer for 1 hour 45 minutes 20 seconds:
```bash
timer -title="Notification title" -msg="Notification message" -h=1 -m=45 -s=20
```

Sets timer to specific time:
```bash
timer -t=18:05:29

# seconds can be omitted
timer -t=18:05
```

Specify either any of `h`, `m`, `s` flags or `t` flag alone.

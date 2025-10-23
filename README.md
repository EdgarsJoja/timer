# timer

 Simple CLI command to setup timer that triggers notification when done.

## Usage

Sets timer for 1 hour 15 minutes 20 seconds:
```bash
timer -title="Notification title" -message="Notification message" -th=1 -tm=45 -ts=20
```

Sets timer to specific time:
```bash
timer -t=18:05:29
```

Specify either any of `th`, `tm`, `ts` flags or `t` flag alone.

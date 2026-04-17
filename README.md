# timer

 Simple CLI command to setup timer that triggers notification when done.

 ![Timer notification](docs/screenshot-1.png "Timer notification")

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


## Run in background

Run program in background and disown it from the shell.

### zsh
```zsh
timer -m 1 &!
```

### bash
```bash
timer -m 1 & disown
```

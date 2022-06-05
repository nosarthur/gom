# gom: job monitor



## design

The CUI

```
gom start
gom show [-v] [-t type]
gom quit
gom config
```

- Here `start` spins up the server if it's not up, or resets the stats for a
  running server
- `-t` picks running fail success
- `config` shows the current settings, including
  - refresh interval
  - cmd.status
  - cmd.fail
  - port

## TODO

- individual sessions, maybe with secret key for each session

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

The config file is assumed to be at home dir.

## TODO

- [x] define .yaml file for server config
- [ ] subprocess call to the commands
- [ ] parse status command and populate data store
- [ ] individual sessions, maybe with secret key for each session

package monitor

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

// Update job status
func heartbeat() {
	c := GetConf()

	ticker := time.NewTicker(time.Duration(c.Interval) * time.Second)
	defer ticker.Stop()
	cmd := strings.Fields(c.CmdStatus)
	for range ticker.C {
		// no need to return error
		out, err := exec.Command(cmd[0], cmd[1:]...).Output()
		if err != nil {
			log.Println(err)
		}
		got := parseStatus(string(out))
		updateStore(got)
	}
}

func parseState(s string) string {
	//14:01 eigen3 ~ $ sacct -bn -j 14722002
	// 14722002        RUNNING      0:0
	fields := strings.Fields(s)
	return fields[1]
}

func parseStatus(s string) map[string]bool {
	IDs := map[string]bool{}
	lines := strings.Split(s, "\n")
	// Skip first title line
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line != "" {
			IDs[strings.Fields(lines[i])[0]] = true
		}
	}
	return IDs
}

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
	for range ticker.C {
		log.Println("beat===========================")
		// no need to return error
		out, err := exec.Command(c.CmdStatus).Output()
		if err != nil {
			log.Println(err)
		}
		got := parseStatus(string(out))
		updateStore(got)
	}
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

package monitor

import (
	"io"
	"log"
	"os/exec"
	"strings"
)

// TODO: if we need to distinguish session, use map[uint][uint]jobInfo
var store map[string]jobInfo

type jobInfo struct {
	Status string
	// TODO: add timing?
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// Add new jobs and tag finished jobs
func updateStore(newIDs map[string]bool) {
	c := GetConf()

	// Finish old jobs
	cmd := strings.Fields(c.CmdFail)
	args := cmd[1:]
	for id, info := range store {
		if info.Status == "RUNNING" && !newIDs[id] {
			store[id] = jobInfo{"DONE"}
			println("done", id, store[id].Status, "|", len(store))

			out, err := exec.Command(cmd[0], append(args, id)...).Output()
			if err != nil {
				log.Println(err)
			}
			state := parseState(string(out))
			// Only save failed jobs
			if state == "COMPLETED" {
				delete(store, id)
			} else {
				store[id] = jobInfo{state}
			}
		}
	}

	// Add new jobs
	for id := range newIDs {
		if _, ok := store[id]; !ok {
			// FIXME: This is wrong, they could be PENDING or CONFIG
			store[id] = jobInfo{"RUNNING"}
			println(id, store[id].Status, "|", len(store))
		}
	}
}

// Reset the store
func Reset() {
	// FIXME: This doesn't work. It needs to be done in handleConn
	// possibly use this: https://github.com/golang/go/issues/47649
	println(len(store))
	store = make(map[string]jobInfo)
	println("reset stats")
	println(len(store))
}

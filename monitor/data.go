package monitor

import (
	"io"
	"log"
)

// TODO: if we need to distinguish session, use map[uint][uint]jobInfo
var store map[string]jobInfo

type jobStatus int

const (
	RUNNING jobStatus = iota
	DONE
	SUCCESS
	FAIL
)

type jobInfo struct {
	Status jobStatus
	// Status string
	// TODO: add timing?
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func updateStore(newIDs map[string]bool) {

	// Add new jobs
	for id := range newIDs {
		if _, ok := store[id]; !ok {
			store[id] = jobInfo{RUNNING}
			println(id, store[id].Status, "|", len(store))
		}
	}
	// Finish old jobs
	for id, info := range store {
		if info.Status == RUNNING && !newIDs[id] {
			store[id] = jobInfo{DONE}
			println("done", id, store[id].Status, "|", len(store))
		}
	}
}

// Reset the store
func Reset() {
	// FIXME: It's not working, it needs to be done in handleConn
	// possibly use this: https://github.com/golang/go/issues/47649
	println(len(store))
	store = make(map[string]jobInfo)
	println("reset stats")
	println(len(store))
}

package monitor

import (
	"io"
	"log"
)

// TODO: if we need to distinguish session, use map[uint][uint]jobInfo
var store map[uint]jobInfo

type jobStatus int

const (
	SUCCESS jobStatus = iota
	RUNNING
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

func Reset() {
	// possibly use this: https://github.com/golang/go/issues/47649
	store = make(map[uint]jobInfo)
	println("reset stats")
}

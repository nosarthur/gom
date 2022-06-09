package monitor

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"time"
)

// Main goroutine for the server
func Spinup() {
	cf := GetConf()

	listener, err := net.Listen("tcp", ":"+cf.Port)

	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	store = make(map[uint]jobInfo)

	// store[99] = jobInfo{"RUNNING"}
	store[99] = jobInfo{RUNNING}

	println("spin up successful")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		println("client comes")
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		bytes, err := json.MarshalIndent(store, "", "\t")
		if err != nil {
			log.Println(err)
			return
		}
		_, err = io.WriteString(c, string(bytes))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(5 * time.Second)
	}
}

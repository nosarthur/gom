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
	println(cf.Host, cf.Port, cf.CmdStatus, cf.Interval)

	listener, err := net.Listen("tcp", ":"+cf.Port)

	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	store = make(map[string]jobInfo)

	println("spin up successful")
	go heartbeat()

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
		// TODO: not json, use csv
		bytes, err := json.MarshalIndent(store, "", "\t")
		if err != nil {
			log.Println(err)
			return
		}
		// TODO: both read and write
		_, err = io.WriteString(c, string(bytes))
		if err != nil {
			log.Println(err)
			return
		}
		// TODO: no need to refresh?
		time.Sleep(5 * time.Second)
	}
}

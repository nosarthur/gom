package monitor

import (
	"bytes"
	"log"
	"net"
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
	// TODO: not json, use csv
	var buf bytes.Buffer

	for k, v := range store {

		buf.WriteString(k)
		buf.WriteByte(':')
		buf.WriteString(v.Status)
		buf.WriteByte('\n')
	}
	_, err := buf.WriteTo(c)
	if err != nil {
		log.Println(err)
		return
	}
	// TODO: both read and write
	// _, err = io.WriteString(c, string(bytes))
}

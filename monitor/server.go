package monitor

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"time"

	"github.com/nosarthur/gom/config"
)

func Spinup() {
	listener, err := net.Listen("tcp", config.HOST+":"+config.PORT)

	if err != nil {
		log.Fatal(err)
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
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		bytes, err := json.Marshal(store)
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

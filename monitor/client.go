package monitor

import (
	"log"
	"net"
	"os"
	"time"
)

func Connect() {
	cf := GetConf()
	conn, err := net.Dial("tcp", cf.Host+":"+cf.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)

}

func IsRunning() bool {
	timeout := time.Second
	cf := GetConf()
	conn, err := net.DialTimeout("tcp", cf.Host+":"+cf.Port, timeout)
	if err != nil {
		log.Println(err)
		return false
	}
	defer conn.Close()
	println("port in use")
	return true
}

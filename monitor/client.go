package monitor

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/nosarthur/gom/config"
)

func Connect() {
	conn, err := net.Dial("tcp", config.HOST+":"+config.PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)

}

func IsRunning() bool {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", config.HOST+":"+config.PORT, timeout)
	if err != nil {
		log.Println(err)
		println("unused port")
		return false
	}
	defer conn.Close()
	println("port in use")
	return true
}

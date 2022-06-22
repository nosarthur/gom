package monitor

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Host      string
	Port      string
	Interval  uint
	CmdStatus string
	CmdFail   string
}

// Singleton
var cf *Conf

func (c *Conf) String() string {
	return fmt.Sprintf("Host: %s:%s\nRefresh interval: %d seconds\nCommands: %s\n\t%s",
		c.Host, c.Port, c.Interval, c.CmdStatus, c.CmdFail)
}

func GetConf() *Conf {
	if cf != nil {
		return cf
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configPath := filepath.Join(home, "gom.yaml")

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	c := Conf{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	cf = &c
	return cf
}

// Show server configurations
func ShowConfig() {
	fmt.Println(GetConf())

}

package monitor

import (
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

// TODO: make it stringer

func GetConf() *Conf {
	if cf != nil {
		return cf
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	println(configDir)
	configPath := filepath.Join(configDir, "gom.yaml")

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	c := Conf{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	println(c.Host, c.Port, c.CmdStatus)

	return &c
}

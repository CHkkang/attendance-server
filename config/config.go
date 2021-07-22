package config

import (
	"attendance-server/database"
	"attendance-server/router"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)


type Config struct {
	DatabaseConfig 	*database.DBConfig 		`yaml:"mysql"`
	ServerConfig 	*router.ServerConfig 	`yaml:"server"`
}

func Load(path string) *Config {
	var file []byte
	var err error
	var config *Config

	if file, err = ioutil.ReadFile(filepath.Clean(path)); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	if err = yaml.Unmarshal(file, &config); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	return config
}

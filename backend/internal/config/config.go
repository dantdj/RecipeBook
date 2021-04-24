package config

import (
	"encoding/json"
	"io/ioutil"
)

var Configuration Config

type Config struct {
	Mongo MongoConfig `json:"mongo"`
}

type MongoConfig struct {
	ConnectionString string `json:"connectionString"`
}

func LoadConfig() {
	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &Configuration)
	if err != nil {
		panic(err)
	}
}

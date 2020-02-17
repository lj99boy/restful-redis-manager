package repo

import (
	"encoding/json"
	"io/ioutil"
	"restful-redis-manager/paramDict"
)

//type SingleDefaultConfig struct {
//	Addr     string `yaml:"Addr"`
//	Password string `yaml:"Password"`
//	DB       int    `yaml:"DB"`
//}

func GetClusterDefaultService() *paramDict.ClusterInputSource {
	configFileData, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	jsonContent := struct {
		Cluster paramDict.ClusterInputSource `json:"Cluster"`
	}{}

	err = json.Unmarshal(configFileData, &jsonContent)

	if err != nil {
		panic(err)
	}

	return &jsonContent.Cluster
}

func GetSingleDefaultService() *paramDict.SingleInputSource {
	configFileData, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	jsonContent := struct {
		Single paramDict.SingleInputSource `json:"Single"`
	}{}

	err = json.Unmarshal(configFileData, &jsonContent)

	if err != nil {
		panic(err)
	}

	return &jsonContent.Single
}

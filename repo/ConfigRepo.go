package repo

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"restful-redis-manager/ParamDict"
)

//type SingleDefaultConfig struct {
//	Addr     string `yaml:"Addr"`
//	Password string `yaml:"Password"`
//	DB       int    `yaml:"DB"`
//}

func GetSingleDefaultService() *ParamDict.SingleInputSource {
	configFileData, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	t := &ParamDict.SingleInputSource{}

	err = yaml.Unmarshal(configFileData, t)
	if err != nil {
		panic(err)
	}

	return t
}

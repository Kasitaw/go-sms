package configs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name string  `yaml:"name,omitempty"`
	Username string  `yaml:"username,omitempty"`
	Password string  `yaml:"password,omitempty"`
	Url string  `yaml:"url,omitempty"`
}

type Provider struct {
	Drivers []Config
	Default string
}

func Parse() Provider {
	credential, err := ioutil.ReadFile("configs/credential.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var drivers Provider
	err = yaml.Unmarshal(credential, &drivers)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return drivers
}

func GetDriver(drivers []Config, defaultDriver string) Config {
	var config Config
	for _, driver := range drivers {
		if defaultDriver == driver.Name {
			config = driver
		}
	}
	return config
}
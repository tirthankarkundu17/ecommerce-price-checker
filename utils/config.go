package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Selector struct {
	Selector  string
	Type      string
	Attribute string
}

type ConfigSelector struct {
	Name   Selector `yaml:"name"`
	Price  Selector `yaml:"price"`
	Images Selector `yaml:"images"`
	Rating Selector `yaml:"rating"`
}

func GetConf(path string) *ConfigSelector {

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := &ConfigSelector{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

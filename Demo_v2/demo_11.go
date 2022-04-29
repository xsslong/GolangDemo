package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
	"path"
)

type conf struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

func (conf *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return conf
}

func main() {
	var conf conf
	conf.getConf()

	fmt.Println(conf)


	folder := path.Join(conf.getConf().Path, "sitemap")
	fmt.Println(folder)
}

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Redis struct {
	Addr           string `yaml:"Addr"`
	Port           string `yaml:"Port"`
	Password       string `yaml:"PassWord"`
	DbNumber int    `yaml:"DataBaseNumber"`
}

func (this *Redis) DefaultRedisConfig() {
	this.Addr = "localhost"
	this.Port = "6379"
	this.DbNumber = 0
	this.Password = ""
}

func (this *Redis) InitRedisConfig(path string) {
	this.DefaultRedisConfig()
	file, _ := ioutil.ReadFile(path)
	if err := yaml.Unmarshal(file, this); err != nil {
		log.Println("ERROR", err)
	}
}

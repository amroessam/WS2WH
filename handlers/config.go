package handlers

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//Config - Base configuration
type Config struct {
	PushbulletAccounts []SingleConfig `yaml:"pushbullet_accounts"`
}

//SingleConfig - Base configuration
type SingleConfig struct {
	Name        string   `yaml:"name"`
	ID          string   `yaml:"id"`
	WebhookUrls []string `yaml:"webhook_urls"`
}

//ReadConfig reads yaml configuration
func ReadConfig() Config {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Println("Unable to read file...")
		log.Fatalln(err)
	}
	var c Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Println("Unable to parse config")
		log.Fatalln(err)
	}
	return c
}

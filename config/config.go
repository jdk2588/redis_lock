package config

import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
)

type Settings struct {
  Servers []string `yaml:"servers"`
  Mock bool `yaml:"mock"`
  NoOfClients int `yaml:"numberofclients"`
  PartitionType string `yaml:"partitiontype"`
  MessageCount int `yaml:"messages"`
  FailAnyRedis bool `yaml:"failanyredis"`
  FailRedisNode []string `yaml:"failredisnodes"`
  Debug bool `yaml:"debug"`
}

func DefaultConfig() Settings {
  var settings Settings
  settings.Servers = []string{}
  settings.Mock = true
  settings.NoOfClients = 3
  settings.PartitionType = "rack"
  settings.MessageCount = 1
  settings.FailAnyRedis = false
  settings.Debug = true
  return settings

}

func LoadConfig(confPath string) (Settings, error) {
  var settings Settings

  configFile, err := ioutil.ReadFile(confPath)

  if err != nil {
    return settings, err
  }

  err = yaml.Unmarshal([]byte(configFile), &settings)

  if err != nil {
    return settings, err
  }

  return settings, nil
}

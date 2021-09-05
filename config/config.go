package config

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

const JwtExpirationPeriod = time.Hour * 6

var NumberOfTopEntriesToConsider int
var NumberOfLastEntriesToConsiderWhenSearchingByTag int
var Salt string
var CachePeriod int64
var CommonCacheTag string
var PostgresConfig string
var Port string

type ConfigYaml struct {
	PostgresConfig                                  string `yaml:"postgresConfig"`
	Port                                            string `yaml:"port"`
	NumberOfTopEntriesToConsider                    int    `yaml:"numberOfTopEntriesToConsider"`
	NumberOfLastEntriesToConsiderWhenSearchingByTag int    `yaml:"numberOfLastEntriesToConsiderWhenSearchingByTag"`
	Salt                                            string `yaml:"Salt"`
	CachePeriod                                     int64  `yaml:"cachePeriod"`
	CommonCacheTag                                  string `yaml:"commonCacheTag"`
}

func UpdateConfigsFor(env string) {
	configYamlFileName := ""
	if env == "local" {
		configYamlFileName = "config/local/local.yaml"
	}

	configYamlFile, err := ioutil.ReadFile(configYamlFileName)
	if err != nil {
		panic("Error reading config file " + err.Error())
	}

	var configYaml ConfigYaml
	err = yaml.Unmarshal(configYamlFile, &configYaml)
	if err != nil {
		panic("Error parsing config")
	}

	PostgresConfig = configYaml.PostgresConfig
	Port = configYaml.Port
	NumberOfTopEntriesToConsider = configYaml.NumberOfTopEntriesToConsider
	NumberOfLastEntriesToConsiderWhenSearchingByTag = configYaml.NumberOfLastEntriesToConsiderWhenSearchingByTag
	Salt = configYaml.Salt
	CachePeriod = configYaml.CachePeriod
	CommonCacheTag = configYaml.CommonCacheTag
}

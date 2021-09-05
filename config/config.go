package config

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

const JwtExpirationPeriod = time.Hour * 6

var NumberOfTopEntriesToConsider int
var NumberOfLastEntriesToConsiderWhenSearchingByTag int
var Salt = "eoriu69045trhgeo4t5780ejr9t78340gjtu9eu9034oij490et74564564asdas8dgyas8d68743r584frth437ry49r58478r5984r984rewhf8943hf349yfg34921eu23ru3efr4eftu43t784tf980eufg8u4f4uft"
var CachePeriod int
var CommonCacheTag string
var PostgresConfig string
var Port string

type ConfigYaml struct {
	PostgresConfig                                  string `yaml:"postgresConfig"`
	Port                                            string `yaml:"port"`
	NumberOfTopEntriesToConsider                    int    `yaml:"numberOfTopEntriesToConsider"`
	NumberOfLastEntriesToConsiderWhenSearchingByTag int    `yaml:"numberOfLastEntriesToConsiderWhenSearchingByTag"`
	Salt                                            string `yaml:"Salt"`
	CachePeriod                                     int    `yaml:"cachePeriod"`
	CommonCacheTag                                  string `yaml:"commonCacheTag"`
}

func UpdateConfigsFor(env string) {
	configYamlFileName := ""
	if env == "local" {
		configYamlFileName = "config/local.yaml"
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

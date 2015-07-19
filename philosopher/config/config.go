package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	PhilosopherNames []string `yaml:"philosopher_names"`
	MinThinkTimeMs int `yaml:"min_think_time_ms"`
	MaxThinkTimeMs int `yaml:"max_think_time_ms"`
	MinEatTimeMs int `yaml:"min_eat_time_ms"`
	MaxEatTimeMs int `yaml:"max_eat_time_ms"`
	RefreshStateTimeMs int `yaml:"refresh_state_time_ms"`
}

func GetConfig(filepath string) (Config, error) {

	yamlFileByteArray, err := getYamlFile(filepath)

	if err != nil {
		return Config{[]string{"Zeno", "Plato", "Epicurus", "Locke", "Aristotle"}, 10, 50, 200, 400, 500}, nil
	}

	return unmarshalFileByteArrayIntoConfig(yamlFileByteArray)
}

func getYamlFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

func unmarshalFileByteArrayIntoConfig(yamlFileByteArray []byte) (Config, error) {
	config := Config{}
	err := yaml.Unmarshal(yamlFileByteArray, &config)
	return config, err
}
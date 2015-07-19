package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestItParseYamlFile(t *testing.T) {
	expectedConfig  := Config{[]string{"Zeno", "Plato", "Epicurus", "Locke", "Aristotle"}, 10, 50, 200, 400, 500}
	config, err := GetConfig("fixtures/confTest.yml")

	assert.Nil(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestItGetDefaultConfIfConfIsMissed(t *testing.T) {
	expectedConfig  := Config{[]string{"Zeno", "Plato", "Epicurus", "Locke", "Aristotle"}, 10, 50, 200, 400, 500}
	config, err := GetConfig("fixtures/missed.yml")

	assert.Nil(t, err)
	assert.Equal(t, expectedConfig, config)
}

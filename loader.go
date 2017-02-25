package caracolines

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// LoadConfig loads config structure from a yaml file
func LoadConfig(filename string) (Config, error) {
	yamlFile, err := getFileContent(filename)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

// LoadServices loads services structure from a yaml file
func LoadServices(filename string) (Services, error) {
	yamlFile, err := getFileContent(filename)
	if err != nil {
		return Services{}, err
	}

	var services Services
	err = yaml.Unmarshal(yamlFile, &services)
	if err != nil {
		return Services{}, err
	}

	return services, nil
}

func getFileContent(filename string) ([]byte, error) {
	filepath, _ := filepath.Abs(filename)
	return ioutil.ReadFile(filepath)
}

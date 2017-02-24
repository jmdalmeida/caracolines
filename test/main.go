package main

import (
	"io/ioutil"
	"path/filepath"

	"fmt"

	"gopkg.in/yaml.v2"
)

type structure struct {
	items map[string]interface{} `yaml:"parameters"`
}

func main() {
	params, _ := getFileContent("config.yaml")

	fmt.Println(params)
}

func getFileContent(filename string) (structure, error) {
	filepath, _ := filepath.Abs(filename)
	yamlFile, err := ioutil.ReadFile(filepath)

	if err != nil {
		return structure{}, err
	}

	var parameters structure
	err = yaml.Unmarshal(yamlFile, &parameters)
	if err != nil {
		return structure{}, err
	}

	return parameters, nil
}

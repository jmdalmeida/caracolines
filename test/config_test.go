package test

import (
	"reflect"
	"testing"

	"github.com/jmdalmeida/caracolines"
)

func TestConfig(t *testing.T) {
	_, err := caracolines.LoadConfig("asd.yaml")
	if err == nil {
		t.Error("should have returned Error")
	}

	config, err := caracolines.LoadConfig("config.yaml")
	if err != nil {
		t.Error("should not give an error")
	}
	parameters := config.Parameters
	if len(parameters) != 6 {
		t.Error("should have len of 6")
	}
	if typeof(parameters["stringTest"]) != "string" {
		t.Error("should be string")
	}
	if typeof(parameters["intTest"]) != "int" {
		t.Error("should be int")
	}
	if typeof(parameters["floatTest"]) != "float64" {
		t.Error("should be float64")
	}
	if typeof(parameters["mapTest"]) != "map[interface {}]interface {}" {
		t.Error("should be map[interface {}]interface {}")
	}
	if typeof(parameters["arrayTest"]) != "[]interface {}" {
		t.Error("should be []interface {}")
	}
	if typeof(parameters["boolTest"]) != "bool" {
		t.Error("should be bool")
	}

	if parameters["stringTest"] != "world" {
		t.Error("should have value 'world'")
	}

	mapTest := parameters["mapTest"].(map[interface{}]interface{})
	if typeof(mapTest["name"]) != "string" {
		t.Error("should be string")
	}
	if typeof(mapTest["age"]) != "int" {
		t.Error("should be int")
	}
	if mapTest["name"] != "chines" {
		t.Error("should have value 'string'")
	}
	if mapTest["age"] != 24 {
		t.Error("should have value '24'")
	}

	arrayTest := parameters["arrayTest"].([]interface{})
	if len(arrayTest) != 4 {
		t.Error("should have len 4")
	}
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

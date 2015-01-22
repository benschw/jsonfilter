package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var _ = log.Print

func parseBytes(b []byte) (interface{}, error) {
	var obj interface{}

	err := json.Unmarshal(b, &obj)

	return obj, err
}

func getMap(in interface{}) (map[string]interface{}, error) {
	return in.(map[string]interface{}), nil
}

func getSlice(in interface{}) ([]interface{}, error) {
	return in.([]interface{}), nil
}

func selectValue(obj interface{}, selector string) (string, error) {
	parts := strings.Split(selector, ".")

	if len(parts) == 1 && parts[0] == "" {
		return fmt.Sprintf("%s", obj), nil

	} else if i, err := strconv.Atoi(parts[0]); err == nil {
		s, err := getSlice(obj)
		if err != nil {
			return "", err
		}
		return selectValue(s[i], strings.Join(parts[1:], "."))
	} else {
		s, err := getMap(obj)
		if err != nil {
			return "", err
		}
		return selectValue(s[parts[0]], strings.Join(parts[1:], "."))
	}

}

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	obj, err := parseBytes(b)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", obj)
}

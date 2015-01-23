package main

import (
	"encoding/json"
	"errors"
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

func formatForJsonDisplay(i interface{}) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}
func formatForDisplay(i interface{}, asJson bool) (string, error) {
	if asJson {
		return formatForJsonDisplay(i)
	} else {
		switch v := i.(type) {
		case int:
			return formatForJsonDisplay(i)
		case float64:
			return formatForJsonDisplay(i)
		case bool:
			return formatForJsonDisplay(i)
		case string:
			return fmt.Sprintf("%s", v), nil
		case []interface{}:
			return fmt.Sprintf("%v", v), nil
		case map[string]interface{}:
			return fmt.Sprintf("%v", v), nil
		default:
			return "", errors.New(fmt.Sprintf("Display error, unknown type: %+v", v))
		}
	}
}

func selectValue(obj interface{}, selector string) (interface{}, error) {
	parts := strings.Split(selector, ".")

	if len(parts) == 1 && parts[0] == "" {
		return obj, nil
	}

	switch v := obj.(type) {
	case []interface{}:
		i, err := strconv.Atoi(parts[0])
		if err != nil {
			return "", err
		}
		return selectValue(v[i], strings.Join(parts[1:], "."))
	case map[string]interface{}:
		return selectValue(v[parts[0]], strings.Join(parts[1:], "."))
	default:
		return "", errors.New(fmt.Sprintf("Bad selector, unknown type: %+v", v))
	}
}

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	obj, err := parseBytes(b)
	if err != nil {
		// http://www.goinggo.net/2013/11/using-log-package-in-go.html
		panic(err)
	}

	fmt.Printf("%+v", obj)
}

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
func formatForDisplay(i interface{}) (string, error) {
	switch v := i.(type) {
	case int:
		return strconv.Itoa(v), nil
	case float64:
		return strconv.FormatFloat(v, 'f', 6, 64), nil
	case string:
		return v, nil
	case []interface{}:
		return formatForJsonDisplay(i)
	case map[string]interface{}:
		return formatForJsonDisplay(i)
	default:
		return "", errors.New("unknown type")
	}
}

func selectValue(obj interface{}, selector string) (interface{}, error) {
	parts := strings.Split(selector, ".")

	if len(parts) == 1 && parts[0] == "" {
		return obj, nil
	}

	switch v := obj.(type) {
	case int:
		return "", errors.New("selector not valid")
	case float64:
		return "", errors.New("selector not valid")
	case string:
		return "", errors.New("selector not valid")
	case []interface{}:
		i, err := strconv.Atoi(parts[0])
		if err != nil {
			return "", err
		}
		return selectValue(v[i], strings.Join(parts[1:], "."))
	case map[string]interface{}:
		return selectValue(v[parts[0]], strings.Join(parts[1:], "."))
	default:
		return "", errors.New("unknown type")
	}

	// if len(parts) == 1 && parts[0] == "" {
	// 	return fmt.Sprintf("%s", obj), nil

	// } else if i, err := strconv.Atoi(parts[0]); err == nil {
	// 	s, err := getSlice(obj)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	return selectValue(s[i], strings.Join(parts[1:], "."))
	// } else {
	// 	s, err := getMap(obj)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	return selectValue(s[parts[0]], strings.Join(parts[1:], "."))
	// }

}

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	obj, err := parseBytes(b)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", obj)
}

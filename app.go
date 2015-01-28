package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

var _ = log.Print

type App struct {
	AsJson   bool
	Pretty   bool
	AsValues bool
	AsKeys   bool
	Debug    bool
	Input    io.Reader
}

func (a *App) run(selector string) (string, error) {

	raw, err := parseReader(a.Input)
	if err != nil {
		return "", err
	}

	found, err := selectValue(raw, selector)
	if err != nil {
		return "", err
	}
	if found == nil {
		return "", errors.New(fmt.Sprint("Selector didn't find anything"))
	}
	var strs []string
	if a.AsValues {
		strs, err = formatValuesForDisplay(found, a.AsKeys, a.AsJson, a.Pretty)
		if err != nil {
			return "", err
		}
	} else {
		str, err := formatForDisplay(found, a.AsJson, a.Pretty)
		if err != nil {
			return "", err
		}
		strs = []string{str}
	}
	str := strings.Join(strs, "\n") + "\n"

	return str, nil
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
			return nil, err
		}
		return selectValue(v[i], strings.Join(parts[1:], "."))
	case map[string]interface{}:
		return selectValue(v[parts[0]], strings.Join(parts[1:], "."))
	default:
		return nil, errors.New(fmt.Sprintf("Bad selector, unknown type: %+v", v))
	}
}

func formatValuesForDisplay(i interface{}, asKeys bool, asJson bool, pretty bool) ([]string, error) {
	strs := make([]string, 0)
	switch v := i.(type) {
	case []interface{}:
		for idx, val := range v {
			var str string
			if asKeys {
				str = strconv.Itoa(idx)
			} else {
				var err error
				str, err = formatForDisplay(val, asJson, pretty)
				if err != nil {
					return strs, err
				}
			}
			strs = append(strs, str)
		}
	case map[string]interface{}:
		for key, val := range v {
			var str string
			var err error
			if asKeys {
				str, err = formatForDisplay(key, asJson, pretty)
			} else {
				str, err = formatForDisplay(val, asJson, pretty)
			}
			if err != nil {
				return strs, err
			}
			strs = append(strs, str)
		}

	default:
		return strs, errors.New(fmt.Sprintf("Bad selector, unknown type: %+v", v))
	}

	return strs, nil
}
func formatForDisplay(i interface{}, asJson bool, pretty bool) (string, error) {
	var str string
	var err error = nil
	if asJson {
		if pretty {
			str, err = formatForPrettyJsonDisplay(i)
		} else {
			str, err = formatForJsonDisplay(i)
		}
	} else {
		switch v := i.(type) {
		case int:
			str, err = formatForJsonDisplay(i)
		case float64:
			str, err = formatForJsonDisplay(i)
		case bool:
			str, err = formatForJsonDisplay(i)
		case string:
			str = fmt.Sprintf("%s", v)
		case []interface{}:
			str = fmt.Sprintf("%v", v)
		case map[string]interface{}:
			str = fmt.Sprintf("%v", v)
		default:
			err = errors.New(fmt.Sprintf("Display error, unknown type: %+v", v))
		}
	}
	if err != nil {
		return "", nil
	} else {
		return str, nil
	}
}

func formatForJsonDisplay(i interface{}) (string, error) {
	b, err := json.Marshal(i)
	return string(b[:]), err
}

func formatForPrettyJsonDisplay(i interface{}) (string, error) {
	b, err := json.MarshalIndent(i, "", "    ")
	return string(b[:]), err
}

func parseReader(in io.Reader) (interface{}, error) {
	var obj interface{}

	err := json.NewDecoder(in).Decode(&obj)

	return obj, err
}

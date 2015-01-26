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

	var str string
	if a.AsValues {
		str, err = formatValuesForDisplay(found, a.AsJson, a.Pretty)
	} else {
		str, err = formatForDisplay(found, a.AsJson, a.Pretty)
	}
	if err != nil {
		return "", err
	}

	return str, nil
}

func parseReader(in io.Reader) (interface{}, error) {
	var obj interface{}

	err := json.NewDecoder(in).Decode(&obj)

	return obj, err
}

func formatForJsonDisplay(i interface{}, pretty bool) (string, error) {
	var b []byte
	var err error

	if pretty {
		b, err = json.MarshalIndent(i, "", "    ")
	} else {
		b, err = json.Marshal(i)
	}
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}
func formatValuesForDisplay(i interface{}, asJson bool, pretty bool) (string, error) {
	return formatForDisplay(i, asJson, pretty)
}
func formatForDisplay(i interface{}, asJson bool, pretty bool) (string, error) {
	var str string
	var err error = nil
	if asJson {
		str, err = formatForJsonDisplay(i, pretty)
	} else {
		switch v := i.(type) {
		case int:
			str, err = formatForJsonDisplay(i, false)
		case float64:
			str, err = formatForJsonDisplay(i, false)
		case bool:
			str, err = formatForJsonDisplay(i, false)
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
		return fmt.Sprintf("%s\n", str), nil
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

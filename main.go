package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var _ = log.Print

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
func formatForDisplay(i interface{}, asJson bool, asValues bool, pretty bool) (string, error) {
	if asJson {
		return formatForJsonDisplay(i, pretty)
	} else {
		switch v := i.(type) {
		case int:
			return formatForJsonDisplay(i, false)
		case float64:
			return formatForJsonDisplay(i, false)
		case bool:
			return formatForJsonDisplay(i, false)
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
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	flag.Usage = func() {
		fmt.Printf("Usage: jsonfilter [options] selector-string\n\nOptions:\n")
		flag.PrintDefaults()
	}

	asJson := flag.Bool("json", false, "display output as json")
	pretty := flag.Bool("pretty", false, "display pretty json (sets -json=true)")
	asValues := flag.Bool("each", false, "display each value of selected structure on its own line")
	verbose := flag.Bool("v", false, "display errors")
	debug := flag.Bool("vv", false, "display errors and info")

	flag.Parse()

	if *pretty {
		asJson = pretty
	}
	if *debug {
		verbose = debug
	}

	selector := ""
	for i := 0; i < flag.NArg(); i++ {
		selector += flag.Arg(i)
	}

	raw, err := parseReader(os.Stdin)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	found, err := selectValue(raw, selector)
	if err != nil {
		if *verbose {
			log.Println(err)
		}
		os.Exit(1)
	}
	_ = asValues
	str, err := formatForDisplay(found, *asJson, *asValues, *pretty)
	if err != nil {
		if *verbose {
			log.Println(err)
		}
		os.Exit(1)
	}

	fmt.Println(str)

}

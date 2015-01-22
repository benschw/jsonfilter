package main

import (
	"reflect"
	"testing"
)

var b = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"],"Address":{"Street":"0001 Cemetery Lane"}}`)

func Test_parseJson(t *testing.T) {
	// given

	expected := map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
	}

	// when
	actual, err := parseBytes(b)

	// then
	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	m, _ := getMap(actual)

	if expected["Name"] != m["Name"] {
		t.Errorf("unexpected parsed obj:\n%+v\n%+v", expected, m)
	}

	m["Parents"], _ = getSlice(m["Parents"])

	if !reflect.DeepEqual(expected["Parents"], m["Parents"]) {
		t.Errorf("unexpected parsed obj:\n%+v\n%+v", expected, m)
	}

}

func Test_valueSelector(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := "Wednesday"
	selector := "Name"

	// when
	found, err := selectValue(obj, selector)

	// then

	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s", found)
	}
}
func Test_sliceSelector(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := "Morticia"
	selector := "Parents.1"

	// when
	found, err := selectValue(obj, selector)

	// then

	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s", found)
	}
}

func Test_mapSelector(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := "0001 Cemetery Lane"
	selector := "Address.Street"

	// when
	found, err := selectValue(obj, selector)

	// then

	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s", found)
	}
}

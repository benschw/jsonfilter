package main

import (
	"testing"
)

var b = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"],"Address":{"Street":"0001 Cemetery Lane"}}`)

func Test_valueSelector(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := "Wednesday"
	selector := "Name"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i)
	// then

	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s:%s", expected, found)
	}
}
func Test_sliceSelector(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := "Morticia"
	selector := "Parents.1"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i)

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
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i)

	// then

	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s", found)
	}
}

func Test_jsonSelector(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := `["Gomez","Morticia"]`
	selector := "Parents"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i)

	// then

	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s", found)
	}
}

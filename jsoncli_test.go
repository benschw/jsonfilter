package main

import (
	"testing"
)

var b = []byte(`{"Name":"Wednesday","Age":6,"Height":2.05,"Adult":false,"Parents":["Gomez","Morticia"],"Address":{"Street":"0001 Cemetery Lane"}}`)

func Test_sliceSelector(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := "Morticia"
	selector := "Parents.1"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false)

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
	found, err := formatForDisplay(i, false)

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
	found, err := formatForDisplay(i, true)

	// then

	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s", found)
	}
}

func Test_stringValue(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := "Wednesday"
	selector := "Name"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false)
	// then

	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s:%s", expected, found)
	}
}

func Test_stringJsonValue(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := `"Wednesday"`
	selector := "Name"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, true)
	// then

	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s:%s", expected, found)
	}
}

func Test_intValue(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := "6"
	selector := "Age"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false)
	foundJson, err2 := formatForDisplay(i, true)

	// then
	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}
	if err2 != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s", found)
	}
	if foundJson != expected {
		t.Errorf("Found value wrong: %s", found)
	}
}

func Test_float64Value(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := "2.05"
	selector := "Height"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false)
	foundJson, err2 := formatForDisplay(i, true)

	// then
	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}
	if err2 != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s", found)
	}
	if foundJson != expected {
		t.Errorf("Found value wrong: %s", found)
	}
}

func Test_boolValue(t *testing.T) {
	// given
	obj, _ := parseBytes(b)
	expected := "false"
	selector := "Adult"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false)
	foundJson, err2 := formatForDisplay(i, true)

	// then
	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}
	if err2 != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s", found)
	}
	if foundJson != expected {
		t.Errorf("Found value wrong: %s", found)
	}
}

package main

import (
	"testing"
)

func Test_stringValue(t *testing.T) {
	// given
	obj, _ := parseReader(NewStubReader())
	expected := "asdf"
	selector := "myString"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false, false)
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
	obj, _ := parseReader(NewStubReader())
	expected := `"asdf"`
	selector := "myString"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, true, false)
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
	obj, _ := parseReader(NewStubReader())
	expected := "6"
	selector := "myInt"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false, false)
	foundJson, err2 := formatForDisplay(i, true, false)

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
	obj, _ := parseReader(NewStubReader())
	expected := "2.05"
	selector := "myDecimal"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false, false)
	foundJson, err2 := formatForDisplay(i, true, false)

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
	obj, _ := parseReader(NewStubReader())
	expected := "false"
	selector := "myBool"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false, false)
	foundJson, err2 := formatForDisplay(i, true, false)

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

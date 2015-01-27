package main

import (
	"testing"
)

func Test_sliceSelector(t *testing.T) {
	// given
	obj, _ := parseReader(NewStubReader())
	expected := "bar"
	selector := "myArray.1"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false, false)

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
	obj, _ := parseReader(NewStubReader())
	expected := "bah"
	selector := "myObj.baz"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false, false)

	// then

	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s", found)
	}
}
func Test_deepSelector(t *testing.T) {
	// given
	obj, _ := parseReader(NewStubReader())
	expected := "efoo"
	selector := "a.bc.2.1.ea"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, false, false)

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
	obj, _ := parseReader(NewStubReader())
	expected := `["foo","bar"]`
	selector := "myArray"

	// when
	i, _ := selectValue(obj, selector)
	found, err := formatForDisplay(i, true, false)

	// then

	if err != nil {
		t.Errorf("Parse Error: %s", err)
	}

	if found != expected {
		t.Errorf("Found value wrong: %s", found)
	}
}

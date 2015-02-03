package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_ArrayValues(t *testing.T) {
	// given
	obj, _ := parseReader(NewStubReader())
	expected := []string{"foo", "bar"}
	expectedJson := []string{`"foo"`, `"bar"`}
	selector := "myArray"

	// when
	i, _ := selectValue(obj, selector)
	found, _ := formatValuesForDisplay(i, false, false, false)
	foundJson, _ := formatValuesForDisplay(i, false, true, false)
	// then

	if !reflect.DeepEqual(found, expected) {
		t.Errorf("Found value wrong: %s:%s", expected, found)
	}
	if !reflect.DeepEqual(foundJson, expectedJson) {
		t.Errorf("Found value wrong: %s:%s", expectedJson, foundJson)
	}
}

func Test_ArrayKeys(t *testing.T) {
	// given
	obj, _ := parseReader(NewStubReader())
	expected := []string{"0", "1"}
	selector := "myArray"

	// when
	i, _ := selectValue(obj, selector)
	found, _ := formatValuesForDisplay(i, true, false, false)
	foundJson, _ := formatValuesForDisplay(i, true, true, false)
	// then

	if !reflect.DeepEqual(found, expected) {
		t.Errorf("Found value wrong: %s:%s", expected, found)
	}
	if !reflect.DeepEqual(foundJson, expected) {
		t.Errorf("Found value wrong: %s:%s", expected, foundJson)
	}
}

func Test_MapValues(t *testing.T) {
	// given
	obj, _ := parseReader(NewStubReader())
	expected := []string{"bah", "cah"}
	expectedJson := []string{`"bah"`, `"cah"`}
	sort.Strings(expected)
	sort.Strings(expectedJson)
	selector := "myObj"

	// when
	i, _ := selectValue(obj, selector)
	found, _ := formatValuesForDisplay(i, false, false, false)
	foundJson, _ := formatValuesForDisplay(i, false, true, false)

	// then
	sort.Strings(found)
	sort.Strings(foundJson)

	if !reflect.DeepEqual(found, expected) {
		t.Errorf("Found value wrong: %s:%s", expected, found)
	}
	if !reflect.DeepEqual(foundJson, expectedJson) {
		t.Errorf("Found value wrong: %s:%s", expectedJson, foundJson)
	}
}

func Test_MapKeys(t *testing.T) {
	// given
	obj, _ := parseReader(NewStubReader())
	expected := []string{"baz", "caz"}
	expectedJson := []string{`"baz"`, `"caz"`}
	sort.Strings(expected)
	sort.Strings(expectedJson)
	selector := "myObj"

	// when
	i, _ := selectValue(obj, selector)
	found, _ := formatValuesForDisplay(i, true, false, false)
	foundJson, _ := formatValuesForDisplay(i, true, true, false)

	// then
	sort.Strings(found)
	sort.Strings(foundJson)
	if !reflect.DeepEqual(found, expected) {
		t.Errorf("Found value wrong: %s:%s", expected, found)
	}
	if !reflect.DeepEqual(foundJson, expectedJson) {
		t.Errorf("Found value wrong: %s:%s", expectedJson, foundJson)
	}
}

package main

import (
	"bytes"
	"io"
	"testing"
)

func NewStubReader() io.Reader {
	return bytes.NewReader([]byte(`
		{
			"myString":"asdf",
			"myInt":6,
			"myDecimal":2.05,
			"myBool":false,
			"myArray":[
				"foo",
				"bar"
			],
			"myObj":{
				"baz":"bah"
			},
			"a":{
				"ba":"afoo",
				"bb":"abar",
				"bc": [
					"ca",
					"cb",
					[
						"da",
						{
							"ea": "efoo"
						}
					]

				]
			}
		}
	`))
}

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

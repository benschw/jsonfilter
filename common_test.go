package main

import (
	"bytes"
	"io"
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
				"baz":"bah",
				"caz":"cah"
			},
			"myMapCollection":{
				"a":{
					"id": "a",
					"value": "foo"
				},
				"b":{
					"id": "b",
					"value": "bar"
				}
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

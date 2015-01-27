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

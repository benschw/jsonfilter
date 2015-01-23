# jsonfilter

cli tool to grab values out of a json block



## example

	$ cat data.json | jsonfilter Hobbies.0
	homicide
	$ cat data.json | jsonfilter Address.Street
	0001 Cemetery Lane


### data.json


	{
		"Name": "Wednesday",
		"Age": 6,
		"Parents": [{
			"Relation": "father",
			"Name": "Gomez"
		},{
			"Relation": "mother",
			"Name": "Morticia"
		}],
		"Address": {
			"Street": "0001 Cemetery Lane",
			"City": "New York",
			"State": "New York"
		},
		"Hobbies": [
			"homicide",
			"playing with her headless Marie Antoinette doll",
			"spiders"
		]

	}


## todo

	$ cat data.json | jsonfilter Address
	{"Street":"0001 Cemetery Lane","City": "New York","State":"New York"}

	$ cat data.json | jsonfilter -pretty Address
	{
		"Street": "0001 Cemetery Lane",
		"City": "New York",
		"State": "New York"
	}


	$ cat data.json | jsonfilter Hobbies
	["homicide","playing with her headless Marie Antoinette doll","spiders"]

	$ cat data.json | jsonfilter -values Hobbies
	homicide
	playing with her headless Marie Antoinette doll
	spiders


	$ cat data.json | jsonfilter Address | jsonfilter Street
	0001 Cemetery Lane


	$ cat data.json | jsonfilter Address && echo found
	found

	$ cat data.json | jsonfilter Foo || echo not found
	not found




### notes
- multi line input?
- error handling


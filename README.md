# jsonfilter

`jsonfilter` is an utility for filtering and selecting values from a json object.




### wednesday.json


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


## suite: jsonfilter examples
### test: should exit 1 when selector is NOT found
#### when:
	cat ./wednesday.json | ./jsonfilter Flub > /dev/null || echo not found

#### then:
	not found


### test: should drill down when using compound selector
#### when:
	cat ./wednesday.json | ./jsonfilter Address.Street

#### then:
	0001 Cemetery Lane

### test: -values should output values in an array
#### when:
	cat ./wednesday.json | ./jsonfilter -values Hobbies

#### then:
	homicide
	playing with her headless Marie Antoinette doll
	spiders

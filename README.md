[![Build Status](https://drone.io/github.com/benschw/jsonfilter/status.png)](https://drone.io/github.com/benschw/jsonfilter/latest)

# jsonfilter

`jsonfilter` is an utility for filtering and selecting values from a json object.

## Install

download the bin to `/usr/local/bin`

	wget -qO- \
	https://drone.io/github.com/benschw/jsonfilter/files/build/output/jsonfilter.gz \
	| gunzip > /usr/local/bin/jsonfilter


or get the `.deb`

	wget https://drone.io/github.com/benschw/jsonfilter/files/build/output/jsonfilter.deb
	sudo dpkg --install jsonfilter.deb



## Examples 
examples guaranteed up to date by [cli-unit](https://github.com/benschw/cli-unit).
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
### test: should exit 0 when selector is found
#### when:
	cat ./wednesday.json | ./jsonfilter Address.Street > /dev/null && echo found

#### then:
	found


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

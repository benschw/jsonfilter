[![Build Status](https://drone.io/github.com/benschw/jsonfilter/status.png)](https://drone.io/github.com/benschw/jsonfilter/latest)

# jsonfilter

`jsonfilter` is an utility for filtering and selecting values from a json object.

## Install

download the bin to `/usr/local/bin`

	wget -qO- \
	https://drone.io/github.com/benschw/jsonfilter/files/build/output/jsonfilter-`uname`-`uname -m`.gz \
	| gunzip > /usr/local/bin/jsonfilter


or get the `.deb`

	wget https://drone.io/github.com/benschw/jsonfilter/files/build/output/jsonfilter-amd64.deb
	sudo dpkg --install jsonfilter.deb


or from the Package Repository (for debian machines like ubuntu)

	sudo add-apt-repository "deb http://dl.bintray.com/benschw/deb wheezy main"
	sudo apt-get update
	sudo apt-get install jsonfilter


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
	cat ./wednesday.json | ./build/output/jsonfilter Address.Street > /dev/null && echo found

#### then:
	found


### test: should drill down when using compound selector
#### when:
	cat ./wednesday.json | ./build/output/jsonfilter Address.Street

#### then:
	0001 Cemetery Lane

### test: -values should output values in an array
#### when:
	cat ./wednesday.json | ./build/output/jsonfilter -values Hobbies

#### then:
	homicide
	playing with her headless Marie Antoinette doll
	spiders

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


or from the Package Repository (for debian machines like ubuntu)

	echo "deb http://dl.bintray.com/benschw/deb wheezy main" | sudo tee -a /etc/apt/sources.list.d/benschw.list
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

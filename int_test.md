### test: -json should work with objects
#### when:
	cat ./wednesday.json | ./jsonfilter -json Address

#### then:
	{"City":"New York","State":"New York","Street":"0001 Cemetery Lane"}


### test: -json should work with arrays
#### when:
	cat ./wednesday.json | ./jsonfilter -json Hobbies

#### then:
	["homicide","playing with her headless Marie Antoinette doll","spiders"]


### test: -pretty should make pretty json output
#### when:
	cat ./wednesday.json | ./jsonfilter -pretty Address

#### then:
	{
	    "City": "New York",
	    "State": "New York",
	    "Street": "0001 Cemetery Lane"
	}


### test: -values should output values in an array
#### when:
	cat ./wednesday.json | ./jsonfilter -values Hobbies

#### then:
	homicide
	playing with her headless Marie Antoinette doll
	spiders

### test: -values -json should output json values in an array
#### when:
	cat ./wednesday.json | ./jsonfilter -values -json Hobbies

#### then:
	"homicide"
	"playing with her headless Marie Antoinette doll"
	"spiders"


### test: should be able to filter object multiple times
#### when:

	JSON=$(cat ./wednesday.json | ./jsonfilter -values -json Parents)

	while read -r parent; do
		echo "$parent" | ./jsonfilter Name
	done <<< "$JSON"

#### then:
	Gomez
	Morticia


### test: should be able to filter object multiple times with pipes
#### when:

	cat ./wednesday.json | ./jsonfilter -json Address | ./jsonfilter Street

#### then:
	0001 Cemetery Lane


### test: should exit 0 when selector is found
#### when:

	cat ./wednesday.json | ./jsonfilter Address > /dev/null && echo found

#### then:
	found

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


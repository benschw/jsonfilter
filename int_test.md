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
	cat ./wednesday.json | ./build/output/jsonfilter -values Hobbies

#### then:
	homicide
	playing with her headless Marie Antoinette doll
	spiders

### test: -values -json should output json values in an array
#### when:
	cat ./wednesday.json | ./build/output/jsonfilter -values -json Hobbies

#### then:
	"homicide"
	"playing with her headless Marie Antoinette doll"
	"spiders"


### test: should be able to filter object multiple times
#### when:
	JSON=$(cat ./wednesday.json | ./build/output/jsonfilter -values -json Parents)
	
	while read -r parent; do
		echo "$parent" | ./build/output/jsonfilter Name
	done <<< "$JSON"

#### then:
	Gomez
	Morticia


### test: should be able to filter object multiple times with pipes
#### when:
	cat ./wednesday.json | ./build/output/jsonfilter -json Address | ./build/output/jsonfilter Street

#### then:
	0001 Cemetery Lane


### test: should exit 0 when selector is found
#### when:
	cat ./wednesday.json | ./build/output/jsonfilter Address > /dev/null && echo found

#### then:
	found

### test-skip: should exit 1 when selector is NOT found 
make on drone is failing because it inherits this code... 
#### when:
	cat ./wednesday.json | ./build/output/jsonfilter Flub > /dev/null || echo not found

#### then:
	not found


### test: should drill down when using compound selector
#### when:
	cat ./wednesday.json | ./build/output/jsonfilter Address.Street

#### then:
	0001 Cemetery Lane


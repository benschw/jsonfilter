### test: -json should work with objects
### shell:
cat ./wednesday.json | ./jsonfilter -json Address

### output:
{"City":"New York","State":"New York","Street":"0001 Cemetery Lane"}


### test: -json should work with arrays
### shell:
cat ./wednesday.json | ./jsonfilter -json Hobbies

### output:
["homicide","playing with her headless Marie Antoinette doll","spiders"]


### test: -pretty should make pretty json output
### shell:
cat ./wednesday.json | ./jsonfilter -pretty Address

### output:
{
    "City": "New York",
    "State": "New York",
    "Street": "0001 Cemetery Lane"
}


### test: -values should output values in an array
### shell:
cat ./wednesday.json | ./jsonfilter -values Hobbies

### output:
homicide
playing with her headless Marie Antoinette doll
spiders

### test: -values -json should output json values in an array
### shell:
cat ./wednesday.json | ./jsonfilter -values -json Hobbies

### output:
"homicide"
"playing with her headless Marie Antoinette doll"
"spiders"


### test: should be able to filter object multiple times
### shell:

JSON=$(cat ./wednesday.json | ./jsonfilter -values -json Parents)

while read -r parent; do
	echo "$parent" | ./jsonfilter Name
done <<< "$JSON"

### output:
Gomez
Morticia


### test: should be able to filter object multiple times with pipes
### shell:

cat ./wednesday.json | ./jsonfilter -json Address | ./jsonfilter Street

### output:
0001 Cemetery Lane


### test: should exit 0 when selector is found
### shell:

cat ./wednesday.json | ./jsonfilter Address > /dev/null && echo found

### output:
found

### test: should exit 1 when selector is NOT found
### shell:

cat ./wednesday.json | ./jsonfilter Flub > /dev/null || echo not found

### output:
not found


### test: should drill down when using compound selector
### shell:

cat ./wednesday.json | ./jsonfilter Address.Street

### output:
0001 Cemetery Lane


# jsoncli

cli tool to grab values out of a json block



## example

	$ cat data.json | jsoncli Parents.1
	Morticia
	$ cat data.json | jsoncli Address.Street
	0001 Cemetery Lane


### data.json


	{
		"Name":"Wednesday",
		"Age":6,
		"Parents":[
			"Gomez",
			"Morticia"
		],
		"Address":{
			"Street":"0001 Cemetery Lane"
		}
	}


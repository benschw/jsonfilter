#!/bin/bash



getBlock() {
	FILE=$1
	TYPE=$2

	start="0"
	while read p; do
		if [ "$start" == "1" ]; then
			echo -e "$p"
		fi
		if [[ $p == $TYPE ]]; then
			start="1"
		fi

	done <$FILE
}

getShell() {
	FILE=$1
	echo -e $(getBlock $FILE "#shell")
}
getOutput() {
	FILE=$1
	echo -e $(getBlock $FILE "#output")
}

for f in $(ls $1); do
	EXPECTED=$(getOutput $f)
	SHELL=$(getShell $f)
	FOUND=$(eval $SHELL)
	if [[ $EXPECTED != $FOUND ]]; then
		echo -e "expected: \n\t$EXPECTED \nnot equal to found: \n\t$FOUND"
		echo FAIL
		exit 1
	fi
done

echo OK




# getOutput() {
# 	file=$1
# 	start="0"
# 	while read p; do
# 		if [ "$start" == "1" ]; then
# 			echo -e $p | cut -c 3-
# 		fi
# 		if [[ $p == "# "Output:* ]]; then
# 			start="1"
# 		fi

# 	done <$file

# }

# for f in $(ls $1); do
# 	EXPECTED=$(getOutput $f)
# 	FOUND=$(/bin/bash $f)

# 	if [ "$EXPECTED" != "$FOUND" ]; then
# 		echo -e "expected: \n\t\t$EXPECTED \n\tnot equal to found:\n\t\t$FOUND"
# 		echo FAIL
# 		exit 1
# 	fi
# done

# echo OK



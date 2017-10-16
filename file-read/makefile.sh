#!/bin/bash
base=1
end=10
file="./file"
touch $file 
while true
do
	echo $base $end
	if [ $base -gt $end ]
	then
		echo if $base
		break
	fi
	echo $base
	echo $base >> $file
	base=$((base + 1))
done 

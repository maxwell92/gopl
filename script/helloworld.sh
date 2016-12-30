#! /bin/bash


base=0
end=10

while [ $base -lt $end ]
do
    echo "Hello" $1 $base
    base=$((base + 1))
    sleep 2 
done

# echo "Hello World"

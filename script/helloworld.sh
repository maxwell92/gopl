#! /bin/bash


base=0
end=10

while [ $base -lt $end ]
do
    echo "Hello" $1 $base
    base=$((base + 1))
    SLEEPTIME=$((RANDOM % 20))
    sleep $SLEEPTIME 
done

# echo "Hello World"

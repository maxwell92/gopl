#! /bin/bash

SLEEPTIME=$((RANDOM % 20 + 1))
echo Hello $1, sleep $SLEEPTIME
sleep $SLEEPTIME 

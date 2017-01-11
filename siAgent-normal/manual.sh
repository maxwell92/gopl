#! /bin/bash

#DEFAULTHOST: 10.151.30.227
DEFAULTHOST=127.0.0.1
DEFAULTPORT=8081
curl -XPOST -d "{\"image\":\"$1\", \"tag\":\"$2\"}" $DEFAULTHOST:$DEFAULTPORT/sync

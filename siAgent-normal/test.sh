#!/bin/bash

curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v1\"}" localhost:8081/api/v1/sync


curl -XGET localhost:8081/api/v1/list

#!/bin/bash

curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v1\"}" localhost:8081/api/v1/sync
curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v2\"}" localhost:8081/api/v1/sync
curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v3\"}" localhost:8081/api/v1/sync
curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v4\"}" localhost:8081/api/v1/sync
curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v5\"}" localhost:8081/api/v1/sync
curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v6\"}" localhost:8081/api/v1/sync
curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v7\"}" localhost:8081/api/v1/sync
curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v8\"}" localhost:8081/api/v1/sync
curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v9\"}" localhost:8081/api/v1/sync
curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v10\"}" localhost:8081/api/v1/sync
curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v11\"}" localhost:8081/api/v1/sync
# curl -XPOST -d "{\"image\":\"busybox\", \"tag\":\"v12\"}" localhost:8081/api/v12/sync


# curl -XGET localhost:8081/api/v1/list

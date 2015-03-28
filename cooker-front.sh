#!/bin/sh
DIR=$(cd $(dirname $0); pwd)
/usr/bin/go run $DIR/main.go $DIR/redis.go

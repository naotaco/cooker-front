#!/bin/sh
DIR=$(cd $(dirname $0); pwd)
go run $DIR/main.go $DIR/redis.go

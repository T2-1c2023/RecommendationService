#!/bin/sh
export PATH=$PATH:$(go env GOPATH)/bin
swag init -dir ./app
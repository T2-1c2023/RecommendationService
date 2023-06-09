#!/bin/sh
go mod download
go fmt ./...
go vet ./...
go test ./__test__ -cover -coverpkg=./app/controller,./app/routes,./app/model,./app/validation,./app/utilities -coverprofile=coverage.out
go tool cover -html=coverage.out -o ./coverage/index.html
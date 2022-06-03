#!/usr/bin/env sh

mkdir -p bin
GOOS=linux   GOARCH=amd64 go build -v -o bin/luxapisrv     main.go
GOOS=windows GOARCH=amd64 go build -v -o bin/luxapisrv.exe main.go

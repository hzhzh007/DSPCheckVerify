#!/bin/bash
mkdir  bin

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/check_linux -ldflags "-s -w"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/check_windows.exe -ldflags "-s -w"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/check_mac -ldflags "-s -w"

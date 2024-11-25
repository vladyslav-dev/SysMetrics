#!/bin/bash
# Build the main.go file
# go build -o core-build main.go
GOOS=darwin GOARCH=arm64 go build -o sysmetrics main.go 
echo "Build completed. Executable: ./sysmetrics"
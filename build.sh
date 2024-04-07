#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o dns-lookup-latency-linux-amd64 main.go
GOOS=linux GOARCH=arm64 go build -o dns-lookup-latency-linux-arm64 main.go
GOOS=windows GOARCH=amd64 go build -o dns-lookup-latency-windows-amd64 main.go
GOOS=windows GOARCH=arm64 go build -o dns-lookup-latency-windows-arm64 main.go

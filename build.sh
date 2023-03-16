#!/bin/bash

set -e

export GOOS=linux
export GOARCH=386

mkdir -p build
go build -o build/pcpu main.go
scp build/pcpu homer:
ssh homer ./pcpu 2866941

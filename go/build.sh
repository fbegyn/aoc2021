#! /usr/bin/env sh

DAY=$1
DIR=day$(printf "%02d" ${DAY})

mkdir -p ./bin
go build -o ./bin/${DIR} -ldflags "-w -s" ./cmd/${DIR}

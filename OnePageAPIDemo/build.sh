#!/bin/sh
#go generate
GOOS=js GOARCH=wasm go build -o web/apiDemo.wasm
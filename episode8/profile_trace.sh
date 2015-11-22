#!/bin/bash

# see line 72 of https://golang.org/src/net/http/pprof/pprof.go
curl localhost:8080/debug/pprof/trace

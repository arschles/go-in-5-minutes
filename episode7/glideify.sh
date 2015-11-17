#!/bin/bash

# Attention: run clean.sh before running this script. Otherwise, it will fail if you run it against this repo as-is, or you've already successfully run it once.

# Details: This script creates a glide.yaml file to track dependencies, and a vendor/ folder to store them. Neither operation is idempotent, so they must be manually removed.

# Also ensure that the GO15VENDOREXPERIMENT env var is set to 1 before running go build

glide create # create the glide.yaml file

# install first level dependencies
glide get github.com/gorilla/mux # download the latest gorilla mux code into vendor/ and add the dep to glide.yaml
glide get github.com/google/go-github/github # download the latest github client code into vendor/ and add the dep to glide.yaml
glide get github.com/arschles/testsrv # download the latest testsrc code into vendor/ and add the dep to glide.yaml

# install transitive dependencies
glide get github.com/google/go-querystring/query
glide get github.com/gorilla/context # note: running this command was elided from the screencast for brevity

# you can now run go build


# Special note: The following are optional commands for project maintenance
glide install # run this command to install updates for all dependencies listed in glide.yaml
glide up github.com/gorilla/mux # run this command to only install updates for gorilla mux

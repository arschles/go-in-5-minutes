#!/bin/bash

# Attention: remove the vendor/ directory and the glide.yaml file before running this script. It will fail to run if you run it against this repository as-is, or you've already successfully run it once.

# Details: This script creates a glide.yaml file to track dependencies, and a vendor/ folder to store them. Neither operation is idempotent, so they must be manually removed.

glide create # create the glide.yaml file

# install first level dependencies
glide get github.com/gorilla/mux # download the latest gorilla mux code into vendor/ and add the dep to glide.yaml
glide get github.com/google/go-github/github # download the latest github client code into vendor/ and add the dep to glide.yaml
glide get github.com/arschles/testsrv # download the latest testsrc code into vendor/ and add the dep to glide.yaml

# install transitive dependencies
glide get github.com/google/go-querystring/query


# Special note: The following are optional commands for project maintenance
glide install # run this command to install updates for all dependencies listed in glide.yaml
glide up github.com/gorilla/mux # run this command to only install updates for gorilla mux

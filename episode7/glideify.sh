#!/bin/bash
# this script won't work if you've already run it once.
# it creates a glide.yaml file to track dependencies and a
# vendor/ folder to house the code to build each dependency
glide create # create the glide.yaml file

# install first level dependencies
glide get github.com/gorilla/mux # download the latest gorilla mux code into vendor/ and add the dep to glide.yaml
glide get github.com/google/go-github/github # download the latest github client code into vendor/ and add the dep to glide.yaml
glide get github.com/arschles/testsrv # download the latest testsrc code into vendor/ and add the dep to glide.yaml

# install transitive dependencies
glide get github.com/google/go-querystring/query


# optional commands for later
glide install # run this command to install updates for all dependencies listed in glide.yaml
glide up github.com/gorilla/mux # run this command to only install updates for gorilla mux

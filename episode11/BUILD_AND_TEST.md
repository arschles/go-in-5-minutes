# Getting Started

This document describes how to build and test this server.

## Dependencies

Before you get started, ensure you have the following software installed:

1. Go 1.5 or higher
2. GNU Make - many Unix or Linux based systems already have this. Type `make` on your command line to see if you already do
3. CUrl - many Unix or Linux based systems already have this. Type `curl` on your command line to see if you already do
4. [Glide](https://github.com/Masterminds/glide) 0.8.0 or higher

## Install and Test

The Makefile included herein has everything you need to run and test a server. Simply type `make bootstrap run` to install code dependencies, build and run the server. Once it's running, open a new terminal and run `make test-create test-get test-delete` to run `curl` commands that test the server.

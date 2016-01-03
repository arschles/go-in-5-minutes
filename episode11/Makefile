bootstrap:
	glide up

build:
	GO15VENDOREXPERIMENT=1 go build -o episode11

run: build
	./episode11

test-create:
	curl -XPOST -d "{\"name\":\"test\"}" localhost:8080/apps

test-get:
	curl -XGET localhost:8080/app/test

test-delete:
	curl -XDELETE localhost:8080/app/test

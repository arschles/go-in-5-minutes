bootstrap:
	glide up

build:
	GO15VENDOREXPERIMENT=1 go build -o episode11

run: build
	./episode11

test-create:
	curl -v -XPOST -d "{\"name\":\"test\"}" localhost:8080/apps

test-get:
	curl -v -XGET localhost:8080/app/test

test-delete:
	curl -v -XDELETE localhost:8080/app/test

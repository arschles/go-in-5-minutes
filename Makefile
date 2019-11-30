.PHONY: api-build
api-build:
	cd api && go build -o bin/api .

.PHONY: api-test
api-test:
	cd api && go test ./...

.PHONY: build
build:
	cd ./episode0 && go build && cd ..
	cd ./episode1 && go build && cd ..
	cd ./episode2 && go build && cd ..
	cd ./episode3 && go build && cd ..
	cd ./episode4 && go build && cd ..
	cd ./episode5 && go build && cd ..
	cd ./episode6 && go build && cd ..
	cd ./episode7 && ../glide install && go build && cd ..
	cd ./episode8 && go build && cd ..
	cd ./episode9 && go build && cd ..
	cd ./episode10 && go build && cd ..
	cd ./episode11 && ../glide install && go build && cd ..
	cd ./episode12 && go build && cd ..
	cd ./episode13 && ../glide install && go build && cd ..
	cd ./episode14 && go build && cd ..
	cd ./episode15 && go build && cd ..
	cd ./episode16 && go build && cd ..
	cd ./episode17 && go build && cd ..
	cd ./episode18 && go build && cd ..
	cd ./episode19 && go build && cd ..
	cd ./episode20 && go build && cd ..
	cd ./episode21 && go build && cd ..
	cd ./episode22 && go build && cd ..
	cd ./episode23 && go build && cd ..
	cd ./episode24 && go build && cd ..
	cd ./episode25 && go build && cd ..

.PHONY: build-modules
build-modules:
	cd ./episode25 && go build
	cd ./episode26 && go build
	cd ./episode27 && go build
	cd ./episode28 && go build
	cd ./episode29 && go build
	cd ./episode30 && go build
	cd ./episode31 && go build

.PHONY: site-dev
site-static:
	cd site/static && netlify dev

.PHONY: site-api
site-api:
	cd site/api && buffalo dev

site-elm-compile:
	cd site/static && ./elm-build.sh

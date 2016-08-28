GLIDE_TARURL := https://github.com/Masterminds/glide/releases/download/v0.12.0/glide-v0.12.0-linux-amd64.tar.gz
GLIDE_TARBALL := glide-v0.12.0-linux-amd64.tar.gz

get-glide:
	curl -L -o ${GLIDE_TARBALL} ${GLIDE_TARURL}
	tar -xvf ${GLIDE_TARBALL}
	mv linux-amd64/glide ./glide

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

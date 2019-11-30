.PHONY: install-caddy
install-caddy:
	curl https://getcaddy.com | bash -s personal

#####
# The proxy
#####
.PHONY: proxy
proxy:
	caddy

#####
# The static site
#####
.PHONY: hugo
hugo:
	cd static && hugo server

#####
# API server
#####
.PHONY: api
api:
	cd api && buffalo dev

.PHONY: api-build
api-build:
	cd api && go build -o bin/api .

.PHONY: api-test
api-test:
	cd api && go test ./...

#####
# The elm
#####
.PHONY: elm-compile
elm-compile:
	cd static && ./elm-build.sh

build:
	go build -o episode18 .
run:
	./episode18
test-flush:
	curl -v localhost:8080/flush
test-client-fail:
	curl -v localhost:8080/client
test-client-succ:
	curl -v localhost:8080/client?timeout=15

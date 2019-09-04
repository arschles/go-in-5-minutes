build:
	go build -o episode31

run:
	./episode31

CURL_PREAMBLE := curl -v --compressed

test-pluralize:
	@echo "singular 'boat':"
	${CURL_PREAMBLE} localhost:8080/pluralize/boat?num=1
	@echo ""
	@echo "plural 'boat':"
	${CURL_PREAMBLE} localhost:8080/pluralize/boat?num=2
	@echo ""

test-middleware:
	${CURL_PREAMBLE} localhost:8080/request_count
	@echo ""

test-json:
	${CURL_PREAMBLE} -H "Content-Type: application/json" -XPOST -d '{"x":123,"y":456}'  localhost:8080/json
	@echo ""

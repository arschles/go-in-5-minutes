build:
	go build -o episode19

run:
	./episode19

test-pluralize:
	@echo "singular 'boat':"
	curl localhost:8080/pluralize/boat?num=1
	@echo ""
	@echo "plural 'boat':"
	curl localhost:8080/pluralize/boat?num=2
	@echo ""

test-middleware:
	curl localhost:8080/request_count
	@echo ""

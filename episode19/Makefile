build:
	go build -o episode19

run:
	./episode19

test-pluralize:
	@echo "singular 'boat':"
	curl localhost:8080/pluralize?singular=boat\&num=1
	@echo ""
	@echo "plural 'boat':"
	curl localhost:8080/pluralize?singular=boat\&num=2
	@echo ""

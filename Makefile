# add this top the top
help: ## help
	@grep -E '(^##)|(\s##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = "(:.*?)?## "}; {if ($$1) { printf "  %-30s %s\n", $$1, $$2 } else { printf "\n%s\n", $$2 }}'


serve: ## run server
	go run main.go server.go csv.go --mode server

convert: ## convert submission to csv
	go run main.go server.go csv.go --mode csv



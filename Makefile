.PHONY: run

run:
	rm -rf ./main && go build cmd/app/main.go && ./main

build:
	go build cmd/app/main.go
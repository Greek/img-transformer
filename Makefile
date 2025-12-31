.PHONY: run

build:
	go build -o out/main .

run:
	go run .

run-prod:
	make build && out/main
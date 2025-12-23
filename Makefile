.PHONY: run

build:
	go build -o out/main .

run:
	make build && out/main
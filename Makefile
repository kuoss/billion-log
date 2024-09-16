BINARY_NAME=log-loop

.PHONY: build
build:
	go build -o bin/$(BINARY_NAME) .

.PHONY: run
run: build
	bin/$(BINARY_NAME)

.PHONY: docker
docker:
	docker build -t log-loop .
	docker run log-loop

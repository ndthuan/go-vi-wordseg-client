.PHONY: examples test

COMPOSE_FILE = ./build/ci/docker-compose.yml

examples:
	docker-compose -f ${COMPOSE_FILE} run --rm tests sh -c "go run examples/main.go"

lint:
	gofmt -s -w .
	golint ./...

test:
	docker-compose -f ${COMPOSE_FILE} build
	docker-compose -f ${COMPOSE_FILE} run --rm tests

clean:
	docker-compose -f ${COMPOSE_FILE} down -v

it: lint test clean

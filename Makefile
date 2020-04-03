.PHONY: setup build up down logs test db create-test-db

BASE_DOCKER_COMPOSE = ./docker/docker-compose.yml
COMPOSE_OPTS        = -f "$(BASE_DOCKER_COMPOSE)" -p sample
FILE				= create_sample

setup: build wait-db create-test-db migration migration-test
clean: down delete-volume

build: go-build
	docker-compose $(COMPOSE_OPTS) up -d --build

go-build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./.artifacts/sample-linux-amd64 ./cmd/sample

up:
	docker-compose $(COMPOSE_OPTS) up -d

down:
	docker-compose $(COMPOSE_OPTS) down

logs:
	docker-compose $(COMPOSE_OPTS) logs -f

new-migration:
	docker-compose $(COMPOSE_OPTS) exec -T sample-api sh -c "./third_party/bin/sql-migrate new ${FILE} -config=dbconfig.yml"

migration:
	docker-compose $(COMPOSE_OPTS) exec -T sample-api sh -c "./third_party/bin/sql-migrate up -config=dbconfig.yml"

migration-test:
	docker-compose $(COMPOSE_OPTS) exec -T sample-api sh -c "./third_party/bin/sql-migrate up -config=dbconfig_test.yml"

db:
	docker-compose $(COMPOSE_OPTS) exec sample-api-db mysql -uroot -ppassword -Dsample

create-test-db:
	docker-compose $(COMPOSE_OPTS) exec -T sample-api-db mysql -uroot -ppassword -e "CREATE DATABASE IF NOT EXISTS sample_test"

delete-volume:
	rm -rf ./db/mysql_data

wait-db:
	sh ./scripts/wait.sh

test:
	go fmt ./...
	golangci-lint run
	go vet ./...
	go test -v ./...

image:
	docker build . --target dev -t starwars-data-proxy:api-dev

build:
	docker compose run --rm api go build -o api

deps:
	docker compose run --rm api go mod tidy
	docker compose run --rm api go mod vendor

setup: image deps build

local: local-down build
	docker compose up api

local-down:
	docker compose rm -sf

test:
	docker compose run --rm api go test -cover ./... -tags nojira

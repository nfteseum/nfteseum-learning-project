.PHONY: postgres createdb dropdb migrate_up migrate_down migrate_create sqlc_generate run proto

postgres:
	docker run --name skyweaverpostgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

createdb:
	docker exec -it skyweaverpostgres createdb --username=postgres --owner=postgres nfteseum

dropdb:
	docker exec -it skyweaverpostgres dropdb -U postgres nfteseum

migrate_up:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate -source file://data/migrations -database postgres://postgres:postgres@localhost:5432/nfteseum?sslmode=disable up

migrate_down:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate  -source file://data/migrations -database postgres://postgres:postgres@localhost:5432/nfteseum?sslmode=disable down 1

migrate_create:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate create -ext sql -dir data/migrations -seq $(FILE_NAME)



ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

sqlc_generate:
	docker run --rm -v $(ROOT_DIR):/src -w /src kjconroy/sqlc generate

run:
	@go run github.com/VojtechVitek/rerun/cmd/rerun -watch ./ -run \
		sh -c 'GOGC=off go build -o ./bin/api-server ./cmd/api-server/main.go && ./bin/api-server --config=./etc/example.conf'

proto:
	go generate ./proto/...

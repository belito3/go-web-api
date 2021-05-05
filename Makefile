postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

build:
	export CGO_ENABLED=0 && go build -o main .

test:
	go test -v -cover ./app/repository/impl

mock_test:
	go test -timeout 30s -coverprofile=/tmp/vscode-goUJWYLL/go-code-cover github.com/belito3/go-web-api/app/api -v -count=1

docker_build:
	make build && docker build -t api-codebase-go:latest .

docker_run:
	docker-compose up

server:
	go run main.go

mock:
	mockgen -package mockdb -destination app/repository/mock/store.go github.com/belito3/go-web-api/app/repository/impl IStore

.PHONY: postgres createdb dropdb migrateup migratedown build test docker_build docker_run server mock

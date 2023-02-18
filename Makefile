build:
	go build -o auction ./cmd/app/main.go

build-goose:
	go build -o goose ./cmd/goose/*.go

migrate:
	GOOSE_DRIVER="mysql" GOOSE_DBSTRING="root:secret@tcp(localhost:3306)/auction" ./goose -dir internal/db/migrations up

migrate-rollback:
	GOOSE_DRIVER="mysql" GOOSE_DBSTRING="root:secret@tcp(localhost:3306)/auction" ./goose -dir internal/db/migrations reset

run-on-server:
	go build -o auction app/main.go \
	&& ./auction

docker-build:
	docker-compose build

docker-migrate:
	docker compose exec auction /bin/sh -c "make migrate"

docker-migrate-rollback:
	docker compose exec auction /bin/sh -c "make migrate-rollback"

docker-run:
	docker-compose -f docker-compose.yml up -d

docker-stop:
	docker-compose -f docker-compose.yml down

docker-clean:
	docker system prune
build:
	go build -o auction ./cmd/app/main.go

build-goose:
	go build -o goose ./cmd/goose/*.go

local-migrate:
	GOOSE_DRIVER="mysql" GOOSE_DBSTRING="root:secret_for_root@tcp(localhost:33061)/auction" ./goose -dir internal/db/migrations up

local-migrate-rollback:
	GOOSE_DRIVER="mysql" GOOSE_DBSTRING="root:secret_for_root@tcp(localhost:33061)/auction" ./goose -dir internal/db/migrations reset

migrate:
	GOOSE_DRIVER="mysql" GOOSE_DBSTRING="doadmin:AVNS_7-mXt03HXbHH3mpoJF6@tcp(db-mysql-fra1-42653-do-user-13642568-0.b.db.ondigitalocean.com:25060)/auction?" ./goose -dir internal/db/migrations up

migrate-rollback:
	GOOSE_DRIVER="mysql" GOOSE_DBSTRING="doadmin:AVNS_7-mXt03HXbHH3mpoJF6@tcp(db-mysql-fra1-42653-do-user-13642568-0.b.db.ondigitalocean.com:25060)/auction" ./goose -dir internal/db/migrations reset

run-on-server:
	LOG_DRIVER=file LOG_LEVEL=info DB_DRIVER=mysql DB_STRING="doadmin:AVNS_7-mXt03HXbHH3mpoJF6@tcp(db-mysql-fra1-42653-do-user-13642568-0.b.db.ondigitalocean.com:25060)/auction?parseTime=true" REDIS_ADDRESS="db-redis-fra1-04092-do-user-13642568-0.b.db.ondigitalocean.com" REDIS_PASSWORD="AVNS_LfGrNA2fS8ib2ujBKJU" REDIS_PORT=25061 REDIS_USERNAME="default" ./auction

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


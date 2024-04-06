.PHONY: start-and-migrate
start-and-migrate: stop-and-delete networkcreate postgres-image postgres wait-for-postgres createdb migrateup

.PHONY: stop-and-delete
stop-and-delete:
	-docker stop postgres16
	-docker rm postgres16
	-docker network rm ev-network

.PHONY: postgres-image
postgres-image:
	docker pull postgres:16-alpine

.PHONY: postgres
postgres:
	docker run --name postgres16 --network ev-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

.PHONY: wait-for-postgres
wait-for-postgres:postgres
	sleep 10  # wait time to run make createdb

.PHONY: networkcreate
networkcreate:
	-docker network create ev-network

.PHONY: delete-network
delete-network:
	-docker network rm ev-network

.PHONY: createdb
createdb: wait-for-postgres
	docker exec -it postgres16 createdb --username=root --owner=root ev-db

.PHONY: dropdb
dropdb:
	docker exec -it postgres16 dropdb ev-db

.PHONY: migrateup
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@127.0.0.1:5432/ev-db?sslmode=disable" -verbose up

.PHONY: server
server:
	go run app/cmd/main.go
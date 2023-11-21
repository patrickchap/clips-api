
postgres:
	docker run --name postres16-clips -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine
createdb:
	docker exec -it postres16-clips createdb --username=root --owner=root clips 
dropdb:
	docker exec -it postres16-clips dropdb clips
migrateup:
	migrate -path db/migrations/ -database "postgresql://root:secret@localhost:5432/clips?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations/ -database "postgresql://root:secret@localhost:5432/clips?sslmode=disable" -verbose down
tests:
	go test -v -cover ./...
sqlc:
	sqlc generate
server:
	go run main.go
mockdb:
	mockgen -package mockdb  -destination db/mock/store.go github.com/patrickchap/clipsapi/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown server sqlc tests mockdb

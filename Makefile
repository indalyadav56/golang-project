# Load environment variables
include .env
export

init:
	echo "Docker shouldb be installed in your system $(MONGODB_URL)"

postgresinit:
	docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=indal_db -p 5432:5432 -v pgdata:/var/lib/postgresql/data --restart=always -d postgres

postgres:
	docker exec --it postgres psql

createdb:
	docker exec --it postgres createdb --username=postgres --owner=postgres db_name

dropdb:
	docker exec --it postgres dropdb db_name

run:
	nodemon --exec go run main.go --signal SIGTERM

redistdb:
	docker run -d --name my-redis -p 6379:6379 -e REDIS_PASSWORD=mysecretpassword redis:latest --requirepass mysecretpassword

.PHONY: clean test deploy  # set target to cleaning the project
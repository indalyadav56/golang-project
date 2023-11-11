# Load environment variables
include .env
export


run:
	nodemon --exec go run main.go --signal SIGTERM

init:
	echo "Docker shouldb be installed in your system $(MONGODB_URL)"

test:
	go test -cover

kubectl-postgres-start:
	kubectl apply -f ./k8s/db/postgres.yaml

postgres-start:
	docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=indal_db -p 5432:5432 -v pgdata:/var/lib/postgresql/data -d postgres

postgres-remove:
	docker rm -f postgres

postgres-shell:
	docker exec -it postgres psql -U postgres

redis-start:
	docker run -d --name redis-container -p 6379:6379 redis:latest --requirepass redis@123

redis-remove:
	docker rm -f redis-container

redis-cli:
	docker exec -it redis-container redis-cli

.PHONY: clean test deploy  # set target to cleaning the project
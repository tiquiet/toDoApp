.SILENT:

run:
	go run cmd/main.go

runBd:
	docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

	docker exec -it d84c44930d5d /bin/bash

	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
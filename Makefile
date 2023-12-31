setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g cmd/main.go

build:
	docker compose up --build

run:
	docker compose up

migrate:
	migrate -path ./schema -database 'postgres://test_user:password@localhost:5435/postgres?sslmode=disable' up

down:
	docker compose down

restart:
	docker compose restart

rebuild:
	docker compose up --build

clean:
	docker stop back
	docker stop db
	docker rm back
	docker rm db
	docker image rm avito-dynamic-segment-back-back
	rm -rf postgres_data
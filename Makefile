db-up:
	docker compose up -d db adminer

dev:
	go run ./cmd/server/

build:
	docker build -t catinator-backend:latest .

docker-run:
	docker compose up backend

setup:
	mkdir .local
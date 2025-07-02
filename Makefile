default: dev

start:
	go run ./cmd/main.go
dev:
	air
up:
	docker-compose up --build

down:
	docker-compose down

down-v:
	docker-compose down -v
logs:
	docker-compose logs -f
ps:
	docker-compose ps
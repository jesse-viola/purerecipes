build-dev:
	docker compose -f docker-compose.dev.yml build

dev:
	docker compose -f docker-compose.dev.yml up -d
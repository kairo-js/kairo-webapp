up:
	docker compose -f infra/docker-compose.yaml up --build

down:
	docker compose -f infra/docker-compose.yaml down

logs:
	docker compose -f infra/docker-compose.yaml logs -f
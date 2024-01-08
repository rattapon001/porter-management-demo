.PHONY: build up down

up: build
	docker compose up --build -d

down:
	docker compose down
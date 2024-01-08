.PHONY: build up down

up: build
	docker-compose up -d

down:
	docker-compose down
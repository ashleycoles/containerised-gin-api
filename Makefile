DOCKER_COMPOSE = docker compose

.DEFAULT_GOAL := up

.PHONY: up
up:
	$(DOCKER_COMPOSE) up --build

.PHONY: start
start:
	$(DOCKER_COMPOSE) up

.PHONY: stop
stop:
	$(DOCKER_COMPOSE) down

.PHONY: rebuild
rebuild:
	$(DOCKER_COMPOSE) up --build --force-recreate
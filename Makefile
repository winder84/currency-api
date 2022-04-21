THIS_FILE := $(lastword $(MAKEFILE_LIST))
.PHONY: start stop restart logs ps currency_db migrate
start:
	docker-compose -f docker/docker-compose.yaml up -d currency_db
	sleep 3
	make migrate
stop:
	docker-compose -f docker/docker-compose.yaml down --volumes
restart:
	make stop
	make start
logs:
	docker-compose -f docker/docker-compose.yaml logs --tail=100 -f currency_db
ps:
	docker-compose -f docker/docker-compose.yaml ps -a
currency_db:
	docker-compose -f docker/docker-compose.yaml exec currency_db /bin/bash
migrate:
	docker-compose -f docker/docker-compose.yaml up -d migrate
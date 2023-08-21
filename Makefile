.PHONY: serve restart kill before start 

PID = /tmp/serving.pid
DOCKER=@docker-compose -f ./docker-compose.yml

install: .out_docker build setup

setup: up database
	@echo "\n \e[1;42m Application disponible: \e[0m\n http://localhost:8080\n"

build: .out_docker
	$(DOCKER) build  --pull --no-cache

up: .out_docker
	$(DOCKER) up -d --remove-orphans --force-recreate

database: .out_docker
	sleep 2
	$(DOCKER) exec db mysql -e "DROP DATABASE IF EXISTS welsh"
	$(DOCKER) exec db mysql -e "CREATE DATABASE welsh"
	$(DOCKER) exec golang go run cmd/migration/migration.go

container: .out_docker
	$(DOCKER) exec golang bash -c "cd src && wire"

stop: .out_docker
	$(DOCKER) stop
	$(DOCKER) down

logs: .out_docker
	$(DOCKER) logs --tail 20 -f

shell: .out_docker
	$(DOCKER) exec golang bash

serve: start
	fswatch -or --event=Updated -d src/ -d cmd/api/ | xargs -n1 -I {} make restart

kill:
	-kill `pstree -p \`cat $(PID)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"`

before:
	@echo "STOPPED" && printf '%*s\n' "40" '' | tr ' ' -

start:
	go run cmd/api/api.go & echo $$! > $(PID)

restart: kill before start
	@echo "STARTED" && printf '%*s\n' "40" '' | tr ' ' -

.out_docker:
ifeq (, $(shell which docker))
	$(error "You must run this command outside the docker container")
endif


DOCKER=@docker compose -f ./docker-compose.yml
PID = /tmp/serving.pid

install: build setup

setup: up
	sleep 2
	cp .env.dist .env
	make database
	@echo "\n \e[1;42m Application disponible: \e[0m\n http://localhost:8080\n"

build:
	$(DOCKER) build  --pull --no-cache

up: 
	$(DOCKER) up -d --remove-orphans --force-recreate

database: 
	$(DOCKER) exec db mysql -e "DROP DATABASE IF EXISTS welsh"
	$(DOCKER) exec db mysql -e "CREATE DATABASE welsh"
	$(DOCKER) exec golang go run cmd/migration/migration.go

container:
	$(DOCKER) exec golang bash -c "cd src && wire"

logs: 
	$(DOCKER) logs --tail 20 -f

shell: 
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

tests: database
	$(DOCKER) exec golang go test -v ./...



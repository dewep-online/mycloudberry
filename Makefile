SHELL=/bin/bash
COMPOSE_FLAGS=-f deployments/docker-compose.yaml -p dev_mycloudberry


lint:
	devtool lint

run_local_build:
	go build -race -o /tmp/mycloudberry.bin ./cmd/mycloudberry/main.go
	sudo /tmp/mycloudberry.bin --config=./configs/config.dev.yaml

run_local:
	go run -race ./cmd/mycloudberry/main.go --config=./configs/config.dev.yaml

dev_up: dev_down
	docker-compose $(COMPOSE_FLAGS) pull
	docker-compose $(COMPOSE_FLAGS) up -d

dev_down:
	docker-compose $(COMPOSE_FLAGS) down -v --remove-orphans

#dev_log:
#	docker-compose $(COMPOSE_FLAGS) logs --follow backend

start: dev_up dump_restore

vendors:
	go mod tidy && go generate ./... && go mod vendor
	#cd web && yarn install --force --ignore-scripts

dump_restore:
	docker-compose $(COMPOSE_FLAGS) exec -T mysql \
		mysql -utest -ptest --default-character-set=utf8mb4 test < ./sources/db/setup.sql

dump_create:
	docker-compose $(COMPOSE_FLAGS) exec -T mysql \
		mysqldump -utest -ptest --default-character-set=utf8mb4 test > ./sources/db/setup.sql
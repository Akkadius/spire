#----------------------
# Parse makefile arguments
#----------------------
RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

#----------------------
# Silence GNU Make
#----------------------
ifndef VERBOSE
MAKEFLAGS += --no-print-directory
endif

#----------------------
# Load .env file
#----------------------
ifneq ("$(wildcard .env)","")
include .env
export
else
endif

DRUNPREFIX=
ifeq ($(OS),Windows_NT)
    DRUNPREFIX = winpty
endif

#----------------------
# Terminal
#----------------------

GREEN  := $(shell tput -Txterm setaf 2)
WHITE  := $(shell tput -Txterm setaf 7)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)

#------------------------------------------------------------------
# - Add the following 'help' target to your Makefile
# - Add help text after each target name starting with '\#\#'
# - A category can be added with @category
#------------------------------------------------------------------

.PHONY: build test

HELP_FUN = \
	%help; \
	while(<>) { \
		push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
		print "\n"; \
		for (sort keys %help) { \
			print "${WHITE}$$_${RESET \
		}\n"; \
		for (@{$$help{$$_}}) { \
			$$sep = " " x (32 - length $$_->[0]); \
			print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
		}; \
		print ""; \
	}

help: ##@other Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

#----------------------
# dev
#----------------------

bash: ##@dev Bash into workspace container
	docker-compose up -d workspace
	$(DRUNPREFIX) docker-compose exec workspace bash

mc: ##@dev Create MySQL shell
	$(DRUNPREFIX) docker-compose exec mysql sh -c "mysql -uroot -p${MYSQL_ROOT_PASSWORD} -h localhost"

rc: ##@dev Create Redis shell
	$(DRUNPREFIX) docker-compose exec redis sh -c "redis-cli"

up: ##@dev Brings up environment
	docker-compose up -d mysql workspace

down: ##@dev Shuts down environment
	docker-compose down -t 1

logs: ##@dev Follow container logs
	docker-compose logs --follow

test: ##@dev Runs local tests
	$(DRUNPREFIX) docker-compose exec workspace bash -c "go test -count=1 -cover ./... | grep -v 'no test files'"

#----------------------
# build
#----------------------

build-prod: ##@build Runs frontend watcher
	packr clean
	cd ./frontend && npm run build
	cp ./frontend/dist/* ./public/ -R
	packr build

#----------------------
# dev-watch
#----------------------

watch-fe: ##@dev-watch Runs frontend watcher
	cd frontend && npm run dev

watch-be: ##@dev-watch Runs backend watcher
	docker-compose exec workspace bash -c "air -c .air.toml"

#----------------------
# seed
#----------------------

seed-peq-database: ##@seed
	docker-compose up -d workspace
	$(DRUNPREFIX) docker-compose exec workspace bash -c "curl http://db.projecteq.net/api/v1/dump/latest -o /tmp/db.zip"
	$(DRUNPREFIX) docker-compose exec workspace bash -c "unzip -o /tmp/db.zip -d /tmp/db/"
	$(DRUNPREFIX) docker-compose exec workspace bash -c "mysql -h mysql -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} ${MYSQL_EQEMU_DATABASE} -e 'DROP DATABASE ${MYSQL_EQEMU_DATABASE}; CREATE DATABASE ${MYSQL_EQEMU_DATABASE};'"
	$(DRUNPREFIX) docker-compose exec workspace bash -c "cd /tmp/db/peq-dump/ && mysql -h mysql -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} ${MYSQL_EQEMU_DATABASE} < ./create_all_tables.sql"
	$(DRUNPREFIX) docker-compose exec workspace bash -c "rm -rf /tmp/db/"

# This needs to get cleaned up later
seed-peq-database-prod: ##@seed
	docker-compose exec prod bash -c "curl http://db.projecteq.net/api/v1/dump/latest -o /tmp/db.zip"
	docker-compose exec prod bash -c "unzip -o /tmp/db.zip -d /tmp/db/"
	docker-compose exec prod bash -c "mysql -h mysql -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} ${MYSQL_EQEMU_DATABASE} -e 'DROP DATABASE ${MYSQL_EQEMU_DATABASE}; CREATE DATABASE ${MYSQL_EQEMU_DATABASE};'"
	docker-compose exec prod bash -c "cd /tmp/db/peq-dump/ && mysql -h mysql -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} ${MYSQL_EQEMU_DATABASE} < ./create_all_tables.sql"
	docker-compose exec prod bash -c "rm -rf /tmp/db/"

seed-spire-tables: ##@seed
	$(DRUNPREFIX) docker-compose exec workspace bash -c "go run main.go spire:migrate"

#----------------------
# generate
#----------------------

generate-axios-client: ##@generate Generate Axios client (Run outside workspace container)
	$(DRUNPREFIX) docker run --rm -v "$${PWD}:/local" openapitools/openapi-generator-cli:v5.0.0 generate \
        -i /local/docs/swagger.yaml \
        -g typescript-axios \
        -o /local/frontend/src/app/api/ \
        -c /local/openapi-generator-config.yaml

generate-axios-client-local: ##@generate Generate Axios client (Run outside workspace container)
	# sudo npm install @openapitools/openapi-generator-cli -g
	openapi-generator-cli version-manager set 5.0.0
	openapi-generator-cli generate --enable-post-process-file \
        -i ./docs/swagger.yaml \
        -g typescript-axios \
        -o ./frontend/src/app/api/ \
        -c ./openapi-generator-config.yaml

generate-swagger: ##@generate Generate swagger docs (Run in workspace container)
	./scripts/generate-swagger.sh

#----------------------
# install
#----------------------

install: ##@install Runs installer
	$(DRUNPREFIX) docker-compose build
	make mysql-init
	make seed-peq-database
	make seed-spire-tables
	make install-assets
	@./scripts/banner.sh "Environment Initialized!"
	@cat ./scripts/install-message.txt

install-assets: ##@install Installs assets
	$(DRUNPREFIX) docker-compose exec workspace bash -c 'curl --compressed -o /tmp/assets.zip -L https://github.com/Akkadius/eq-asset-preview/archive/refs/heads/master.zip'
	$(DRUNPREFIX) docker-compose exec workspace bash -c 'unzip -o /tmp/assets.zip -d /tmp/assets'
	$(DRUNPREFIX) docker-compose exec workspace bash -c 'cp -R /tmp/assets/eq-asset-preview-master/ ./frontend/public/'

#----------------------
# mysql
#----------------------

mysql-init: ##@mysql Initialize database
	@./scripts/banner.sh "Initializing MySQL Database..."
	docker-compose kill mysql
	docker-compose build mysql
	docker-compose up -d mysql
	$(DRUNPREFIX) docker-compose run --user=root mysql bash -c "chown -R mysql:mysql /var/lib/mysql && exit"
	docker-compose up -d mysql
	$(DRUNPREFIX) docker-compose exec -T mysql sh -c 'while ! mysqladmin ping -h "mysql" --silent; do sleep .5; done'
	$(DRUNPREFIX) docker-compose exec mysql sh -c "mysql -h localhost -uroot -p${MYSQL_ROOT_PASSWORD} -e 'CREATE DATABASE IF NOT EXISTS ${MYSQL_EQEMU_DATABASE}'"
	$(DRUNPREFIX) docker-compose exec mysql sh -c "mysql -h localhost -uroot -p${MYSQL_ROOT_PASSWORD} -e 'GRANT ALL PRIVILEGES ON ${MYSQL_EQEMU_DATABASE}.* TO \"${MYSQL_USERNAME}\"@\"%\"'"
	$(DRUNPREFIX) docker-compose exec mysql sh -c "mysql -h localhost -uroot -p${MYSQL_ROOT_PASSWORD} -e 'CREATE DATABASE IF NOT EXISTS ${MYSQL_SPIRE_DATABASE}';"
	$(DRUNPREFIX) docker-compose exec mysql sh -c "mysql -h localhost -uroot -p${MYSQL_ROOT_PASSWORD} -e 'GRANT ALL PRIVILEGES ON ${MYSQL_SPIRE_DATABASE}.* TO \"${MYSQL_USERNAME}\"@\"%\"'"

init-strip-mysql-remote-root: ##@mysql Strips MySQL remote root user
	$(DRUNPREFIX) docker-compose exec mysql bash -c "mysql -uroot -p${MYSQL_ROOT_PASSWORD} -h localhost -e \"delete from mysql.user where User = 'root' and Host = '%'; FLUSH PRIVILEGES\""

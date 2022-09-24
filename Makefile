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

COMPOSE_COMMAND=docker-compose
ifeq ($(APP_ENV),production)
	COMPOSE_COMMAND=docker-compose -f docker-compose.yml -f docker-compose.prod.yml
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
	$(COMPOSE_COMMAND) up -d workspace
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash

mc: ##@dev Create MySQL shell
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec mysql sh -c "mysql -uroot -p${MYSQL_ROOT_PASSWORD} -h localhost"

rc: ##@dev Create Redis shell
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec redis sh -c "redis-cli"

up: ##@dev Brings up environment
	$(COMPOSE_COMMAND) up -d mysql workspace
ifeq ($(APP_ENV),production)
	$(COMPOSE_COMMAND) up -d mysql workspace
endif

down: ##@dev Shuts down environment
	$(COMPOSE_COMMAND) down -t 1

logs: ##@dev Follow container logs
	$(COMPOSE_COMMAND) logs --follow

test: ##@dev Runs local tests
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash -c "go test -count=1 -cover ./... | grep -v 'no test files'"

#----------------------
# dev-watch
#----------------------

watch-fe: ##@dev-watch Runs frontend watcher
	$(COMPOSE_COMMAND) exec workspace bash -c "cd frontend && npm run dev"

watch-be: ##@dev-watch Runs backend watcher
	$(COMPOSE_COMMAND) exec workspace bash -c "air -c .air.toml"

#----------------------
# seed
#----------------------

seed-peq-database: ##@seed
	$(COMPOSE_COMMAND) up -d workspace
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash -c "curl http://db.projecteq.net/api/v1/dump/latest -o /tmp/db.zip"
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash -c "unzip -o /tmp/db.zip -d /tmp/db/"
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash -c "mysql -h mysql -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} ${MYSQL_EQEMU_DATABASE} -e 'DROP DATABASE ${MYSQL_EQEMU_DATABASE}; CREATE DATABASE ${MYSQL_EQEMU_DATABASE};'"
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash -c "cd /tmp/db/peq-dump/ && mysql -h mysql -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} ${MYSQL_EQEMU_DATABASE} < ./create_all_tables.sql"
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash -c "rm -rf /tmp/db/"

# This needs to get cleaned up later
seed-peq-database-prod: ##@seed
	$(COMPOSE_COMMAND) exec prod bash -c "curl http://db.projecteq.net/api/v1/dump/latest -o /tmp/db.zip"
	$(COMPOSE_COMMAND) exec prod bash -c "unzip -o /tmp/db.zip -d /tmp/db/"
	$(COMPOSE_COMMAND) exec prod bash -c "mysql -h mysql -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} ${MYSQL_EQEMU_DATABASE} -e 'DROP DATABASE ${MYSQL_EQEMU_DATABASE}; CREATE DATABASE ${MYSQL_EQEMU_DATABASE};'"
	$(COMPOSE_COMMAND) exec prod bash -c "cd /tmp/db/peq-dump/ && mysql -h mysql -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} ${MYSQL_EQEMU_DATABASE} < ./create_all_tables.sql"
	$(COMPOSE_COMMAND) exec prod bash -c "rm -rf /tmp/db/"

seed-spire-tables: ##@seed
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash -c "go run main.go spire:migrate"

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

build: ##@install Build
	$(DRUNPREFIX) $(COMPOSE_COMMAND) build

publish: ##@images
	docker push akkadius/spire:go-workspace

install: ##@install Runs installer
	$(DRUNPREFIX) $(COMPOSE_COMMAND) build
	make mysql-init
	make seed-peq-database
	make seed-spire-tables
	make install-assets
	@./scripts/banner.sh "Environment Initialized!"
	@cat ./scripts/install-message.txt

install-assets: ##@install Installs assets
	@./scripts/banner.sh "Initializing eq-asset-preview assets..."
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash -c 'curl --compressed -o /tmp/assets.zip -L https://github.com/Akkadius/eq-asset-preview/archive/refs/heads/master.zip'
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash -c 'unzip -o /tmp/assets.zip -d /tmp/assets'
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash -c 'cp -R /tmp/assets/eq-asset-preview-master/ ./frontend/public/'

install-frontend: ##@install Install and initialize frontend packages
	cp frontend/.env.example frontend/.env
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec workspace bash -c 'cd frontend && npm install'

#----------------------
# mysql
#----------------------

mysql-init: ##@mysql Initialize database
	@./scripts/banner.sh "Initializing MySQL Database..."
	$(COMPOSE_COMMAND) kill mysql
	$(COMPOSE_COMMAND) build mysql
	$(COMPOSE_COMMAND) up -d mysql
	$(DRUNPREFIX) $(COMPOSE_COMMAND) run --user=root mysql bash -c "chown -R mysql:mysql /var/lib/mysql && exit"
	$(COMPOSE_COMMAND) up -d mysql
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec -T mysql sh -c 'while ! mysqladmin ping -h "mysql" --silent; do sleep .5; done'
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec mysql sh -c "mysql -h localhost -uroot -p${MYSQL_ROOT_PASSWORD} -e 'CREATE DATABASE IF NOT EXISTS ${MYSQL_EQEMU_DATABASE}'"
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec mysql sh -c "mysql -h localhost -uroot -p${MYSQL_ROOT_PASSWORD} -e 'GRANT ALL PRIVILEGES ON ${MYSQL_EQEMU_DATABASE}.* TO \"${MYSQL_USERNAME}\"@\"%\"'"
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec mysql sh -c "mysql -h localhost -uroot -p${MYSQL_ROOT_PASSWORD} -e 'CREATE DATABASE IF NOT EXISTS ${MYSQL_SPIRE_DATABASE}';"
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec mysql sh -c "mysql -h localhost -uroot -p${MYSQL_ROOT_PASSWORD} -e 'GRANT ALL PRIVILEGES ON ${MYSQL_SPIRE_DATABASE}.* TO \"${MYSQL_USERNAME}\"@\"%\"'"

init-strip-mysql-remote-root: ##@mysql Strips MySQL remote root user
	$(DRUNPREFIX) $(COMPOSE_COMMAND) exec mysql bash -c "mysql -uroot -p${MYSQL_ROOT_PASSWORD} -h localhost -e \"delete from mysql.user where User = 'root' and Host = '%'; FLUSH PRIVILEGES\""

#----------------------
# build
#----------------------

build-assets: ##@build Builds static assets before packing into binary
	curl --compressed -o /tmp/assets.zip -L https://github.com/Akkadius/eq-asset-preview/archive/refs/heads/master.zip
	unzip -qq -o /tmp/assets.zip -d /tmp/assets
	cp -R /tmp/assets/eq-asset-preview-master/ ./frontend/public/

strip-extra-assets: ##@build Strips extra assets not needed to packet into binary
	rm -rf frontend/public/eq-asset-preview-master/assets/npc_models
	rm -rf frontend/public/eq-asset-preview-master/assets/objects
	rm -rf frontend/public/eq-asset-preview-master/assets/item_icons

build-frontend: ##@build Builds frontend to be packed into binary
	cd frontend && npm install && npm run build

build-binary: ##@build Build and packs release binary
	packr clean
	packr --compress
	GOOS=linux GOARCH=amd64 go build -o spire-linux-amd64
	GOOS=windows GOARCH=amd64 go build -o spire-windows-amd64.exe
	zip spire-linux-amd64.zip spire-linux-amd64
	zip spire-windows-amd64.exe.zip spire-windows-amd64.exe

release-binary: ##@build Releases binary
	gh-release --assets=spire-linux-amd64.zip,spire-windows-amd64.exe.zip -y

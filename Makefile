# Executables (local)
DOCKER_COMP = docker-compose
DOCKER = docker

# Executables
COMPOSER      = $(PHP_CONT) composer

# Misc
.DEFAULT_GOAL = help
.PHONY        = help build up start down logs sh composer vendor sf cc
PWD   = $(shell pwd)

##
##-—— 💡 MAKEFILE ——————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'


##
##-—— 💡 Container ——————————————————————————————————
con: ## connect with app container
	docker exec -it app /bin/sh


reset: ## reset the containers
	docker rm -f app
	docker rmi -f go-app
	docker-compose up -d



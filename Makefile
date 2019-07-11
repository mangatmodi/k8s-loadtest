# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

HOST=$(shell hostname)

.PHONY: check-app

ifndef tag
override tag = latest
endif

check-app:
ifeq (,$(filter $(app),master slave))
	$(error USAGE: make [command] app=[master|slave])
else
	@echo "app:${app}"	
endif 

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: check-app ## build the docker image with latest and the provided tag
	@docker build -t mangatmodi/k8s-loadtest-$(app):latest -t mangatmodi/k8s-loadtest-$(app):$(tag) $(app)/

push: check-app ## push the docker image with latest tag
	@docker push mangatmodi/k8s-loadtest-$(app):latest

pushTag: check-app ## push the docker image with given tag
	@docker push mangatmodi/k8s-loadtest--$(app):$(tag)

composeUp: check-app  ## compose up the given app
	$(shell HOST=$(HOST) docker-compose -f $(app)/docker-compose.yml up -d)

composeDown: check-app ## compose up the given app
	$(shell HOST=$(HOST) docker-compose -f $(app)/docker-compose.yml down)

apply: check-app  ## apply config for given app for latest or the tagged image
	@echo "$(shell DOCKER_IMAGE=mangatmodi/k8s-loadtest-$(app):$(tag) ./deploy-k8s.sh $(app) apply)"

delete: check-app ## delete config for the given app
	@echo "$(shell DOCKER_IMAGE=mangatmodi/k8s-loadtest-$(app):$(tag) ./deploy-k8s.sh $(app) delete)"

buildApply: check-app build push apply ## build docker image and deploy the latest version of the image

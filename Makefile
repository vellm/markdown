.DEFAULT_GOAL := help
PORT=8080
ENV=development
API_ENDPOINT ?= api.vellm.io
API_STAGE ?= v1

export PORT
export ENV
export API_ENDPOINT
export API_STAGE

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

build:
	docker-compose build

serve: build ## serve local development
	docker-compose up

deploy: ## deploy new version to heroku
	git push heroku master

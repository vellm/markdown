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

setup: ## setup local development
	go get -u github.com/kataras/rizla

serve: ## serve local development
	rizla run main.go

deploy: ## deploy new version to heroku
	git push heroku master

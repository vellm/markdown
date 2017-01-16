.DEFAULT_GOAL := help
PORT=8080

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

serve: ## serve local development
	PORT=$(PORT) docker-compose up

deploy: ## deploy new version to heroku
	git push heroku master

.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

build: ## Build docker container
	docker-compose build

serve: build ## serve backend instance
	docker-compose up

build-aws:
	docker build -t markdown-backend .

tag-aws:
	docker tag markdown-backend:latest 906967097666.dkr.ecr.eu-central-1.amazonaws.com/markdown-backend:latest

push-aws:
	docker push 906967097666.dkr.ecr.eu-central-1.amazonaws.com/markdown-backend:latest

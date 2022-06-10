# Variables
LATEST_COMMIT := $$(git rev-parse HEAD)
VERSION ?= latest
HOST_FOR_DOCKER_IMAGE ?= dcloud.hicrystal.com
ENVIRONMENT ?=

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
%:
	@:

build_dist: ## Build dist folder that needed for frontend.
	cd web/console && npm i --force && npm run build && cd .. && cd nftdrop/landing && npm i --force && npm run build

build_nft_signer: ## Build NFT Signer docker image.
	docker build -f ./deploy/nftsigner.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_signer:$(LATEST_COMMIT) . && docker build -f ./deploy/nftsigner.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_signer$(ENVIRONMENT):$(VERSION) .

push_nft_signer: ## Push NFT Signer docker image.
	docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_signer:$(LATEST_COMMIT) && docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_signer$(ENVIRONMENT):$(VERSION)

build_currency_signer: ## Build currency signer docker image.
	docker build -f ./deploy/currencysigner.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_currency_signer:$(LATEST_COMMIT) . && docker build -f ./deploy/currencysigner.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_currency_signer$(ENVIRONMENT):$(VERSION) .

push_currency_signer: ## Push currency signer docker image.
	docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_currency_signer:$(LATEST_COMMIT) && docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_currency_signer$(ENVIRONMENT):$(VERSION)

build_card_generator: ## Build Card Generator docker image.
	docker build -f ./deploy/cardgenerator.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_card_generator:$(LATEST_COMMIT) . && docker build -f ./deploy/cardgenerator.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_card_generator$(ENVIRONMENT):$(VERSION) .

push_card_generator: ## Push Card Generator docker image.
	docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_card_generator:$(LATEST_COMMIT) && docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_card_generator$(ENVIRONMENT):$(VERSION)

build_app: ## Build Application docker image.
	make build_dist && docker build -f ./deploy/ultimatedivision.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division:$(LATEST_COMMIT) . && docker build -f ./deploy/ultimatedivision.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division$(ENVIRONMENT):$(VERSION) .

push_app: ## Push Application docker image.
	docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division:$(LATEST_COMMIT) && docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division$(ENVIRONMENT):$(VERSION)

build: ## Build all necessary docker images.
	make build_app build_nft_signer build_currency_signer build_card_generator

push: ## Push all necessary docker images.
	make push_app push_nft_signer push_currency_signer push_card_generator

docker: ## Build and push all necessary docker images.
	make build push

run_local: ## Build and run app locally.
	make build_dist && cd deploy/locall && docker-compose up

stop_local: ## Stop app locally.
	cd deploy/locall && docker-compose down
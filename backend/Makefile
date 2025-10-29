GCP_PROJECT ?= samhumphreys-website

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: run
run: ## Run the project
	PORT=8080 go run .

.PHONY: build
build: ## Build the project and output to server binary
	@echo "Building frontend..."
	@cd ../frontend && npm ci && npm run build
	@mkdir -p web/dist
	@rsync -a --delete ../frontend/dist/ web/dist/
	@echo "Building backend..."
	@go build -o server
	@echo "Build complete."

.PHONY: build-run
build-run: build ## Builds and runs the project using the server binary
	./server

.PHONY: clean
clean: ## Clean build artifacts
	@rm -rf web/dist server

##@ Deployment

.PHONY: deploy-deps
deploy-deps: ## Prepare deployment dependencies
	@gcloud config set project $(GCP_PROJECT)
	@gcloud auth application-default set-quota-project $(GCP_PROJECT)

.PHONY: deploy-no-promote
deploy-no-promote: deploy-deps ## Deploy the application without promoting
	@echo "Deploying to GCP project without promotion: $(GCP_PROJECT)"
	@gcloud app deploy --quiet --no-promote
	@echo "Deployment complete."

.PHONY: promote
promote: ## Promote a deployed version of the application
	@read -r -p "Version to promote: " V; \
	echo "Promoting version: $$V"; \
	gcloud app deploy --version "$$V" --quiet

.PHONY: deploy
deploy: deploy ## Deploy the application
	@echo "Deploying to GCP project: $(GCP_PROJECT)"
	@gcloud app deploy --quiet
	@echo "Deployment complete."

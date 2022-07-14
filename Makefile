#!make
-include .env

HUB ?= b0rr3g0
IMAGE ?= golang-hello-world
VERSION ?= 0.1
MAIN ?= app.go
CONTAINER ?= Dockerfile
.DEFAULT_GOAL := help

.PHONY: help 
help: ## Show opstions and short description
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build 
build: ## Build the golang application and generate the image
	@podman build -t ${HUB}/${IMAGE}:${VERSION} . -f ${CONTAINER}

.PHONY: image
image-push: build  ## Clean, build, generate the image and upload it
	@podman tag ${HUB}/${IMAGE}:${VERSION} ${HUB}/${IMAGE}:latest
	@podman push ${HUB}/${IMAGE}:${VERSION}
	@podman push ${HUB}/${IMAGE}:latest
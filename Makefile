############################################################
# Copyright (c) 2024 wsadza 
# Released under the MIT license
# ----------------------------------------------------------
#
############################################################

APP_NAME := kubeforge
APP_PATH := apps/$(APP_NAME)

# PATHS
PATH_MAIN       := $(DIR_CURRENT)/$(DIR_BUILD)
PATH_DOCKERFILE := $(DIR_CURRENT)/$(DIR_BUILD)/Dockerfile

# -----------------------------------------

.PHONY: local-run
local-run: 
	-@go -C $(APP_PATH) run cmd/main.go run \
			--kubernetesConfig ~/.kube/config \
			--sourceConfiguration "${PWD}/$(APP_PATH)/test/k8s/sourceConfiguration.yml"
	-@kubectl apply -f "${PWD}/$(APP_PATH)/test/k8s/crd.yaml
	-@kubectl apply -f "${PWD}/$(APP_PATH)/test/k8s/overlay.yaml

# -----------------------------------------

.PHONY: docker-build
docker-build:
	-@docker build \
			--file $(APP_PATH)/build/Dockerfile.alpine \
			--tag $(APP_NAME) \
			$(APP_PATH)

.PHONY: docker-run
docker-run:
	-@docker run --rm "$(APP_NAME):latest"

# -----------------------------------------

.PHONY: helm-uninstall
helm-uninstall:
	-@helm uninstall kubeforge

.PHONY: helm-install
helm-install: 
	-@cd charts/kubeforge && helm upgrade --install kubeforge -f values.yaml .

.PHONY: helm-test
helm-test: 
	-@helm test kubeforge

# -----------------------------------------

.PHONY: publish-k3s
publish-k3s: docker-build
	-@docker save $(APP_NAME):latest | k3s ctr images import -

# -----------------------------------------

.PHONY: test
test: docker-build publish-k3s helm-uninstall helm-install helm-test

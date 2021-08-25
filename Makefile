TAG=latest
ENV_TAG=latest

build-image:
	@docker build --rm -t ${REGISTRY}/${PROJECT_NAMESPACE}/${PROJECT_NAME}:${TAG} .
	@docker tag ${REGISTRY}/${PROJECT_NAMESPACE}/${PROJECT_NAME}:${TAG} ${REGISTRY}/${PROJECT_NAMESPACE}/${PROJECT_NAME}:${ENV_TAG}

push-image:
	@docker push ${REGISTRY}/${PROJECT_NAMESPACE}/${PROJECT_NAME}:${TAG}
	@docker push ${REGISTRY}/${PROJECT_NAMESPACE}/${PROJECT_NAME}:${ENV_TAG}

swag-init:
	swag init -g api/main.go

.PHONY: build

build:
	@eval $(minikube docker-env)
	@docker build -t fibonacci .

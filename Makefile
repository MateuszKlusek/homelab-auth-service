VERSION=v1
REGISTRY=api.mateuszklusek.com

build:
	docker buildx build --platform linux/amd64 -t $(REGISTRY)/auth-service:$(VERSION) --push .
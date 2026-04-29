VERSION=v1
REGISTRY=registry.mateuszklusek.com

build-and-push:
	docker buildx build --platform linux/amd64 -t $(REGISTRY)/auth-service:$(VERSION) --push .
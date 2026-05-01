.PHONY: build-container
build-container:
	docker build -t fleet-management-api-builder .

.PHONY: buf-generate
# Builds the local image first; otherwise Docker tries to pull fleet-management-api-builder from a registry.
buf-generate: build-container
	docker run --rm -v $(CURDIR):/api fleet-management-api-builder

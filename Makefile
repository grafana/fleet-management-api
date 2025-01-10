.PHONY: build-container
build-container:
	docker build -t fleet-management-api-builder .

.PHONY: buf-generate
buf-generate:
	docker run -v $(CURDIR):/api fleet-management-api-builder

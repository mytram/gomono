O = out

$(O):
	@mkdir -p $@

# Develop
.PHONY: lint
lint:
	golangci-lint run

# Build
.PHONY: build
build: | $(O)
	go build -o $(O) ./cmd/...

# Docker
.PHONY: docker-build-tetra
docker-build-tetra:
	docker build --build-arg SERVICE=tetra \
		--tag tetra \
		-f docker/go.Dockerfile .

docker-run-tetra: docker-build-tetra
	docker run --rm -p 8080:8080 tetra

# Test

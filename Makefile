# ------------------------------------------------------------------------------------------------------
#  Deployment/Build/Docker arguments

COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date "+%Y-%m-%d-%T:%N%Z")
VERSION := $(shell git describe --tags --always)

# ------------------------------------------------------------------------------------------------------


# ------------------------------------------------------------------------------------------------------
#  For development only

start:
	docker compose up

stop:
	docker compose down

tests:
	docker compose run \
		--rm --entrypoint="" \
		api sh -c "go test -v -timeout 500ms -cover ./..."
	$(MAKE) stop

fmt:
	docker compose run \
		--rm --entrypoint="" \
		api sh -c "go fmt ./..."
	$(MAKE) stop

# ------------------------------------------------------------------------------------------------------


# ------------------------------------------------------------------------------------------------------
#  For deployment only

build:
	docker build \
		--no-cache \
		--build-arg BUILD_DATE="$(BUILD_DATE)" \
		--build-arg COMMIT_HASH="$(COMMIT_HASH)" \
		--build-arg VERSION="$(VERSION)" \
		-t paymentsapi:$(VERSION) -f env/Dockerfile.prd .

# ------------------------------------------------------------------------------------------------------


.PHONY: start stop tests fmt build

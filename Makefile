# Values used in the build steps
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date "+%Y-%m-%d-%T:%N%Z")


start:
	docker compose up

stop:
	docker compose down

tests:
	docker compose run \
		--rm --entrypoint="" \
		api sh -c "go test -v -timeout 500ms -cover ./..."
	$(MAKE) stop

build:
	docker build \
		--no-cache \
		--build-arg BUILD_DATE="$(BUILD_DATE)" \
		--build-arg COMMIT_HASH="$(COMMIT_HASH)" \
		-t paymentsapi:$(COMMIT_HASH) -f env/Dockerfile.prd .


.PHONY: start stop tests build

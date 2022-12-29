.PHONY: test security run stop

export WAC_MAJOR_VERSION = 2021
export WAC_MINOR_VERSION = 1
export WAC_BUILD_VERSION = 0
export WAC_SESSION_PATH="./storage/sessions"
export WAC_UPLOAD_PATH="./storage/uploads"

export SERVER_NAME="GOWA v$(WAC_MAJOR_VERSION).$(WAC_MINOR_VERSION).$(WAC_BUILD_VERSION)"
export SERVER_DESCRIPTION="GOWA Service v$(WAC_MAJOR_VERSION).$(WAC_MINOR_VERSION).$(WAC_BUILD_VERSION)"
export SERVER_PORT="8080"
export SERVER_URL="0.0.0.0:$(SERVER_PORT)"
export SERVER_ENV=0 # 0 for debug 1 for release 2 for tests
export SERVER_READ_TIMEOUT=60
export SERVER_UPLOAD_LIMIT=1

BUILD_DIR = $(PWD)/bin/build/app

#security:
#	gosec -quiet ./...

#test: security
#	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
#	go tool cover -func=cover.out

swag: tests
	@ echo "Re-generate Swagger File (API Spec docs)"
	@ swag init --parseDependency --parseInternal \
		--parseDepth 4 -g ./cmd/service/main.go
	@ echo "done"


tests: $(GOTESTSUM) lint
	@ echo "Run tests"
	@ gotestsum --format pkgname-and-test-fails \
		--hide-summary=skipped \
		-- -coverprofile=cover.out ./...
	@ rm cover.out

lint: $(GOLANGCI)
	@ echo "Applying linter"
	@ golangci-lint cache clean
	@ golangci-lint run -c .golangci.yaml ./...

run:
	@echo "Run App"
	go mod tidy -compat=1.19
	go run ./cmd/app/main.go

#docker_build_image:
#	docker build -t gowa .

#docker_app: docker_build_image
#	docker run -d \
#        		--name gowa-c \
#        		-p $(SERVER_PORT):$(SERVER_PORT) \
#        		-e SERVER_NAME=$(SERVER_NAME) \
#        		-e SERVER_DESCRIPTION=$(SERVER_DESCRIPTION) \
#        		-e SERVER_UPLOAD_LIMIT=$(SERVER_UPLOAD_LIMIT) \
#        		-e SERVER_URL=$(SERVER_URL) \
#        		-e SERVER_ENV=$(SERVER_ENV) \
#        		-e SERVER_READ_TIMEOUT=$(SERVER_READ_TIMEOUT) \
#        		-e WAC_MAJOR_VERSION=$(WAC_MAJOR_VERSION) \
#        		-e WAC_MINOR_VERSION=$(WAC_MINOR_VERSION) \
#        		-e WAC_BUILD_VERSION=$(WAC_BUILD_VERSION) \
#        		-e WAC_SESSION_PATH=$(WAC_SESSION_PATH) \
#        		-e WAC_UPLOAD_PATH=$(WAC_UPLOAD_PATH) \
#        		gowa

#run: swag docker_app

#stop:
#	docker container stop gowa-c
#	docker container rm gowa-c
#	docker rmi gowa
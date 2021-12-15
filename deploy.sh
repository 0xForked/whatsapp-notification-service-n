## Whatsapp Client version
export WAC_MAJOR_VERSION=2021
export WAC_MINOR_VERSION=1
export WAC_BUILD_VERSION=0
## Whatsapp Client Paths
export WAC_SESSION_PATH="./storage/sessions"
export WAC_UPLOAD_PATH="./storage/uploads"
# Server Environment settings:
export SERVER_NAME="GOWA v${WAC_MAJOR_VERSION}.${WAC_MINOR_VERSION}.${WAC_BUILD_VERSION}"
export SERVER_DESCRIPTION="Whatsapp Web API with Golang (Gin Gonic)"
export SERVER_URL="0.0.0.0:8080"
export SERVER_ENV="debug" # debug | release
export SERVER_READ_TIMEOUT=60
export SERVER_UPLOAD_LIMIT=1
# Download all the dependencies that are required in your source files and update go.mod file with that dependency and
# remove all dependencies from the go.mod file which are not required in the source files.
go mod tidy
# Run Swagger
swag init --parseDependency --parseInternal --parseDepth 1
# Build the app
go build -o ./bin/build/app/gowa_"v${WAC_MAJOR_VERSION}.${WAC_MINOR_VERSION}.${WAC_BUILD_VERSION}" main.go
# Copy needed folders
cp -r ./storage ./bin/build/app/
cp -r ./docs ./bin/build/app/
# Run the app
#./bin/build/app/gowa_"v${WAC_MAJOR_VERSION}.${WAC_MINOR_VERSION}.${WAC_BUILD_VERSION}"
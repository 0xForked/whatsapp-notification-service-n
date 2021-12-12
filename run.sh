# Environment settings:
export SERVER_SHORT_NAME="GOWA v2021.1.0"
export SERVER_LONG_NAME="Whatsapp Web API with Golang (Gin Gonic)"
export SERVER_URL="0.0.0.0:8080"
export SERVER_ENV="debug" # debug | release
export SERVER_READ_TIMEOUT=60
export SERVER_UPLOAD_LIMIT=1
## WhatsApp Configuration
export WHATSAPP_CLIENT_VERSION_MAJOR=2021
export WHATSAPP_CLIENT_VERSION_MINOR=1
export WHATSAPP_CLIENT_VERSION_BUILD=0
export WHATSAPP_CLIENT_SESSION_PATH="./storage/temps"
# Download all the dependencies that are required in your source files and update go.mod file with that dependency and
# remove all dependencies from the go.mod file which are not required in the source files.
go mod tidy
# Run Swagger
swag init --parseDependency --parseInternal --parseDepth 1
# Run app
go run main.go
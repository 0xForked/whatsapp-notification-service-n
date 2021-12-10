# Environment settings:
export SERVER_URL="0.0.0.0:8080"
# debug | release
export SERVER_ENV="debug"
export SERVER_READ_TIMEOUT=60
## JWT Configuration
export JWT_SECRET_KEY="GOWA_JWT_SECRET:base64(string):amNjx+OkIltCJU3aTYhO3A=="
export JWT_SECRET_KEY_EXPIRE_MINUTES=15
## WhatsApp Configuration
export WHATSAPP_CLIENT_VERSION_MAJOR=2
export WHATSAPP_CLIENT_VERSION_MINOR=2126
export WHATSAPP_CLIENT_VERSION_BUILD=11
export WHATSAPP_CLIENT_SESSION_PATH="./storage/temps"
# Download all the dependencies that are required in your source files and update go.mod file with that dependency and
# remove all dependencies from the go.mod file which are not required in the source files.
go mod tidy
# Run Swagger
swag init
# Run app
go run main.go
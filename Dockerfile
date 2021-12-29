#FROM golang:1.x-alpine AS builder
FROM golang:alpine as builder

RUN apk update && apk add --no-cache

# Move to working directory (/app).
WORKDIR /bin/build

# Copy the code into the container.
COPY . .

# Download all the dependencies that are required in your source files and update go.mod file with that dependency.
# Remove all dependencies from the go.mod file which are not required in the source files.
RUN go mod tidy

# Build the application server.
RUN go build -o gowa .

## Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --no-cache --update add bash tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 8080

COPY --from=builder /bin/build/gowa /app
COPY --from=builder /bin/build/docs/swagger.json /app/docs/swagger.json
COPY --from=builder /bin/build/docs/swagger.yaml /app/docs/swagger.yaml
COPY --from=builder /bin/build/storage /app/storage

RUN chmod 755 /app/storage
RUN chmod 755 /app/storage/logs
RUN chmod 755 /app/storage/sessions
RUN chmod 755 /app/storage/uploads

# Command to run when starting the container.
CMD /app/gowa

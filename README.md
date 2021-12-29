# WhatsApp Notification Service with Golang

`gowa` Implements the WhatsApp Web API using [Gin Gonic](https://github.com/gin-gonic/gin) web framework,
and also [Gin Swagger](https://github.com/swaggo/gin-swagger) to generate Swagger Documentation (2.0).

## Getting Started
These instructions will get you a copy of the project up and running on docker container and on your local machine.

### Prerequisites
Perquisites package:
* [Docker](https://www.docker.com/get-started) - for developing, shipping, and running applications (Application Containerization).
* [Go](https://golang.org/) - Go Programming Language
* [Swag](https://github.com/swaggo/gin-swagger) - converts Go annotations to [Swagger Documentation 2.0](https://swagger.io/docs/specification/2-0/basic-structure/)
* [Make](https://www.gnu.org/software/make/manual/make.html) - Automated Execution using Makefile

### Running On Docker Container
1. Run project by using following command:
```bash
$ make run

# Process:
#   - Generate API docs by Swagger
#   - Build and run Docker containers
```
Stop application by using following command:
```bash
$ make stop

# Process:
#   - Stop and remove app container
#   - remove image
```

### Running On Local Machine
Below is the instructions to run this project on your local machine:
1. Open new `terminal`.
2. Set `run.sh` file permission.
```bash
$ chmod +x ./run.sh
```
4. Run application from terminal by using following command:
```bash
$ ./run.sh
```

### API Access
Go to your API Docs page: [127.0.0.1:8080/docs/index.html](http://127.0.0.1:8080/swagger/index.html)

### Important

To avoid getting banned from whatsapp, 
please don't send a spam message and make sure 
receiver is on your contacts list.
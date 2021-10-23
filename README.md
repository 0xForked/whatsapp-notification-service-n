# WhatsApp REST API Implementation Using Golang and Fiber Framework

`wagorf` Implements the WhatsApp Web API using [Fiber](https://github.com/gofiber/fiber) web framework,
and also [Swag](https://github.com/gofiber/fiber) to generate Swagger Documentation (2.0).
<br>This repository contains example of implementation [Rhymen/go-whatsapp](https://github.com/Rhymen/go-whatsapp) package.

## Getting Started
These instructions will get you a copy of the project up and running on docker container and on your local machine.

### Prerequisites
Prequisites package:
* [Docker](https://www.docker.com/get-started) - for developing, shipping, and running applications (Application Containerization).
* [Go](https://golang.org/) - Go Programming Language
* [Swag](https://github.com/swaggo/swag) - converts Go annotations to [Swagger Documentation 2.0](https://swagger.io/docs/specification/2-0/basic-structure/)
* [Make](https://www.gnu.org/software/make/manual/make.html) - Automated Execution using Makefile

### Running On Docker Container
1. Rename `Makefile.example` to `Makefile` and fill it with your make setting.
2. Run project by using following command:
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
1. Rename `run.sh.example` to `run.sh` and fill it with your environment values.
2. Open new `terminal`.
3. Set `run.sh` file permission.
```bash
$ chmod +x ./run.sh
```
4. Run application from terminal by using following command:
```bash
$ ./run.sh
```

### API Access
Go to your API Docs page: [127.0.0.1:8888/swagger/index.html](http://127.0.0.1:8888/swagger/index.html)


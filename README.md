# WhatsApp Notification Service

`Gowans` Implement the WhatsApp Web API powered by [rhymen/go-whatsapp](https://github.com/Rhymen/go-whatsapp).

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


### API Access
Go to your API Docs page: []()

### Important

To avoid getting banned from whatsapp, 
please don't send a spam message and make sure 
receiver is on your contacts list.

### TODO

- [ ] Refactor
- [ ] Tests
- [ ] Impl Auto-reply Bot
- [ ] Improve Handler
- [ ] Improve Service
- [ ] Impl Tracing, Logging, etc.
- [ ] Events handling
  - [ ] Message Forwarding 
  - [ ] Message for Bot with GPT
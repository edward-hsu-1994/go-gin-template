go-fiber-template
======

This is a Web API application template based on Fiber. This template includes Wire and Swagger, allowing you to quickly develop Web API applications.


## Setup

### Install Go-Wire

To install Go-Wire, run the following command:

```bash
go install github.com/google/wire/cmd/wire@latest
```

### Install swag

Since this project template uses Swagger to automatically generate API documents, you need to install swag before using it. You can use the following command to install swag:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Install gofiber-swagger

Since this project template uses the Swagger middleware of Fiber to provide Swagger UI, you need to install go-fiber-swagger before using it. You can use the following command to install go-fiber-swagger:

```bash
go get -u github.com/gofiber/swagger
```

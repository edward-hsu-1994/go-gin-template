﻿go-gin-template
======

This is a Web API application template based on Gin. This template includes Wire and Swagger, allowing you to quickly develop Web API applications.


## Setup

### Install Go-Wire

To install Go-Wire, run the following command:

```bash
go install github.com/google/wire/cmd/wire@latest
go mod tidy # This command is used to update the go.mod file
```

### Install swag

Since this project template uses Swagger to automatically generate API documents, you need to install swag before using it. You can use the following command to install swag:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Install gin-swagger

Since this project template uses the Swagger middleware of Gin to provide Swagger UI, you need to install gin-swagger before using it. You can use the following command to install gin-swagger:

```bash
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

## How dependencies injection works in this template

This project uses Wire to manage dependencies. You can use the following command to generate the dependency injection code:

In the wire.go file, we can see that three WireSets are defined, namely repoSet, serviceSet, routesSet,
which respectively define the dependency relationship of the Repository layer, Service layer, and Route layer.

It is worth noting that the route layer in this project defines an interface to describe the routing configuration for GinApp. All routing configurations must implement this interface.

```go
type GinRouter interface {
	ConfigureRoutes(app *Gin.Engine)
}


type PostRouter struct {
    _postService *services.PostService
}

// The PostRouter requires a PostService object to be injected into the constructor.
func NewPostRouter(postService *services.PostService) *PostRouter {
	return &PostRouter{
        _postService: postService,
    }
}
```

After defining the routing configuration, you can add the objects you need to Dependency injection in the constructor of the Route type.

```go
var routesSet = wire.NewSet(
	routes.NewNewsRouter,
	routes.NewPostRouter, // Add the Router object to the dependency injection set
	
	AssembleGinRouters,
)

// This method is used to assemble all routing configurations and return a list of GinRouters.
// This list will be used in main.go to configure the routes of GinApp.
func AssembleGinRouters(
	newsRouter *routes.NewsRouter, 
	postRouter *routes.PostRouter, // Inject the Router object into the AssembleGinRouters method
) []routes.GinRouter {
	return []routes.GinRouter{newsRouter, postRouter}
}
```

This AssembleGinRouters method is used to assemble all routing configurations and return a list of GinRouters. 
This list will be used in main.go to configure the routes of GinApp.

The dependency injection relationship diagram of the entire project is as follows:

![image](https://www.plantuml.com/plantuml/png/VP71JiCm38RlUOfez_44KpM4n0s4u0WSqimTXZIE71V4swCcKboqmCt-_Tdzy_UOnR4iSp019h52bl7y9lQ435wGeg7n7RpOsM6hTuU33oxdONY9jpW2NrsdjBCkMvVI1eAupC1k3B2Ipw-5VQH5eC2yLlaVVYgtRoXEU2uRlfGz6m-KHI-theVUmrTMj7L_NNq2_aHVOUstE4PXc9o7PWGIPHJYxxQbKkyxT-G_EekNiZ76fMJt-w7-a1f8wTVeQw8wRZaKgKvDU_5Mr9TLMHFyRQ0E5JRgyG9HoHmXdfOv_000)

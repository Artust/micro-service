# Services Structure

```
├── api // OpenAPI/Swagger specs, JSON schema files protocol definition files.                   
│   └── grpc // gRPC spec
│       └── example.proto
├── go.mod // Go mod for all microservice
├── go.sum
├── Makefile // Scripts to perform various build, install, analysis, etc operations.
├── pkg // Shared package for microservice, contain logger, monitor,...
│   └── logger
│       └── logger.go
├── README.md
└── services // Contain list microservices
    └── account_management // microservice name
        ├── cmd   // Main applications for service
        │   └── main.go
        ├── config // Contain configuration
        │   ├── config.go // environment variables
        │   └── constants.go // constants variables
        ├── Dockerfile
        ├── domain // Domain layer, interface for database layer
        │   ├── entity // model interface
        │   └── repository // repository interface
        ├── handler // controller, handle request 
        ├── infra // Database layer
        │   ├── mongodb
        │   └── neo4j
        │       ├── model  // database model
        │       └── repository // database repository
        ├── middleware // put the before & after logic of handle request
        ├── pkg // utils for individual service
        ├── protos // generated file contain gRPC method
        ├── router // for services use REST API
        └── usecase // business logic
```# micro-service
# micro-service
# micro-service
# micro-service

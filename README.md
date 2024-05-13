# Go Hexagonal Architecture Bootstrap
A simple repository with several getting started patterns.

## Releases

> **Note:** Each Branch in this repository contains a different release version. The main branch will always contain the latest release. Each release should work stand alone but features may vary between releases.

| Release    | Branch | Description |
| -------- | ------- |  ------- |
| v0.0.1-a.1  | release/v0.0.1-a.1 | HTTP Rest Server (Primary Adapter) with Core Adapter (Business Logic), No External Dependencies (No Secondary Adapters) |     

## Release v0.0.1-a.1
Go Hexagonal Architecture Bootstrap containing a Single Primary and **No Secondary** Adapters. 

### Primary Adapters/Port:  
_Contains the primary adapters of your application, representing the interfaces to infrastructure_
- **HTTP Web Server**: 
    - HTTP Adapter (services/bootstrap/internal/adapters/primary/http/server.go)
    - HTTP Port (services/bootstrap/internal/ports/primary.go)


### Application (App) Adapters/Port: 
_Combine the primary and secondary adapters of your core application_
- **API**: 
    - App/API Adapter (services/bootstrap/internal/adapters/app/api/api.go)
    - API Port (services/bootstrap/internal/ports/app.go)

### Core/Domain (Business Logic) Adapters/Port: 
_Contains the business logic and use cases of your application_
- **API**: 
    - Core Adapter (services/bootstrap/internal/adapters/core/core.go)
    - CorePort Port (services/bootstrap/internal/ports/core.go)


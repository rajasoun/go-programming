layout: true

.signature[@algogrit]

---

class: center, middle

# Go Web Apps

Gaurav Agarwal

---
class: center, middle

## `net/http` package

---

- Supports both `http` & `https`

- Also has support for certificate management

- Simple and powerful

---

- `http.ListenAndServe(<address>, <router>)`

- `r.Handle(<route>, <router>)`

- `r.HandleFunc(<route>, <handlerFunc>)`

---
class: center, middle

### Introducing [`gorilla` toolkit](https://www.gorillatoolkit.org/)

---
class: center, middle

#### `mux` [package](https://github.com/gorilla/mux)

---

#### [Other packages](https://www.gorillatoolkit.org/) of interest

- `gorilla/handlers`
- `gorilla/csrf`
- `gorilla/sessions`

- `gorilla/websocket`
- `gorilla/rpc`

---
class: center, middle

### [Project Layout](https://github.com/golang-standards/project-layout)

---
class: center, middle

### Clean Code Architecture

---
class: center, middle

![Clean Architecture](assets/images/CleanArchitecture.jpg)

---

#### Layers

- Entities (entities)

  Defines all the Models in the application

- Repository

  Encapsulates the interaction with the database. This is the lowest layer in the application.

- Services

  Orchestrates the interaction with repository and other services.

---

#### Layers (continued)

- Transport

  Encapsulates the interaction with application over an http API

- Binaries

  The code in `cmd/server` ties all the layers together, in order the start the app.

---
class: center, middle

Sample App: [YAES Server](https://github.com/algogrit/yaes-server)

---
class: center, middle

### Advance Testing

---
class: center, middle

#### Mocking

---
class: center, middle

*Go is a statically typed language.*

---
class: center, middle

[gomock](https://github.com/golang/mock) is a mocking framework for the Go programming language.

---
class: center, middle

`mockgen`

---
class: center, middle

`go:generate`

---
class: center, middle

#### HTTP Testing

---
class: center, middle

[`net/http/httptest`](https://golang.org/pkg/net/http/httptest/)

---
class: center, middle

#### Consistency Testing

---
class: center, middle

Failing the build due to *race* conditions...

---
class: center, middle

#### Fuzz testing

.content-credits[https://github.com/google/gofuzz]

---
class: center, middle

## gRPC

---

- Works on top of HTTP 2

- Usually, uses [protobuf](https://developers.google.com/protocol-buffers) for messaging, JSON is optional

  - Uses *proto3* syntax

- Auto-generates client and server side stubs

.content-credits[https://developers.google.com/protocol-buffers/docs/proto3]

---

- Supports both a unary (request-response) rpc as well as streaming (two-way communication)

- Middleware are known as interceptors

- Isn't well supported in the browser

---
class: center, middle

### Install protobuf for your platform

---

#### For Mac

```bash
brew install protobuf
```

---

### Install go gRPC plugin

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest # For generating message stubs
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest # For generating client & server stubs
```

.content-credits[https://grpc.io/docs/languages/go/quickstart/#prerequisites]

---
class: center, middle

### Writing your first `.proto` file

---

### Generating stubs

```bash
protoc \
  --go_out=. --go_opt=paths=source_relative \ # For generating message stubs
  --go-grpc_out=. --go-grpc_opt=paths=source_relative  \ # For generating client & server stubs
  api/<filename>.proto
```

---

### Steps

- Define a new `.proto` file
  - Define the req & response `message`s
  - Define the `service`
- Generate stubs
- Implement the server
- Wire up the server for connections
- Dial the client to the server

---

```golang
import "google.golang.org/grpc"

func main() {
    grpc.NewServer(...) // Server side

    grpc.Dial(...) // Client side
}
```

.content-credits[https://pkg.go.dev/google.golang.org/grpc]

---
class: center, middle

### Interceptors

---
class: center, middle

[go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)

---
class: center, middle

## Alternatives to ReST or gRPC?

---
class: center, middle

### [GraphQL](https://github.com/graphql-go/graphql)

.content-credits[https://pkg.go.dev/github.com/graphql-go/graphql]

<!-- ---
class: center, middle

### [RSocket](https://rsocket.io/)

.content-credits[https://www.youtube.com/watch?v=jf0xhnoezC4] -->

---
class: center, middle

## Working with DB

---
class: center, middle

`database/sql`

---

### Interfaces

- `sql.Scanner`
- `driver.Valuer`

---
class: center, middle

### Third-party packages

---
class: center, middle

[`sqlx`](https://github.com/jmoiron/sqlx)

.content-credits[http://jmoiron.github.io/sqlx/]

---
class: center, middle

[`gorm`](https://github.com/go-gorm/gorm)

.content-credits[https://gorm.io/index.html]

---
class: center, middle

### What about NoSQL?

---
class: center, middle

[`mongo-go-driver`](https://github.com/mongodb/mongo-go-driver)

.content-credits[https://www.mongodb.com/languages/golang]

---
class: center, middle

Code
https://github.com/algogrit/presentation-go-web-apps

Slides
https://go-web-apps.slides.algogrit.com

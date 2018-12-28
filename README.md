[![Build Status](https://travis-ci.org/Didstopia/tinyserver.svg?branch=master)](https://travis-ci.org/Didstopia/tinyserver)
[![codecov](https://codecov.io/gh/Didstopia/tinyserver/branch/master/graph/badge.svg)](https://codecov.io/gh/Didstopia/tinyserver)

# tinyserver

A web server for Go, with the primary purpose of being small, fast and extremely light on resources.  

**WARNING:** _Work in progress, do not use in production yet!_  

## Installation

```sh
go get -u github.com/Didstopia/tinyserver
```

## Usage

```go
package main

import (
  "github.com/Didstopia/tinyserver"
)

func main() {
  // Create a new server
  server, err := NewServer()
  if err != nil { panic(err) }

  // Create a new route
  route := func(w http.ResponseWriter, r *http.Request) { w.Write("Hello World!") }
  if err = server.AddRoute("/route", route); err != nil { panic(err) }

  // Start the server, listening on port 8080
  // NOTE: You can also use "ListenAsync" to run the server asynchronously (non-blocking)
  if err = server.Listen("8080"); err != nil { panic(err) }
}
```

## Benchmarks

_TODO_

## License

See [LICENSE](LICENSE).  

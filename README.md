# expvar

[![Run Tests](https://github.com/elliotxx/expvar/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/elliotxx/expvar/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/elliotxx/expvar/branch/master/graph/badge.svg)](https://codecov.io/gh/elliotxx/expvar)
[![Go Report Card](https://goreportcard.com/badge/github.com/elliotxx/expvar)](https://goreportcard.com/report/github.com/elliotxx/expvar)
[![GoDoc](https://godoc.org/github.com/elliotxx/expvar?status.svg)](https://godoc.org/github.com/elliotxx/expvar)

A expvar handler for gin framework, [expvar](https://golang.org/pkg/expvar/) provides a standardized interface to public variables.

## Usage

### Start using it

Download and install it:

```sh
go get github.com/elliotxx/expvar
```

Import it in your code:

```go
import "github.com/elliotxx/expvar"
```

### Canonical example

```go
package main

import (
  "log"

  "github.com/elliotxx/expvar"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/debug/vars", expvar.Handler())

  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

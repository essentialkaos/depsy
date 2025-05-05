<p align="center"><a href="#readme"><img src=".github/images/card.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/depsy"><img src=".github/images/godoc.svg"/></a>
  <a href="https://kaos.sh/r/depsy"><img src="https://kaos.sh/r/depsy.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/y/depsy"><img src="https://kaos.sh/y/e582190205934d328a62649aae1ed52b.svg" alt="Codacy Badge" /></a>
  <a href="https://kaos.sh/c/depsy"><img src="https://kaos.sh/c/depsy.svg" alt="Coverage Status" /></a>
  <a href="https://kaos.sh/w/depsy/ci"><img src="https://kaos.sh/w/depsy/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/depsy/codeql"><img src="https://kaos.sh/w/depsy/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src=".github/images/license.svg"/></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage-example">Usage example</a> • <a href="#ci-status">CI Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`depsy` is simple Go package with minimal dependencies for parsing dependencies info in from go.mod files.

### FAQ

**Q: Why should I  use `depsy` instead of using [`runtime/debug.ReadBuildInfo()`](https://pkg.go.dev/runtime/debug@go1.20.1#ReadBuildInfo) for reading already embedded info?**

**A:** First of all — size:

| Binary | Size (in bytes) |
|--------|-----------------|
| Original binary | `1819094` |
| depsy + embedded `go.mod` | `1861531` (+ 42,437) |
| `runtime/debug` | `1891541` (+ 72,447) |

Second reason — with debug package, you can't print only direct dependencies.

### Installation

Make sure you have a working [Go 1.23+](https://github.com/essentialkaos/.github/blob/master/GO-VERSION-SUPPORT.md) workspace (_[instructions](https://go.dev/doc/install)_), then:

````bash
go get github.com/essentialkaos/depsy
````

### Usage example

```go
package main

import (
  _ "embed"

  "fmt"

  "github.com/essentialkaos/depsy"
)

//go:embed go.mod
var modules []byte

func main() {
  deps := depsy.Extract(modules, false)

  for _, dep := range deps {
    fmt.Println(dep)
  }
}
```

### CI Status

| Branch | Status |
|--------|----------|
| `master` | [![CI](https://kaos.sh/w/depsy/ci.svg?branch=master)](https://kaos.sh/w/depsy/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/depsy/ci.svg?branch=develop)](https://kaos.sh/w/depsy/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/.github/blob/master/CONTRIBUTING.md).

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://kaos.dev"><img src="https://raw.githubusercontent.com/essentialkaos/.github/refs/heads/master/images/ekgh.svg"/></a></p>

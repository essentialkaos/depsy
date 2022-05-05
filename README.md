<p align="center"><a href="#readme"><img src="https://gh.kaos.st/depsy.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/depsy.v1"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev" /></a>
  <a href="https://kaos.sh/w/depsy/ci"><img src="https://kaos.sh/w/depsy/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/r/depsy"><img src="https://kaos.sh/r/depsy.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/c/depsy"><img src="https://kaos.sh/c/depsy.svg" alt="Coverage Status" /></a>
  <a href="https://kaos.sh/b/depsy"><img src="https://kaos.sh/b/d2067e6e-8722-4f20-8274-4398ffa09e97.svg" alt="Codebeat badge" /></a>
  <a href="https://kaos.sh/w/depsy/codeql"><img src="https://kaos.sh/w/depsy/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#ci-status">CI Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`depsy` is simple Go package for parsing dependencies info from go.mod files with minimal dependencies.

### Installation

Make sure you have a working Go 1.17+ workspace (_[instructions](https://golang.org/doc/install)_), then:

````bash
go get -d github.com/essentialkaos/depsy
````

### Usage example

```go
package main

import (
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

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>

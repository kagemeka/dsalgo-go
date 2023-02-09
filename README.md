# dsalgo

Datastructures and Algorithms.

[![Go Reference](https://pkg.go.dev/badge/github.com/kagemeka/dsalgo-go.svg)](https://pkg.go.dev/github.com/kagemeka/dsalgo-go)
[![package][ci-badge]][ci-url]
[![License: MIT][mit-badge]][mit-url]
[![pre-commit][pre-commit-badge]][pre-commit-url]

[ci-badge]: https://github.com/kagemeka/dsalgo-go/actions/workflows/ci.yml/badge.svg
[ci-url]: https://github.com/kagemeka/dsalgo-go/actions/workflows/ci.yml
[pre-commit-badge]: https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=yellow
[pre-commit-url]: https://github.com/pre-commit/pre-commit
[mit-badge]: https://img.shields.io/badge/License-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT

## install

```sh
go get -u github.com/kagemeka/dsalgo-go/...
```

## example

```go
package main

import "github.com/kagemeka/dsalgo-go/src/<pkg/path>" // dsalgo packages
import "github.com/kagemeka/dsalgo-go/cp" // cp package

func main() {
}
```

## development

```sh
docker compose up -d
```

(enter the container)

```sh
./ci.sh
```

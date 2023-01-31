# dsalgo Go

Datastructures and Algorithms written in Go.

## how to use

### install

```sh
go get -u github.com/kagemeka/dsalgo-go/...
```

### import to your codes

```go
package main

import "github.com/kagemeka/dsalgo-go/src/<pkg/path>" // dsalgo packages
import "github.com/kagemeka/dsalgo-go/cp" // cp package

func main() {
}
```

## document

<https://pkg.go.dev/github.com/kagemeka/dsalgo-go>

## structure

- cp
  - easy package for competitive programming.
  - it can be bundled into one submission file.
  - simple or no tests.

- src
  - packages for deep and wide range of datastructures/algorithms.
  - multiple implementation for each algorithm.
  - well tested.
  - package naming is not satisfying some standard convention to avoid namespace conflict.
    - voiolating a rule "abbreviated or short name".
    - each package name is same to the directory and file name.
    - snake_case

## development

### use docker

```sh
docker compose up -d
```

### CI locally before commit & push

```sh
./ci.sh
```

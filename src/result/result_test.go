package result

import (
	"testing"

	. "github.com/kagemeka/dsalgo-go/src/assert"
	. "github.com/kagemeka/dsalgo-go/src/unreachable"
)

type A struct {
	x int
}

func (a A) IsOk() bool {
	return a.x < 1
}
func (a A) Ok() int {
	if !a.IsOk() {
		Unreachable()
	}
	Assert(a.IsOk())
	return a.x
}

type Err struct{}

func (e Err) Error() string {
	return "error occured"
}

func (a A) Err() Err {
	Assert(!a.IsOk())
	return Err{}
}

func unwrapResult[T any](res Result[T, Err]) T {
	return res.Ok()
}

func Test(t *testing.T) {
	res := A{-1}
	Assert(unwrapResult[int](res) == -1)
}

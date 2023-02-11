package result

type Result[T, E any] interface {
	IsOk() bool
	Ok() T
	Err() E
}

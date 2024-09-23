package result

type Result[T any] struct {
	t   *T
	err error
}

func NewResult[T any](t *T, err error) *Result[T] {
	return &Result[T]{t, err}
}

func ErrorResult[T any](err error) *Result[T] {
	return &Result[T]{err: err}
}

func OkResult[T any](t *T) *Result[T] {
	return &Result[T]{t: t}
}

func (r *Result[T]) Error(fun func(error)) *Result[T] {
	if r.err != nil {
		fun(r.err)
	}

	return r
}

func (r *Result[T]) Ok(fun func(t *T)) {
	if r.err == nil {
		fun(r.t)
	}
}

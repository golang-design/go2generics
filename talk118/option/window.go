package main

type Window struct {
	cfg1, cfg2 any
}

func NewWindow(opts ...Option) *Window {
	w := &Window{}
	for _, opt := range opts {
		opt(w)
	}
	return w
}

type Option func(w *Window)

func Config1(cfg1 any) Option {
	return func(w *Window) {
		w.cfg1 = cfg1
	}
}

func Config2(cfg2 any) Option {
	return func(w *Window) {
		w.cfg2 = cfg2
	}
}

type GenericWindow[T1, T2 any] struct {
	cfg1 T1
	cfg2 T2
}

type GenericOption[T1, T2 any] func(w *GenericWindow[T1, T2])

func GenericConfig1[T1, T2 any](cfg1 T1) GenericOption[T1, T2] {
	return func(w *GenericWindow[T1, T2]) {
		w.cfg1 = cfg1
	}
}

func GenericConfig2[T1, T2 any](cfg2 T2) GenericOption[T1, T2] {
	return func(w *GenericWindow[T1, T2]) {
		w.cfg2 = cfg2
	}
}

func NewGenericWindow[T1, T2 any](opts ...GenericOption[T1, T2]) *GenericWindow[T1, T2] {
	w := &GenericWindow[T1, T2]{}
	for _, opt := range opts {
		opt(w)
	}
	return w
}

func main() {
	_ = NewWindow(Config1(1), Config2("2"))

	_ = NewGenericWindow[int, string](
		GenericConfig1[int, string](1),
		GenericConfig2[int, string]("2"),
	)
}

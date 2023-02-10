package cp

func assert(ok bool) {
	if !ok {
		panic("not true")
	}
}

func unreachable() {
	panic("unreachable")
}

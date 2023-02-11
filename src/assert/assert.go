package assert

func Assert(ok bool) {
	if !ok {
		panic("not true")
	}
}

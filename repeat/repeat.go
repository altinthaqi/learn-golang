package repeat

func Repeat(s string, c int) (res string) {
	for i := 0; i < c; i++ {
		res += s
	}

	return
}

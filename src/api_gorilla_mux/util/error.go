package util

func ShowError(e error) {
	if e != nil {
		panic(e)
	}
}

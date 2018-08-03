package main

func showError(e error) {
	if e != nil {
		panic(e)
	}
}

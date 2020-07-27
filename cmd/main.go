package main

import (
	"fmt"

	"github.com/martinkunc/gowasmer/exec"
)

func main() {
	w, close, err := exec.NewWebAssembly()
	if err != nil {
		panic(err)
	}
	_ = close
	//defer close()

	if err != nil {
		panic(err)
	}
	result, err := w.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("main: %s \n", result)
}

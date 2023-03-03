package main

import (
	"os"
)

func (w world) world() string {
	w.string = os.Getenv("HELLO_OUTPUT")

	return "Hello" + w.string
}

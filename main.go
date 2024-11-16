package main

import "os"

func main() {
	addr := os.Getenv("GH2ALTSTORE_ADDR")
	if len(addr) == 0 {
		addr = ":8000"
	}

	serve(addr)
}

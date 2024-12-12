package main

import (
	"fmt"
	"io"
	"os"

	"github.com/minio/sha256-simd"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	shaWriter := sha256.New()
	file, err := os.Open(os.Args[1])
	check(err)
	io.Copy(shaWriter, file)
	fmt.Printf("%x", shaWriter.Sum(nil))
}

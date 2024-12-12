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

	switch os.Args[1] {
	case "simple":
		file, err := os.Open(os.Args[2])
		check(err)
		io.Copy(shaWriter, file)
		fmt.Printf("%x", shaWriter.Sum(nil))
	case "server":
		server := sha256.NewAvx512Server()
		h512 := sha256.NewAvx512(server)
		h512.Write(fileBlock)
		fmt.Printf("%x", h512.Sum([]byte{}))
	}
}

package main

import (
	"fmt"
	"io"
	"log"
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
	file, err := os.Open(os.Args[2])
	check(err)

	switch os.Args[1] {
	case "simple":
		io.Copy(shaWriter, file)
		fmt.Printf("%x", shaWriter.Sum(nil))
	case "server-avx512":
		server := sha256.NewAvx512Server()
		h512 := sha256.NewAvx512(server)
		buf := make([]byte, 2<<20)
		for {
			n, err := file.Read(buf)
			// fmt.Printf("read %d, %v\n", n, err)
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			h512.Write(buf[:n])
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
		}
		fmt.Printf("%x", h512.Sum([]byte{}))
	}
}

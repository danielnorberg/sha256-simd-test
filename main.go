package main

import (
	"fmt"
	"hash"
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
		fmt.Printf("%x\n", shaWriter.Sum(nil))
	case "server-avx512":
		server := sha256.NewAvx512Server()
		h512 := sha256.NewAvx512(server)
		for {
			buf := make([]byte, 2<<20)
			n, err := file.Read(buf)
			if n == 0 {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			h512.Write(buf[0:n])
		}
		fmt.Printf("%x\n", h512.Sum([]byte{}))
	case "server-avx512-parallel-16":
		server := sha256.NewAvx512Server()
		h512s := make([]hash.Hash, 16)
		for i := 0; i < len(h512s); i++ {
			h512s[i] = sha256.NewAvx512(server)
		}
		for {
			buf := make([]byte, 2<<20)
			n, err := file.Read(buf)
			if n == 0 {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			for i := 0; i < len(h512s); i++ {
				h512s[i].Write(buf[0:n])
			}
		}
		for i := 0; i < len(h512s); i++ {
			fmt.Printf("%x\n", h512s[i].Sum([]byte{}))
		}
	}
}

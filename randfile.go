// +build ignore

package main

import (
	"crypto/rand"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(log.Lshortfile)

	if len(os.Args) < 3 {
		log.Fatal("not enough parameter")
	}

	lenStr := os.Args[2]
	length, err := strconv.Atoi(lenStr)
	if err != nil {
		log.Fatal("error length: %v: %v", lenStr)
	}
	oname := os.Args[1]
	fout, err := os.Create(oname)
	if err != nil {
		log.Fatalf("cannot create file %v: %v", oname, err)
	}
	defer fout.Close()
	log.Printf("creating file of length %v byte(s)", length)

	//
	blockSize := 65536
	buffer := make([]byte, blockSize)
	cc := length / blockSize
	for i := 0; i < cc; i++ {
		rand.Read(buffer)
		fout.Write(buffer)
	}
	remains := length % blockSize
	rand.Read(buffer[:remains])
	fout.Write(buffer[:remains])
}

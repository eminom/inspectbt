package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/eminom/gobencode"
)

var (
	fVerbose    = flag.Bool("v", false, "verbose")
	fSingleHash = flag.Bool("s", false, "single hash enabled")
)

func init() {
	flag.Parse()

	bencode.SetSingleHashEnabled(*fSingleHash)
}

func main() {

	// log.Printf("%v", bencode.BNodeStringHex)
	// os.Exit(-1)
	log.SetFlags(log.Lshortfile)

	if len(flag.Args()) < 1 {
		log.Fatal("error: parameter")
	}

	ib := flag.Args()[0]
	chunk, err := ioutil.ReadFile(ib)
	if err != nil {
		log.Fatalf("cannot read %v", ib)
	}

	node, left := bencode.Scan(chunk)
	if len(left) > 0 {
		log.Printf("warnings: still some bytes left: %v", left)
	}

	if *fVerbose {
		bencode.PrintNode(node, node.Cat, 2)
	}

	if len(flag.Args()) < 2 {
		os.Exit(0)
	}

	rawFile := flag.Args()[1]
	t := bencode.NewTorrent(node.AsMap()["info"].AsMap())

	log.Printf("Start verifying. %v", strings.Repeat("#", 20))

	verified, err := t.VerifyFile(rawFile)
	if err != nil {
		log.Fatalf("error verifying: %v", err)
	}

	log.Printf("verified: %v", verified)
}

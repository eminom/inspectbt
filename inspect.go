package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eminom/gobencode"
)

var (
	fVerbose    = flag.Bool("v", false, "verbose")
	fVerboseTwo = flag.Bool("vv", false, "verbose more")
	fSingleHash = flag.Bool("s", false, "single hash enabled")
)

func init() {
	flag.Parse()

	bencode.SetSingleHashEnabled(*fSingleHash)
}

func main() {

	// log.Printf("%v", bencode.BNodeStringHex)
	// os.Exit(-1)
	log.SetFlags(log.Lshortfile | log.Ltime)
	startTS := time.Now()

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

	if *fVerbose || *fVerboseTwo {
		if *fVerboseTwo {
			bencode.SetVerboseLevel(2)
		}
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
	log.Printf("time elapsed: %v", time.Now().Sub(startTS))
}

package main

import (
	"flag"
	"log"
)

var (
	docs      = flag.String("d", "docs", "Path to Markdown docs directory")
	artifacts = flag.String("a", "artifacts", "Path to artifacts")
)

func main() {
	flag.Parse()
	release, err := NewRelease(*docs)
	if err != nil {
		log.Fatal(err)
	}

	for _, fn := range []Exporter{ToJSON, ToCSV} {
		if err := fn(*artifacts, release); err != nil {
			log.Fatal(err)
		}
	}
}

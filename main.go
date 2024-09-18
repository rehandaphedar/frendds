package main

import (
	"fmt"
	"log"
	"os"

	"git.sr.ht/~rehandaphedar/frendds/pkg/relations"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: frendds [domain]")
	}
	root := os.Args[1]

	rootRelations := relations.GetRelations(root)

	for _, relation := range rootRelations {
		fmt.Printf("\"%s\" -> \"%s\"\n", relation.Source, relation.Target)
	}
}

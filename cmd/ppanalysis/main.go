package main

import (
	"log"
	"os"

	"github.com/akito0107/ppanalysis"
)

func main() {
	body, err := ppanalysis.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	ppanalysis.Print(os.Stdout, body)
}

package main

import (
	"bytes"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Failed to parse args, require >= 3: %d, %s", len(os.Args), os.Args)
	}
	patternArg := os.Args[1]

	patternList := strings.Split(patternArg, "@")
	if len(patternList) < 2 {
		log.Fatalf("Failed to parse pattern list, require >=2: %d", len(patternList))
	}
	find := []byte(patternList[0])
	replace := []byte(patternList[1])

	for i := 2; i < len(os.Args); i++ {
		fname := os.Args[i]
		b, err := os.ReadFile(fname)
		if err != nil {
			log.Fatalf("Failed to read %s: %v", fname, err)
		}

		newB := bytes.ReplaceAll(b, find, replace)
		os.WriteFile(fname, newB, 0644)
	}
}

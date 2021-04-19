// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fmt.Println(os.Args[:1])
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		filesRepeated := make([]string, 0, 10)
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// We check if we have some lines repeated
			hasFileRepeated := countLines(f, counts)
			// If so, we save it in the slice
			if hasFileRepeated {
				filesRepeated = append(filesRepeated, arg)
			}
			f.Close()
		}
		// Finally we print all the names that has repeated lines
		for _, nameFile := range filesRepeated {
			fmt.Println(nameFile)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(f)
	hasFileRepeated := false
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if counts[text] > 1 {
			hasFileRepeated = true
		}
	}
	return hasFileRepeated
	// NOTE: ignoring potential errors from input.Err()
}

//!-

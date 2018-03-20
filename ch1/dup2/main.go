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
	foundIn:=make(map[string]string)
	// fmt.Printf(files)
	if len(files) == 0 {
		countLines(os.Stdin, counts,foundIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts,foundIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n < 2 && n>0 {
			fmt.Printf("%d\t%s\t \n", n, line)
			
		}
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\t \n", n, line,foundIn[line])
		}
		
	}
}

func countLines(f *os.File, counts map[string]int,foundIn map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		
		counts[input.Text()]++
		if(counts[input.Text()]>1){
			foundIn[input.Text()]=f.Name()
			// fmt.Printf("again:%s\n")
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-

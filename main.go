package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
)

// last modified 2024-01-10.0945

/*
Cyclone's RegX: A Flexible Potfile Parsing Tool

GNU GENERAL PUBLIC LICENSE 2.0
https://github.com/cyclone-github/regx?tab=GPL-2.0-1-ov-file#

Features:
	Versatile CLI utility designed for parsing text files using regular expressions
	Built in set of hash regex using popular hashcat modes, md5 example: -m 0 or -m md5
	Allows custom RE2 compatible regular expressions, ex: -r '^[a-fA-F0-9]{32}'

Version:
	v0.1.0; initial release
*/

func main() {
	mode := flag.String("m", "", "Hash mode, ex: -m md5")
	customRegex := flag.String("r", "", "Custom regex, ex: -r '^[a-fA-F0-9]{32}'")
	hashFile := flag.String("f", "", "File containing hashes")
	cycloneFlag := flag.Bool("cyclone", false, "")
	versionFlag := flag.Bool("version", false, "Version info")
	helpFlag := flag.Bool("help", false, "Prints help")
	flag.Parse()

	// run sanity checks for special flags
	if *versionFlag {
		versionFunc()
		os.Exit(0)
	}
	if *cycloneFlag {
		cycloneFunc()
		os.Exit(0)
	}

	if *helpFlag {
		helpFunc()
		os.Exit(0)
	}

	if *hashFile == "" {
		fmt.Fprintln(os.Stderr, "Usage: missing -f {file}")
		os.Exit(1)
	}

	if *mode != "" && *customRegex != "" {
		fmt.Fprintln(os.Stderr, "Usage: cannot use -m {mode} and -r {regex} at the same time")
		os.Exit(1)
	}

	if *mode == "" && *customRegex == "" {
		fmt.Fprintln(os.Stderr, "Usage: missing -m {mode} or -r {regex}")
		fmt.Fprintln(os.Stderr, "Examples:\n-m 0\n-m md5\n-r '^[a-fA-F0-9]{32}'")
		os.Exit(1)
	}

	file, err := os.Open(*hashFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	numWorkers := runtime.NumCPU() * 10 // load CPU by creating more workers
	bufferSize := numWorkers * 1000     // channel buffer size
	workChan := make(chan []byte, bufferSize)
	resultsChan := make(chan []byte, numWorkers)
	var wg sync.WaitGroup      // worker goroutine sync group
	var writeWg sync.WaitGroup // writer sync group

	// compile regex and check compile errors
	regex, err := compileRegex(*mode, *customRegex)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error compiling regex:", err)
		os.Exit(1)
	}

	// start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(regex, workChan, resultsChan, &wg)
	}

	// start reader goroutine
	go func() {
		if err := startReader(file, workChan); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		}
	}()

	// start writer goroutine
	writeWg.Add(1)
	go startWriter(resultsChan, &writeWg)

	// wait for all workers to finish
	wg.Wait()
	close(resultsChan)
	writeWg.Wait()
}

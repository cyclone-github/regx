package main

import (
	"regexp"
	"sync"
)

// last modified 2024-01-10.0945

// worker goroutine
func worker(regex *regexp.Regexp, workChan <-chan []byte, resultsChan chan<- []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for lineBytes := range workChan {
		if regex.Match(lineBytes) {
			resultsChan <- lineBytes
		}
	}
}

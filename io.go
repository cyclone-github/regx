package main

import (
	"bufio"
	"io"
	"os"
	"sync"
	"time"
)

// last modified 2024-01-10.0945

// file reader
func startReader(file *os.File, workChan chan<- []byte) error {
	bufferSize := 10 * 1024 * 1024 // 10MB read buffer
	reader := bufio.NewReaderSize(file, bufferSize)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break // end of file
			}
			return err // return error
		}
		workChan <- line
	}
	close(workChan)
	return nil
}

// output writer
func startWriter(resultsChan <-chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	outputBuffer := bufio.NewWriterSize(os.Stdout, 1*1024*1024) // 1MB write buffer

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case resultBytes, ok := <-resultsChan:
			if !ok {
				outputBuffer.Flush()
				return
			}
			outputBuffer.Write(resultBytes)

		case <-ticker.C:
			// flush the buffer every ticker
			outputBuffer.Flush()
		}
	}
}

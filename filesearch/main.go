package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches   []string
	waitgroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
)

func fileSearch(directory string, filename string) {
	fmt.Println("Searching in", directory)
	files, error := ioutil.ReadDir(directory)
	if error != nil {
		fmt.Println(error)
		return
	}
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(directory, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			waitgroup.Add(1)
			go fileSearch(filepath.Join(directory, file.Name()), filename)
		}
	}
	waitgroup.Done()
}

func main() {
	waitgroup.Add(1)
	go fileSearch("/home/daniel", "README.md")
	waitgroup.Wait()
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
}

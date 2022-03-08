package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency *[26]int32, waitGroup *sync.WaitGroup) {
	resp, error := http.Get(url)
	if error != nil {
		fmt.Printf("An error occurred while requesting the content for %s\n", url)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		index := strings.Index(letters, c)
		if index >= 0 {
			atomic.AddInt32(&frequency[index], 1)
		}
	}
	waitGroup.Done()
}

func main() {
	var frequency [26]int32
	waitGroup := sync.WaitGroup{}
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		waitGroup.Add(1)
		go countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency, &waitGroup)
	}
	waitGroup.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)
	for i, f := range frequency {
		fmt.Printf("%s -> %d\n", string(letters[i]), f)
	}
}

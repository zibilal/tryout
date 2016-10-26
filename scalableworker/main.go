package main

import (
	"sync"
	"bufio"
	"os"
	"log"
	"fmt"
	"net/http"
)

type task interface {
	process()
	output()
}

type factory interface {
	create(line string) task
}

func run(f factory) {
	var wg sync.WaitGroup

	in := make(chan task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.create(s.Text())
		}

		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan task)

	for i:= 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for t:= range in {
				t.process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	} ()

	for t:= range out {
		t.output()
	}
}

type HTTPTask struct {
	url string
	ok bool
}

func (h *HTTPTask) process() {
	resp, err := http.Get(h.url)
	if err != nil {
		h.ok = false
		return
	}

	if resp.StatusCode == http.StatusOK {
		h.ok = true
		return
	}

	h.ok = false
}

func (h *HTTPTask) output() {
	fmt.Printf("%s %t\n", h.url, h.ok)
}

type Factory struct {

}
func (f *Factory) create(line string) task {
	h := &HTTPTask{}
	h.url = line
	return h
}

func main() {
	f := &Factory{}
	run(f)
}

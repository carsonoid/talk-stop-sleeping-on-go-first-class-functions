package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// START OMIT
type printer struct {
	sync.Mutex
	w io.Writer
}

func (p *printer) print(s string) {
	p.Lock()
	defer p.Unlock()
	fmt.Fprintln(p.w, s)
}

// END OMIT

// START MAIN OMIT
func main() {
	p := &printer{w: os.Stdout}
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			p.print(fmt.Sprintf("Hello, World! %d", i))
		}()
	}
	wg.Wait()
}

// END MAIN OMIT

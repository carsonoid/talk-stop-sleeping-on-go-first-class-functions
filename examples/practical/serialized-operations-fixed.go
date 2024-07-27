package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// START OMIT
func main() {
	p := newPrinter(3)
	go p.processOps()
	for i := 0; i < 10; i++ {
		i := i // copy loop variable
		p.queue(
			func(w io.Writer) {
				fmt.Fprint(w, "Hello, ")
				fmt.Fprintf(w, "World! %d\n", i)
			},
		)
	}
	p.Wait()
}

// END OMIT

// START PRINTER1 OMIT
type printFunc func(w io.Writer)

type printer struct {
	sync.WaitGroup
	w       io.Writer
	opsChan chan printFunc
}

func newPrinter(n int) *printer {
	return &printer{
		w:       os.Stdout,
		opsChan: make(chan printFunc, n),
	}
}

// END PRINTER1 OMIT
// START PRINTER2 OMIT

func (p *printer) queue(fn printFunc) {
	p.Add(1)
	select {
	case p.opsChan <- fn:
	default:
		go func() { p.opsChan <- fn }()
	}
}

func (p *printer) processOps() {
	time.Sleep(250 * time.Millisecond) // test sleep to ensure queue func hits default case
	for fn := range p.opsChan {
		fn(p.w)
		p.Done()
	}
}

// END PRINTER2 OMIT

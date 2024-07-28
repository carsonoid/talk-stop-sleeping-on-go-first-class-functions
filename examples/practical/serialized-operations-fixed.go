package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func main() {
	p := newPrinter(3)
	go p.processOps()
	for i := 0; i < 10; i++ {
		i := i // copy loop variable
		p.queue(func() { fmt.Printf("Hello, World! %d\n", i) })
	}
	p.Wait()
}

// END OMIT

// START PRINTER OMIT
type printer struct {
	sync.WaitGroup
	opsChan chan func()
}

func newPrinter(n int) *printer {
	return &printer{opsChan: make(chan func(), n)}
}

func (p *printer) queue(op func()) {
	p.Add(1)
	select {
	case p.opsChan <- op:
	default:
		go func() { p.opsChan <- op }()
	}
}

func (p *printer) processOps() {
	time.Sleep(250 * time.Millisecond) // test sleep to ensure queue func hits default case
	for op := range p.opsChan {
		op()
		p.Done()
	}
}

// END PRINTER OMIT

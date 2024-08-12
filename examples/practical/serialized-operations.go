package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	p := &printer{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			p.print(fmt.Sprintf("Hello, World! %d", i))
		}()
	}
	wg.Wait()
}

type printer struct {
	sync.Mutex
}

func (p *printer) print(s string) {
	p.Lock()
	defer p.Unlock()
	fmt.Println(s)
}

// END OMIT

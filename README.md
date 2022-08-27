# Go Observer

Implementation of the observer pattern using generics.

## Installation

    go get github.com/ltunc/go-observer@latest

## Usage

```go
package main

import (
	"fmt"
	"github.com/ltunc/go-observer/observer"
)

func main() {
	es := &observer.Subject[Event]{}
	p := &Printer{}
	es.Subscribe(p)
	f := &Fancy{}
	es.Subscribe(f)
	for i := 1; i < 4; i++ {
		e := Event{name: fmt.Sprintf("e%d", i)}
		es.Fire(e)
	}
}

type Event struct {
	name string
}

func (e Event) Name() string {
	return e.name
}

type Printer struct {
}

func (p *Printer) Notify(ev Event) {
	fmt.Println(ev.Name())
}

type Fancy struct {
}

func (r *Fancy) Notify(ev Event) {
	fmt.Printf("Event '%s', hurray!\n", ev.Name())
}

```

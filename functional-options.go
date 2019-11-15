package main

import (
	"fmt"
)

// Funtional options
// See https://commandcenter.blogspot.nl/2014/01/self-referential-functions-and-design.html for 'undo' variant.
// https://play.golang.org/p/cmGxtlgyAn7

// Option function.
type option func(*Foo)

// Foo is the object to manipulate.
type Foo struct {
    header    string
    footer    string
    verbosity int
}

// Verbosity sets Foo's verbosity level to v.
func Verbosity(v int) option {
    return func(f *Foo) {
        f.verbosity = v
    }
}

// Option sets the options specified.
func (f *Foo) Option(opts ...option) {
    for _, opt := range opts {
        opt(f)
    }
}

func main() {
    f := &Foo{}

    f.Option(
        Verbosity(3),
    )

    fmt.Printf("%+#v\n", f)
}
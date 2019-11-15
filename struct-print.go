package main

/* Print and compare structs while debugging.

Printing: https://github.com/davecgh/go-spew
Diff, Equal: https://github.com/google/go-cmp
 */

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
)

func main() {
	var x = struct {
		Name string
	}{
		Name: "jacobus",
	}
	spew.Dump(x)

	var y = struct {
		Name string
	}{
		Name: "anna",
	}
	spew.Dump(y)

	fmt.Println("Diff structs:\n", cmp.Diff(x, y))
}

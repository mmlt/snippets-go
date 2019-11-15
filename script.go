//usr/bin/env go run "$0" "$@" ;  exit "$?"
package main

/*
Execute a .go file like a script (Linux).

Option 1:
Add the following as a first line to the go file
	//usr/bin/env go run "$0" "$@" ;  exit "$?"

Prereq: go tools in path
Con: non zero exit codes are printed (not returned)
	 executed by shell

Option 2:
As 1 but instead of 'go 'run' use gorun.
Add first line:
	#!/usr/bin/env gorun

Prereq: go get github.com/erning/gorun and make gorun available (in paht or /usr/local/bin)
Pre: exit code is properly returned (and more, see repo)
	 #! is interpreted by the kernel, works without a shell
Con: first line is not valid go

Option 2a:
As 2 but with first line:
	//usr/bin/env go run "$0" "$@" ;  exit "$?"
Pre: first line is valid go

Option 4:
Register GO as binary format
	echo ':GO:E::go::/usr/local/bin/gorun:' | sudo tee /proc/sys/fs/binfmt_misc/register
(see https://www.kernel.org/doc/html/latest/admin-guide/binfmt-misc.html)
*/

import (
	"fmt"
	"os"
)


func main() {
	fmt.Println("Hello", os.Args[1])
	os.Exit(45)
}


// Related:
// Mix bash and GO: https://github.com/progrium/go-basher
// Shelling out from GO: https://github.com/progrium/go-shell


package main

// Show how to use 'goto' to dedup error condition code.

import (
	"fmt"
	"strings"
)

// Goal: Replace / with tab and // with newline.
// See:  Structured programming with go to statements page 271
// https://play.golang.org/p/ht2rkS2wrgF

func main() {
	var err error
	var ch rune
	
	r := strings.NewReader("test /1//test///")
	
	for {
		ch,_,err = r.ReadRune()
		if err != nil {
			goto error
		}

		if ch == '/' {
			ch,_,err = r.ReadRune()
			if err != nil {
				goto error
			}
			
			if ch == '/' {
				fmt.Print("\n")
				continue
			}
			
			fmt.Print("\t")
		}
		fmt.Print(string(ch))
	}

	return

error:
	fmt.Println("error or eof", err)
}
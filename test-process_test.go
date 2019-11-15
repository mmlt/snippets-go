package main

import (
	"fmt"
	"os"
	"testing"
	"os/exec"
)

// https://talks.golang.org/2014/testing.slide#23
// https://www.youtube.com/watch?v=ndmB0bj7eyw

// Crasher is the function to test in its own process.
func Crasher() {
	fmt.Println("Going down in flames!")
	os.Exit(1)
}

// TestCrasher invokes the test binary itself as a subprocess.
func TestCrasher(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		Crasher()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCrasher")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}


package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	cmd := "ls"
	args := []string{"-ll"}
	out, err := exec.Command(cmd, args...).Output()
	spew.Dump(string(out), err)
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Successfully halved image in size")
}

package goroutine

import (
	"fmt"
	"strings"
	"testing"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		fmt.Println("filter ", item)

		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

func printGopher(upstream chan string) {
	for item := range upstream {
		fmt.Println("print ", item)
	}
}

// Pipeline source->filter->print
func TestPipeline(t *testing.T) {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}

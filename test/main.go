package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		// Note if you set a break point on the line above you can't reach it using the debugger
		// but if you just run it, then the code is executed.
		fmt.Println("we got a line")
		counts[line]++
		if len(line) == 0 {
			break
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
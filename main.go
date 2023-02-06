package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	queenNum := 15
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "single":
			start := time.Now()
			backtrack([]queen{}, queenNum, 1)
			elapsed := time.Since(start)
			fmt.Println(elapsed.Microseconds())
		case "multi":
			start := time.Now()
			backtrack([]queen{}, queenNum, runtime.NumCPU())
			elapsed := time.Since(start)
			fmt.Println(elapsed.Microseconds())
		}
	} else {
		fmt.Printf("Usage: %v single|multi\n", os.Args[0])
		os.Exit(1)
	}
}

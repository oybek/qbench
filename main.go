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
		case "single", "multi":
			parn := 1
			if os.Args[1] == "multi" {
				parn = runtime.NumCPU()
			}
			start := time.Now()
			backtrack([]queen{}, queenNum, parn)
			elapsed := time.Since(start)
			fmt.Println(elapsed.Microseconds())
			os.Exit(0)
		}
	}
	fmt.Printf("Usage: %v single|multi\n", os.Args[0])
	os.Exit(1)
}

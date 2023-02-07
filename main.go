package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
			modelName, err := exec.Command("bash", "-c", "lscpu | grep 'Model name'").Output()
			info := strings.Split(string(modelName), ":")
			if err != nil || len(info) != 2 {
				fmt.Printf("Can't detect CPU model name %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("|%v|%v|\n", elapsed.Microseconds(), strings.Trim(info[1], " \n"))
			os.Exit(0)
		}
	}
	fmt.Printf("Usage: %v single|multi\n", os.Args[0])
	os.Exit(1)
}

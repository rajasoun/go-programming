package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("      GOOS : ", runtime.GOOS)
	fmt.Println("    GOARCH : ", runtime.GOARCH)
	fmt.Println(" CPU Cores : ", runtime.NumCPU())
	fmt.Println("Go Vesrion : ", runtime.Version())
}

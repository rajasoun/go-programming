package main

import (
	"fmt"
	"os"
	"runtime"
)

func printRuntimeProfile() {
	fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))
}

func printCurrentWorkingDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  cwd = %s\n", cwd)
}

func printCmdLineArgs() {
	fmt.Printf("  os.Args = %#v\n", os.Args)
}

func printRuntimeDetails() {
	fmt.Println("  GOARCH = ", runtime.GOARCH)
	fmt.Println("  GOOS = ", runtime.GOOS)
	fmt.Println("  CPU Cores = ", runtime.NumCPU())
	fmt.Println("  Go Vesrion = ", runtime.Version())
}

func main() {
	printRuntimeProfile()
	printCurrentWorkingDir()
	printCmdLineArgs()
	printRuntimeDetails()
}

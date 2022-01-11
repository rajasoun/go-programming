package main

import (
	"fmt"
	"os"
)

func main() {
	printRuntimeProfile()
	printCurrentWorkingDir()
	printCmdLineArgs()
	printEnvDetails()
}

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

func printEnvDetails() {
	for _, ev := range []string{"GOARCH", "GOOS", "GOFILE", "GOLINE", "GOPACKAGE", "DOLLAR"} {
		fmt.Println("  ", ev, "=", os.Getenv(ev))
	}
}

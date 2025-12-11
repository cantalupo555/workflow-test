package main

import (
	"fmt"
	"runtime"
)

var (
	appName    = "workflow-test"
	appVersion = "dev"
)

func main() {
	fmt.Printf("%s v%s\n", appName, appVersion)
	fmt.Printf("OS: %s, Arch: %s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Go Version: %s\n", runtime.Version())
}

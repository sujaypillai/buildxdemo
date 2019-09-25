package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	hostname, _ := os.Hostname()
	fmt.Printf("Hello gitaction, I'm %s running on %s/%s\n", hostname, runtime.GOOS, runtime.GOARCH)
}

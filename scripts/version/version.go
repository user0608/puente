package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	version, _ := os.ReadFile("./version")
	fmt.Println("Version: "+string(version), time.Now().UTC().Format(time.DateTime), "UTC")
}

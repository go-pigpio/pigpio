package main

import (
	"fmt"

	"github.com/wtsi-hgi/pigpio-go/pigpio"
)

func main() {
	version := pigpio.Version()
	fmt.Printf("pigpio version %d\n", version)
}

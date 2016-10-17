package main

import (
	"fmt"

	"github.com/wtsi-hgi/pigpio-go/pigpio"
)

func main() {
	pigpioVersion := pigpio.Version()
	fmt.Printf("pigpio version %d\n", pigpioVersion)

	hwRevision := pigpio.HardwareRevision()
	fmt.Printf("hardware revision %x\n", hwRevision)
}

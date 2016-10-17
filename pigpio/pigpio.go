/*
Package pigpio provides bindings to the pigpio C API.
*/
package pigpio

/*
#cgo CFLAGS: -pthread -W -Wall -Wno-unused-parameter -Wno-format-extra-args -Wbad-function-cast -O2 -g
#cgo LDFLAGS: -lpigpio -lrt
#include <pigpio.h>
*/
import "C"

// Version returns the version of the pigpio C library as a uint
func Version() (version uint) {
	version = uint(C.gpioVersion())
	return
}

// HardwareRevision returns the hardware revision of the Raspberry Pi as a uint
func HardwareRevision() (version uint) {
	version = uint(C.gpioHardwareRevision())
	return
}

/*******************************************************************************
 * Copyright (c) 2016 Genome Research Ltd.
 *
 * Author: Joshua C. Randall <jcrandall@alum.mit.edu>
 *
 * This file is part of pigpio-go.
 *
 * pigpio-go is free software: you can redistribute it and/or modify it under
 * the terms of the GNU Affero General Public License as published by the Free
 * Software Foundation; either version 3 of the License, or (at your option) any
 * later version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
 * FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more
 * details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 ******************************************************************************/

/*
Package pigpio provides bindings to the pigpio C library.

pigpio is a C library for the Raspberry which allows control of the GPIO.
*/
package pigpio

/*
#cgo CFLAGS: -pthread -W -Wall -Wno-unused-parameter -Wno-format-extra-args -Wbad-function-cast -Wno-unused-variable -O2 -g
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

// Initialise returns the pigpio version number if OK, otherwise sets err
func Initialise() (version uint, err error) {
	versErr := int(C.gpioInitialise())
	if versErr >= 0 {
		version = uint(versErr)
	} else {
		err = Errno(versErr)
	}
	return
}

// Terminates the library (call before program exit)
func Terminate() {
	C.gpioTerminate()
}

// SetPullUpDown sets the GPIO pull-up/pull-down resistor mode
func SetPullUpDown(gpio uint, pud uint) (err error) {
	cErr := int(C.gpioSetPullUpDown(C.unsigned(gpio), C.unsigned(pud)))
	if cErr != 0 {
		err = Errno(cErr)
	}
	return
}

// Read reads the GPIO level (on or off)
func Read(gpio uint) (level uint, err error) {
	levelErr := int(C.gpioRead(C.unsigned(gpio)))
	if levelErr >= 0 {
		level = uint(levelErr)
	} else {
		err = Errno(levelErr)
	}
	return
}

// Write sets the GPIO level (on or off)
func Write(gpio uint, level uint) (err error) {
	cErr := int(C.gpioWrite(C.unsigned(gpio), C.unsigned(level)))
	if cErr != 0 {
		err = Errno(cErr)
	}
	return
}

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

Documentation for the C library should be referenced in conjunction with
godocs below: http://abyz.co.uk/rpi/pigpio/cif.html

*/
package pigpio

/*
#cgo CFLAGS: -pthread -W -Wall -Wno-unused-parameter -Wno-format-extra-args -Wbad-function-cast -Wno-unused-variable -O2 -g
#cgo LDFLAGS: -lpigpio -lrt
#include <pigpio.h>
*/
import "C"
import "time"

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

// PWM starts PWM on the specified GPIO, with a dutycycle between
// 0 (off) and PWMrange (use SetPWMrange to set the range, which
// defaults to 255)
func PWM(gpio uint, dutycycle uint) (err error) {
	cErr := int(C.gpioPWM(C.unsigned(gpio), C.unsigned(dutycycle)))
	if cErr != 0 {
		err = Errno(cErr)
	}
	return
}

// GetPWMdutycycle returns the PWM dutycycle setting for the specified GPIO
func GetPWMdutycycle(gpio uint) (dutycycle uint, err error) {
	dutycycleErr := int(C.gpioGetPWMdutycycle(C.unsigned(gpio)))
	if dutycycleErr >= 0 {
		dutycycle = uint(dutycycleErr)
	} else {
		err = Errno(dutycycleErr)
	}
	return
}

// Servo starts servo pulses on the specified GPIO.AlertFunc
// pulsewidth values: 0 (off), 500 (most anti-clockwise) - 2500 (most clockwise)
// A value of 1500 should always be safe and represents the mid-point of rotation.
// WARNING: you can DAMAGE a servo if you command it to move beyond its limits.
func Servo(gpio uint, pulsewidth uint) (err error) {
	cErr := int(C.gpioServo(C.unsigned(gpio), C.unsigned(pulsewidth)))
	if cErr != 0 {
		err = Errno(cErr)
	}
	return
}

// GetServoPulsewidth returns the servo pulsewidth setting for the GPIO
func GetServoPulsewidth(gpio uint) (pulsewidth uint, err error) {
	pulsewidthErr := int(C.gpioGetServoPulsewidth(C.unsigned(gpio)))
	if pulsewidthErr >= 0 {
		pulsewidth = uint(pulsewidthErr)
	} else {
		err = Errno(pulsewidthErr)
	}
	return
}

// Delay delays for at least the specified duration
// A delay of 100 microseconds or less will use a busy wait.
// If the requested delay duration is longer than C.UINT32_MAX in Microseconds
// (which is ~4295s or a little under 72m), the underlying gpioDelay function
// will be called multiple times to make up the requested duration.
func Delay(delay time.Duration) (actualDelay time.Duration) {
	delayMicros := uint(delay / time.Microsecond)
	actualDelayMicros := uint(0)
	for delayMicros > uint(C.UINT32_MAX) {
		actualDelayMicros += uint(C.gpioDelay(C.UINT32_MAX))
		delayMicros -= uint(C.UINT32_MAX)
		if delayMicros < 100 {
			// do not busy-wait the last 100 microseconds of a long delay
			delayMicros = 0
		}
	}
	if delayMicros > 0 {
		actualDelayMicros += uint(C.gpioDelay(C.uint32_t(delayMicros)))
	}
	actualDelay = time.Duration(actualDelayMicros) * time.Microsecond
	return
}

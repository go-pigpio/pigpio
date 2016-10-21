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
#include <pigpio.h>
*/
import "C"
import "fmt"

type Mode int

// mode
const (
	Input  Mode = C.PI_INPUT
	Output Mode = C.PI_OUTPUT
	Alt0   Mode = C.PI_ALT0
	Alt1   Mode = C.PI_ALT1
	Alt2   Mode = C.PI_ALT2
	Alt3   Mode = C.PI_ALT3
	Alt4   Mode = C.PI_ALT4
	Alt5   Mode = C.PI_ALT5
)

func (mode Mode) String() (s string) {
	modeStr := ""
	switch mode {
	case Input:
		modeStr = "Input"
	case Output:
		modeStr = "Output"
	case Alt0:
		modeStr = "Alt0"
	case Alt1:
		modeStr = "Alt1"
	case Alt2:
		modeStr = "Alt2"
	case Alt3:
		modeStr = "Alt3"
	case Alt4:
		modeStr = "Alt4"
	case Alt5:
		modeStr = "Alt5"
	}
	s = fmt.Sprintf("Mode(%s)", modeStr)
	return
}

// SetMode sets the GPIO mode
func SetMode(gpio uint, mode Mode) (err error) {
	cErr := int(C.gpioSetMode(C.unsigned(gpio), C.unsigned(mode)))
	if cErr != 0 {
		err = Errno(cErr)
	}
	return
}

// GetMode gets the GPIO mode
func GetMode(gpio uint) (mode Mode, err error) {
	modeErr := int(C.gpioGetMode(C.unsigned(gpio)))
	if modeErr >= 0 {
		mode = Mode(modeErr)
	} else {
		err = Errno(modeErr)
	}
	return
}

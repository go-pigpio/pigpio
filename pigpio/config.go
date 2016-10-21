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
import "time"

// CfgBufferSize configures pigpio to buffer duration worth of GPIO samples
func CfgBufferSize(duration time.Duration) (err error) {
	cfgMillis := uint(duration / time.Millisecond)
	cErr := int(C.gpioCfgBufferSize(C.unsigned(cfgMillis)))
	if cErr != 0 {
		err = Errno(cErr)
	}
	return
}

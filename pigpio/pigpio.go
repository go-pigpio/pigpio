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

extern int goSetAlertFunc(unsigned userGpio, int cbi);
extern int goSetTimerFunc(unsigned timer, unsigned millis, int cbi);
*/
import "C"
import (
	"fmt"
	"time"
)

/*******************************************************************************
* Essential
*******************************************************************************/

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

/*******************************************************************************
* Beginner
*******************************************************************************/

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

// SetPullUpDown sets the GPIO pull-up/pull-down resistor mode
func SetPullUpDown(gpio uint, pud uint) (err error) {
	cErr := int(C.gpioSetPullUpDown(C.unsigned(gpio), C.unsigned(pud)))
	if cErr != 0 {
		err = Errno(cErr)
	}
	return
}

// Read reads the GPIO level (on or off)
func Read(gpio uint) (level bool, err error) {
	levelErr := int(C.gpioRead(C.unsigned(gpio)))
	if levelErr >= 0 {
		level = levelErr != 0
	} else {
		err = Errno(levelErr)
	}
	return
}

// Write sets the GPIO level (on or off)
func Write(gpio uint, level bool) (err error) {
	levelUint := uint(0)
	if level {
		levelUint = 1
	}
	cErr := int(C.gpioWrite(C.unsigned(gpio), C.unsigned(levelUint)))
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

// SetAlertFunc registers a callback function to be called when the
// specified GPIO changes state.
func SetAlertFunc(gpio uint, alertFunc AlertFunc) (err error) {
	cbi := registerAlertFunc(alertFunc)
	cErr := C.goSetAlertFunc(C.unsigned(gpio), C.int(cbi))
	if cErr != 0 {
		err = fmt.Errorf("SetAlertFunc Error: %d", cErr)
	}
	return
}

// SetTimerFunc registers a callback function to be called periodically
// period can vary from 10ms to 60s - durations outside that range will
// result in a BadMs error
func SetTimerFunc(timer uint, period time.Duration, timerFunc TimerFunc) (err error) {
	cbi := registerTimerFunc(timerFunc)
	cErr := C.goSetTimerFunc(C.unsigned(timer), C.unsigned(period/time.Millisecond), C.int(cbi))
	if cErr != 0 {
		err = fmt.Errorf("SetTimerFunc Error: %d", cErr)
	}
	return
}

/*******************************************************************************
* Intermediate
*******************************************************************************/

// Trigger sends a trigger pulse to the specified GPIO
// GPIO is set to level for duration and then reset to not level
func Trigger(gpio uint, pulseLen time.Duration, level bool) (err error) {
	levelUint := uint(0)
	if level {
		levelUint = 1
	}
	cErr := int(C.gpioTrigger(C.unsigned(gpio), C.unsigned(pulseLen/time.Microsecond), C.unsigned(levelUint)))
	if cErr != 0 {
		err = Errno(cErr)
	}
	return
}

// SetWatchdog sets a watchdog on a GPIO
// Duration of timeout can range from 0-60s
// Set timeout to 0 to disable watchdog
func SetWatchdog(gpio uint, timeout time.Duration) (err error) {
	cErr := int(C.gpioSetWatchdog(C.unsigned(gpio), C.unsigned(timeout/time.Millisecond)))
	if cErr != 0 {
		err = Errno(cErr)
	}
	return
}

/*
gpioSetPWMrangeConfigure PWM range for a GPIO
gpioGetPWMrangeGet configured PWM range for a GPIO
gpioSetPWMfrequencyConfigure PWM frequency for a GPIO
gpioGetPWMfrequencyGet configured PWM frequency for a GPIO
gpioRead_Bits_0_31Read all GPIO in bank 1
gpioRead_Bits_32_53Read all GPIO in bank 2
gpioWrite_Bits_0_31_ClearClear selected GPIO in bank 1
gpioWrite_Bits_32_53_ClearClear selected GPIO in bank 2
gpioWrite_Bits_0_31_SetSet selected GPIO in bank 1
gpioWrite_Bits_32_53_SetSet selected GPIO in bank 2
gpioStartThreadStart a new thread
gpioStopThreadStop a previously started thread
*/

/*******************************************************************************
* Advance
*******************************************************************************/

/*
gpioGetPWMrealRangeGet underlying PWM range for a GPIO
gpioSetAlertFuncExRequest a GPIO change callback, extended
gpioSetISRFuncRequest a GPIO interrupt callback
gpioSetISRFuncExRequest a GPIO interrupt callback, extended
gpioSetSignalFuncRequest a signal callback
gpioSetSignalFuncExRequest a signal callback, extended
gpioSetGetSamplesFuncRequests a GPIO samples callback
gpioSetGetSamplesFuncExRequests a GPIO samples callback, extended
gpioSetTimerFuncExRequest a regular timed callback, extended
gpioNotifyOpenRequest a notification handle
gpioNotifyOpenWithSizeRequest a notification handle with sized pipe
gpioNotifyBeginStart notifications for selected GPIO
gpioNotifyPausePause notifications
gpioNotifyCloseClose a notification
gpioSerialReadOpenOpens a GPIO for bit bang serial reads
gpioSerialReadInvertConfigures normal/inverted for serial reads
gpioSerialReadReads bit bang serial data from a GPIO
gpioSerialReadCloseCloses a GPIO for bit bang serial reads
gpioHardwareClockStart hardware clock on supported GPIO
gpioHardwarePWMStart hardware PWM on supported GPIO
gpioGlitchFilterSet a glitch filter on a GPIO
gpioNoiseFilterSet a noise filter on a GPIO
gpioGetPadGets a pads drive strength
gpioSetPadSets a pads drive strength
shellExecutes a shell command
*/

/*******************************************************************************
* Scripts
*******************************************************************************/

/*
gpioStoreScriptStore a script
gpioRunScriptRun a stored script
gpioScriptStatusGet script status and parameters
gpioStopScriptStop a running script
gpioDeleteScriptDelete a stored script
*/

/*******************************************************************************
* Waves
*******************************************************************************/

/*
gpioWaveClearDeletes all waveforms
gpioWaveAddNewStarts a new waveform
gpioWaveAddGenericAdds a series of pulses to the waveform
gpioWaveAddSerialAdds serial data to the waveform
gpioWaveCreateCreates a waveform from added data
gpioWaveDeleteDeletes a waveform
gpioWaveTxSendTransmits a waveform
gpioWaveChainTransmits a chain of waveforms
gpioWaveTxAtReturns the current transmitting waveform
gpioWaveTxBusyChecks to see if the waveform has ended
gpioWaveTxStopAborts the current waveform
gpioWaveGetMicrosLength in microseconds of the current waveform
gpioWaveGetHighMicrosLength of longest waveform so far
gpioWaveGetMaxMicrosAbsolute maximum allowed micros
gpioWaveGetPulsesLength in pulses of the current waveform
gpioWaveGetHighPulsesLength of longest waveform so far
gpioWaveGetMaxPulsesAbsolute maximum allowed pulses
gpioWaveGetCbsLength in control blocks of the current waveform
gpioWaveGetHighCbsLength of longest waveform so far
gpioWaveGetMaxCbsAbsolute maximum allowed control blocks
*/

/*******************************************************************************
* I2C
*******************************************************************************/

/*
i2cOpenOpens an I2C device
i2cCloseCloses an I2C device
i2cWriteQuickSMBus write quick
i2cWriteByteSMBus write byte
i2cReadByteSMBus read byte
i2cWriteByteDataSMBus write byte data
i2cWriteWordDataSMBus write word data
i2cReadByteDataSMBus read byte data
i2cReadWordDataSMBus read word data
i2cProcessCallSMBus process call
i2cWriteBlockDataSMBus write block data
i2cReadBlockDataSMBus read block data
i2cBlockProcessCallSMBus block process call
i2cWriteI2CBlockDataSMBus write I2C block data
i2cReadI2CBlockDataSMBus read I2C block data
i2cReadDeviceReads the raw I2C device
i2cWriteDeviceWrites the raw I2C device
i2cSwitchCombinedSets or clears the combined flag
i2cSegmentsPerforms multiple I2C transactions
i2cZipPerforms multiple I2C transactions
bbI2COpenOpens GPIO for bit banging I2C
bbI2CCloseCloses GPIO for bit banging I2C
bbI2CZipPerforms multiple bit banged I2C transactions
*/

/*******************************************************************************
* SPI
*******************************************************************************/

/*
spiOpenOpens a SPI device
spiCloseCloses a SPI device
spiReadReads bytes from a SPI device
spiWriteWrites bytes to a SPI device
spiXferTransfers bytes with a SPI device
bbSPIOpenOpens GPIO for bit banging SPI
bbSPICloseCloses GPIO for bit banging SPI
bbSPIXferPerforms multiple bit banged SPI transactions
*/

/*******************************************************************************
* Serial
*******************************************************************************/

/*
serOpenOpens a serial device
serCloseCloses a serial device
serReadByteReads a byte from a serial device
serWriteByteWrites a byte to a serial device
serReadReads bytes from a serial device
serWriteWrites bytes to a serial device
serDataAvailableReturns number of bytes ready to be read
*/

/*******************************************************************************
* Files
*******************************************************************************/

/*
fileOpenOpens a file
fileCloseCloses a file
fileReadReads bytes from a file
fileWriteWrites bytes to a file
fileSeekSeeks to a position within a file
fileListList files which match a pattern
*/

/*******************************************************************************
* Configuration
*******************************************************************************/

/*
gpioCfgBufferSizeConfigure the GPIO sample buffer size
gpioCfgClockConfigure the GPIO sample rate
gpioCfgDMAchannelConfigure the DMA channel (DEPRECATED)
gpioCfgDMAchannelsConfigure the DMA channels
gpioCfgPermissionsConfigure the GPIO access permissions
gpioCfgInterfacesConfigure user interfaces
gpioCfgSocketPortConfigure socket port
gpioCfgMemAllocConfigure DMA memory allocation mode
gpioCfgNetAddrConfigure allowed network addresses
gpioCfgInternalsConfigure miscellaneous internals (DEPRECATED)
gpioCfgGetInternalsGet internal configuration settings
gpioCfgSetInternalsSet internal configuration settings
*/

/*******************************************************************************
* Custom
*******************************************************************************/

/*
gpioCustom1User custom function 1
gpioCustom2User custom function 2
*/

/*******************************************************************************
* Utilities
*******************************************************************************/

// Tick returns the current system tick
func Tick() (tick uint) {
	tick = uint(C.gpioTick())
	return
}

// HardwareRevision returns the hardware revision of the Raspberry Pi as a uint
func HardwareRevision() (version uint) {
	version = uint(C.gpioHardwareRevision())
	return
}

// Version returns the version of the pigpio C library as a uint
func Version() (version uint) {
	version = uint(C.gpioVersion())
	return
}

/*
getBitInBytesGet the value of a bit
putBitInBytesSet the value of a bit
gpioTimeGet current time
gpioSleepSleep for specified time
time_sleepSleeps for a float number of seconds
time_timeFloat number of seconds since the epoch
*/

/*******************************************************************************
* Expert
*******************************************************************************/
/*
rawWaveAddSPINot intended for general use
rawWaveAddGenericNot intended for general use
rawWaveCBNot intended for general use
rawWaveCBAdrNot intended for general use
rawWaveGetOOLNot intended for general use
rawWaveSetOOLNot intended for general use
rawWaveGetOutNot intended for general use
rawWaveSetOutNot intended for general use
rawWaveGetInNot intended for general use
rawWaveSetInNot intended for general use
rawWaveInfoNot intended for general use
rawDumpWaveNot intended for general use
rawDumpScriptNot intended for general useful
*/

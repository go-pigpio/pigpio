package pigpio

/*
#cgo CFLAGS: -std=gnu99
#include <stdint.h>
#include <pigpio.h>

extern int goSetAlertFunc(unsigned userGpio, int cbi);
extern int goSetTimerFunc(unsigned timer, unsigned millis, int cbi);
*/
import "C"

import (
	"fmt"
	"sync"
	"time"
)

type AlertFunc func(gpio int, level int, tick uint32)

var alertFuncMu sync.Mutex
var alertFuncs = make(map[int]AlertFunc)
var alertFuncIndex int

func lookupAlertFunc(cbIndex int) AlertFunc {
	alertFuncMu.Lock()
	defer alertFuncMu.Unlock()
	return alertFuncs[cbIndex]
}

func registerAlertFunc(alertFunc AlertFunc) int {
	alertFuncMu.Lock()
	defer alertFuncMu.Unlock()
	alertFuncIndex++
	for alertFuncs[alertFuncIndex] != nil {
		alertFuncIndex++
	}
	alertFuncs[alertFuncIndex] = alertFunc
	return alertFuncIndex
}

func SetAlertFunc(userGpio uint, alertFunc AlertFunc) (err error) {
	cbi := registerAlertFunc(alertFunc)
	cErr := C.goSetAlertFunc(C.unsigned(userGpio), C.int(cbi))
	if cErr != 0 {
		err = fmt.Errorf("SetAlertFunc Error: %d", cErr)
	}
	return
}

//export goAlertFunc
func goAlertFunc(cbIndex C.int, gpio C.int, level C.int, tick C.uint32_t) {
	fn := lookupAlertFunc(int(cbIndex))
	fn(int(gpio), int(level), uint32(tick))
}

type TimerFunc func()

var timerFuncMu sync.Mutex
var timerFuncs = make(map[int]TimerFunc)
var timerFuncIndex int

func lookupTimerFunc(cbIndex int) TimerFunc {
	timerFuncMu.Lock()
	defer timerFuncMu.Unlock()
	return timerFuncs[cbIndex]
}

func registerTimerFunc(timerFunc TimerFunc) int {
	timerFuncMu.Lock()
	defer timerFuncMu.Unlock()
	timerFuncIndex++
	for timerFuncs[timerFuncIndex] != nil {
		timerFuncIndex++
	}
	timerFuncs[timerFuncIndex] = timerFunc
	return timerFuncIndex
}

// SetTimerFunc registers a callback function to be called periodically
// period can vary from 10ms to 60s - durations outside that range will result in a BadMs error
func SetTimerFunc(timer uint, period time.Duration, timerFunc TimerFunc) (err error) {
	cbi := registerTimerFunc(timerFunc)
	cErr := C.goSetTimerFunc(C.unsigned(timer), C.unsigned(period/time.Millisecond), C.int(cbi))
	if cErr != 0 {
		err = fmt.Errorf("SetTimerFunc Error: %d", cErr)
	}
	return
}

//export goTimerFunc
func goTimerFunc(cbIndex C.int) {
	fn := lookupTimerFunc(int(cbIndex))
	fn()
}

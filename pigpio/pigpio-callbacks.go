package pigpio

/*
#cgo CFLAGS: -std=gnu99
#include <stdint.h>
#include <pigpio.h>
*/
import "C"

import "sync"

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
	if alertFunc == nil {
		return -1
	}
	alertFuncMu.Lock()
	defer alertFuncMu.Unlock()
	alertFuncIndex++
	for alertFuncs[alertFuncIndex] != nil {
		alertFuncIndex++
	}
	alertFuncs[alertFuncIndex] = alertFunc
	return alertFuncIndex
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
	if timerFunc == nil {
		return -1
	}
	timerFuncMu.Lock()
	defer timerFuncMu.Unlock()
	timerFuncIndex++
	for timerFuncs[timerFuncIndex] != nil {
		timerFuncIndex++
	}
	timerFuncs[timerFuncIndex] = timerFunc
	return timerFuncIndex
}

//export goTimerFunc
func goTimerFunc(cbIndex C.int) {
	fn := lookupTimerFunc(int(cbIndex))
	fn()
}

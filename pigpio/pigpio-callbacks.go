package pigpio

/*
#cgo CFLAGS: -std=gnu99
#include <stdint.h>
#include <pigpio.h>

extern int goSetAlertFunc(unsigned userGpio, int cbi);
*/
import "C"

import (
	"fmt"
	"sync"
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

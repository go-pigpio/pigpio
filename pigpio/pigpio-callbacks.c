#include <stdlib.h>
#include <pigpio.h>
#include "_cgo_export.h"

void goAlertFunc_cgo(int gpio, int level, uint32_t tick, void *userdata);
int goSetAlertFunc(unsigned userGpio, int cbi);

typedef struct {
    int cbi;
} goAlertFunc_userdata;

void goAlertFunc_cgo(int gpio, int level, uint32_t tick, void *userdata) {
    goAlertFunc_userdata *myUserdata = (goAlertFunc_userdata*)userdata;
    goAlertFunc(myUserdata->cbi, gpio, level, tick);
}

int goSetAlertFunc(unsigned userGpio, int cbi) {
    goAlertFunc_userdata *myUserdata;
    myUserdata = malloc(sizeof(goAlertFunc_userdata));
    myUserdata->cbi = cbi;
    return gpioSetAlertFuncEx(userGpio, goAlertFunc_cgo, myUserdata);
}

void goTimerFunc_cgo(void *userdata);
int goSetTimerFunc(unsigned timer, unsigned millis, int cbi);

typedef struct {
    int cbi;
} goTimerFunc_userdata;

void goTimerFunc_cgo(void *userdata) {
    goTimerFunc_userdata *myUserdata = (goTimerFunc_userdata*)userdata;
    goTimerFunc(myUserdata->cbi);
}

int goSetTimerFunc(unsigned timer, unsigned millis, int cbi) {
    goTimerFunc_userdata *myUserdata;
    myUserdata = malloc(sizeof(goTimerFunc_userdata));
    myUserdata->cbi = cbi;
    return gpioSetTimerFuncEx(timer, millis, goTimerFunc_cgo, myUserdata);
}

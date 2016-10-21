package pigpio_test

import (
	"testing"
	"time"

	"github.com/wtsi-hgi/pigpio-go/pigpio"
)

func TestCallbacks(t *testing.T) {
	_, err := pigpio.Initialise()
	defer pigpio.Terminate()
	if err != nil {
		t.Error("failed to initialise: ", err)
	}
	err = pigpio.SetMode(7, pigpio.Input)
	if err != nil {
		t.Error("error setting mode: ", err)
	}
	err = pigpio.SetPullUpDown(7, pigpio.PudDown)
	if err != nil {
		t.Error("error setting gpio 7 pull down: ", err)
	}
	t.Run("TestAlertFunc", func(t *testing.T) {
		alertChan := make(chan int, 2)
		err = pigpio.SetAlertFunc(7, func(gpio int, level int, tick uint32) {
			alertChan <- level
		})
		if err != nil {
			t.Error("error setting alert func: ", err)
		}
		time.Sleep(200 * time.Millisecond)
		err = pigpio.SetPullUpDown(7, pigpio.PudUp)
		if err != nil {
			t.Error("error setting gpio 7 pull up: ", err)
		}
		time.Sleep(200 * time.Millisecond)
		err = pigpio.SetPullUpDown(7, pigpio.PudDown)
		if err != nil {
			t.Error("error setting gpio 7 pull up: ", err)
		}
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(1 * time.Second)
			timeout <- true
		}()
		select {
		case high := <-alertChan:
			if high != 1 {
				t.Error("first alert was not high:", high)
			}
		case <-timeout:
			t.Error("no alert on channel before timeout")
		}
		go func() {
			time.Sleep(1 * time.Second)
			timeout <- true
		}()
		select {
		case low := <-alertChan:
			if low != 0 {
				t.Error("first alert was not low:", low)
			}
		case <-timeout:
			t.Error("no alert on channel before timeout")
		}
	})
	t.Run("TestTimerFunc", func(t *testing.T) {
		timerChan := make(chan int, 2)
		err = pigpio.SetTimerFunc(0, 100*time.Millisecond, func() {
			timerChan <- 1
		})
		if err != nil {
			t.Error("error setting timer func: ", err)
		}
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(200 * time.Millisecond)
			timeout <- true
		}()
		select {
		case <-timerChan:
			// ok
		case <-timeout:
			t.Error("no timer on channel before timeout")
		}
		go func() {
			time.Sleep(200 * time.Millisecond)
			timeout <- true
		}()
		select {
		case <-timerChan:
			// ok
		case <-timeout:
			t.Error("no timer on channel before timeout")
		}
	})
}

package pigpio_test

import (
	"testing"
	"time"

	"github.com/wtsi-hgi/pigpio-go/pigpio"
)

func TestEssential(t *testing.T) {
	t.Run("Initialise", func(t *testing.T) {
		_, err := pigpio.Initialise()
		if err != nil {
			t.Error("failed to initialise: ", err)
		}
		t.Run("Terminate", func(t *testing.T) {
			pigpio.Terminate()
			t.Run("InitialiseAfterTerminate", func(t *testing.T) {
				_, err := pigpio.Initialise()
				defer pigpio.Terminate()
				if err != nil {
					t.Error("failed to initialise: ", err)
				}
			})
		})
	})

}

func TestBeginner(t *testing.T) {
	_, err := pigpio.Initialise()
	defer pigpio.Terminate()
	if err != nil {
		t.Error("failed to initialise: ", err)
	}
	t.Run("SetGetMode", func(t *testing.T) {
		t.Run("Input", func(t *testing.T) {
			err = pigpio.SetMode(7, pigpio.Input)
			if err != nil {
				t.Error("error setting mode: ", err)
			}
			mode, err := pigpio.GetMode(7)
			if err != nil {
				t.Error("error getting mode: ", err)
			}
			if mode != pigpio.Input {
				t.Error("mode was not set to input: ", mode)
			}
		})
		t.Run("Output", func(t *testing.T) {
			err = pigpio.SetMode(7, pigpio.Output)
			if err != nil {
				t.Error("error setting mode: ", err)
			}
			mode, err := pigpio.GetMode(7)
			if err != nil {
				t.Error("error getting mode: ", err)
			}
			if mode != pigpio.Output {
				t.Error("mode was not set to output: ", mode)
			}
		})
	})
	t.Run("BeginnerCallbacks", func(t *testing.T) {
		err = pigpio.SetMode(7, pigpio.Input)
		if err != nil {
			t.Error("error setting mode: ", err)
		}
		err = pigpio.SetPullUpDown(7, pigpio.PudDown)
		if err != nil {
			t.Error("error setting gpio 7 pull down: ", err)
		}
		t.Run("AlertFunc", func(t *testing.T) {
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
			firstTimeout := make(chan bool, 1)
			go func() {
				time.Sleep(1 * time.Second)
				firstTimeout <- true
			}()
			select {
			case high := <-alertChan:
				if high != 1 {
					t.Error("first alert was not high:", high)
				}
			case <-firstTimeout:
				t.Error("no high level alert on channel before firstTimeout")
			}
			secondTimeout := make(chan bool, 1)
			go func() {
				time.Sleep(1 * time.Second)
				secondTimeout <- true
			}()
			select {
			case low := <-alertChan:
				if low != 0 {
					t.Error("first alert was not low:", low)
				}
			case <-secondTimeout:
				t.Error("no low level alert on channel before secondTimeout")
			}
		})
		t.Run("TimerFunc", func(t *testing.T) {
			timerChan := make(chan int, 2)
			err = pigpio.SetTimerFunc(0, 100*time.Millisecond, func() {
				timerChan <- 1
			})
			if err != nil {
				t.Error("error setting timer func: ", err)
			}
			firstTimeout := make(chan bool, 1)
			go func() {
				time.Sleep(200 * time.Millisecond)
				firstTimeout <- true
			}()
			select {
			case <-timerChan:
				// ok
			case <-firstTimeout:
				t.Error("no timer on channel before first firstTimeout")
			}
			secondTimeout := make(chan bool, 1)
			go func() {
				time.Sleep(200 * time.Millisecond)
				secondTimeout <- true
			}()
			select {
			case <-timerChan:
				// ok
			case <-secondTimeout:
				t.Error("no timer on channel before second secondTimeout")
			}
		})
	})
}

func TestIntermediate(t *testing.T) {
	_, err := pigpio.Initialise()
	defer pigpio.Terminate()
	if err != nil {
		t.Error("failed to initialise: ", err)
	}
	t.Run("WatchdogAlert", func(t *testing.T) {
		alertChan := make(chan int, 1)
		err = pigpio.SetAlertFunc(7, func(gpio int, level int, tick uint32) {
			alertChan <- level
		})
		if err != nil {
			t.Error("error setting alert func: ", err)
		}
		err = pigpio.SetWatchdog(7, 200*time.Millisecond)
		if err != nil {
			t.Error("error setting watchdog: ", err)
		}
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(1 * time.Second)
			timeout <- true
		}()
		select {
		case level := <-alertChan:
			if level != pigpio.Timeout {
				t.Error("alert was not set to Timeout:", level)
			}
		case <-timeout:
			t.Error("no alert on channel before timeout")
		}
		err = pigpio.SetWatchdog(7, time.Duration(0))
		if err != nil {
			t.Error("error cancelling watchdog: ", err)
		}
	})
}

func TestUtility(t *testing.T) {
	// HardwareRevision and Version can be called before Initialise
	t.Run("HardwareRevision", func(t *testing.T) {
		revision := pigpio.HardwareRevision()
		if revision <= 0 {
			t.Error("invalid hardware revision number: ", revision)
		}
	})
	t.Run("Version", func(t *testing.T) {
		version := pigpio.Version()
		if version <= 0 {
			t.Error("invalid version number: ", version)
		}
	})
	version, err := pigpio.Initialise()
	defer pigpio.Terminate()
	if err != nil {
		t.Error("failed to initialise: ", err)
	}
	t.Run("InitialiseVersionMatch", func(t *testing.T) {
		if version != pigpio.Version() {
			t.Error("version mismatch on initialise: ", version)
		}
	})

}

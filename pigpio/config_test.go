package pigpio_test

import (
	"testing"
	"time"

	"github.com/wtsi-hgi/pigpio-go/pigpio"
)

func TestCfgAfterInitialise(t *testing.T) {
	_, err := pigpio.Initialise()
	defer pigpio.Terminate()
	if err != nil {
		t.Error("failed to initialise: ", err)
	}
	err = pigpio.CfgBufferSize(time.Duration(120) * time.Millisecond)
	if err != pigpio.Initialised {
		t.Error("call to CfgBufferSize after Initialise did not result in Inititialised error")
	}
}

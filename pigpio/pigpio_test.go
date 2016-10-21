package pigpio_test

import (
	"testing"

	"github.com/wtsi-hgi/pigpio-go/pigpio"
)

func TestVersion(t *testing.T) {
	version := pigpio.Version()
	if version <= 0 {
		t.Error("invalid version number: ", version)
	}
}

func TestHardwareRevision(t *testing.T) {
	revision := pigpio.HardwareRevision()
	if revision <= 0 {
		t.Error("invalid hardware revision number: ", revision)
	}
}

func TestInitialise(t *testing.T) {
	version, err := pigpio.Initialise()
	defer pigpio.Terminate()
	if err != nil {
		t.Error("failed to initialise: ", err)
	}
	if version != pigpio.Version() {
		t.Error("version mismatch on initialise: ", version)
	}
}

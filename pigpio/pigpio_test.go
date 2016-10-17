package pigpio

import "testing"

func TestVersion(t *testing.T) {
	version := Version()
	if version <= 0 {
		t.Error("invalid version number: ", version)
	}
}

func TestHardwareRevision(t *testing.T) {
	revision := HardwareRevision()
	if revision <= 0 {
		t.Error("invalid hardware revision number: ", revision)
	}
}

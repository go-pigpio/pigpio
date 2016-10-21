package pigpio

import "fmt"

type Errno int

func (err Errno) Error() string {
	s, ok := errorString(err)
	if !ok {
		if err >= PigifErr0 && err <= PigifErr99 {
			s = fmt.Sprintf("pigif error %d", err)
		} else if err >= CustomErr0 && err <= CustomErr999 {
			s = fmt.Sprintf("custom error %d", err)
		} else {
			s = fmt.Sprintf("error %d", err)
		}
	}
	return fmt.Sprintf("pigpio: %s", s)
}

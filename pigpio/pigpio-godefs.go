// +build ignore

package pigpio

// #include <pigpio.h>
import "C"

type Header C.gpioHeader_t

const Sizeof_Header = C.sizeof_gpioHeader_t

type Extent C.gpioExtent_t

const Sizeof_Extent = C.sizeof_gpioExtent_t

type Sample C.gpioSample_t

const Sizeof_Sample = C.sizeof_gpioSample_t

type Report C.gpioReport_t

const Sizeof_Report = C.sizeof_gpioReport_t

type Pulse C.gpioPulse_t

const Sizeof_Pulse = C.sizeof_gpioPulse_t

type RawWave C.rawWave_t

const Sizeof_RawWave = C.sizeof_rawWave_t

type RawWaveInfo C.rawWaveInfo_t

const Sizeof_RawWaveInfo = C.sizeof_rawWaveInfo_t

type RawSPI C.rawSPI_t

const Sizeof_RawSPI = C.sizeof_rawSPI_t

type RawCbs C.rawCbs_t

const Sizeof_RawCbs = C.sizeof_rawCbs_t

type I2CMsg C.pi_i2c_msg_t

const Sizeof_I2CMsg = C.sizeof_pi_i2c_msg_t

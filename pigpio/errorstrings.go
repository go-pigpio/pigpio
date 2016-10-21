package pigpio

func errorString(err Errno) (s string, ok bool) {
	ok = true
	switch err {
	case InitFailed:
		s = "gpioInitialise failed"
	case BadUserGpio:
		s = "GPIO not 0-31"
	case BadGpio:
		s = "GPIO not 0-53"
	case BadMode:
		s = "mode not 0-7"
	case BadLevel:
		s = "level not 0-1"
	case BadPud:
		s = "pud not 0-2"
	case BadPulsewidth:
		s = "pulsewidth not 0 or 500-2500"
	case BadDutycycle:
		s = "dutycycle outside set range"
	case BadTimer:
		s = "timer not 0-9"
	case BadMs:
		s = "ms not 10-60000"
	case BadTimetype:
		s = "timetype not 0-1"
	case BadSeconds:
		s = "seconds < 0"
	case BadMicros:
		s = "micros not 0-999999"
	case TimerFailed:
		s = "gpioSetTimerFunc failed"
	case BadWdogTimeout:
		s = "timeout not 0-60000"
	case BadClkPeriph:
		s = "clock peripheral not 0-1"
	case BadClkMicros:
		s = "clock micros not 1, 2, 4, 5, 8, or 10"
	case BadBufMillis:
		s = "buf millis not 100-10000"
	case BadDutyrange:
		s = "dutycycle range not 25-40000"
	case BadSignum:
		s = "signum not 0-63"
	case BadPathname:
		s = "can't open pathname"
	case NoHandle:
		s = "no handle available"
	case BadHandle:
		s = "unknown handle"
	case BadIfFlags:
		s = "ifFlags > 3"
	case BadChannel:
		s = "DMA channel not 0-14"
	case BadSocketPort:
		s = "socket port not 1024-32000"
	case BadFifoCommand:
		s = "unrecognized fifo command"
	case BadSecoChannel:
		s = "DMA secondary channel not 0-6"
	case NotInitialised:
		s = "function called before gpioInitialise"
	case Initialised:
		s = "function called after gpioInitialise"
	case BadWaveMode:
		s = "waveform mode not 0-3"
	case BadCfgInternal:
		s = "bad parameter in gpioCfgInternals call"
	case BadWaveBaud:
		s = "baud rate not 50-250K(RX)/50-1M(TX)"
	case TooManyPulses:
		s = "waveform has too many pulses"
	case TooManyChars:
		s = "waveform has too many chars"
	case NotSerialGpio:
		s = "no bit bang serial read on GPIO"
	case BadSerialStruc:
		s = "bad (null) serial structure parameter"
	case BadSerialBuf:
		s = "bad (null) serial buf parameter"
	case NotPermitted:
		s = "GPIO operation not permitted"
	case SomePermitted:
		s = "one or more GPIO not permitted"
	case BadWvscCommnd:
		s = "bad WVSC subcommand"
	case BadWvsmCommnd:
		s = "bad WVSM subcommand"
	case BadWvspCommnd:
		s = "bad WVSP subcommand"
	case BadPulselen:
		s = "trigger pulse length not 1-100"
	case BadScript:
		s = "invalid script"
	case BadScriptId:
		s = "unknown script id"
	case BadSerOffset:
		s = "add serial data offset > 30 minutes"
	case GpioInUse:
		s = "GPIO already in use"
	case BadSerialCount:
		s = "must read at least a byte at a time"
	case BadParamNum:
		s = "script parameter id not 0-9"
	case DupTag:
		s = "script has duplicate tag"
	case TooManyTags:
		s = "script has too many tags"
	case BadScriptCmd:
		s = "illegal script command"
	case BadVarNum:
		s = "script variable id not 0-149"
	case NoScriptRoom:
		s = "no more room for scripts"
	case NoMemory:
		s = "can't allocate temporary memory"
	case SockReadFailed:
		s = "socket read failed"
	case SockWritFailed:
		s = "socket write failed"
	case TooManyParam:
		s = "too many script parameters (> 10)"
	case ScriptNotReady:
		s = "script initialising"
	case BadTag:
		s = "script has unresolved tag"
	case BadMicsDelay:
		s = "bad MICS delay (too large)"
	case BadMilsDelay:
		s = "bad MILS delay (too large)"
	case BadWaveId:
		s = "non existent wave id"
	case TooManyCbs:
		s = "No more CBs for waveform"
	case TooManyOol:
		s = "No more OOL for waveform"
	case EmptyWaveform:
		s = "attempt to create an empty waveform"
	case NoWaveformId:
		s = "no more waveforms"
	case I2cOpenFailed:
		s = "can't open I2C device"
	case SerOpenFailed:
		s = "can't open serial device"
	case SpiOpenFailed:
		s = "can't open SPI device"
	case BadI2cBus:
		s = "bad I2C bus"
	case BadI2cAddr:
		s = "bad I2C address"
	case BadSpiChannel:
		s = "bad SPI channel"
	case BadFlags:
		s = "bad i2c/spi/ser open flags"
	case BadSpiSpeed:
		s = "bad SPI speed"
	case BadSerDevice:
		s = "bad serial device name"
	case BadSerSpeed:
		s = "bad serial baud rate"
	case BadParam:
		s = "bad i2c/spi/ser parameter"
	case I2cWriteFailed:
		s = "i2c write failed"
	case I2cReadFailed:
		s = "i2c read failed"
	case BadSpiCount:
		s = "bad SPI count"
	case SerWriteFailed:
		s = "ser write failed"
	case SerReadFailed:
		s = "ser read failed"
	case SerReadNoData:
		s = "ser read no data available"
	case UnknownCommand:
		s = "unknown command"
	case SpiXferFailed:
		s = "spi xfer/read/write failed"
	case BadPointer:
		s = "bad (NULL) pointer"
	case NoAuxSpi:
		s = "no auxiliary SPI on Pi A or B"
	case NotPwmGpio:
		s = "GPIO is not in use for PWM"
	case NotServoGpio:
		s = "GPIO is not in use for servo pulses"
	case NotHclkGpio:
		s = "GPIO has no hardware clock"
	case NotHpwmGpio:
		s = "GPIO has no hardware PWM"
	case BadHpwmFreq:
		s = "hardware PWM frequency not 1-125M"
	case BadHpwmDuty:
		s = "hardware PWM dutycycle not 0-1M"
	case BadHclkFreq:
		s = "hardware clock frequency not 4689-250M"
	case BadHclkPass:
		s = "need password to use hardware clock 1"
	case HpwmIllegal:
		s = "illegal, PWM in use for main clock"
	case BadDatabits:
		s = "serial data bits not 1-32"
	case BadStopbits:
		s = "serial (half) stop bits not 2-8"
	case MsgToobig:
		s = "socket/pipe message too big"
	case BadMallocMode:
		s = "bad memory allocation mode"
	case TooManySegs:
		s = "too many I2C transaction segments"
	case BadI2cSeg:
		s = "an I2C transaction segment failed"
	case BadSmbusCmd:
		s = "SMBus command not supported by driver"
	case NotI2cGpio:
		s = "no bit bang I2C in progress on GPIO"
	case BadI2cWlen:
		s = "bad I2C write length"
	case BadI2cRlen:
		s = "bad I2C read length"
	case BadI2cCmd:
		s = "bad I2C command"
	case BadI2cBaud:
		s = "bad I2C baud rate, not 50-500k"
	case ChainLoopCnt:
		s = "bad chain loop count"
	case BadChainLoop:
		s = "empty chain loop"
	case ChainCounter:
		s = "too many chain counters"
	case BadChainCmd:
		s = "bad chain command"
	case BadChainDelay:
		s = "bad chain delay micros"
	case ChainNesting:
		s = "chain counters nested too deeply"
	case ChainTooBig:
		s = "chain is too long"
	case Deprecated:
		s = "deprecated function removed"
	case BadSerInvert:
		s = "bit bang serial invert not 0 or 1"
	case BadEdge:
		s = "bad ISR edge value, not 0-2"
	case BadIsrInit:
		s = "bad ISR initialisation"
	case BadForever:
		s = "loop forever must be last command"
	case BadFilter:
		s = "bad filter parameter"
	case BadPad:
		s = "bad pad number"
	case BadStrength:
		s = "bad pad drive strength"
	case FilOpenFailed:
		s = "file open failed"
	case BadFileMode:
		s = "bad file mode"
	case BadFileFlag:
		s = "bad file flag"
	case BadFileRead:
		s = "bad file read"
	case BadFileWrite:
		s = "bad file write"
	case FileNotRopen:
		s = "file not open for read"
	case FileNotWopen:
		s = "file not open for write"
	case BadFileSeek:
		s = "bad file seek"
	case NoFileMatch:
		s = "no files match pattern"
	case NoFileAccess:
		s = "no permission to access file"
	case FileIsADir:
		s = "file is a directory"
	case BadShellStatus:
		s = "bad shell return status"
	case BadScriptName:
		s = "bad script name"
	case BadSpiBaud:
		s = "bad SPI baud rate, not 50-500k"
	case NotSpiGpio:
		s = "no bit bang SPI in progress on GPIO"
	default:
		ok = false
	}
	return
}

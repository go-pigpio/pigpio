package pigpio

/*
 #include <pigpio.h>
*/
import "C"

const Inpfifo = C.PI_INPFIFO
const Outfifo = C.PI_OUTFIFO
const Errfifo = C.PI_ERRFIFO
const Envport = C.PI_ENVPORT
const Envaddr = C.PI_ENVADDR
const Lockfile = C.PI_LOCKFILE
const I2cCombined = C.PI_I2C_COMBINED

const WaveFlagRead = C.WAVE_FLAG_READ
const WaveFlagTick = C.WAVE_FLAG_TICK

// gpio
const MinGpio = C.PI_MIN_GPIO
const MaxGpio = C.PI_MAX_GPIO
const MaxUserGpio = C.PI_MAX_USER_GPIO

// level
const Off = C.PI_OFF
const On = C.PI_ON
const Clear = C.PI_CLEAR
const Set = C.PI_SET
const Low = C.PI_LOW
const High = C.PI_HIGH
const Timeout = C.PI_TIMEOUT

// pud
const PudOff = C.PI_PUD_OFF
const PudDown = C.PI_PUD_DOWN
const PudUp = C.PI_PUD_UP

// dutycycle
const DefaultDutycycleRange = C.PI_DEFAULT_DUTYCYCLE_RANGE
const MinDutycycleRange = C.PI_MIN_DUTYCYCLE_RANGE
const MaxDutycycleRange = C.PI_MAX_DUTYCYCLE_RANGE

// pulsewidth
const ServoOff = C.PI_SERVO_OFF
const MinServoPulsewidth = C.PI_MIN_SERVO_PULSEWIDTH
const MaxServoPulsewidth = C.PI_MAX_SERVO_PULSEWIDTH

// hardware PWM
const HwPwmMinFreq = C.PI_HW_PWM_MIN_FREQ
const HwPwmMaxFreq = C.PI_HW_PWM_MAX_FREQ
const HwPwmRange = C.PI_HW_PWM_RANGE

// hardware clock
const HwClkMinFreq = C.PI_HW_CLK_MIN_FREQ
const HwClkMaxFreq = C.PI_HW_CLK_MAX_FREQ

const NotifySlots = C.PI_NOTIFY_SLOTS

const WaveBlocks = C.PI_WAVE_BLOCKS
const WaveMaxPulses = C.PI_WAVE_MAX_PULSES
const WaveMaxChars = C.PI_WAVE_MAX_CHARS

const BbI2cMinBaud = C.PI_BB_I2C_MIN_BAUD
const BbI2cMaxBaud = C.PI_BB_I2C_MAX_BAUD

const BbSpiMinBaud = C.PI_BB_SPI_MIN_BAUD
const BbSpiMaxBaud = C.PI_BB_SPI_MAX_BAUD

const BbSerMinBaud = C.PI_BB_SER_MIN_BAUD
const BbSerMaxBaud = C.PI_BB_SER_MAX_BAUD

const BbSerNormal = C.PI_BB_SER_NORMAL
const BbSerInvert = C.PI_BB_SER_INVERT

const WaveMinBaud = C.PI_WAVE_MIN_BAUD
const WaveMaxBaud = C.PI_WAVE_MAX_BAUD

const SpiMinBaud = C.PI_SPI_MIN_BAUD
const SpiMaxBaud = C.PI_SPI_MAX_BAUD

const MinWaveDatabits = C.PI_MIN_WAVE_DATABITS
const MaxWaveDatabits = C.PI_MAX_WAVE_DATABITS

const MinWaveHalfstopbits = C.PI_MIN_WAVE_HALFSTOPBITS
const MaxWaveHalfstopbits = C.PI_MAX_WAVE_HALFSTOPBITS

const WaveMaxMicros = C.PI_WAVE_MAX_MICROS

const MaxWaves = C.PI_MAX_WAVES

const MaxWaveCycles = C.PI_MAX_WAVE_CYCLES
const MaxWaveDelay = C.PI_MAX_WAVE_DELAY

const WaveCountPages = C.PI_WAVE_COUNT_PAGES

// wave tx mode
const WaveModeOneShot = C.PI_WAVE_MODE_ONE_SHOT
const WaveModeRepeat = C.PI_WAVE_MODE_REPEAT
const WaveModeOneShotSync = C.PI_WAVE_MODE_ONE_SHOT_SYNC
const WaveModeRepeatSync = C.PI_WAVE_MODE_REPEAT_SYNC

// special wave at return values
const WaveNotFound = C.PI_WAVE_NOT_FOUND
const NoTxWave = C.PI_NO_TX_WAVE

// Files, I2C, SPI, SER
const FileSlots = C.PI_FILE_SLOTS
const I2cSlots = C.PI_I2C_SLOTS
const SpiSlots = C.PI_SPI_SLOTS
const SerSlots = C.PI_SER_SLOTS

const MaxI2cAddr = C.PI_MAX_I2C_ADDR

const NumAuxSpiChannel = C.PI_NUM_AUX_SPI_CHANNEL
const NumStdSpiChannel = C.PI_NUM_STD_SPI_CHANNEL

const MaxI2cDeviceCount = C.PI_MAX_I2C_DEVICE_COUNT
const MaxSpiDeviceCount = C.PI_MAX_SPI_DEVICE_COUNT

// max pi_i2c_msg_t per transaction
const I2cRdrwIoctlMaxMsgs = C.PI_I2C_RDRW_IOCTL_MAX_MSGS

// flags for i2cTransaction
const I2cMWr = C.PI_I2C_M_WR
const I2cMRd = C.PI_I2C_M_RD
const I2cMTen = C.PI_I2C_M_TEN
const I2cMRecvLen = C.PI_I2C_M_RECV_LEN
const I2cMNoRdAck = C.PI_I2C_M_NO_RD_ACK
const I2cMIgnoreNak = C.PI_I2C_M_IGNORE_NAK
const I2cMRevDirAddr = C.PI_I2C_M_REV_DIR_ADDR
const I2cMNostart = C.PI_I2C_M_NOSTART

// I2CZip commands
const I2cEnd = C.PI_I2C_END
const I2cEsc = C.PI_I2C_ESC
const I2cStart = C.PI_I2C_START
const I2cCombinedOn = C.PI_I2C_COMBINED_ON
const I2cStop = C.PI_I2C_STOP
const I2cCombinedOff = C.PI_I2C_COMBINED_OFF
const I2cAddr = C.PI_I2C_ADDR
const I2cFlags = C.PI_I2C_FLAGS
const I2cRead = C.PI_I2C_READ
const I2cWrite = C.PI_I2C_WRITE

// Longest busy delay
const MaxBusyDelay = C.PI_MAX_BUSY_DELAY

// Timeout
const MinWdogTimeout = C.PI_MIN_WDOG_TIMEOUT
const MaxWdogTimeout = C.PI_MAX_WDOG_TIMEOUT

// Timer
const MinTimer = C.PI_MIN_TIMER
const MaxTimer = C.PI_MAX_TIMER

// Millis
const MinMs = C.PI_MIN_MS
const MaxMs = C.PI_MAX_MS

const MaxScripts = C.PI_MAX_SCRIPTS

const MaxScriptTags = C.PI_MAX_SCRIPT_TAGS
const MaxScriptVars = C.PI_MAX_SCRIPT_VARS
const MaxScriptParams = C.PI_MAX_SCRIPT_PARAMS

// Script status
const ScriptIniting = C.PI_SCRIPT_INITING
const ScriptHalted = C.PI_SCRIPT_HALTED
const ScriptRunning = C.PI_SCRIPT_RUNNING
const ScriptWaiting = C.PI_SCRIPT_WAITING
const ScriptFailed = C.PI_SCRIPT_FAILED

// Signum
const MinSignum = C.PI_MIN_SIGNUM
const MaxSignum = C.PI_MAX_SIGNUM

// TimeType
const TimeRelative = C.PI_TIME_RELATIVE
const TimeAbsolute = C.PI_TIME_ABSOLUTE

const MaxMicsDelay = C.PI_MAX_MICS_DELAY
const MaxMilsDelay = C.PI_MAX_MILS_DELAY

// CfgMillis
const BufMillisMin = C.PI_BUF_MILLIS_MIN
const BufMillisMax = C.PI_BUF_MILLIS_MAX

// CfgPeripheral
const ClockPwm = C.PI_CLOCK_PWM
const ClockPcm = C.PI_CLOCK_PCM

// DMA Channel
const MinDmaChannel = C.PI_MIN_DMA_CHANNEL
const MaxDmaChannel = C.PI_MAX_DMA_CHANNEL

// Port
const MinSocketPort = C.PI_MIN_SOCKET_PORT
const MaxSocketPort = C.PI_MAX_SOCKET_PORT

// IfFlags
const DisableFifoIf = C.PI_DISABLE_FIFO_IF
const DisableSockIf = C.PI_DISABLE_SOCK_IF
const LocalhostSockIf = C.PI_LOCALHOST_SOCK_IF

// MemallocMode
const MemAllocAuto = C.PI_MEM_ALLOC_AUTO
const MemAllocPagemap = C.PI_MEM_ALLOC_PAGEMAP
const MemAllocMailbox = C.PI_MEM_ALLOC_MAILBOX

// Filters
const MaxSteady = C.PI_MAX_STEADY
const MaxActive = C.PI_MAX_ACTIVE

// CfgInternals
const CfgDbgLevel = C.PI_CFG_DBG_LEVEL
const CfgAlertFreq = C.PI_CFG_ALERT_FREQ

// ISR
const RisingEdge = C.RISING_EDGE
const FallingEdge = C.FALLING_EDGE
const EitherEdge = C.EITHER_EDGE

// Pads
const MaxPad = C.PI_MAX_PAD

const MinPadStrength = C.PI_MIN_PAD_STRENGTH
const MaxPadStrength = C.PI_MAX_PAD_STRENGTH

// Files
const FileNone = C.PI_FILE_NONE
const FileMin = C.PI_FILE_MIN
const FileRead = C.PI_FILE_READ
const FileWrite = C.PI_FILE_WRITE
const FileRw = C.PI_FILE_RW
const FileAppend = C.PI_FILE_APPEND
const FileCreate = C.PI_FILE_CREATE
const FileTrunc = C.PI_FILE_TRUNC
const FileMax = C.PI_FILE_MAX

const FromStart = C.PI_FROM_START
const FromCurrent = C.PI_FROM_CURRENT
const FromEnd = C.PI_FROM_END

// Allowed socket connect addresses
const MaxConnectAddresses = C.MAX_CONNECT_ADDRESSES

// Socket command codes
const CmdModes = C.PI_CMD_MODES
const CmdModeg = C.PI_CMD_MODEG
const CmdPud = C.PI_CMD_PUD
const CmdRead = C.PI_CMD_READ
const CmdWrite = C.PI_CMD_WRITE
const CmdPwm = C.PI_CMD_PWM
const CmdPrs = C.PI_CMD_PRS
const CmdPfs = C.PI_CMD_PFS
const CmdServo = C.PI_CMD_SERVO
const CmdWdog = C.PI_CMD_WDOG
const CmdBr1 = C.PI_CMD_BR1
const CmdBr2 = C.PI_CMD_BR2
const CmdBc1 = C.PI_CMD_BC1
const CmdBc2 = C.PI_CMD_BC2
const CmdBs1 = C.PI_CMD_BS1
const CmdBs2 = C.PI_CMD_BS2
const CmdTick = C.PI_CMD_TICK
const CmdHwver = C.PI_CMD_HWVER
const CmdNo = C.PI_CMD_NO
const CmdNb = C.PI_CMD_NB
const CmdNp = C.PI_CMD_NP
const CmdNc = C.PI_CMD_NC
const CmdPrg = C.PI_CMD_PRG
const CmdPfg = C.PI_CMD_PFG
const CmdPrrg = C.PI_CMD_PRRG
const CmdHelp = C.PI_CMD_HELP
const CmdPigpv = C.PI_CMD_PIGPV
const CmdWvclr = C.PI_CMD_WVCLR
const CmdWvag = C.PI_CMD_WVAG
const CmdWvas = C.PI_CMD_WVAS
const CmdWvgo = C.PI_CMD_WVGO
const CmdWvgor = C.PI_CMD_WVGOR
const CmdWvbsy = C.PI_CMD_WVBSY
const CmdWvhlt = C.PI_CMD_WVHLT
const CmdWvsm = C.PI_CMD_WVSM
const CmdWvsp = C.PI_CMD_WVSP
const CmdWvsc = C.PI_CMD_WVSC
const CmdTrig = C.PI_CMD_TRIG
const CmdProc = C.PI_CMD_PROC
const CmdProcd = C.PI_CMD_PROCD
const CmdProcr = C.PI_CMD_PROCR
const CmdProcs = C.PI_CMD_PROCS
const CmdSlro = C.PI_CMD_SLRO
const CmdSlr = C.PI_CMD_SLR
const CmdSlrc = C.PI_CMD_SLRC
const CmdProcp = C.PI_CMD_PROCP
const CmdMics = C.PI_CMD_MICS
const CmdMils = C.PI_CMD_MILS
const CmdParse = C.PI_CMD_PARSE
const CmdWvcre = C.PI_CMD_WVCRE
const CmdWvdel = C.PI_CMD_WVDEL
const CmdWvtx = C.PI_CMD_WVTX
const CmdWvtxr = C.PI_CMD_WVTXR
const CmdWvnew = C.PI_CMD_WVNEW

const CmdI2co = C.PI_CMD_I2CO
const CmdI2cc = C.PI_CMD_I2CC
const CmdI2crd = C.PI_CMD_I2CRD
const CmdI2cwd = C.PI_CMD_I2CWD
const CmdI2cwq = C.PI_CMD_I2CWQ
const CmdI2crs = C.PI_CMD_I2CRS
const CmdI2cws = C.PI_CMD_I2CWS
const CmdI2crb = C.PI_CMD_I2CRB
const CmdI2cwb = C.PI_CMD_I2CWB
const CmdI2crw = C.PI_CMD_I2CRW
const CmdI2cww = C.PI_CMD_I2CWW
const CmdI2crk = C.PI_CMD_I2CRK
const CmdI2cwk = C.PI_CMD_I2CWK
const CmdI2cri = C.PI_CMD_I2CRI
const CmdI2cwi = C.PI_CMD_I2CWI
const CmdI2cpc = C.PI_CMD_I2CPC
const CmdI2cpk = C.PI_CMD_I2CPK

const CmdSpio = C.PI_CMD_SPIO
const CmdSpic = C.PI_CMD_SPIC
const CmdSpir = C.PI_CMD_SPIR
const CmdSpiw = C.PI_CMD_SPIW
const CmdSpix = C.PI_CMD_SPIX

const CmdSero = C.PI_CMD_SERO
const CmdSerc = C.PI_CMD_SERC
const CmdSerrb = C.PI_CMD_SERRB
const CmdSerwb = C.PI_CMD_SERWB
const CmdSerr = C.PI_CMD_SERR
const CmdSerw = C.PI_CMD_SERW
const CmdSerda = C.PI_CMD_SERDA

const CmdGdc = C.PI_CMD_GDC
const CmdGpw = C.PI_CMD_GPW

const CmdHc = C.PI_CMD_HC
const CmdHp = C.PI_CMD_HP

const CmdCf1 = C.PI_CMD_CF1
const CmdCf2 = C.PI_CMD_CF2

const CmdBi2cc = C.PI_CMD_BI2CC
const CmdBi2co = C.PI_CMD_BI2CO
const CmdBi2cz = C.PI_CMD_BI2CZ

const CmdI2cz = C.PI_CMD_I2CZ

const CmdWvcha = C.PI_CMD_WVCHA

const CmdSlri = C.PI_CMD_SLRI

const CmdCgi = C.PI_CMD_CGI
const CmdCsi = C.PI_CMD_CSI

const CmdFg = C.PI_CMD_FG
const CmdFn = C.PI_CMD_FN

const CmdNoib = C.PI_CMD_NOIB

const CmdWvtxm = C.PI_CMD_WVTXM
const CmdWvtat = C.PI_CMD_WVTAT

const CmdPads = C.PI_CMD_PADS
const CmdPadg = C.PI_CMD_PADG

const CmdFo = C.PI_CMD_FO
const CmdFc = C.PI_CMD_FC
const CmdFr = C.PI_CMD_FR
const CmdFw = C.PI_CMD_FW
const CmdFs = C.PI_CMD_FS
const CmdFl = C.PI_CMD_FL

const CmdShell = C.PI_CMD_SHELL

const CmdBspic = C.PI_CMD_BSPIC
const CmdBspio = C.PI_CMD_BSPIO
const CmdBspix = C.PI_CMD_BSPIX

// Pseudo commands
const CmdScript = C.PI_CMD_SCRIPT

const CmdAdd = C.PI_CMD_ADD
const CmdAnd = C.PI_CMD_AND
const CmdCall = C.PI_CMD_CALL
const CmdCmdr = C.PI_CMD_CMDR
const CmdCmdw = C.PI_CMD_CMDW
const CmdCmp = C.PI_CMD_CMP
const CmdDcr = C.PI_CMD_DCR
const CmdDcra = C.PI_CMD_DCRA
const CmdDiv = C.PI_CMD_DIV
const CmdHalt = C.PI_CMD_HALT
const CmdInr = C.PI_CMD_INR
const CmdInra = C.PI_CMD_INRA
const CmdJm = C.PI_CMD_JM
const CmdJmp = C.PI_CMD_JMP
const CmdJnz = C.PI_CMD_JNZ
const CmdJp = C.PI_CMD_JP
const CmdJz = C.PI_CMD_JZ
const CmdTag = C.PI_CMD_TAG
const CmdLd = C.PI_CMD_LD
const CmdLda = C.PI_CMD_LDA
const CmdLdab = C.PI_CMD_LDAB
const CmdMlt = C.PI_CMD_MLT
const CmdMod = C.PI_CMD_MOD
const CmdNop = C.PI_CMD_NOP
const CmdOr = C.PI_CMD_OR
const CmdPop = C.PI_CMD_POP
const CmdPopa = C.PI_CMD_POPA
const CmdPush = C.PI_CMD_PUSH
const CmdPusha = C.PI_CMD_PUSHA
const CmdRet = C.PI_CMD_RET
const CmdRl = C.PI_CMD_RL
const CmdRla = C.PI_CMD_RLA
const CmdRr = C.PI_CMD_RR
const CmdRra = C.PI_CMD_RRA
const CmdSta = C.PI_CMD_STA
const CmdStab = C.PI_CMD_STAB
const CmdSub = C.PI_CMD_SUB
const CmdSys = C.PI_CMD_SYS
const CmdWait = C.PI_CMD_WAIT
const CmdX = C.PI_CMD_X
const CmdXa = C.PI_CMD_XA
const CmdXor = C.PI_CMD_XOR

// Defaults
const DefaultBufferMillis = C.PI_DEFAULT_BUFFER_MILLIS
const DefaultClkMicros = C.PI_DEFAULT_CLK_MICROS
const DefaultClkPeripheral = C.PI_DEFAULT_CLK_PERIPHERAL
const DefaultIfFlags = C.PI_DEFAULT_IF_FLAGS
const DefaultDmaChannel = C.PI_DEFAULT_DMA_CHANNEL
const DefaultDmaPrimaryChannel = C.PI_DEFAULT_DMA_PRIMARY_CHANNEL
const DefaultDmaSecondaryChannel = C.PI_DEFAULT_DMA_SECONDARY_CHANNEL
const DefaultSocketPort = C.PI_DEFAULT_SOCKET_PORT
const DefaultSocketPortStr = C.PI_DEFAULT_SOCKET_PORT_STR
const DefaultSocketAddrStr = C.PI_DEFAULT_SOCKET_ADDR_STR
const DefaultUpdateMaskUnknown = C.PI_DEFAULT_UPDATE_MASK_UNKNOWN
const DefaultUpdateMaskB1 = C.PI_DEFAULT_UPDATE_MASK_B1
const DefaultUpdateMaskAB2 = C.PI_DEFAULT_UPDATE_MASK_A_B2
const DefaultUpdateMaskAplusBplus = C.PI_DEFAULT_UPDATE_MASK_APLUS_BPLUS
const DefaultUpdateMaskZero = C.PI_DEFAULT_UPDATE_MASK_ZERO
const DefaultUpdateMaskPi2b = C.PI_DEFAULT_UPDATE_MASK_PI2B
const DefaultUpdateMaskPi3b = C.PI_DEFAULT_UPDATE_MASK_PI3B
const DefaultUpdateMaskCompute = C.PI_DEFAULT_UPDATE_MASK_COMPUTE
const DefaultMemAllocMode = C.PI_DEFAULT_MEM_ALLOC_MODE

const DefaultCfgInternals = C.PI_DEFAULT_CFG_INTERNALS

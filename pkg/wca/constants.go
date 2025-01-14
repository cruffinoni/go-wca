package wca

const (
	AUDCLNT_SESSIONFLAGS_EXPIREWHENUNOWNED       = 0x10000000
	AUDCLNT_SESSIONFLAGS_DISPLAY_HIDE            = 0x20000000
	AUDCLNT_SESSIONFLAGS_DISPLAY_HIDEWHENEXPIRED = 0x40000000
)

const (
	AUDCLNT_STREAMOPTIONS_NONE = iota
	AUDCLNT_STREAMOPTIONS_RAW
	AUDCLNT_STREAMOPTIONS_MATCH_FORMAT
)

const (
	WAVE_FORMAT_EXTENSIBLE = 0xFFFE
)

const (
	AudioCategory_Other = iota
	AudioCategory_ForegroundOnlyMedia
	AudioCategory_BackgroundCapableMedia
	AudioCategory_Communications
	AudioCategory_Alerts
	AudioCategory_SoundEffects
	AudioCategory_GameEffects
	AudioCategory_GameMedia
	AudioCategory_GameChat
	AudioCategory_Speech
	AudioCategory_Movie
	AudioCategory_Media
)

const (
	WAVE_FORMAT_PCM = 0x1
)

const (
	INFINITE = 0xFFFFFFFF
)

const (
	EConsole = iota
	EMultimedia
	ECommunications
	ERole_enum_count
)

const (
	DELETE       = 0x00010000
	READ_CONTROL = 0x00020000
	SYNCHRONIZE  = 0x00100000
	WRITE_DAC    = 0x00040000
	WRITE_OWNER  = 0x00080000
)

const (
	EVENT_ALL_ACCESS   = 0x1F0003
	EVENT_MODIFY_STATE = 0x0002
)

const (
	CREATE_EVENT_INITIAL_SET  = 0x00000002
	CREATE_EVENT_MANUAL_RESET = 0x00000001
)

const (
	AUDCLNT_BUFFERFLAGS_DATA_DISCONTINUITY = 0x1
	AUDCLNT_BUFFERFLAGS_SILENT             = 0x2
	AUDCLNT_BUFFERFLAGS_TIMESTAMP_ERROR    = 0x4
)

const (
	AUDCLNT_STREAMFLAGS_CROSSPROCESS        = 0x00010000
	AUDCLNT_STREAMFLAGS_LOOPBACK            = 0x00020000
	AUDCLNT_STREAMFLAGS_EVENTCALLBACK       = 0x00040000
	AUDCLNT_STREAMFLAGS_NOPERSIST           = 0x00080000
	AUDCLNT_STREAMFLAGS_RATEADJUST          = 0x00100000
	AUDCLNT_STREAMFLAGS_AUTOCONVERTPCM      = 0x80000000
	AUDCLNT_STREAMFLAGS_SRC_DEFAULT_QUALITY = 0x08000000
)

const (
	AUDCLNT_SHAREMODE_SHARED = iota
	AUDCLNT_SHAREMODE_EXCLUSIVE
)

const (
	ENDPOINT_SYSFX_ENABLED = iota
	ENDPOINT_SYSFX_DISABLED
)

const (
	DEVICE_STATE_ACTIVE     = 0x00000001
	DEVICE_STATE_DISABLED   = 0x00000002
	DEVICE_STATE_NOTPRESENT = 0x00000004
	DEVICE_STATE_UNPLUGGED  = 0x00000008
	DEVICE_STATEMASK_ALL    = 0x0000000F
)

const (
	ERender = iota
	ECapture
	EAll
	EDataFlow_enum_count
)

const (
	STGM_READ       = 0x0
	STGM_WRITE      = 0x1
	STGM_READ_WRITE = 0x2
)

const (
	CLSCTX_INPROC_SERVER          = 0x1
	CLSCTX_INPROC_HANDLER         = 0x2
	CLSCTX_LOCAL_SERVER           = 0x4
	CLSCTX_INPROC_SERVER16        = 0x8
	CLSCTX_REMOTE_SERVER          = 0x10
	CLSCTX_INPROC_HANDLER16       = 0x20
	CLSCTX_RESERVED1              = 0x40
	CLSCTX_RESERVED2              = 0x80
	CLSCTX_RESERVED3              = 0x100
	CLSCTX_RESERVED4              = 0x200
	CLSCTX_NO_CODE_DOWNLOAD       = 0x400
	CLSCTX_RESERVED5              = 0x800
	CLSCTX_NO_CUSTOM_MARSHAL      = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD   = 0x2000
	CLSCTX_NO_FAILURE_LOG         = 0x4000
	CLSCTX_DISABLE_AAA            = 0x8000
	CLSCTX_ENABLE_AAA             = 0x10000
	CLSCTX_FROM_DEFAULT_CONTEXT   = 0x20000
	CLSCTX_ACTIVATE_32_BIT_SERVER = 0x40000
	CLSCTX_ACTIVATE_64_BIT_SERVER = 0x80000
	CLSCTX_ENABLE_CLOAKING        = 0x100000
	CLSCTX_APPCONTAINER           = 0x400000
	CLSCTX_ACTIVATE_AAA_AS_IU     = 0x800000
	CLSCTX_PS_DLL                 = 0x80000000
	CLSCTX_ALL                    = CLSCTX_INPROC_SERVER | CLSCTX_INPROC_HANDLER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
)

const (
	AUDCLNT_E_NOT_INITIALIZED              = 0x001
	AUDCLNT_E_ALREADY_INITIALIZED          = 0x002
	AUDCLNT_E_WRONG_ENDPOINT_TYPE          = 0x003
	AUDCLNT_E_DEVICE_INVALIDATED           = 0x004
	AUDCLNT_E_NOT_STOPPED                  = 0x005
	AUDCLNT_E_BUFFER_TOO_LARGE             = 0x006
	AUDCLNT_E_OUT_OF_ORDER                 = 0x007
	AUDCLNT_E_UNSUPPORTED_FORMAT           = 0x008
	AUDCLNT_E_INVALID_SIZE                 = 0x009
	AUDCLNT_E_DEVICE_IN_USE                = 0x00a
	AUDCLNT_E_BUFFER_OPERATION_PENDING     = 0x00b
	AUDCLNT_E_THREAD_NOT_REGISTERED        = 0x00c
	AUDCLNT_E_EXCLUSIVE_MODE_NOT_ALLOWED   = 0x00e
	AUDCLNT_E_ENDPOINT_CREATE_FAILED       = 0x00f
	AUDCLNT_E_SERVICE_NOT_RUNNING          = 0x010
	AUDCLNT_E_EVENTHANDLE_NOT_EXPECTED     = 0x011
	AUDCLNT_E_EXCLUSIVE_MODE_ONLY          = 0x012
	AUDCLNT_E_BUFDURATION_PERIOD_NOT_EQUAL = 0x013
	AUDCLNT_E_EVENTHANDLE_NOT_SET          = 0x014
	AUDCLNT_E_INCORRECT_BUFFER_SIZE        = 0x015
	AUDCLNT_E_BUFFER_SIZE_ERROR            = 0x016
	AUDCLNT_E_CPUUSAGE_EXCEEDED            = 0x017
	AUDCLNT_E_BUFFER_ERROR                 = 0x018
	AUDCLNT_E_BUFFER_SIZE_NOT_ALIGNED      = 0x019
	AUDCLNT_E_INVALID_DEVICE_PERIOD        = 0x020
)

const (
	AudioSessionStateInactive = iota
	AudioSessionStateActive
	AudioSessionStateExpired
)

const (
	SPEAKER_FRONT_LEFT            = 0x1
	SPEAKER_FRONT_RIGHT           = 0x2
	SPEAKER_FRONT_CENTER          = 0x4
	SPEAKER_LOW_FREQUENCY         = 0x8
	SPEAKER_BACK_LEFT             = 0x10
	SPEAKER_BACK_RIGHT            = 0x20
	SPEAKER_FRONT_LEFT_OF_CENTER  = 0x40
	SPEAKER_FRONT_RIGHT_OF_CENTER = 0x80
	SPEAKER_BACK_CENTER           = 0x100
	SPEAKER_SIDE_LEFT             = 0x200
	SPEAKER_SIDE_RIGHT            = 0x400
	SPEAKER_TOP_CENTER            = 0x800
	SPEAKER_TOP_FRONT_LEFT        = 0x1000
	SPEAKER_TOP_FRONT_CENTER      = 0x2000
	SPEAKER_TOP_FRONT_RIGHT       = 0x4000
	SPEAKER_TOP_BACK_LEFT         = 0x8000
	SPEAKER_TOP_BACK_CENTER       = 0x10000
	SPEAKER_TOP_BACK_RIGHT        = 0x20000
)

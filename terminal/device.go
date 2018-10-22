package terminal

type Device int

const (
	DeviceStdout               Device = iota >> 1 // stdout
	DeviceStderr                                  // stderr
	DeviceStdin                                   // stdin
	DevicePipeline                                // |
	DeviceRedirectOutput                          // >
	DeviceRedirectOutputAppend                    // >>
	DeviceRedirectInput                           // <
)

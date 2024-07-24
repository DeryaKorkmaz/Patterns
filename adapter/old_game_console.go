package adapter

// OldGameConsole is a device with AV output
type OldGameConsole struct{}

func (c *OldGameConsole) OutputAV() string {
	return "AV signal from Old Game Console"
}

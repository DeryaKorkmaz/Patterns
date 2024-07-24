package adapter

// AVDevice represents devices with AV output
type AVDevice interface {
	OutputAV() string
}

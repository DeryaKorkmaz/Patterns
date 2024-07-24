package adapter

// HDMIDevice represents devices with HDMI input
type HDMIDevice interface {
	InputHDMI(data string)
}

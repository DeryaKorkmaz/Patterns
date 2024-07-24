package adapter

import "fmt"

// AVtoHDMIAdapter converts AV signal to HDMI
type AVtoHDMIAdapter struct {
	AVDevice AVDevice
}

func (adapter *AVtoHDMIAdapter) InputHDMI(data string) {
	avSignal := adapter.AVDevice.OutputAV()
	hdmiSignal := "Converted to HDMI: " + data + " " + avSignal
	fmt.Println(hdmiSignal)
}

package adapter

import "fmt"

// NewTelevision is a device with HDMI input
type NewTelevision struct{}

func (tv *NewTelevision) InputHDMI(data string) {
	fmt.Println("Displaying:", data)
}

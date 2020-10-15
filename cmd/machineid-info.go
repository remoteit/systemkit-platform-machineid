package main

import (
	"fmt"

	machineid "github.com/remoteit/systemkit-platform-machineid"
)

func main() {
	fmt.Println(machineid.MachineID())
}

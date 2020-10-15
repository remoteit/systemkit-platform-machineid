// +build linux
// +build !android

package tests

import (
	"fmt"
	"testing"

	machineid "github.com/remoteit/systemkit-platform-machineid"
)

func Test_Linux(t *testing.T) {
	fmt.Println(machineid.MachineID())
}

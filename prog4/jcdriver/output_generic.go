// +build !linux

package main

import (
	"joycon/prog4/jcpc"
	"joycon/prog4/output"
)

func getOutputFactory() jcpc.OutputFactory {
	return func(t jcpc.JoyConType, playerNum int, remap InputRemappingOptions) (jcpc.Output, error) {
		return output.NewConsole(t, playerNum)
	}
}

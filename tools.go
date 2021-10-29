package Toxel

import (
	"time"
)

//Here are located helpfull and simple functions

//SaveSliceAccess checks if you entered valid index for given slice, if you didn't it returns 0
func SaveSliceAccess(index, sliceLength int) int {
	if sliceLength < index {
		return 0
	}
	return index
}

//Stupid and simple implementation of FPS count
func FPSCounter(currentFrames *float64, returnFPS func(float64)) {
	go func(currentFrames *float64) {
		for {
			time.Sleep(time.Second)
			returnFPS(*currentFrames)
			*currentFrames = 0
		}
	}(currentFrames)
}

// Copyright 2013, David Howden
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raspicam

import (
	"strings"
	"testing"
)

func TestDefaultParams(t *testing.T) {
	const paramsOut = "--output -"

	testNames := [...]string{"Still", "StillYUV", "Vid"}
	testCases := [...]CaptureCommand{NewStill(), NewStillYUV(), NewVid()}

	for i, test := range testNames {
		paramString := strings.Join(testCases[i].params(), " ")
		if paramString != paramsOut {
			t.Errorf("%v: param() returned %v, expected %v", test, paramString, paramsOut)
		}
	}
}

func TestBasicParams(t *testing.T) {
	const paramsOut = "--output - --timeout 10 --width 100 --height 1000"

	testNames := [...]string{"Still", "StillYUV", "Vid"}

	still := NewStill()
	still.Timeout = 10
	still.Width = 100
	still.Height = 1000

	stillYUV := NewStillYUV()
	stillYUV.Timeout = 10
	stillYUV.Width = 100
	stillYUV.Height = 1000

	vid := NewVid()
	vid.Timeout = 10
	vid.Width = 100
	vid.Height = 1000

	testCases := [...]CaptureCommand{still, stillYUV, vid}

	for i, test := range testNames {
		paramString := strings.Join(testCases[i].params(), " ")
		if paramString != paramsOut {
			t.Errorf("%v: param() returned %v, expected %v", test, paramString, paramsOut)
		}
	}
}

func TestCameraParams(t *testing.T) {
	const paramsOut = "--output - --timeout 10 --sharpness 11 --contrast 13 --brightness 12"

	testNames := [...]string{"Still", "StillYUV", "Vid"}

	still := NewStill()
	still.Timeout = 10
	still.Camera.Sharpness = 11
	still.Camera.Brightness = 12
	still.Camera.Contrast = 13

	stillYUV := NewStillYUV()
	stillYUV.Timeout = 10
	stillYUV.Camera.Sharpness = 11
	stillYUV.Camera.Brightness = 12
	stillYUV.Camera.Contrast = 13

	vid := NewVid()
	vid.Timeout = 10
	vid.Camera.Sharpness = 11
	vid.Camera.Brightness = 12
	vid.Camera.Contrast = 13

	testCases := [...]CaptureCommand{still, stillYUV, vid}

	for i, test := range testNames {
		paramString := strings.Join(testCases[i].params(), " ")
		if paramString != paramsOut {
			t.Errorf("%v: param() returned %v, expected %v", test, paramString, paramsOut)
		}
	}
}
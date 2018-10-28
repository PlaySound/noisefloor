package engine

import . "github.com/jacoblister/noisefloor/common"

import "fmt"

// Engine - DSP processing engine
type Engine struct {
}

var midiinput MIDIInput
var patch Patch

// Start initilized the engine, with a specified sampling rate
func Start(sampleRate int) {
	println("do DSP start, sample rate:", sampleRate)
	midiinput.Start()
	patch.Start(sampleRate)
}

// Stop suspends the engine
func Stop() {
	println("do DSP stop")
}

// Process processes a block of samples and midi events
func Process(samplesIn [][]AudioFloat, samplesOut [][]AudioFloat, midiIn []MidiEvent, midiOut []MidiEvent) {
	midiinput.ProcessMIDI(midiIn)
	freqs := midiinput.Process()
	fmt.Println(freqs)

	var len = len(samplesOut[0])
	for i := 0; i < len; i++ {
		var sample = patch.Process(freqs[0][0], freqs[0][1], freqs[0][2])
		samplesOut[0][i] = sample
		samplesOut[1][i] = sample
	}
}

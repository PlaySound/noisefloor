package engine

import . "github.com/jacoblister/noisefloor/common"
import "github.com/jacoblister/noisefloor/engine/processor"

// Patch is a simple minimal example patch
type Patch struct {
	oscillator processor.Oscillator
	envelope   processor.Envelope
	gain       processor.Gain
}

// Start - init patch
func (p *Patch) Start(sampleRate int) {
	p.oscillator.Start(sampleRate)
	p.envelope.Start(sampleRate)
	p.gain.Start(sampleRate)
}

// Process - produce next sample
func (p *Patch) Process(freq AudioFloat, gate AudioFloat, trigger AudioFloat) (output AudioFloat) {
	p.oscillator.Freq = freq
	sample := p.oscillator.Process()
	env := p.envelope.Process(gate, trigger)

	output = p.gain.Process(sample, env)
	return
}

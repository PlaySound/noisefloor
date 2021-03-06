// +build windows

package nf

import (
	"fmt"
	"unsafe"

	"github.com/jacoblister/noisefloor/app/audiomodule"
)

//#include <string.h>
import "C"

type driverAudioASIO struct {
	asioDriver     *ASIODriver
	audioProcessor audiomodule.AudioProcessor
	driverMidi     driverMidi
}

//export goAudioASIOCallback
func goAudioASIOCallback(arg unsafe.Pointer, blockLength int,
	channelInCount int, channelIn unsafe.Pointer,
	channelOutCount int, channelOut unsafe.Pointer) {

	driverAudio := (*driverAudioASIO)(arg)

	goAudioCallback(driverAudio, blockLength, channelInCount, channelIn, channelOutCount, channelOut)
}

func (d *driverAudioASIO) getDriverMidi() driverMidi {
	return d.driverMidi
}

func (d *driverAudioASIO) setDriverMidi(driverMidi driverMidi) {
	d.driverMidi = driverMidi
}

func (d *driverAudioASIO) getAudioProcessor() audiomodule.AudioProcessor {
	return d.audioProcessor
}

func (d *driverAudioASIO) setAudioProcessor(audioProcessor audiomodule.AudioProcessor) {
	d.audioProcessor = audioProcessor
}

func (d *driverAudioASIO) start() {
	println("ASIO start")

	CoInitialize(0)

	drivers, err := ListDrivers()
	if err != nil {
		panic("ASIO cannot load drivers")
	}

	d.asioDriver = drivers["GP-10"]

	err = d.asioDriver.Open()
	if err != nil {
		panic("ASIO cannot open driver")
	}

	drv := d.asioDriver.ASIO

	// getChannels
	in, out, err := drv.GetChannels()
	if err != nil {
		panic("ASIO cannot get channels")
	}
	fmt.Printf("getChannels():        %d, %d\n", in, out)

	// getBufferSize
	minSize, maxSize, preferredSize, granularity, err := drv.GetBufferSize()
	if err != nil {
		panic("ASIO cannot get buffer size")
	}
	fmt.Printf("getBufferSize():      %d, %d, %d, %d\n", minSize, maxSize, preferredSize, granularity)

	bufferDescriptors := make([]BufferInfo, 0, in+out)
	for i := 0; i < in; i++ {
		bufferDescriptors = append(bufferDescriptors, BufferInfo{
			Channel: i,
			IsInput: true,
		})
		cinfo, err := drv.GetChannelInfo(i, true)
		if err != nil {
			panic("ASIO cannot get input channel info")
		}
		fmt.Printf(" IN%-2d: active=%v, group=%d, type=%d, name=%s\n", i+1, cinfo.IsActive, cinfo.ChannelGroup, cinfo.SampleType, cinfo.Name)
	}
	for i := 0; i < out; i++ {
		bufferDescriptors = append(bufferDescriptors, BufferInfo{
			Channel: i,
			IsInput: false,
		})
		cinfo, err := drv.GetChannelInfo(i, false)
		if err != nil {
			panic("ASIO cannot get output channel info")
		}
		fmt.Printf("OUT%-2d: active=%v, group=%d, type=%d, name=%s\n", i+1, cinfo.IsActive, cinfo.ChannelGroup, cinfo.SampleType, cinfo.Name)
	}

	err = drv.CreateBuffers(bufferDescriptors, preferredSize)
	if err != nil {
		panic("ASIO cannot create buffers")
	}

	drv.SetBufferChannels(unsafe.Pointer(d), in, out, preferredSize)
	for i := 0; i < in; i++ {
		drv.SetBufferPtr(1, i, unsafe.Pointer(bufferDescriptors[i].Buffers[0]), unsafe.Pointer(bufferDescriptors[i].Buffers[1]))
	}
	for i := 0; i < out; i++ {
		drv.SetBufferPtr(0, i, unsafe.Pointer(bufferDescriptors[i+in].Buffers[0]), unsafe.Pointer(bufferDescriptors[i+in].Buffers[1]))
	}

	fmt.Printf("Output Ready %t\n", drv.OutputReady())

	err = drv.Start()
	if err != nil {
		panic("ASIO cannot start driver")
	}

	fmt.Printf("ASIO Started\n")
}

func (d *driverAudioASIO) stop() {
	println("ASIO stop")

	drv := d.asioDriver.ASIO
	err := drv.Stop()
	if err != nil {
		panic("ASIO cannot stop driver")
	}

	drv.DisposeBuffers()

	d.asioDriver.Close()

	CoUninitialize()
}

func (d *driverAudioASIO) samplingRate() int {
	return 44100
}

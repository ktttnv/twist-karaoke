package audio

import (
	"fmt"
	"math"
	"sync/atomic"

	"github.com/gen2brain/malgo"
)

type Capture struct {
	ctx    *malgo.AllocatedContext
	device *malgo.Device
	rms    atomic.Uint32
}

func NewCapture() (*Capture, error) {
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	if err != nil {
		return nil, fmt.Errorf("init context: %w", err)
	}

	c := &Capture{ctx: ctx}

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Capture)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 1
	deviceConfig.SampleRate = 16000
	deviceConfig.Alsa.NoMMap = 1

	onData := func(_, in []byte, framecount uint32) {
		if len(in) < 2 {
			return
		}
		var sum float64
		for i := 0; i < len(in)-1; i += 2 {
			sample := int16(in[i]) | int16(in[i+1])<<8
			f := float64(sample) / 32768.0
			sum += f * f
		}
		numSamples := len(in) / 2
		if numSamples > 0 {
			rms := math.Sqrt(sum / float64(numSamples))
			c.rms.Store(math.Float32bits(float32(rms)))
		}
	}

	device, err := malgo.InitDevice(ctx.Context, deviceConfig, malgo.DeviceCallbacks{
		Data: onData,
	})
	if err != nil {
		_ = ctx.Uninit()
		ctx.Free()
		return nil, fmt.Errorf("init device: %w", err)
	}
	c.device = device

	if err := device.Start(); err != nil {
		device.Uninit()
		_ = ctx.Uninit()
		ctx.Free()
		return nil, fmt.Errorf("start device: %w", err)
	}

	return c, nil
}

func (c *Capture) RMS() float32 {
	return math.Float32frombits(c.rms.Load())
}

func (c *Capture) Stop() {
	if c.device != nil {
		c.device.Uninit()
	}
	if c.ctx != nil {
		_ = c.ctx.Uninit()
		c.ctx.Free()
	}
}

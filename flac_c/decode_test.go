package flac_c

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/faiface/beep/speaker"
)

func TestFlacC(t *testing.T) {

	f, err := os.Open("/Users/anhoder/Desktop/a.flac")
	if err != nil {
		panic(err)
	}

	streamer, format, _ := Decode(f)

	log.Printf("Convert audio sample rate: %d, channels: %d\n", format.SampleRate, format.NumChannels)

	_ = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))

	speaker.Play(streamer)

	select {}
}

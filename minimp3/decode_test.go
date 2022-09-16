package mp3

import (
	"github.com/faiface/beep/speaker"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestMp3(t *testing.T) {

	resp, _ := http.Get("http://m801.music.126.net/20220916220332/344de4d472ba8a4a49dc8f8637d646b5/jdymusic/obj/wo3DlMOGwrbDjj7DisKw/9444694535/c8bb/65da/db54/9e257ff0610f84e4242cf0127dac6005.mp3")

	streamer, format, _ := Decode(resp.Body)

	log.Printf("Convert audio sample rate: %d, channels: %d\n", format.SampleRate, format.NumChannels)

	_ = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))

	speaker.Play(streamer)

	select {}
}

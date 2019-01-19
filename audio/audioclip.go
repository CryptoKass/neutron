package audio

import (
	"github.com/veandco/go-sdl2/mix"
)

type AudioClip struct {
	SdlChunk *mix.Chunk
	SdlType  mix.MusicType
	Volume   float32
	Channel  int
}

func LoadAudioClip(filename string) *AudioClip {
	mus, err := mix.LoadWAV(filename)
	if err != nil {
		panic(err)
	}
	return &AudioClip{
		SdlChunk: mus,
		Volume:   100,
	}
}

func (ac *AudioClip) Play() {
	ac.Channel, _ = ac.SdlChunk.Play(-1, 0)
}

func (ac *AudioClip) Stop() {
	mix.HaltChannel(ac.Channel)
}

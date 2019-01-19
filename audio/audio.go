package audio

import (
	"github.com/veandco/go-sdl2/mix"
)

func Init() {
	err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096)
	if err != nil {
		panic(err)
	}
}

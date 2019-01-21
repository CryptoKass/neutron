package console

import (
	"fmt"
	"log"
	"os"

	"github.com/CryptoKass/neutron/core"
	"github.com/veandco/go-sdl2/ttf"
)

func Init() {
	var err error
	DebugLogger = log.New(os.Stdout, "neutron", 1)
	gameloglines = make([]string, 0)

	Font, err = ttf.OpenFont("./res/main.ttf", 12)
	if err != nil {
		panic(err)
	}

	Color = core.Color{R: 255, G: 255, B: 255, A: 255}
}

func Log(v ...interface{}) {
	DebugLog(v...)
	GameLog(v...)
}

func LogF(format string, v ...interface{}) {
	DebugLog(fmt.Sprintf(format, v...))
	GameLog(fmt.Sprintf(format, v...))
}

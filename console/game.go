package console

import "fmt"

var gameloglines []string

func GameLog(v ...interface{}) {
	addgamelogline(fmt.Sprint(v))
}

func GameLogf(format string, v ...interface{}) {
	addgamelogline(fmt.Sprintf(format, v))
}

func addgamelogline(str string) {
	if len(gameloglines) < GameLogLineCount {
		gameloglines = append(gameloglines, str)
		return
	}
	gameloglines = append(gameloglines[1:], str)
}

package input

import "github.com/veandco/go-sdl2/sdl"

func IsKeyDown(keycode int8) bool {
	keys := sdl.GetKeyboardState()
	return keys[keycode] == 1
}

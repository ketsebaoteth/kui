package window

import rl "github.com/gen2brain/raylib-go/raylib"
//this file will contain common window operation functions
//sets target fps
func SetFps(fps int32) {
	rl.SetTargetFPS(fps)
}

package window

import rl "github.com/gen2brain/raylib-go/raylib"

//clears screen veryframe and returns a callback for an everyframe function
func WindowLoop(f func(), k *Kui) {
	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{k.BackgroundColor[0], k.BackgroundColor[1], k.BackgroundColor[2], k.BackgroundColor[3]})
		f()
		rl.EndDrawing()
	}
}

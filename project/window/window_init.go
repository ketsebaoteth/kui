package window

import rl "github.com/gen2brain/raylib-go/raylib"

//creates a window needs a pointer to the window struct
func CreateWindow(k *Kui) {
	rl.InitWindow(k.Dimension[0], k.Dimension[1], k.Windowtitle)
	rl.SetWindowPosition(k.Position[0], k.Position[1])

}

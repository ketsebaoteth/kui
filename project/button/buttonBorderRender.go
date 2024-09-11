package button

import rl "github.com/gen2brain/raylib-go/raylib"

func drawButtonBorder(b Button) {
	if b.BorderStyle == "solid" {
		borderRectangle := rl.Rectangle{X: float32(b.Pos[0] + b.Margin[0]), Y: float32(b.Pos[1] + b.Margin[1]), Width: float32(b.Dim[0] - b.Margin[2]), Height: float32(b.Dim[1] - b.Margin[3])}
		rl.DrawRectangleLinesEx(borderRectangle, float32(b.BorderWidth), rl.Color{R: uint8(b.BorderColor[0]), G: uint8(b.BorderColor[1]), B: uint8(b.BorderColor[2]), A: uint8(b.BorderColor[3])})
	}
}

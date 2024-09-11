package button

import rl "github.com/gen2brain/raylib-go/raylib"

//draws basic rectangle frame for button
//Todo: handle buttons with border radius
func drawButtonSkeleton(b Button) {
	rl.DrawRectangle(b.Pos[0]+b.Margin[0], b.Pos[1]+b.Margin[1], b.Dim[0]-b.Margin[2], b.Dim[1]-b.Margin[3], rl.RayWhite)

}

package button

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// handles flexbox alignItems for buttons text element
func textPositionAlign(alignitems string, b Button) int32 {
	var verticalExtras int32
	textMeasurment := rl.MeasureTextEx(b.Fontfamily, b.Text, 20, 1)
	switch alignitems {
	case "center":
		verticalExtras = (b.Dim[1]/2 - int32(textMeasurment.Y/2))
	case "flex-start":
		verticalExtras = (b.Padding[1] + b.Margin[1])
	case "flex-end":
		verticalExtras = (b.Dim[1] - int32(textMeasurment.Y) - b.Padding[3] + b.Margin[0])
	case "baseline":
		verticalExtras = (b.Dim[1] - int32(textMeasurment.Y) - b.Padding[3] - b.Margin[3])
	}

	return verticalExtras
}

// hadnles flexbox JustifyContent for button text
func textPositionJustify(JustifyContent string, b Button) int32 {
	var horizontalExtras int32
	textMeasurment := rl.MeasureTextEx(b.Fontfamily, b.Text, 20, 1)
	switch JustifyContent {
	case "center":
		horizontalExtras = (b.Dim[0]/2 - int32(textMeasurment.X/2))
	case "flex-start":
		horizontalExtras = (b.Padding[0] + b.Margin[0])
	case "flex-end":
		horizontalExtras = (b.Dim[0] - int32(textMeasurment.X) - b.Padding[3] + b.Margin[0])
	case "baseline":
		horizontalExtras = (b.Dim[0] - int32(textMeasurment.X) - b.Padding[3] - b.Margin[3])
	}
	return horizontalExtras
}

// draws buttons text inside them
func drawButtonText(b Button) {
	textposition := rl.Vector2{X: float32(b.Pos[0] + b.Dim[0]/2), Y: float32(b.Pos[1] + b.Dim[1]/2)}
	textMeasurment := rl.MeasureTextEx(b.Fontfamily, b.Text, 20, 1)
	if b.Display == "flex" {
		//everything that needs to be added to put the text properly considering the alignitems justicontent and margin and padding
		verticalExtras := textPositionAlign(b.AlignItems, b)
		horizontalExtras := textPositionJustify(b.JustifyContent, b)
		//final textposition calculated
		textposition = rl.Vector2{X: float32(b.Pos[0] + horizontalExtras), Y: float32(b.Pos[1] + verticalExtras)}
	} else if b.Display == "block" {
		textposition = rl.Vector2{X: float32(b.Pos[0] + b.Dim[0]/2 - int32(textMeasurment.X/2)), Y: float32(b.Pos[1] + b.Dim[1]/2 - int32(textMeasurment.Y/2))}
	}
	//drawText
	rl.DrawTextEx(b.Fontfamily, b.Text, textposition, 20, 1, rl.Black)
	//change text antialiasing style
	rl.SetTextureFilter(b.Fontfamily.Texture, rl.FilterBilinear)
}

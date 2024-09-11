package main

import (
	"main/project/button"
	elementtree "main/project/elementTree"
	"main/project/window"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var mywindow *window.Kui
var tree elementtree.ElementTree

func everyframe() {
	// This function will be called every frame
	tree.RenderTree()

}
func sayHelloWorld() {
	println("Hello, World!")
}

func main() {
	mywindow = &window.Kui{
		Windowtitle:     "Hello, World!",
		Position:        [2]int{100, 100},
		Dimension:       [2]int32{800, 450},
		BackgroundColor: [4]uint8{125, 125, 125, 255},
	}

	window.CreateWindow(mywindow)
	font := rl.LoadFont("./fonts/Roboto/Roboto-Regular.ttf")
	defer rl.UnloadFont(font)

	if font.Texture.ID == 0 {
		println("Error loading font")
	}
	window.SetFps(60)
	tree = elementtree.ElementTree{}
	mybtn := button.Button{Text: "Hello, World!", Pos: [2]int32{100, 100}, Dim: [2]int32{200, 100}, ClickCallBack: sayHelloWorld, Fontfamily: font}
	tree.AddElement(mybtn)
	window.WindowLoop(everyframe, mywindow)
}

package window

//a window type contains window properties
type Kui struct {
	Windowtitle     string
	Position        [2]int
	Dimension       [2]int32
	BackgroundColor [4]uint8
	//onresizeeventlistners []func
}

//func WindowResizeEventListener(fn func){
//	onresizeeventlistners.put(fn)
//}

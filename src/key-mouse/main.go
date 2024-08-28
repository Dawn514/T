package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"gocv.io/x/gocv"
)

var screenWidth int
var screenHeight int
func init() {
	screenWidth, screenHeight = robotgo.GetScreenSize()
}
func transX(x1 int) (x2 int) {
	x2 = x1*screenWidth/1920
	return
}
func transY(y1 int) (y2 int) {
	y2 = y1*screenHeight/1080
	return
}
func main() {
	//robotgo.Sleep(1)
	//x := 1190
	//y := 939
	//
	//
	//for i:=0; i <5; i++ {
	//	robotgo.Sleep(1)
	//	col := robotgo.GetPixelColor(transX(x), transY(y))
	//	fmt.Println("color = ", col)
	//
	//	robotgo.Sleep(1)
	//	robotgo.Move(transX(x), transY(y))
	//	robotgo.Sleep(1)
	//
	//}

	fmt.Println("OpenCV version:", gocv.Version())

	// 测试加载图像
	img := gocv.IMRead("D：桌面/1.jpg", gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("Failed to read image")
		return
	}
	defer img.Close()

	fmt.Printf("Image dimensions: %d x %d\n", img.Cols(), img.Rows())






}

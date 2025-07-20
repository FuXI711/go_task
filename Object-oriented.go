package main

// // 定义 Shape 接口
// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// // Rectangle 结构体
// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.Width * r.Height
// }

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.Width + r.Height)
// }

// // Circle 结构体
// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	return c.Radius * c.Radius
// }

// func (c Circle) Perimeter() float64 {
// 	return 2 * c.Radius
// }

// func main() {
// 	rect := Rectangle{Width: 5, Height: 3}

// 	circle := Circle{Radius: 4}

// 	var shape Shape

// 	shape = rect
// 	fmt.Printf("矩形 - 宽度: %.2f, 高度: %.2f\n", rect.Width, rect.Height)
// 	fmt.Printf("面积: %.2f\n", shape.Area())
// 	fmt.Printf("周长: %.2f\n\n", shape.Perimeter())

// 	shape = circle
// 	fmt.Printf("圆形 - 半径: %.2f\n", circle.Radius)
// 	fmt.Printf("面积: %.2f\n", shape.Area())
// 	fmt.Printf("周长: %.2f\n", shape.Perimeter())
// }

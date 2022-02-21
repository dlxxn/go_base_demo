package examples

import (
	"fmt"
)

/**
To implement an interface in Go, we just need to implement all the methods in the interface. Here we
implement geometry on rects.

go语言中实现接口无需使用implement关键字，只要实现接口中定义的所有方法就是实现了接口
*/
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func GetRect() rect {
	return rect{
		width: 3, height: 5}
}

func Measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

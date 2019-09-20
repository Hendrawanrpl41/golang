package main

import (
	"fmt"
	"math"
)

func main(){

	circle := Circle{x : 1.0, y : 1.0, radius : 10.5}
	rectangle := Rectangle{width : 10, height : 5}

	fmt.Println("Area dari Circle = ", getArea(circle))
	fmt.Println("Area dari Rectangle = ", getArea(rectangle))

}

type Shape interface{
	Area() float64

}

type Circle struct{
	x 		float64
	y 		float64
	radius 	float64
}

func ( c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

type Rectangle struct{
	width 	float64
	height 	float64

}

func (r Rectangle) Area() float64{
	return r.width * r.height
}

func getArea(s Shape)float64{
	return s.Area()
}
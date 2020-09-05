package main


import "fmt"

type (
	shape interface {
		getArea() float64
	}
)

type triangle struct {
	height float64
	base float64
}

type square struct {
	length float64
}


func main() {

	t:= triangle{ 8, 8 }
	s:=square{8}
	printArea(t)
	printArea(s)



}


func (t triangle) getArea() float64{



	 return 0.5 *t.height* t.base

}

func (s square) getArea() float64{



	return s.length * s.length



}

func printArea (s shape){

	fmt.Println(s.getArea())


}
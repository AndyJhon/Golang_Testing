package main

import "fmt"
import "os"
import "math"

type Person struct{
	Code string
	Name string
	Age int
	Average float64
	Sex bool
}

func main() {
	//Structures()
	//FunctionStructure()
	//callMovesPoint()
	///callStarter()
	callSumaAreas()
}

/****************************************************************/
func Structures() {
	type Person struct{
		Code string
		Name string
		Age int
		Average float64
		Sex bool
	}
	
	s1 := Person{ "1", "Andy", 29, 15.5, true}
	fmt.Println(s1)
	fmt.Printf("%+v\n", s1)

	//-----------
	s2 := Person{
		Code: "1", 
		Name: "Pepe", 
		Age: 18, 
		Average: 15.5, 
		Sex: true,
	}

	fmt.Printf("%+v\n", s2)

	//-----------
	s3 := Person{}
	fmt.Printf("%+v\n", s3)
}

/****************************************************************/

func (p *Person) functionValue() float64{
	value := float64(p.Average) * 5.0
	return value
}

func FunctionStructure(){
	s4 := Person{"","", 29, 25.0, true}
	fmt.Println(s4.functionValue())
}

/****************************************************************/
// Why is method struct by reference
type Point struct {
	x int
	y int
}

func (p Point) movesPoint(dx int, dy int) {// * for pointer ference removed
	p.x = dx
	p.y = dy
}

func callMovesPoint(){
	p := Point {1,2} //-->Printed this one
	p.movesPoint(2,1) //--> Instead of this one
	fmt.Printf("%+v\n", p)
}

/****************************************************************/
// New Structs with function 
// To Replace a constructor in order to start a struc
/*
Code string
	Name string
	Age int
	Average float64
	Sex bool
*/
func NewPerson(code string, name string, age int, average float64, sex bool) (*Person, error){
	
	if code == "" {
		return nil, fmt.Errorf("Code can not be empty")
	}

	if age < 30 {
		return nil, fmt.Errorf("Person must be older than 29`s")
	}

	person1 := &Person{
		Code: code, 
		Name: name, 
		Age: age, 
		Average: average, 
		Sex: sex}
	return person1, nil
}

func callStarter(){

	p1, err := NewPerson("1", "Andy", 30, 12.0, true)
	if err != nil {
		fmt.Printf("Can not create a person %+s\n", err)
		os.Exit(1)
	}
	fmt.Println(p1)
}

/****************************************************************/
// Interfaces

type Shape interface {
	getArea() float64
}

type Square struct {
	lenght float64
}

func (s *Square) getArea() float64{
	return s.lenght * s.lenght
}

type Circle struct {
	radious float64
}

func (r *Circle) getArea() float64{
	return math.Pi * r.radious * r.radious
}

//func sum area
func sumAreas(shapes []Shape) float64{
	total := 0.0

	for _, shape := range shapes {
		total += shape.getArea()
	}

	return total
}

func callSumaAreas(){
	s1 := &Square{5.0}
	fmt.Println(s1.getArea())

	c1 := &Circle{5.0}
	fmt.Println(c1.getArea())

	shapes := []Shape{s1, c1}
	sa := sumAreas(shapes)
	fmt.Println(sa)
}
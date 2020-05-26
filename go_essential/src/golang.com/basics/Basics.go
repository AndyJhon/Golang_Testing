package main

import (
	"fmt"
	"strings"
)

func main() {	
	//Numbers()
	//Conditional()
	//Loops()
	//FizBuzz()
	//Strings()
	//Evenended()
	//Slices()
	//MapFunctions()
	MapChallenge()
}

func Numbers() {
	/*
	var x float64
	var y float64
	x = 1
	y = 2
	*/

	/*x := 1.0
	y := 2.0*/

	x, y := 1.0, 2.0

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	/*var mean float64 
	mean = (x+ y) / 2*/
	mean := (x+ y) / 2.0
	
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}

func Conditional(){
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}
	if x > 5 && x < 15 { // || for or conditions
		fmt.Println("x is just right")
	} else {
		fmt.Println("x is not just right")
	}

	a := 11.0
	b := 20.0

	if frac := a/b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	//Switch
	i := 2

	switch i {
	case 1:
		fmt.Println("One")
	
	case 2:
		fmt.Println("Two")
	
	case 3:
		fmt.Println("Three")
	
	default:
		fmt.Println("Many")
	}
	// You can also specify conditional in swith as java
}

func Loops() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	
	fmt.Println("-------")

	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("-------")

	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("-------")
	fmt.Println("While Behaviour")
	a := 0

	for a < 3 {
		fmt.Println(a)
		a++
	}

	fmt.Println("-------")
	fmt.Println("While true")
	b := 0

	for {
		if b > 2 {
			break
		}

		fmt.Println(b)
		b++
	}
}

func FizBuzz() {
	for i := 0; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func Strings() {
	book := "The colour of magic"
	
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Println(book[0])
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	//String in go are inmutable

	//Slice(start,end) 0 based, 1/2 empty range
	fmt.Println(book[4:11])

	//Slice(no end)
	fmt.Println(book[4:])

	//Slice(no start)
	fmt.Println(book[:3])

	//Concatenate
	fmt.Println("Concatenated: " + book[:3])

	//Unicode

	//Multiline
	multi := `
	This is a line
	This is a second line
	Final line
	`
	fmt.Println(multi);

}

func Evenended() {
	// count = 0
	count := 0

	//for every pair of 4 digit numbers

	for a := 1000; a<=9999; a++ {
	
		for b := a; b<=9999; b++ {
			n := a * b

			// check if a * b even ended
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {

				count++
			}
		}		
	}
	fmt.Println(count);
}

func Slices() {
	//Same type
	loons := []string {"bugs","daffy","taz"}
	fmt.Printf("loons =%v (type %T)\n", loons, loons)
	
	//Length
	fmt.Println(len(loons))//3

	//0 indexing
	fmt.Println(loons[1])// daffy

	//Slices
	fmt.Println(loons[1:])// daffy taz

	//Iterate 
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	// Iterate 2: Single value range
	for i := range loons {
		fmt.Println(i)
	}	

	// Iterate 2: Double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	// Iterate 2: Double value range ingnoring index by using _
	for _, name := range loons {
		fmt.Println(name)
	}

	//Append
	loons = append(loons, "elmer")
	fmt.Println(loons)

	//Max value
	nums := []int{16, 15,25,4,96,12,100}

	_max := nums[0] // Initialize max with the first value

	for _, value := range nums[1:]{
		if value > _max {
			_max = value
		}
	}

	fmt.Println(_max)
}

func MapFunctions(){

	stock := map[string]float64{
		"UNO": 1.0,
		"DOS": 2.0,
		"TRES": 3.0,
	}

	//Number of Items
	fmt.Println(len(stock))

	//Get a value
	fmt.Println(stock["DOS"])

	//Get 0 if it is not found
	fmt.Println(stock["CUATRO"])

	//Use two value to see if found
	value, ok := stock["CUATRO"]
	if !ok {
		fmt.Println("CUATRO not found")
	}else {
		fmt.Println(value)
	}

	//Set
	stock["DOS"] = 22.0
	fmt.Println(stock["DOS"])

	//Delete
	delete(stock, "DOS")
	fmt.Println(stock)

	//Add
	stock["DOS"] = 222.0

	fmt.Println("------------------------------")
	fmt.Println("Single values for is on keys")
	for key := range stock {
		fmt.Println(key)
	}

	fmt.Println("------------------------------")
	fmt.Println("Double value for is key, value")
	for key,value := range stock {
		fmt.Printf("%s -> %.2f\n", key, value)
	}

}

func MapChallenge(){
	// Count words
	text := `
	Rapto
	Deja Vu
	Tu cicatriz en mi Rapto
	De musica ligera
	`
	words := strings.Fields(text)
	counts := map[string]int{} //Empty map
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)
}
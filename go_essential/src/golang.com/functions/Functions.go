package main

import (
	"fmt"
	"math"
	"net/http"
)

func main(){
	//basicFunctions()
	//parameterPassing()
	//ErrorReturn()
	//Defer()
	DeferChallenge()
}

func add(a int, b int) int{
	return a + b
}

func divmod(a int, b int) (int, int) {
return a /b , a % b
}

func basicFunctions(){
val := add(1, 2)
fmt.Println(val);

div, mod := divmod(7, 2)
fmt.Printf("div=%d, mod=%d\n", div, mod)
}

func doubleAt(values []int, i int) {
	values[i] *=2
}

func double (n int){
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func parameterPassing() {
	values := []int{1,2,3,4}
	doubleAt(values, 2)
	fmt.Println(values)

	val1 := 10
	double(val1)
	fmt.Println(val1)

	val2 := 10
	doublePtr(&val2)
	fmt.Println(val2)
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("Sqrt of negative value (%f)", n)
	}
	return math.Sqrt(n), nil
}

func ErrorReturn() {
	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s1)
	}

	s2, err := sqrt(-2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s2)
	}
}

func cleanup(name string) {
	fmt.Printf("Clean up %s\n", name)
}

func worker() {
	defer cleanup("A")
	fmt.Println("worker")
}

func Defer(){
	worker()
}

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() // Make sure we close the body
	
	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("cannot find Content-Type header")
	}
	return ctype, nil
}

func DeferChallenge(){
	ctype, err := contentType("https://google.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}
package main

import "fmt"
import "rsc.io/quote"
import "math"
import "math/rand"

var c, python, java bool
var d, cython, javascript = false, true, 1.23

func main() {
	fmt.Println("Hello World")
	quote_text := quote.Go()
	fmt.Println(quote_text)
	demo()

	for i := 0; i < 5; i++ {
		fmt.Printf("Counter: %d\n", i)
	}

	out := math.Max(1e3, 3)
	out2 := rand.Intn(3) // rand.Float32()
	fmt.Println("Max: ", out, "..", out2*2, "PI", math.Pi, add(2, 32))

	// Primitive types
	var i, j int8 = 10, 20

println()
	fmt.Println(i, c, python, java)
	fmt.Println(d, cython, javascript)
  print(j, "\n")

  var f float32 = 12.34
  println(f)
  fmt.Printf("Val: %v, %T\n", f, f)
  fmt.Printf("Converting to uint..\nVal: %v, %T\n", uint(f), uint(f))



}

// func add(x int, y int) (out1 int) {
func add(x, y int) (out1 int) {
	out1 = x + y
	return
}

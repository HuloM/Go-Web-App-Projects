package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"syscall/js"
)
func calcSimple(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	value3 := js.Global().Get("document").Call("getElementById", i[2].String()).Get("value").String()

	int1, _ := strconv.ParseFloat(value1, 64)
	int2, _ := strconv.ParseFloat(value3, 64)

	var answer float64
	switch {
		case value2 == "+":
			answer = int1 + int2
		case value2 == "-":
			answer = int1 - int2
		case value2 == "/":
			answer = int1 / int2
		case value2 == "*":
			answer = int1 * int2
		case value2 == "%":
			answer = math.Mod(int1, int2)
		default:
			;
	}

	js.Global().Get("document").Call("getElementById", i[3].String()).Set("value", answer)
	fmt.Printf("%s %s %s = ", value1, value2, value3)
	fmt.Print(answer, "\n")
	return js.ValueOf(answer)
}
func calcTrig(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	value3 := js.Global().Get("document").Call("getElementById", i[2].String()).Get("value").String()

	int1, _ := strconv.ParseFloat(value1, 64)

	var answer float64
	switch {
		case value2 == "sin":
			answer = math.Sin(int1)
		case value2 == "cos":
			answer = math.Cos(int1)
		case value2 == "tan":
			answer = math.Tan(int1)
		case value2 == "asin":
			answer = math.Asin(int1)
		case value2 == "acos":
			answer = math.Acos(int1)
		case value2 == "atan":
			answer = math.Atan(int1)
		default:
			;
	}

	if  strings.Contains(strings.ToLower(value3), "deg"){
		answer *= (180.0/math.Pi)
	}

	js.Global().Get("document").Call("getElementById", i[3].String()).Set("value", answer)
	fmt.Printf("%s(%s) =  " ,value2, value1)
	fmt.Print(answer, "\n")
	return js.ValueOf(answer)
}

func registerCallbacks() {
	js.Global().Set("calcSimple", js.FuncOf(calcSimple))
	js.Global().Set("calcTrig", js.FuncOf(calcTrig))
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("Go Initialized")
	registerCallbacks()

	<-c
}

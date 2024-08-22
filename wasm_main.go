package main

import (
	"fmt"
	add "go_project_template/math"
	"strconv"
	"syscall/js"
)

// Gets a HTML element given an ID.
func getElementById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

// Gets a numeric value from an HTML element body given an ID.
func getNumberFromElement(id string) (int, error) {
	element := getElementById(id)
	valueStr := element.Get("value").String()
	return strconv.Atoi(valueStr)
}

// Adds the two input elements and sets the sum to another element.
func addAndSetResult() {
	sumNode := getElementById("sum")
	number1, err := getNumberFromElement("number1")
	if err != nil {
		sumNode.Set("innerHTML", err.Error())
		fmt.Println(err)
		return
	}
	number2, err := getNumberFromElement("number2")
	if err != nil {
		sumNode.Set("innerHTML", err.Error())
		fmt.Println(err)
		return
	}
	sumNode.Set("innerHTML", add.Add(number1, number2))
}

// Start of execution.
func main() {
	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Call("setAttribute", "id", "sum")
	p.Set("innerHTML", "Enter the two numbers to get the sum.")
	document.Get("body").Call("appendChild", p)
	println("Hello, JS console")
	number1Input := document.Call("getElementById", "number1")
	number2Input := document.Call("getElementById", "number2")
	fmt.Println("number 2:", number2Input.Get("value"))
	var callBack js.Func = js.FuncOf(func(this js.Value, args []js.Value) any {
		addAndSetResult()
		return nil
	})
	number1Input.Call("addEventListener", "change", callBack)
	number2Input.Call("addEventListener", "change", callBack)
	select {}
}

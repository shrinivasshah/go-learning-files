package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()
	print("smth else")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("not messatsu!")
		}
		fmt.Println(i)
	}
}

func foo() {
	fmt.Println("hello i am in foo")
}

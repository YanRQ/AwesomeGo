package main

import "fmt"

func swap2(x *int, y *int){
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main(){
    fmt.Println("Hello, World!")
}

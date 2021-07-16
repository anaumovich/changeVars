package main

import (
	"fmt"
	"strconv"
)

var a  = 11
var b = 45
var c int

func ChangeVars1( a int,b int)(string string){
	c = a
	a = b
	b = c
	return "Число А = " + strconv.Itoa(a) + "\nЧисло B = " + strconv.Itoa(b)
}
func ChangeVars2( a int,b int) (string string){
	a = a + b
	b = b - a
	b = -b
	a = a - b
	return "Число А = " + strconv.Itoa(a) + "\nЧисло B = " + strconv.Itoa(b)
}
func ChangeVars3( a int,b int)(string string){
	a,b = b,a
	return "Число А = " + strconv.Itoa(a) + "\nЧисло B = " + strconv.Itoa(b)
}
func ChangeVars4( a int,b int)(string string){
	a = a * b
	b = a / b
	a = a / b
	return "Число А = " + strconv.Itoa(a) + "\nЧисло B = " + strconv.Itoa(b)
}

func main()  {
	fmt.Print("Вариант 1\n" + ChangeVars1(a,b))
	fmt.Print("\nВариант 2\n" + ChangeVars2(a,b))
	fmt.Print("\nВариант 3\n" + ChangeVars3(a,b))
	fmt.Print("\nВариант 4\n" + ChangeVars4(a,b))
}
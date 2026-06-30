package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myInt int = 777
	var myFloat float64 = 777.7
	var myString string = "sevensevenseven"
	var myBool bool = true
	var mybyte byte = 255
	var myRune rune = 7
	fmt.Printf("%d %f %s %t %c %c", myInt, myFloat, myString, myBool, mybyte, myRune)
	const (
		Pi       = 3.14159
		AppName  = "MyApp"
		Maxusers = 100
	)
	fmt.Println(Pi, AppName, Maxusers)
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	name := "Bekzhan"
	age := 22
	balance := 42500
	fmt.Println(name, age, balance)
	a := float64(777)
	b := int(myFloat)
	c := strconv.Itoa(777)
	fmt.Println(a, b, c)
	type timeMng struct {
		nameTask   string
		timeHour   float64
		status     bool
		user       string
		checked    bool
		codeReview bool
		meeting    bool
		testing    bool
	}
	Task1 := timeMng{
		nameTask: "CLI PROJECT",
		timeHour: float64(3),
		status:   true,
		user:     "Bekzhan",
		checked:  false,
	}
	fmt.Println(Task1)
	Task2 := timeMng{
		codeReview: true,
		meeting:    true,
		testing:    true,
	}
	fmt.Println(Task2)
	Task3 := timeMng{
		nameTask:   "CLI PROJECT",
		timeHour:   float64(3),
		status:     false,
		user:       "Bekzhan",
		checked:    false,
		codeReview: true,
		meeting:    false,
		testing:    true,
	}
	fmt.Println(Task3)
}

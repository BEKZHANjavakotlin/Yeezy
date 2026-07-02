package main

import (
	"errors"
	"fmt"
	"week1/models"
)

func main() {
	/* var myInt int = 777
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
	firstRequest := add(2, 5)
	fmt.Println(firstRequest)
	secondRequest := greet("Bekzhan")
	fmt.Println(secondRequest)
	thirdRequest := isAdult(22)
	fmt.Println(thirdRequest)
	firstDivide, err := divide(8, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(firstDivide)
	secondDivide, err := divide(8, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(secondDivide)

	min, max := minMax([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(min, max)
	area, perimeter := circleStats(5.0)
	fmt.Println(area, perimeter)
	convert, err := parsenAndDouble("21")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(convert)
	convert2, err2 := parsenAndDouble("abc")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(convert2)



	hours := [5]int{1, 2, 3, 4, 5}
	fmt.Println(hours)
	fmt.Println(hours[2])
	fmt.Println(len(hours))
	tasks := make([]string, 0, 3)
	tasks = append(tasks, "утреннее совещание", "code review", "деплой")
	fmt.Println(tasks, len(tasks), cap(tasks))
	backup := make([]string, len(tasks))
	copy(backup, tasks)
	fmt.Println(backup)
	fmt.Println(backup[:2])
	var total int
	for index, _ := range tasks {
		fmt.Println(tasks[index])
		total += len(tasks[index])
		fmt.Println(total)
	}
	start := time.Now()
	fmt.Println(start)
	time.Sleep(time.Second)
	elapsed := time.Since(start)
	fmt.Println(elapsed)*/

	hoursPerTask := map[string]int{
		"разработка": 5,
		"ревью":      2,
		"митинг":     1,
	}
	fmt.Println(hoursPerTask)
	delete(hoursPerTask, "митинг")
	fmt.Println(hoursPerTask)
	value1, ok1 := hoursPerTask["разработка"]
	if ok1 {
		fmt.Println("найдено:", value1)
	} else {
		fmt.Println("Ключ не найден")
	}
	value2, ok2 := hoursPerTask["митинг"]
	if ok2 {
		fmt.Println("найдено:", value2)
	} else {
		fmt.Println("Ключ не найден")
	}
	t := models.Task{Title: "разработка", Hours: 3.5}
	fmt.Println(t.Summary())
	fmt.Println(t.Done)
	t.Complete()
	fmt.Println(t.Complete())
	models.Reset(&t)
	entry := models.NewTimeEntry("разработка")
	fmt.Println(entry.Task)
	fmt.Println(getTaskStatus(1.0))
	fmt.Println(getTaskStatus(4.0))
	fmt.Println(getTaskStatus(8.0))
	fmt.Println(getDayType("пн"))
	fmt.Println(getDayType("вс"))
	fmt.Println(getDayType("день"))
	tasks := []string{"разработка", "ревью", "митинг", "деплой"}
	for index, task := range tasks {
		fmt.Println(index+1, task)
	}
	fmt.Println(validateHours(1.0))
	fmt.Println(validateHours(45.9))
	err := validateHours(-1.0)
	if errors.Is(err, ErrNegativeHours) {
		fmt.Println("поймали отрицательные часы")
	}
	openReport()

}

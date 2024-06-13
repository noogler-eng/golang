package main

// automatically import packages from modules
import "fmt"

func main() {

	x := 5
	if x > 4 {
		fmt.Println("value of x is greater then 4")
	} else if x == 5 {
		fmt.Println("value of x is equal to 4")
	} else {
		fmt.Println("value of x is less then 4")
	}

	a := 2
	b := 3
	c := 1
	if a > b && a > c {
		fmt.Println("a is greatest")
	}
	if b > a && b > c {
		fmt.Println("b is greatest")
	}
	if c > a && c > b {
		fmt.Println("c is greatest")
	}

	// switch statmnet
	day := "Monday"

	switch day {
	case "monday", "Monday":
		fmt.Println("MON")
	case "tuesday":
		fmt.Println("TUE")
	case "wednesday":
		fmt.Println("WED")
	case "thrusday":
		fmt.Println("THU")
	case "friday":
		fmt.Println("FRI")
	case "saturday":
		fmt.Println("SAT")
	default:
		fmt.Println("Unknown")
	}

	// using expression in switch cases
	temperature := 25
	switch {
	case temperature < 25:
		fmt.Println("it cold here!")
	case temperature == 25:
		fmt.Println("normal temperature")
	case temperature > 25:
		fmt.Println("it hot here!")
	}

	// there is only form of loop in go lang -> for
	for i := 0; i < 10; i++ {
		if i == 7 {
			break
		}
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// infinte loop
	counter := 1
	for counter > 0 {
		fmt.Println("loop running")
		if counter == 4 {
			break
		}
		counter++
	}

	numbers := []int{1, 2, 6, 7, 8}
	// range simplifier the numbers array into index and value
	for index, value := range numbers {
		fmt.Printf("index and value are %d and %v \n", index, value)
	}

	data := "Hello World"
	for index, character := range data {
		// single character comes in ''
		if character == ' ' {
			continue
		}
		fmt.Printf("index and character are %d and %c \n", index, character)
	}

	studentGrades := make(map[string]int)
	studentGrades["sharad"] = 100
	studentGrades["poddar"] = 80
	studentGrades["happy"] = 60

	fmt.Println("🎯", studentGrades["happy"])
	studentGrades["happy"] = 100
	fmt.Println("🎯", studentGrades["happy"])

	// print orderwise
	fmt.Println("🎯", studentGrades)
	delete(studentGrades, "happy")
	fmt.Println("🎯", studentGrades)
	// it will show 0
	fmt.Println("🎯", studentGrades["happy"])

	// checking of the key exists ot not
	grade, isExist := studentGrades["happy"]
	fmt.Printf("grade of happy %d and exist or not %v \n", grade, isExist)
	fmt.Println("is happy exist or not:", isExist)

	// index is key and grades is value
	for index, value := range studentGrades {
		fmt.Println("🎯", index, value)
	}

}

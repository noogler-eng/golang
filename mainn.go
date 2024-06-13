package main

// go uses relative import
import (
	"another/another"
	"bufio"
	"fmt"
	"os"
)

// function
func sum(a int, b int) int {
	return a + b
}

// declarin all a, b to int, short form
func sumAdv(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("dominator must not be 0")
	}
	return a / b, nil
}

func main() {

	// we can use var, const same as of js
	var name string = "sharad poddar"
	var version int = 1
	var currency int = 34567
	var distance float64 = 12.234234
	var isDecide bool = false

	// can be used as neither var or const
	// initlially it setted
	person := "tim david"

	fmt.Println("learning golang")
	fmt.Println("print anything .. 🌈")

	another.PrintMsg()
	fmt.Println(name)
	fmt.Println(version)
	fmt.Println("currency: ", currency)
	fmt.Println("dimmension: ", distance)
	fmt.Println("isDeicded: ", isDecide)

	fmt.Println("new person in house: ", person)
	person = "malinga"
	fmt.Println("new person in house: ", person)

	// if you want to export something, letter must be capital
	// format spcifier
	fmt.Printf("my name is %s\n", person)
	fmt.Printf("my data is %.2f\n", distance)
	fmt.Printf("my data type is %T\n", distance)

	// it only take single word until space, scan reads until whitespace
	fmt.Println("what is your name?")
	var fullName string
	fmt.Scan(&fullName)
	fmt.Println(fullName)

	// how we read in golang
	fmt.Println("what is your name?")
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println(os.Stdin)
	// read string until '\n'
	// giving us an input value and error if exists
	whatName, _ := reader.ReadString('\n')
	fmt.Println(whatName)

	ans := sum(2, 3)
	fmt.Println("👀 calling the function", ans)

	ans = sumAdv(2, 3)
	fmt.Println("👀 calling the function", ans)

	// error handling
	// _ is used for discarding the error
	new_ans, err := divide(10, 0)
	fmt.Println("👀 calling the function", new_ans)
	fmt.Println(err)

	// here all elemnts initilized to 0 initially
	// spaces are allocated
	var nums [5]int
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	fmt.Println("array elements", nums)

	corals := [4]string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}
	fmt.Println(corals)
	// printing full aray by %q
	// %q means quoted string
	fmt.Printf("%q\n", corals)

	// finding length of an array
	length := len(corals)
	fmt.Println(length)

	price := [5]int{}
	fmt.Println(price)

	new_price := [5]string{}
	fmt.Println(new_price)
	fmt.Printf("%q\n", new_price)

	// for dynamic array we are using slice
	// here we can append the data
	// slice underlying data structure array in RAM
	// it has pointer, len, capacity
	// here also length becomes double when it will full
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(cap(numbers))

	numbers = append(numbers, 3, 10, 13, 15)
	fmt.Println(numbers)
	fmt.Printf("%T\n", numbers);
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	// empty slice
	cities := []int{};
	fmt.Println(cities)
	fmt.Printf("%T\n", cities);
	fmt.Println(len(cities))
	fmt.Println(cap(cities))

	// using make to set capacites and length
	new_cities := make([]int, 3, 5)
	new_cities = append(new_cities, 10, 20, 30)
	fmt.Println("✅", new_cities)
	fmt.Printf("%T\n", new_cities);
	fmt.Println(len(new_cities))
	fmt.Println(cap(new_cities))

}

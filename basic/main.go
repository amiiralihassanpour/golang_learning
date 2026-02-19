package main

import "fmt"


func main() {
	
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to Go programming,", "Let's learn Go together.")

	var name string = "Alice"
	fmt.Println("My name is", name)

	age := 30
	fmt.Println("I am", age, "years old.")
	
	var a, b int = 5, 10
	fmt.Println("The sum of", a, "and", b, "is", a+b)

	var x, y = 1.5, "Go"
	fmt.Println("The value of x is", x, "and the value of y is", y)

	const pi = 3.14
	fmt.Println("The value of pi is", pi)

	// --------------------------------------------------------------------

	for i:=0; i<5; i++ {
		fmt.Println("Iteration:", i)
	}

	// --------------------------------------------------------------------

	for i:=1; i<=5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	if name:= "Alice"; name == "Alice" {
		fmt.Println("Hello, Alice!")
	} else {
		fmt.Println("Hello, stranger!")
	}

	key := 3

	switch key {
	case 1:
		fmt.Println("key is 1")
	case 2:
		fmt.Println("key is 2")
	case 3: 
		fmt.Println("key is 3")
	default:
		fmt.Println("key is not in range [1,3]")		
	}

	// --------------------------------------------------------------------

	var nums [5]int
	nums[1] = 20
	fmt.Println(nums)
	
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	s := make([]int, 3, 4)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	fmt.Println("Slice:", s, "Length:", len(s), "Capacity:", cap(s))
	fmt.Printf("%p\n", s)
	fmt.Println("Address of slice:", &s[0])

	s = append(s, 40)
	s = append(s, 50)
	fmt.Println("Slice after appending:", s, "Length:", len(s), "Capacity:", cap(s))
	fmt.Printf("%p\n", s)
	fmt.Println("Address of slice after appending:", &s[0])

	slice := []string{"Go", "Python", "Java"}
	fmt.Println("Slice:", slice)

	mapping := map[string]int{"Alice": 30, "Bob": 25}
	fmt.Println("Map:", mapping)

	var TwoDArray [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Two-dimensional array:", TwoDArray)

}
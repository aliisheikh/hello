package main

import (
	"fmt"
	"sort"
)

func greetingMor(n string) {
	fmt.Printf("Good Morning %v \n", n)
}

func greetingEve(n string) {
	fmt.Printf("Good Evening %v \n", n)
}

func main() {

	greetingMor("Ali Sohail")
	greetingEve("Umer Sohail")

	age := 30
	name := " Ali Sohail "

	fmt.Println("Hello Sir!")

	var name1 string = "Ali sohail"
	fmt.Println(name1)

	var name2 string = " Umer Sohail"
	fmt.Println(name2)

	name3 := "bass"
	fmt.Println(name3)

	fmt.Println(name1, name2, name3)

	//	integers

	age_1 := 10
	age_2 := 20
	age_3 := 30
	var age4 int16 = 3456
	fmt.Println(age_1, age_2, age_3, age4)

	fmt.Print("cheque \n")
	fmt.Print(" new line \n")
	//fmt.Println("My name is ", name, " My age is: ", age)
	fmt.Printf("My name is %v and My age is %v", name, age)
	fmt.Println("\n")

	var str = fmt.Sprintf("My name is %v and My age is %v", name, age)
	fmt.Println("The saved string are: ", str)

	//	Arrays:
	var ages = [3]int{23, 45, 67}
	fmt.Println(ages, len(ages))

	scores := []int{100, 20, 45}
	scores[2] = 124
	scores = append(scores, 21)
	fmt.Println(scores)

	var listt = []int{12, 34, 567, 78, 54}
	sort.Ints(listt)
	fmt.Println(listt)
	//
	//names := []string{"josh", "marlin", "pack", "asad", "Ali"}
	//sort.Strings(names)
	//fmt.Println(names)
	//
	//fmt.Println(sort.SearchStrings(names, "pack"))

	//	For loop
	names := []string{"josh", "marlin", "pack", "asad", "Ali"}

	for i := 0; i < len(names); i++ {

		fmt.Println(names[i])
	}

	for _, value := range names {
		fmt.Printf("The Value is %v \n", value)

	}
}

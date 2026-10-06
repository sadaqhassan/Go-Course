package main

import (
	"fmt"
	
)

//main function Go



func main()  {


	// fmt.Println("hello world")

	// //variables 
	// var myname string = "";
	// age := 20;
	// fmt.Println(age); 
	// fmt.Println(myname); 

	// //conditions
	// if age > 15 {
	// 	fmt.Println("You are adult")
	// 	return;
	// }else{
	// 	fmt.Print("you are young")
	// 	return;
	// }


	//loops
	// for i:= 0 ; i < 10 ; i++ {
	// 	fmt.Println(i)
	// }

	// infiniteLoop
	// for {
	// 	fmt.Println("hello")
	// }

	//range
	// names := [] string{};

	// for index , name := range names {
	// 	fmt.Println(index,name)
	// }

	// Arrays


	numbers := [6] int {1,3,4,5,2,5}


	for i:= 0 ; i <= len(numbers); i++{
		fmt.Println("waa socotaa...")
		if i == len(numbers) {
			fmt.Println("waa dhamaty ")
			return
		}
	}


}

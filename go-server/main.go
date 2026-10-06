package main

import (
	"fmt"
	
)

//main function Go

func main()  {
	// fmt.Println("hello world")

	// //variables 
	// var myname string = "sadak";
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

	names := [] string{"sadak","ikhro","abdimalik","hooyoHaawo","aabo" ,"fadumo"};

	for index , name := range names {
		fmt.Println(index,name)
	}


}

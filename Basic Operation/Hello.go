package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main (){
	fmt.Println("aman manwani is the best");

	// variables
	var username string= "aman" 
	fmt.Println(username)
	// %t is for type
	fmt.Printf("variable is of type : %T \n",username);
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn);
	//can be defined only in range
	var smallVal uint = 255
	fmt.Println(smallVal)
	// you can only initialise no need to put value in it 
	// it is set to be default 0 and empty string in the case of int and string
	var anotherVar int
	fmt.Println(anotherVar);
	var anotherString string
	fmt.Println(anotherString) 

	// implicit type
	// it is also valid but if you dont specify the variable type it wil decide automatically like string below
	var name = "aman";
	fmt.Println(name);

	//no var type 
	// you are not allowed to declare this type of varibles globally 
	// means you can only initialise them in 
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)

	//public variable declaration 
	// first letter uppercase rakhna h 
	const Login string = "abhgyt"

	welcome := "welcome to aman's world"
	fmt.Println((welcome))

	// input taking in golang
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the value b/w 1 & 5 : ")

	// comma ok syntax || error ok syntax
	// this the notation like a try catch
	// if it goes right, then value is in input or if error comes, then it will be in errvariable instead of _
	// input,err := reader.ReadString('\n');
	input,_ := reader.ReadString('\n');
	fmt.Println("your input value is :",input)

	// input is of type string and operation needs to be int 
	// numRating := input + 1;
	//below is the operation to convert strings to int type
	// to do this, we use strconv
	// we use strings.trimSpace becoz input is input\n so we have to remove \n from end , to do this we use string.trimSpace()
	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if(err!=nil){
		fmt.Println(err)
	}
	fmt.Println(numRating);

	
}
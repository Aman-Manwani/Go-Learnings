package main

import (
	"bufio"
	"math/rand"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("aman manwani is the best")

	// variables
	var username string = "aman"
	fmt.Println(username)
	// %t is for type
	fmt.Printf("variable is of type : %T \n", username)
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	//can be defined only in range
	var smallVal uint = 255
	fmt.Println(smallVal)
	// you can only initialise no need to put value in it
	// it is set to be default 0 and empty string in the case of int and string
	var anotherVar int
	fmt.Println(anotherVar)
	var anotherString string
	fmt.Println(anotherString)

	// implicit type
	// it is also valid but if you dont specify the variable type it wil decide automatically like string below
	var name = "aman"
	fmt.Println(name)

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
	input, _ := reader.ReadString('\n')
	fmt.Println("your input value is :", input)

	// input is of type string and operation needs to be int
	// numRating := input + 1;
	//below is the operation to convert strings to int type
	// to do this, we use strconv
	// we use strings.trimSpace becoz input is input\n so we have to remove \n from end , to do this we use string.trimSpace()
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(numRating)

	// time initialization
	presentTime := time.Now()
	fmt.Println(presentTime)

	// to get time in a desired format use like this
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)

	// Pointers
	var ptr *int
	// gives nil
	fmt.Println(ptr)

	myNumber := 23
	var ptr2 = &myNumber
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	*ptr2 = *ptr2 + 2
	//gives 25
	fmt.Println(myNumber)

	// Arrays
	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[1] = "banana"
	fruitlist[2] = "cucumber"

	fmt.Println(fruitlist)
	// gives len 4 but we have only 3 values
	fmt.Println(len(fruitlist))
	var veglist = [5]string{"potato", "beans", "mouse"}
	fmt.Println(len(veglist))

	// slices
	var fruitlistt = []string{"apple", "tomato", "Peach"}
	fmt.Printf("%T", fruitlistt)

	fruitlistt = append(fruitlistt, "mango", "banana")
	fmt.Println(fruitlistt)
	// from index 1 to 3 we got the values that's why iits called slices
	fruitlistt = append(fruitlistt[1:3])
	fmt.Println(fruitlistt)

	highScore := make([]int, 4)
	highScore[0] = 1
	highScore[1] = 2
	highScore[2] = 3
	highScore[3] = 4
	// you cant do below operation as it has only space to store 4 values
	// highScore[4] = 5;
	// but you can use append method to do this operation
	highScore = append(highScore, 555, 345)
	fmt.Println(highScore)
	// to sort it
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	// maps in golang
	languages := make(map[string]string)
	languages["aman"] = "val"
	languages["ama"] = "vl"
	languages["amana"] = "vala"
	languages["amanaa"] = "valaa"
	fmt.Println(languages)
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v", key, value)
	}

	//structs properties
	aman := User{"aman","aman...gmail.com",true,23}
	fmt.Println(aman)

	// if else
	loginCount := 23
	var result string
	if(loginCount<10){
		result = "regular user"
	}else if loginCount>10 {
		result = "nice user"
	}else{
		result = "something else"
	}
	fmt.Println(result)
	if num := 3 ; num < 10 {
		fmt.Println("number is less than 10")
	} else{
		fmt.Println("number is greater or equal to 10")
	}

	// switch cases 
	rand.Seed(time.Now().UnixNano());
	diceNumber := rand.Intn((6))+1
	fmt.Println(diceNumber);

	switch diceNumber{
		case 1:
			fmt.Println("value is 1")
		case 2:
			fmt.Println("value is 1")
		case 3:
			fmt.Println("value is 1")
		case 4:
			fmt.Println("value is 1")
		case 5:
			fmt.Println("value is 1")
		case 6:
			fmt.Println("value is 1")
	}
	// it has break automatically unlike other programming languages
	// to continue further statements use fallthrough

	// functions in golang
	days := []string{"monday","tuesday","wednesday","thursday"}
	fmt.Println(days)

	for d:= 0; d<len(days); d++{
		fmt.Println(days[d]);
	}

	for i:= range days{
		fmt.Println(days[i]);
	}

	for index,day := range days {
		fmt.Printf("index is %v & day is %v",index,day);
	}

	temp := 0
	for(temp<10){
		if(temp == 5){
			break;
		}
		if(temp == 3){
			temp ++;
			continue;
		}
		temp++;
		fmt.Println(temp);
	}

	// functions in Golang
	// functions are just same defining as JS but we also have to write the datatype of return type in last also
	resultant := adder(3,5);
	fmt.Println(resultant);
	
	resultant = proAdder(2,4,5,7,9);
	fmt.Println(resultant);
	// we can return more values from a function as following 

	tempp,tempppp := proAdder2(2,5)
	fmt.Printf("%v  \t  %v",tempp, tempppp)


	














}

// structs in golang
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func adder (val1 int, val2 int) int {
	return val1 + val2;
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values{
		total += val;
	}
	return total
}
func proAdder2(values ...int) (int,string) {
	total := 0
	for _, val := range values{
		total += val;
	}
	return total,"aman"
}
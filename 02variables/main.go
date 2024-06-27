package main

import "fmt"

// jwtToken := 100 this is not allowed outside function

//constants capital L makes it public
const LoginToken string = "ddddd"

func main() {
	fmt.Println("variables")
	var username string = "paridhi"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint32 = 2332
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float32 = 244.988888888888888888888888
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	//default values and some aliasses
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type : %T \n", anotherString)

	//implicit type
	var website = "test"
	fmt.Println(website)

	// website = 3 this wont work

	// no var style
	numberOfUser := 100
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}

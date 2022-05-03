package main

import ("fmt"
        "os"
        "strconv")

// In this exercise we are going to demostrate how to implement a function that return multiple value
func main(){
  // This is a simple example that return two strings value
  fmt.Println("\n|------------- Multi return ------------------|")
  value1, value2 := multiReturns()
  fmt.Println("Value 1: "+value1,"\nValue 2: "+ value2)
  fmt.Println("|---------------------------------------------|")

   // This section handles argument that are being passed
   // In order to execute this section enter the following task:
   // 1. Navigate to 'multipleReturn' folder through terminal and enter the following command: 'go run main.go 10'
   argument := os.Args
   if len(argument) != 2 {
     fmt.Println("The program need atleast 1 argument to perform the argument calculation")
     return
   }

   value, err := strconv.Atoi(argument[1]) //Retrieve argument value on index 1 (index 0 is the path to the program)
   if err != nil {  // In case it failed to retrieve argument catch error
     fmt.Println(err) // Error message
     return
   }

   // This example demostrate a simple example of a function two int values on which one is the value that is being passed and the other one will be multiply by 10
   fmt.Println("\n|------------- Multi return int ------------------|")
   valuePassed, valueMultiply := multiplyValue(value)

   valueConvertedToString := strconv.Itoa(valuePassed)
   multipliedValueConvertedToString := strconv.Itoa(valueMultiply)

   fmt.Println("Argument value (int): "+valueConvertedToString)
   fmt.Println("Multiply argument (int): "+multipliedValueConvertedToString)
   fmt.Println("|-------------------------------------------------|")

   // This example demostrate how to handle multiple return values of different types
   fmt.Println("\n|------------ Multi return with multiple types --------------------|")
   passedValue, passedValueMultipliedBy10, taskResult := multiplyValueTypes(value)
   valuePassedConvertedToString := strconv.Itoa(passedValue)
   valueMultiplyConvertedToString := strconv.Itoa(passedValueMultipliedBy10)

   fmt.Println("Passed (int) value : "+valuePassedConvertedToString)
   fmt.Println("Passed multiply (int) value : "+valueMultiplyConvertedToString)
   fmt.Println("Task result (string): "+taskResult)
   fmt.Println("|------------------------------------------------------------------|")
}

// Return multiple string values
func multiReturns() (value1 string, value2 string) {
  value1 = "return string value 1"
  value2 = "return string value 2"
  return
}

// Multiply the int value that is passed by 10
func multiplyValue(value int) (valuePassed int, multiplyValue int) {
  valuePassed = value
  multiplyValue = value * 10
  return
}

// Multiply the int value and return multiple values of different type
func multiplyValueTypes(value int) (int, int, string) {
  return value, value * 10, "calculated"
}

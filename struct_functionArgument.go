package main

import "fmt"

type Person struct {
	Name string
	Age int 
	subject string
	marks int
}

func struckAsFunctionArgument () {
fmt.Println("this code is to use struch as function argument")
person := []Person{
	{Name : "pranto",
     Age: 25,
	subject: "CSE",
    marks: 90,         },
	{
      Name: "shovon",
     Age :30,
	 subject: "EEE",
	 marks: 80,
     },}
	 for _,persons := range person {
        printFunction(persons)
	 }

}


func printFunction (p Person){
  fmt.Println(p)
  
}
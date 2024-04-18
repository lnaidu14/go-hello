package main

import (
	"fmt"
	"hello/helper"
	"sync"
	"time"
)	

type UserData struct {
	firstName string
	age int
	friendsNames []helper.FriendData
}

var wg = sync.WaitGroup{}

func main() {
	var firstName string
	var age int
	var numOfFriends int
	var allUsers []UserData
	var totalNumOfUsers int
	fmt.Printf("Enter the total number of users: ")
	fmt.Scan(&totalNumOfUsers)
	for {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your age: ")
		fmt.Scan(&age)
		fmt.Println("How many friends do you have? ")
		fmt.Scan(&numOfFriends)
		
		fmt.Printf("My name is %v! \n", firstName)
		fmt.Printf("I am %v years old \n", age)

		friendsNames := helper.GetFriends(numOfFriends)
		var userData = UserData {
			firstName: firstName,
			age: age,
			friendsNames: friendsNames,
		}
		allUsers = append(allUsers, userData)
		wg.Add(1)
		go printUsers(allUsers)
		if len(allUsers) == totalNumOfUsers {
			break
		}
	}
	wg.Wait()
}

func printUsers(allUsers []UserData) {
	time.Sleep(20 * time.Second)
	fmt.Println("################################")
	fmt.Printf("All users: %v \n", allUsers )
	fmt.Println("################################")
	wg.Done()
}
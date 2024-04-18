package helper

import "fmt"

type FriendData struct {
	firstName string
	age int
}

func GetFriends(numOfFriends int) ([]FriendData) {
	var friendsName string
	friendsNames := make([]FriendData,0)
	var friendsAge int
	for {
		fmt.Printf("Enter your friends first name: ")
		fmt.Scan(&friendsName)
		fmt.Printf("Enter your friends age: ")
		fmt.Scan(&friendsAge)
		var friendData = FriendData {
			firstName:friendsName,
			age:friendsAge,
		}
		friendsNames = append(friendsNames, friendData)
		fmt.Println(friendsNames)
		if len(friendsNames) == numOfFriends {
			fmt.Println("Max number of friends entered")
			return friendsNames
		}
	}
	
}
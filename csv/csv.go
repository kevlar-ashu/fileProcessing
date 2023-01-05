package csv

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	FILE_NAME = "/csv/students.csv"
)

func ProcessFile() {

	f, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}
	users := scanFile(f)
	// for _, v := range users {
	// 	fmt.Print(v.Id, v.Name, v.LastName, v.Phone, v.Email, v.FriendIds)
	// 	fmt.Println()
	// }

	// sequential processing
	sequentialProcessing(users)

	// fmt.Printf("users %v\n", users)
}

func sequentialProcessing(users []*User) {
	visited := make(map[string]bool)
	for _, user := range users {
		if !visited[user.Id] {
			visited[user.Id] = true
			sendSmsNotification(user)
			for _, friendId := range user.FriendIds {
				friend, err := findUserById(friendId, users)
				if err != nil {
					fmt.Printf("Error %v\n", err)
					continue
				}

				if !visited[friend.Id] {
					visited[friend.Id] = true
					sendSmsNotification(friend)
				}
			}
		}
	}
}

func sendSmsNotification(user *User) {
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("Sending sms notification to %v\n", user.Phone)
}

func findUserById(userId string, users []*User) (*User, error) {
	for _, user := range users {
		if user.Id == userId {
			return user, nil
		}
	}

	return nil, fmt.Errorf("User not found with id %v", userId)
}

func scanFile(f *os.File) []*User {
	s := bufio.NewScanner(f)
	users := []*User{}
	for s.Scan() {
		line := strings.Trim(s.Text(), " ")
		lineArray := strings.Split(line, ",")
		ids := strings.Split(lineArray[5], " ")
		ids = ids[1 : len(ids)-1]
		user := &User{
			Id:        lineArray[0],
			Name:      lineArray[1],
			LastName:  lineArray[2],
			Email:     lineArray[3],
			Phone:     lineArray[4],
			FriendIds: ids,
		}
		users = append(users, user)
	}
	return users
}

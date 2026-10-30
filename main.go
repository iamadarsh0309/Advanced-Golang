package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	now := time.Now()

	id := getUserByName("John")
	fmt.Println(id)

	chats := getUserChats(id)
	friends := getUserFriends(id)

	log.Println(chats)
	log.Println(friends)

	log.Println(time.Since(now))

}

func getUserFriends(id string) []string{
	time.Sleep(time.Second * 2)

	return []string{
		"Jane",
		"Joe",
		"Chandler",
		"Friend-5",
	}

}


func getUserChats(id string) []string {
	time.Sleep(time.Second * 2)

	return []string{
		"Jane",
		"Joe",
		"Chandler",
	}
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("%s-2", name)
}

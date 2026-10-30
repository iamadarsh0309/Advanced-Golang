package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {

	now := time.Now()

	id := getUserByName("John")
	fmt.Println(id)

	wg := sync.WaitGroup{}

	ch := make(chan *Message,2)

	wg.Add(2)

	go getUserChats(id, ch, &wg)
	go getUserFriends(id, ch, &wg)

	wg.Wait()
	close(ch)

	for msg := range ch {
		log.Println(msg)
	}
	// log.Println(chats)
	// log.Println(friends)

	log.Println(time.Since(now))

}

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 2)

	friends := []string{
		"Jane",
		"Joe",
		"Chandler",
		"Friend-5",
	}

	ch <- &Message{
		friends: friends,
	}

}

func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 2)

	chats := []string{
		"Chat-1",
		"Chat-2",
		"Chat-3",
		"Chat-4",
	}

	ch <- &Message{
		chats: chats,
	}
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("%s-2", name)
}

package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func doClientWork(client *rpc.Client) {
	//initial one go routine for watching change
	go func() {
		var keyChanged string

		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("key changed: ", keyChanged)
	}()

	err := client.Call("KVStoreService.Set", [2]string{"abc", "abc-value"}, new(struct{}))
	err = client.Call(
		"KVStoreService.Set", [2]string{"abc", "another-value"},
		new(struct{}),
	)

	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	doClientWork(client)
}

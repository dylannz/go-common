package redis

import (
	"fmt"
)

// Subscribe creates a subscription to the redis publishing system
//
// Messages will be passed into handleResponse.
// Subscribe will block forever, so a goroutine is recommended
func (c Cache) Subscribe(subscription string, handleResponse func(interface{}) ) {
	conn := c.Conn()
	defer conn.Close()

	_, err := conn.Do("PSUBSCRIBE", subscription)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Subscribe loop")
	for {
		reply, err := conn.Receive()
		if err != nil {
			fmt.Println(err)
		}

		handleResponse(reply)
	}
	fmt.Println("Subscribe loop")

}

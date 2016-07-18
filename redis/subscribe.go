package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// Subscribe creates a subscription to the redis publishing system
//
// Messages will be passed into handleResponse.
// Subscribe will block forever, so a goroutine is recommended
func (c Cache) Subscribe(subscription string, handleResponse func(string) ) {
	conn := c.Conn()
	defer conn.Close()

	_, err := conn.Do("PSUBSCRIBE", subscription)
	if err != nil {
		fmt.Println(err)
	}

	for {
		reply, err := redis.Strings(conn.Receive())
		if err != nil {
			fmt.Println(err)
		}

		//The third element of the reply is the redis message
		handleResponse(reply[3])
	}
}

package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// Subscribe a
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

		//The third element of the reply is the actual redis message
		handleResponse(reply[3])
	}
}

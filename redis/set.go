package redis

// Set adds a new key value pair to the redis cache.
func (c Cache) Set(key, val string) error {
	conn := c.Conn()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	if err != nil {
		return err
	}

	return nil
}

// SetEx adds a new key value pair to the redis cache with expire time in seconds
func (c Cache) SetEx(key, val string, expireTime int) error {
	conn := c.Conn()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, expireTime, val)
	if err != nil {
		return err
	}

	return nil
}

package utils

import (
	radix "github.com/mediocregopher/radix/v3"
)

// RedisPing determines if there is a redis server
// sitting and listening at the address provided by
// the REDIS_URL environment variable.
func RedisPing(client radix.Client) (string, error) {
	var response string
	err := client.Do(radix.Cmd(&response, "PING"))

	return response, err
}

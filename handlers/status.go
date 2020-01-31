package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	radix "github.com/mediocregopher/radix/v3"

	"github.com/yashdalfthegray/hue-remote-v2/models"
	"github.com/yashdalfthegray/hue-remote-v2/utils"
)

// Status sends out the status of the server as a health check
func Status(w http.ResponseWriter, req *http.Request) {
	currentStatus := models.NewStatusResponse()
	redisURL, _ := os.LookupEnv("REDIS_URL")
	client, poolErr := radix.NewPool("redis", redisURL, 10)
	if poolErr != nil {
		currentStatus.RedisServerFound = false
	} else {
		response, respErr := utils.RedisPing(client)
		if respErr != nil {
			currentStatus.RedisServerFound = false
		} else {
			currentStatus.RedisServerFound = (response == "PONG")
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(currentStatus); err != nil {
		panic(err)
	}
}

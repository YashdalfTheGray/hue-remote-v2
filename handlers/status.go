package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	radix "github.com/mediocregopher/radix/v3"

	"github.com/yashdalfthegray/hue-remote-v2/models"
	"github.com/yashdalfthegray/hue-remote-v2/utils"
)

// Status sends out the status of the server as a health check
func Status(w http.ResponseWriter, req *http.Request) {
	currentStatus := models.NewStatusResponse()

	currentStatus.RedisServerFound = isRedisServerAvailable()
	currentStatus.BridgeFound = isHueBridgeFound()
	currentStatus.BridgeUserFound = isHueBridgeUserFound()
	currentStatus.APITokenFound = isAPITokenFound()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(currentStatus); err != nil {
		panic(err)
	}
}

func isRedisServerAvailable() bool {
	redisURL, _ := os.LookupEnv(utils.EnvRedisURL)
	urlPieces := strings.Split(redisURL, "://")
	client, poolErr := radix.NewPool("tcp", urlPieces[1], 10)
	if poolErr != nil {
		return false
	}

	response, respErr := utils.RedisPing(client)
	if respErr != nil {
		return false
	}

	return (response == "PONG")
}

func isHueBridgeFound() bool {
	bridgeAddress, _ := os.LookupEnv(utils.EnvHueBridgeAddress)
	return bridgeAddress != ""
}

func isHueBridgeUserFound() bool {
	bridgeUser, _ := os.LookupEnv(utils.EnvHueBridgeUsername)
	return bridgeUser != ""
}

func isAPITokenFound() bool {
	apiToken, _ := os.LookupEnv(utils.EnvHueRemoteToken)
	return apiToken != ""
}

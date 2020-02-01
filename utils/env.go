package utils

import (
	"os"

	"github.com/yashdalfthegray/hue-remote-v2/models"
)

// CheckEnv checks that the right keys are present in the
// environment. Optionally takes a lookup function and falls
// back to os.LookupEnv if one is not provided
func CheckEnv(lookupFunc func(key string) (string, bool)) (err error) {
	if lookupFunc == nil {
		lookupFunc = os.LookupEnv
	}

	envVars := [4]string{
		"HUE_BRIDGE_ADDRESS",
		"HUE_BRIDGE_USERNAME",
		"HUE_REMOTE_TOKEN",
		"REDIS_URL",
	}

	for _, v := range envVars {
		if _, ok := lookupFunc(v); !ok {
			return models.NewMissingEnvError(v)
		}
	}

	return nil
}

// GetPort either reads the port from the environment variable
// called PORT or returns a default, which is 8080
func GetPort(lookupFunc func(key string) (string, bool)) string {
	if lookupFunc == nil {
		lookupFunc = os.LookupEnv
	}

	if port, ok := lookupFunc("PORT"); ok {
		return port
	}

	return "8080"
}

package utils

import (
	"os"

	"github.com/yashdalfthegray/hue-remote-v2/models"
)

// CheckEnv checks that the right keys are present
// in the environment
func CheckEnv() (err error) {
	envVars := [4]string{
		"HUE_BRIDGE_ADDRESS",
		"HUE_BRIDGE_USERNAME",
		"HUE_REMOTE_TOKEN",
		"REDIS_URL",
	}

	for i := 0; i < len(envVars); i++ {
		if _, ok := os.LookupEnv(envVars[i]); !ok {
			return models.NewMissingEnvError(envVars[i])
		}
	}

	return nil
}

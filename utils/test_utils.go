package utils

// Must panics if an error is thrown or returns the value
// from the function call. Added as a test utility
func Must(value interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}

	return value
}

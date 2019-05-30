package models

// HueAction is a struct that represents the request going to
// the Hue bridge from the user
type HueAction struct {
	Effect    string `json:"effect"`
	Colormode string `json:"colormode"`
}

// HueState is a struct that represents the response coming
// back from the Hue bridge
type HueState struct {
	Effect    string `json:"effect"`
	Colormode string `json:"colormode"`
	Reachable string `json:"reachable"`
}

package models

// HueAction is a struct that represents the request going to
// the Hue bridge from the user
type HueAction struct {
	Effect    string    `json:"effect"`
	Colormode string    `json:"colormode"`
	Alert     string    `json:"alert"`
	On        bool      `json:"on"`
	Hue       uint16    `json:"hue"`
	Sat       uint8     `json:"sat"`
	Bri       uint8     `json:"bri"`
	XY        []float64 `json:"xy"`
	CT        uint16    `json:"ct"`
}

// HueState is a struct that represents the response coming
// back from the Hue bridge
type HueState struct {
	Effect    string    `json:"effect"`
	Colormode string    `json:"colormode"`
	Reachable string    `json:"reachable"`
	Alert     string    `json:"alert"`
	On        bool      `json:"on"`
	Hue       uint16    `json:"hue"`
	Sat       uint8     `json:"sat"`
	Bri       uint8     `json:"bri"`
	XY        []float64 `json:"xy"`
	CT        uint16    `json:"ct"`
}

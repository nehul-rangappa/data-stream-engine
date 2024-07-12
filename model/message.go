package model

// Message defines the attributes of the data stream
type Message struct {
	UserID     string `json:"user_id"`
	AppVersion string `json:"app_version"`
	DeviceType string `json:"device_type"`
	IP         string `json:"ip"`
	Locale     string `json:"locale"`
	DeviceID   string `json:"device_id"`
	Timestamp  int64  `json:"timestamp,omitempty"`
	DateTime   string `json:"datetime"`
}

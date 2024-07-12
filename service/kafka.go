package service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nehul-rangappa/data-stream-engine/model"
)

// DataInsights defines the insights obtained from the streaming data
type DataInsights struct {
	IPTraffic              map[string]int `json:"ip_traffic"`
	DeviceTypeDistribution map[string]int `json:"device_type_distribution"`
	AppVersionUsage        map[string]int `json:"app_version_usage"`
	LocaleTraffic          map[string]int `json:"locale_traffic"`
}

// New is a service layer factory function
func New() *DataInsights {
	return &DataInsights{
		IPTraffic:              make(map[string]int),
		DeviceTypeDistribution: make(map[string]int),
		AppVersionUsage:        make(map[string]int),
		LocaleTraffic:          make(map[string]int),
	}
}

// ProcessData method takes a message, process it
// converts the timestamp to human readable version
// filters out the messages from the locale IN
// captures insights with aggregation from the data stream and
// returns the processed message
func (s *DataInsights) ProcessData(message model.Message) []byte {
	// Make the timestamp human-readable
	message.DateTime = time.Unix(message.Timestamp, 0).Format(time.RFC3339)
	message.Timestamp = 0

	s.IPTraffic[message.IP] += 1
	s.DeviceTypeDistribution[message.DeviceType] += 1
	s.AppVersionUsage[message.AppVersion] += 1
	s.LocaleTraffic[message.Locale] += 1

	log.Printf("Data Insights: %v", s)

	// Filter messages for the locale - IN
	if message.Locale != "IN" {
		return nil
	}

	// Convert the processed message back to JSON
	processedMessage, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshalling processed message: %v", err)
		return nil
	}

	return processedMessage
}

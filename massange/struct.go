package massange

import "time"

type state int

// Message format to share between sender and consumer
type Message struct {
	ID              int32      `json:id`
	ExternalID      string     `json: "externalID"`
	Dst             string     `json:"Dst"`
	Message         string     `json: "message"`
	Src             string     `json: "src"`
	State           state      `json: "state"`
	CreatedAt       *time.Time `json: "createdAt"`
	LastUpdatedDate *time.Time `json: "lastUpdatedDate"`
	SMSCMessageID   string     `json: "SMSCMessageID"`
}

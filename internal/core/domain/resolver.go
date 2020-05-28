package domain

import "time"

type Resolver struct {
	PhoneNumber uint32    `json:"phoneNumber"`
	Company     string    `json:"company"`
	XDay        time.Time `json:"xDay"`
}

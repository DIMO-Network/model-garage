package ruptela

import (
	"encoding/json"
	"time"
)

type RuptelaEvent struct {
	DS             string          `json:"ds"`
	Signature      string          `json:"signature"`
	Time           time.Time       `json:"time"`
	Data           json.RawMessage `json:"data"`
	VehicleTokenID *uint32         `json:"vehicleTokenId"`
	DeviceTokenID  *uint32         `json:"deviceTokenId"`
}

type DataContent struct {
	Signals map[string]string `json:"signals"`
}

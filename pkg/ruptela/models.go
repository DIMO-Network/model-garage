package ruptela

import (
	"encoding/json"
)

type RuptelaEvent struct {
	DS             string          `json:"ds"`
	Signature      string          `json:"signature"`
	Time           string          `json:"time"`
	Data           json.RawMessage `json:"data"`
	VehicleTokenID *uint32         `json:"vehicleTokenId"`
	DeviceTokenID  *uint32         `json:"deviceTokenId"`
}

type DataContent struct {
	Signals map[string]string `json:"signals"`
}

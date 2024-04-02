package entity

type Quote struct {
	CarrierID  string  `json:"carrier_id"`
	Name       string  `json:"name"`
	Service    string  `json:"service"`
	FinalPrice float64 `json:"final_price"`
	Deadline   string  `json:"deadline"`
}

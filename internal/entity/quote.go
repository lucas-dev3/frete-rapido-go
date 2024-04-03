package entity

type Quote struct {
	Service    string  `json:"service"`
	FinalPrice float64 `json:"final_price"`
	Deadline   string  `json:"deadline"`
	Carrier    struct {
		CarrierID string `json:"carrier_id"`
		Name      string `json:"name"`
	} `json:"carrier"`
	Volumes []struct {
		Category      int     `json:"category"`
		Amount        int     `json:"amount"`
		Price         float64 `json:"price"`
		Height        float64 `json:"height"`
		Width         float64 `json:"width"`
		Length        float64 `json:"length"`
		UnitaryWeight float64 `json:"unitary_weight"`
	} `json:"volumes"`
	Recipient struct {
		Address struct {
			Zipcode int64 `json:"zipcode"`
		} `json:"address"`
	}
}

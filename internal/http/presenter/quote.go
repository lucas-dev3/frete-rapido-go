package presenter

import (
	"strconv"

	"github.com/lucas-dev3/frete-rapido-go.git/config"
	"github.com/lucas-dev3/frete-rapido-go.git/internal/entity"
)

// QuoteRequest representa a estrutura de dados para uma solicitação de cotação
type QuoteRequest struct {
	Recipient struct {
		Address struct {
			Zipcode string `json:"zipcode"`
		} `json:"address"`
	} `json:"recipient"`
	Volumes []struct {
		Category      int     `json:"category"`
		Amount        int     `json:"amount"`
		UnitaryWeight float64 `json:"unitary_weight"`
		Price         float64 `json:"price"`
		SKU           string  `json:"sku"`
		Height        float64 `json:"height"`
		Width         float64 `json:"width"`
		Length        float64 `json:"length"`
	} `json:"volumes"`
}

// QuoteResponse representa a estrutura de dados para a resposta de uma cotação
type QuoteResponse struct {
	Carrier []struct {
		Name     string  `json:"name"`
		Service  string  `json:"service"`
		Deadline string  `json:"deadline"`
		Price    float64 `json:"price"`
	} `json:"carrier"`
}

type FreightQuoteRequest struct {
	Shipper struct {
		RegisteredNumber string `json:"registered_number"`
		Token            string `json:"token"`
		PlataformCode    string `json:"plataform_code"`
	} `json:"shipper"`
	Recipient struct {
		Type             int    `json:"type"`
		RegisteredNumber string `json:"registered_number"`
		Country          string `json:"country"`
		Zipcode          int64  `json:"zipcode"`
	} `json:"recipient"`
	Dispatchers []struct {
		RegisteredNumber string  `json:"registered_number"`
		Zipcode          int64   `json:"zipcode"`
		TotalPrice       float64 `json:"total_price"`
		Volumes          []struct {
			Amount        int     `json:"amount"`
			Category      string  `json:"category"`
			Height        int     `json:"height"`
			Length        int     `json:"length"`
			UnitaryPrice  float64 `json:"unitary_price"`
			UnitaryWeight float64 `json:"unitary_weight"`
		}
	}
	SimulationType []int16 `json:"simulation_type"`
}

type FreightQuoteResponse struct {
	Offers []struct {
		Carrier struct {
			Name      string `json:"name"`
			Logo      string `json:"logo"`
			Reference string `json:"reference"`
		}
		Service      string `json:"service"`
		DeliveryTime struct {
			Days int `json:"days"`
		}
		FinalPrice float64 `json:"final_price"`
	} `json:"offers"`
}

func NewQuotesResponse(quote []*entity.Quote) []*QuoteResponse {
	var response []*QuoteResponse

	for _, q := range quote {
		newQuoteRespo := &QuoteResponse{}
		newQuoteRespo.Carrier = append(newQuoteRespo.Carrier, struct {
			Name     string  `json:"name"`
			Service  string  `json:"service"`
			Deadline string  `json:"deadline"`
			Price    float64 `json:"price"`
		}{
			Name:     q.Name,
			Service:  q.Service,
			Deadline: q.Deadline,
			Price:    q.FinalPrice,
		})

		response = append(response, newQuoteRespo)
	}

	return response
}

func NewFreightQuoteRequest(request *QuoteRequest) *FreightQuoteRequest {
	freightRequest := &FreightQuoteRequest{}

	zipCode, _ := strconv.ParseInt(request.Recipient.Address.Zipcode, 10, 64)
	freightRequest.Recipient.Zipcode = zipCode
	freightRequest.Recipient.Country = "BR"
	freightRequest.Recipient.RegisteredNumber = config.GetEnvVars().FRRegisteredNum
	freightRequest.Recipient.Type = 0

	freightRequest.Shipper.PlataformCode = config.GetEnvVars().FRPlataformCode
	freightRequest.Shipper.RegisteredNumber = config.GetEnvVars().FRRegisteredNum
	freightRequest.Shipper.Token = config.GetEnvVars().FRToken

	for _, volume := range request.Volumes {
		newVolume := struct {
			Amount        int     `json:"amount"`
			Category      string  `json:"category"`
			Height        int     `json:"height"`
			Length        int     `json:"length"`
			UnitaryPrice  float64 `json:"unitary_price"`
			UnitaryWeight float64 `json:"unitary_weight"`
		}{
			Amount:        volume.Amount,
			Category:      strconv.Itoa(volume.Category),
			Height:        int(volume.Height),
			Length:        int(volume.Length),
			UnitaryPrice:  volume.Price,
			UnitaryWeight: volume.UnitaryWeight,
		}

		if len(freightRequest.Dispatchers) == 0 {
			freightRequest.Dispatchers = append(freightRequest.Dispatchers, struct {
				RegisteredNumber string  `json:"registered_number"`
				Zipcode          int64   `json:"zipcode"`
				TotalPrice       float64 `json:"total_price"`
				Volumes          []struct {
					Amount        int     `json:"amount"`
					Category      string  `json:"category"`
					Height        int     `json:"height"`
					Length        int     `json:"length"`
					UnitaryPrice  float64 `json:"unitary_price"`
					UnitaryWeight float64 `json:"unitary_weight"`
				}
			}{
				RegisteredNumber: config.GetEnvVars().FRRegisteredNum,
				Zipcode:          zipCode,
				TotalPrice:       0,
			})

			freightRequest.Dispatchers[0].Volumes = append(freightRequest.Dispatchers[0].Volumes, newVolume)
		}
	}

	freightRequest.SimulationType = append(freightRequest.SimulationType, 0)

	return freightRequest
}

// func NewFreightQuoteResponse(request *FreightQuoteResponse)

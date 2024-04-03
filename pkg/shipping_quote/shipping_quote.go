package shippingquote

import (
	"context"

	"github.com/lucas-dev3/frete-rapido-go.git/internal/entity"
)

type ShippingQuoteResponse struct {
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

type ShippingQuoteRequest struct {
	Zipcode int64 `json:"zipcode"`
	Volumes []struct {
		Category      int     `json:"category"`
		Amount        int     `json:"amount"`
		Price         float64 `json:"price"`
		Height        float64 `json:"height"`
		Width         float64 `json:"width"`
		Length        float64 `json:"length"`
		UnitaryWeight float64 `json:"unitary_weight"`
	} `json:"volumes"`
}

type ShippingQuoteFetcher interface {
	SimulateQuote(ctx context.Context, request *ShippingQuoteRequest) (*ShippingQuoteResponse, error)
}

func ToShippingQuoteEntity(e *entity.Quote) *ShippingQuoteRequest {
	return &ShippingQuoteRequest{
		Zipcode: e.Recipient.Address.Zipcode,
		Volumes: e.Volumes,
	}
}

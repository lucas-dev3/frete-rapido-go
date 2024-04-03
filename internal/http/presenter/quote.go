package presenter

import (
	"strconv"

	"github.com/lucas-dev3/frete-rapido-go.git/internal/entity"
)

// QuoteRequest representa a estrutura de dados para uma solicitação de cotação
type QuoteRequest struct {
	Recipient struct {
		Address struct {
			Zipcode string `json:"zipcode" validate:"required"`
		} `json:"address" validate:"required"`
	} `json:"recipient" validate:"required"`
	Volumes []struct {
		Category      int     `json:"category" validate:"required"`
		Amount        int     `json:"amount" validate:"required"`
		UnitaryWeight float64 `json:"unitary_weight" validate:"required"`
		Price         float64 `json:"price" validate:"required"`
		SKU           string  `json:"sku" validate:"required"`
		Height        float64 `json:"height" validate:"required"`
		Width         float64 `json:"width" validate:"required"`
		Length        float64 `json:"length" validate:"required"`
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

func (qr *QuoteRequest) ToQuoteEntity() *entity.Quote {
	quote := &entity.Quote{}

	zipcode, err := strconv.ParseInt(qr.Recipient.Address.Zipcode, 10, 64)
	if err != nil {
		return nil
	}
	quote.Recipient.Address.Zipcode = zipcode

	for _, volume := range qr.Volumes {
		newVolume := struct {
			Category      int     `json:"category"`
			Amount        int     `json:"amount"`
			Price         float64 `json:"price"`
			Height        float64 `json:"height"`
			Width         float64 `json:"width"`
			Length        float64 `json:"length"`
			UnitaryWeight float64 `json:"unitary_weight"`
		}{
			Category:      volume.Category,
			Amount:        volume.Amount,
			Price:         volume.Price,
			Height:        volume.Height,
			Width:         volume.Width,
			Length:        volume.Length,
			UnitaryWeight: volume.UnitaryWeight,
		}

		quote.Volumes = append(quote.Volumes, newVolume)
	}

	return quote
}

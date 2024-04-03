package freterapido

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/lucas-dev3/frete-rapido-go.git/config"
	shippingquote "github.com/lucas-dev3/frete-rapido-go.git/pkg/shipping_quote"
)

type FreteRapido struct {
	BaseURL string
}

type FreteRapidoRequest struct {
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

type Response struct {
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

func New() *FreteRapido {
	return &FreteRapido{
		BaseURL: config.GetEnvVars().FRBaseURL,
	}
}

func (frc *FreteRapido) SimulateQuote(ctx context.Context, request *shippingquote.ShippingQuoteRequest) (*shippingquote.ShippingQuoteResponse, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, errors.New("erro ao tratar dados")
	}

	// Convertendo requestBody para string para impressão legível
	fmt.Println("Request body:", string(requestBody))

	req, err := http.NewRequest("POST", frc.BaseURL+"/quote/simulate", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx) // Associa o contexto passado ao cliente HTTP

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Log do status code para debug
	fmt.Println("Response status code:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body) // Usando io.ReadAll aqui
		return nil, fmt.Errorf("erro ao calcular frete: %s", string(responseBody))
	}

	// Deserializar o corpo da resposta
	var quoteResponse shippingquote.ShippingQuoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&quoteResponse); err != nil {
		return nil, errors.New("erro ao deserializar a resposta")
	}

	return &quoteResponse, nil
}

func ToShippingQuoteEntity(quote *shippingquote.ShippingQuoteRequest) *FreteRapidoRequest {
	var frr FreteRapidoRequest

	// Configurações do remetente (Shipper)
	frr.Shipper.RegisteredNumber = config.GetEnvVars().FRRegisteredNum
	frr.Shipper.Token = config.GetEnvVars().FRToken
	frr.Shipper.PlataformCode = config.GetEnvVars().FRPlataformCode

	// Configurações do destinatário (Recipient)
	frr.Recipient.Type = 1
	frr.Recipient.RegisteredNumber = config.GetEnvVars().FRRegisteredNum
	frr.Recipient.Country = "BR"
	frr.Recipient.Zipcode = quote.Zipcode

	// Tipo de Simulação
	frr.SimulationType = []int16{1}

	// Configurações dos Despachantes (Dispatchers)
	for _, v := range quote.Dispatchers {
		newDispatcher := struct {
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
		}{}
	}

	return &frr
}

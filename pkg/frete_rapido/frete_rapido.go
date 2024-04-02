package freterapido

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/lucas-dev3/frete-rapido-go.git/config"
	"github.com/lucas-dev3/frete-rapido-go.git/internal/http/presenter"
)

type FreteRapido struct {
	BaseURL string
}

func NewService() *FreteRapido {
	return &FreteRapido{
		BaseURL: config.GetEnvVars().FRBaseURL,
	}
}

func (frc *FreteRapido) CalculateQuote(request *presenter.FreightQuoteRequest) (*presenter.FreightQuoteResponse, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, errors.New("erro ao tratar dados")
	}

	req, err := http.NewRequest("POST", frc.BaseURL+"/quote/simulate", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}

	// Faz a requisição
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	print(resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("erro ao calcular frete")
	}

	return nil, nil
}

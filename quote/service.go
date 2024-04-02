package quote

import (
	"context"
	"fmt"
	"log"

	"github.com/lucas-dev3/frete-rapido-go.git/internal/entity"
	"github.com/lucas-dev3/frete-rapido-go.git/internal/http/presenter"
	freterapido "github.com/lucas-dev3/frete-rapido-go.git/pkg/frete_rapido"
)

type Service struct {
	repo        Repository
	freterapido *freterapido.FreteRapido
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ProcessQuote(ctx context.Context, quote *presenter.QuoteRequest) error {
	log.Println("[Service] ProcessQuote started")

	fmt.Printf("Quote: %+v\n", *quote)

	resp, err := s.freterapido.CalculateQuote(presenter.NewFreightQuoteRequest(quote))
	if err != nil {
		log.Println("[Service] Error while processing quote")
		return err
	}

	fmt.Printf("Response: %+v\n", resp)

	return nil
}

func (s *Service) GetMetrics(ctx context.Context) ([]*entity.Quote, error) {
	return s.repo.FindAll(ctx)
}

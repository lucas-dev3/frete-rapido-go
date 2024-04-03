package quote

import (
	"context"
	"fmt"
	"log"

	"github.com/lucas-dev3/frete-rapido-go.git/internal/entity"
	shippingquote "github.com/lucas-dev3/frete-rapido-go.git/pkg/shipping_quote"
)

type Service struct {
	repo                 Repository
	ShippingQuoteFetcher shippingquote.ShippingQuoteFetcher
}

func NewService(repo Repository, sqf shippingquote.ShippingQuoteFetcher) *Service {
	return &Service{
		repo:                 repo,
		ShippingQuoteFetcher: sqf,
	}
}

func (s *Service) ProcessQuote(ctx context.Context, quote *entity.Quote) error {
	log.Println("[Service] ProcessQuote started")
	log.Println("estou no service useCase")

	qr := shippingquote.ToShippingQuoteEntity(quote)

	resp, err := s.ShippingQuoteFetcher.SimulateQuote(ctx, qr)
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
